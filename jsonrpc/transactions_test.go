package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransactions_Info(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Transactions.Info("db1aa687737858cc9199bfa336f9b1c035915c30aaee60b1e0f8afadfdb946bd")
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
}

func TestTransactions_Create_Broadcast(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	// Passphrase here is the passphrase of a genesis delegate on the testnet
	res, err := c.Transactions.Create("1", "19LmwF3rG9JehYdcf99zQQv9sd2bBJ83G6",
		"clay harbor enemy utility margin pretty hub comic piece aerobic umbrella acquire")
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
	res2, err := c.Transactions.Broadcast([]string{res.ID})
	assert.NoError(t, err)
	assert.NotNil(t, res2.Result)
}

func TestTransactions_Bip38Create(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Transactions.Bip38Create("1", "AHXtmB84sTZ9Zd35h9Y1vfFvPE2Xzqj8ri",
		"5KWQF4hrHPUoCeKfyojyJZHjePpEJBvNr22qbudcARZUUXhCfDJ", "19LmwF3rG9JehYdcf99zQQv9sd2bBJ83G6")
	assert.NoError(t, err)
	assert.NotNil(t, res.Result)
}
