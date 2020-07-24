package modules

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

const (
	URL_NATS = "http://nats:4222"
)

func PingNats() (err error) {
	nc, err := nats.Connect(URL_NATS, nats.PingInterval(20*time.Second), nats.MaxPingsOutstanding(5))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer nc.Close()
	return err
}

func ConnectNats() (nc *nats.Conn, err error) {
	nc, err = nats.Connect(URL_NATS, nats.Timeout(time.Second*10))
	log.Println("Connected to " + URL_NATS)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer nc.Close()
	return nc, err
}
