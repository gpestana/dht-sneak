module dhtsneak

require (
	github.com/ipfs/go-cid v0.0.1
	github.com/ipfs/go-log v0.0.1
	github.com/libp2p/go-libp2p v0.0.2
	github.com/libp2p/go-libp2p-kad-dht v0.0.4
	github.com/libp2p/go-libp2p-peer v0.0.1
	github.com/libp2p/go-libp2p-peerstore v0.0.1
	github.com/multiformats/go-multiaddr v0.0.2
	github.com/whyrusleeping/go-logging v0.0.0-20170515211332-0457bb6b88fc
)

replace github.com/libp2p/go-libp2p-kad-dht v0.0.4 => ./go-libp2p-kad-dht
