package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWallets_Info(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Wallets.Info("AHXtmB84sTZ9Zd35h9Y1vfFvPE2Xzqj8ri")
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Jsonrpc)
}

func TestWallets_Transactions(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Wallets.Transactions("AHXtmB84sTZ9Zd35h9Y1vfFvPE2Xzqj8ri", 0)
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Jsonrpc)
}

func TestWallets_Create(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Wallets.Create("testytest")
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Jsonrpc)
}

func TestWallets_Bip38Create(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Wallets.Bip38Create("5KWQF4hrHPUoCeKfyojyJZHjePpEJBvNr22qbudcARZUUXhCfDJ",
		"19LmwF3rG9JehYdcf99zQQv9sd2bBJ83G6")
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Jsonrpc)
}

func TestWallets_Bip38Info(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Wallets.Bip38Info("19LmwF3rG9JehYdcf99zQQv9sd2bBJ83G6")
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Jsonrpc)
}
