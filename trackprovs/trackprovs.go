package main

import (
	"context"
	"fmt"
)

type Tracker interface {
	track(ctx context.Context, id string, confs config) error
}

type config struct {
	output       string
	interval     int
	queryTimeout int
}

func main() {
	//tracks IPFS content providers and saves results to output
	// #TODO refactor as CLI
	contentId := "QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv"

	confs := config{
		output:       fmt.Sprintf("provs-%v.out", contentId),
		interval:     10,
		queryTimeout: 180,
	}

	ctx := context.Background()
	t := NewIpfsTracker()
	go func() {
		t.track(ctx, contentId, confs)
	}()

	select {}
}
