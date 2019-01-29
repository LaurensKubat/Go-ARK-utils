package webhooks

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func newTestClient() (*Client, error) {
	return NewClient(os.Getenv("TESTURL"))
}

func TestNewClient(t *testing.T) {
	assert.NotPanics(t, func() {
		c, err := NewClient(os.Getenv("TESTURL"))
		assert.NoError(t, err)
		assert.NotNil(t, c)
	})
}

func TestNewCondition(t *testing.T) {
	assert.NotPanics(t, func() {
		NewCondition("1", Equal, "1")
	})
}

func TestClient_Create(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Create(BlockForged, "fakeurl", "true", []Condition{
		NewCondition("AHXtmB84sTZ9Zd35h9Y1vfFvPE2Xzqj8ri", Equal, "AHXtmB84sTZ9Zd35h9Y1vfFvPE2Xzqj8ri"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, res.Data.Token)
	assert.NotNil(t, res.Data.ID)
}

func TestClient_Delete(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	err = c.Delete("1")
	assert.NoError(t, err)
}

func TestClient_List(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.List(0, 0)
	assert.NoError(t, err)
	assert.NotNil(t, res.Data[0].Target)
	assert.NotNil(t, res.Data[0].ID)
}

func TestClient_Retrieve(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	res, err := c.Retrieve("1")
	assert.NoError(t, err)
	assert.NotNil(t, res.Data.ID)
	assert.NotNil(t, res.Data.Target)
}

func TestClient_Update(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)
	err = c.Update("1", BlockForged, "fakeurl", "true", []Condition{
		NewCondition("AHXtmB84sTZ9Zd35h9Y1vfFvPE2Xzqj8ri", NotEqual, "asdasda"),
	})
	require.NoError(t, err)
}
