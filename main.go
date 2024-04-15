package main

import (
	"log"

	"github.com/bumblebee47293/Distributed-File-Storage-System/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
