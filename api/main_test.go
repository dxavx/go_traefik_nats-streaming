package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const endPointPub = "/v1/pub"

func TestPing(t *testing.T) {

	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s"+endPointPub, ts.URL))
	if err != nil {
		t.Fatalf("error GET request  %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error body %v", body)
	}
	assert.Equal(t, 10, len(body))

}
