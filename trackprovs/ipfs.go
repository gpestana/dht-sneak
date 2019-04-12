package main

import (
	"context"
	"encoding/json"
	"fmt"
	cid "github.com/ipfs/go-cid"
	//	u "github.com/ipfs/go-ipfs-util"
	"github.com/libp2p/go-libp2p"
	kad "github.com/libp2p/go-libp2p-kad-dht"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	//	mh "github.com/multiformats/go-multihash"
	ma "github.com/multiformats/go-multiaddr"
	"io/ioutil"
	"log"
	"time"
)

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

	bpeers := configBootstrapPeers()
	for _, paddr := range bpeers {
		peerinfo, _ := pstore.InfoFromP2pAddr(paddr)
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
	id, err := cid.Parse(contentId)
	if err != nil {
		return err
	}
	log.Println("finding providers for", id)

	for {
		fmt.Printf(".")
		tctx, _ := context.WithTimeout(ctx, time.Duration(c.queryTimeout))
		pinfos, err := t.dht.FindProviders(tctx, id)
		if err != nil {
			log.Println(err)
		}

		for _, pi := range pinfos {
			jpi, err := json.Marshal(pi)
			if err != nil {
				log.Fatal(err)
			}
			err = ioutil.WriteFile(c.output, jpi, 0644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(time.Now(), pi)
		}

		time.Sleep(time.Duration(c.interval) * time.Second)
	}
	return nil
}

func configBootstrapPeers() []ma.Multiaddr {
	var multiaddrs []ma.Multiaddr
	maStr := []string{
		"/dnsaddr/bootstrap.libp2p.io/ipfs/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/ipfs/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/ipfs/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/ipfs/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.236.179.241/tcp/4001/ipfs/QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM",
		"/ip4/104.236.76.40/tcp/4001/ipfs/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64",
		"/ip4/128.199.219.111/tcp/4001/ipfs/QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu",
		"/ip4/178.62.158.247/tcp/4001/ipfs/QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd",
		"/ip6/2400:6180:0:d0::151:6001/tcp/4001/ipfs/QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu",
		"/ip6/2604:a880:1:20::203:d001/tcp/4001/ipfs/QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM",
		"/ip6/2604:a880:800:10::4a:5001/tcp/4001/ipfs/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64",
		"/ip6/2a03:b0c0:0:1010::23:1001/tcp/4001/ipfs/QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd",
	}

	for _, s := range maStr {
		m, err := ma.NewMultiaddr(s)
		if err != nil {
			log.Fatal(err)
		}
		multiaddrs = append(multiaddrs, m)
	}
	return multiaddrs
}
