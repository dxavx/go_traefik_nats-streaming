package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

const END_POINT_PING = "http://localhost/v1/ping"
const TIME_OUT = 10

func TestPing(t *testing.T) {

	client := http.Client{Timeout: time.Second * TIME_OUT}
	var requestBody []byte
	request, err := http.NewRequest("GET", END_POINT_PING, bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
