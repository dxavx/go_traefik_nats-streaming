package modules

import (
	nats "github.com/nats-io/nats.go"
	"log"
)

func Sum(a int, b int) int {
	return a + b
}

func ConnectNats() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
}
