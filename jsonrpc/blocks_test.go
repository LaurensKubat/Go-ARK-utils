package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBlocks_Info(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Blocks.Info("d9acd04bde4234a81addb8482333b4ac906bed7be5a9970ce8ada428bd083192")
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Jsonrpc)
}

func TestBlocks_Latest(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Blocks.Latest()
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Jsonrpc)
}

func TestBlocks_Transactions(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Blocks.Transactions("d9acd04bde4234a81addb8482333b4ac906bed7be5a9970ce8ada428bd083192", 0)
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Jsonrpc)

}
