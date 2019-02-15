package utils

import (
	"fmt"
	cid "github.com/ipfs/go-cid"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	peer "github.com/libp2p/go-libp2p-peer"
)

func Hit(event string, peer peer.ID, msg *pb.Message) {
	id, _ := cid.Cast([]byte(msg.GetKey()))
	fmt.Printf("{ %v: { peer: %s, key: %s }\n", event, peer.Pretty(), id)
}
