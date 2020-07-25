package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"sync"
)

func main() {

	var natsUrl = os.Getenv("NATS_URL")

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Asynchronous Subscriptions
	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("updates", func(m *nats.Msg) {
		fmt.Println(string(m.Data))
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()

	// Synchronous Subscriptions

	//// Subscribe
	//sub, err := nc.SubscribeSync("updates")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Wait for a message
	//msg, err := sub.NextMsg(10 * time.Second)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Use the response
	//log.Printf("Reply: %s", msg.Data)
}
