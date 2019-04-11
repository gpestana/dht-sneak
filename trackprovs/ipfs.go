package main

import (
	"context"
	"fmt"
	cid "github.com/ipfs/go-cid"
	ipfsaddr "github.com/ipfs/go-ipfs-addr"
	u "github.com/ipfs/go-ipfs-util"
	"github.com/libp2p/go-libp2p"
	kad "github.com/libp2p/go-libp2p-kad-dht"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	//	"io/ioutil"
	"log"
	"time"
)

var bootstrapPeers = []string{
	"/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	"/ip4/104.236.179.241/tcp/4001/ipfs/QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM",
	"/ip4/104.236.76.40/tcp/4001/ipfs/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64",
	"/ip4/128.199.219.111/tcp/4001/ipfs/QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu",
	"/ip4/178.62.158.247/tcp/4001/ipfs/QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd",
}

type IpfsTracker struct {
	dht kad.IpfsDHT
}

func NewIpfsTracker() *IpfsTracker {
	ctx := context.Background()

	h, err := libp2p.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ipfsdht, err := kad.New(ctx, h)
	if err != nil {
		log.Fatal(err)
	}

	err = ipfsdht.Bootstrap(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range bootstrapPeers {
		pAddr, _ := ipfsaddr.ParseString(addr)
		peerinfo, _ := pstore.InfoFromP2pAddr(pAddr.Multiaddr())
		if err = h.Connect(ctx, *peerinfo); err != nil {
			log.Println("ERROR: ", err)
		}
	}

	log.Println("dht initialized")

	return &IpfsTracker{
		dht: *ipfsdht,
	}
}

func (t *IpfsTracker) track(ctx context.Context, contentId string, c config) error {
	// string -> multihash -> cid
	mhv := u.Hash([]byte(contentId))
	contId := cid.NewCidV0(mhv)

	log.Println("finding providers of ", contId)
	for {
		log.Println(">>")
		tctx, _ := context.WithTimeout(ctx, time.Duration(10))
		pinfos, err := t.dht.FindProviders(tctx, contId)
		if err != nil {
			return err
		}
		for _, pi := range pinfos {
			// write to output fd as JSON?
			//err := ioutil.WriteFile(c.output, []byte(pi), 0644)
			fmt.Println(time.Now(), pi)
		}

		time.Sleep(time.Duration(c.interval) * time.Second)
	}
	return nil
}
