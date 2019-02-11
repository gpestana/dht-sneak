package main

import (
	"context"
	"fmt"
	kad "github.com/gpestana/go-libp2p-kad-dht"
	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
	golog "github.com/whyrusleeping/go-logging"
	"log"
	"sync"
	"time"
)

func main() {
	logFile := fmt.Sprintf("./%v.log", time.Now().Format(time.RFC3339))
	_ = configLogging(logFile)

	listenAddrs := configListenAddrs()
	ctx := context.Background()
	host, err := libp2p.New(ctx, libp2p.ListenAddrs(listenAddrs...))
	errFatal(err)

	log.Printf("> node init: %v | %v\n", host.ID(), host.Addrs())

	dht, err := kad.New(ctx, host)
	errFatal(err)

	err = dht.Bootstrap(ctx)
	errFatal(err)

	var wg sync.WaitGroup
	for _, peerAddr := range configBootstrapPeers() {
		peerinfo, _ := peerstore.InfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := host.Connect(ctx, *peerinfo); err != nil {
				// log.Println(err)
			} else {
				log.Printf("> dht: connection established with bootstrap node %v [%v]",
					peerinfo.ID, peerinfo.Addrs)
			}
		}()
	}
	wg.Wait()

	log.Printf("\n\n> dht-sneak node running and logging to %v. Press CTRL+C to exit.\n", logFile)
	select {}

}

func configLogging(logFile string) logging.EventLogger {
	logEvent := logging.Logger("dht-sneak")
	golog.SetLevel(golog.INFO, "dht-sneak")
	return logEvent
}

func configListenAddrs() []ma.Multiaddr {
	listenAddrTcp, err := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/0")
	errFatal(err)
	listenAddrUdp, err := ma.NewMultiaddr("/ip4/0.0.0.0/udp/0")
	errFatal(err)
	return []ma.Multiaddr{listenAddrTcp, listenAddrUdp}
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
		errFatal(err)
		multiaddrs = append(multiaddrs, m)
	}
	return multiaddrs
}

func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}