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
	// #TODO refactor for CLI
	contentId := "QmdPtC3T7Kcu9iJg6hYzLBWR5XCDcYMY7HV685E3kH3EcS"

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
