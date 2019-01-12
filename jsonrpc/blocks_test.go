package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBlocks_Info(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Blocks.Info("123141")
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Result)
}

func TestBlocks_Latest(t *testing.T) {

}

func TestBlocks_Transactions(t *testing.T) {

}
