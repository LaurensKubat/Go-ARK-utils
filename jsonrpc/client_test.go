package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testurl = os.Getenv("TESTURL")

// newTestClient is used to get a client for testing purposes
func newTestClient() (*Client, error) {
	c, err := NewClient(testurl)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func Test_NewClient(t *testing.T) {
	assert.NotPanics(t, func() {
		_, err := NewClient(testurl)
		if err != nil {
			panic(err)
		}
	})
}
