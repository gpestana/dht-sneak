package main

import (
	"context"
	"fmt"
)

type Tracker interface {
	track(ctx context.Context, id string, confs config) error
}

type config struct {
	output   string
	interval int
}

func main() {
	//tracks IPFS content providers and saves results to output
	// #TODO refactor for CLI
	contentId := "Qmb1r3Cf1PcU1H3QVhARBa9V4frExfRmAqp1ZVnLXiy24m"

	confs := config{
		output:   fmt.Sprintf("provs-%v.out\n", contentId),
		interval: 10,
	}

	ctx := context.Background()
	t := NewIpfsTracker()
	go func() {
		t.track(ctx, contentId, confs)
	}()

	select {}
}
