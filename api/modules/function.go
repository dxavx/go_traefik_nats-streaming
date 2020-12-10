package modules

import (
	"github.com/nats-io/nats.go"
	"log"
	"math/rand"
	"os"
	"time"
)

var natsUrl = os.Getenv("NATS_URL")

func PingNats() (err error) {
	nc, err := nats.Connect(natsUrl, nats.PingInterval(20*time.Second), nats.MaxPingsOutstanding(5))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer nc.Close()
	return err
}

func ConnectNats() (nc *nats.Conn, err error) {
	nc, err = nats.Connect(natsUrl, nats.Timeout(time.Second*10))
	log.Println("Connected to " + natsUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer nc.Close()
	return nc, err
}

// RandomString
func RandomString(n int) string {

	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	s := make([]byte, n)
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < n; i++ {
		t := rand.Intn(len(letters))
		s[i] = letters[t]
	}
	return string(s)
}
