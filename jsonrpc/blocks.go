package jsonrpc

import (
	"encoding/json"
)

// Blocks is the service supporting all block related api calls in https://docs.ark.io/api/json-rpc/blocks.html
type Blocks Service

// Info gets information about a block.
func (b *Blocks) Info(blockID string) (BlocksInfo, error) {
	r, err := b.client.prepBody("blocks.info", Params{ID: blockID})
	if err != nil {
		return BlocksInfo{}, err
	}
	res, err := b.client.send(contenttype, r)
	if err != nil {
		return BlocksInfo{}, err
	}
	msg, err := b.client.handleResponse(*res)
	if err != nil {
		return BlocksInfo{}, err
	}
	buf := BlocksInfoResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return BlocksInfo{}, err
	}
	return BlocksInfo{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}

// Latest gets the latest block.
func (b *Blocks) Latest() (BlocksLatest, error) {
	r, err := b.client.prepBody("blocks.latest", Params{})
	if err != nil {
		return BlocksLatest{}, err
	}
	res, err := b.client.send(contenttype, r)
	if err != nil {
		return BlocksLatest{}, err
	}
	msg, err := b.client.handleResponse(*res)
	if err != nil {
		return BlocksLatest{}, err
	}
	buf := BlocksLatestResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return BlocksLatest{}, err
	}
	return BlocksLatest{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}

// Transactions gets all Transactions of a block.
func (b *Blocks) Transactions(blockID string, offset int) (BlocksTransactions, error) {
	r, err := b.client.prepBody("blocks.Transactions", Params{ID: blockID, Offset: offset})
	if err != nil {
		return BlocksTransactions{}, err
	}
	res, err := b.client.send(contenttype, r)
	if err != nil {
		return BlocksTransactions{}, err
	}
	msg, err := b.client.handleResponse(*res)
	if err != nil {
		return BlocksTransactions{}, err
	}
	buf := BlocksTransactionsResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return BlocksTransactions{}, err
	}
	return BlocksTransactions{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}
