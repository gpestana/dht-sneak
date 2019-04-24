package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//cid := "QmSJqTBmUD3gLTqFoqmEJ3yXKrxgvXu54fUMjhuTpEeHk2"

	//cmd := exec.Command("/bin/bash", "dht findprovs -n 1"+cid)
	cmd := exec.Command("ipfs", "get QmSJqTBmUD3gLTqFoqmEJ3yXKrxgvXu54fUMjhuTpEeHk2")
	log.Println(cmd)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(out))
}
