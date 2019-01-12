package jsonrpc

import "encoding/json"

// Transactions is the service supporting all transactions related api calls in https://docs.ark.io/api/json-rpc/transactions.html
type Transactions Service

// Create creates a transaction. After the creation, the transaction has to be broadcast using Transactions.Broadcast
// amount is the amount in arktoshis, recipientID is the address of the recipient and passphrase is the passphrase of the sender
func (t *Transactions) Create(amount string, recipientID string, passphrase string) (TransactionsCreate, error) {
	r, err := t.client.prepBody("transactions.create", Params{Amount: amount, RecipientID: recipientID,
		Passphrase: passphrase})
	if err != nil {
		return TransactionsCreate{}, err
	}
	res, err := t.client.send(contenttype, r)
	if err != nil {
		return TransactionsCreate{}, err
	}
	msg, err := t.client.handleResponse(*res)
	if err != nil {
		return TransactionsCreate{}, err
	}
	buf := TransactionsCreateResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return TransactionsCreate{}, err
	}
	return TransactionsCreate{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}

// Broadcast broadcasts transactions. Any transaction that you wish to broadcast have to be created through Transactions.Create.
// The ids of the transaction should be the ids returned by Transactions.Create
func (t *Transactions) Broadcast(ids []string) (TransactionsBroadcast, error) {
	r, err := t.client.prepBody("transactions.broadcast", Params{Txs: ids})
	if err != nil {
		return TransactionsBroadcast{}, err
	}
	res, err := t.client.send(contenttype, r)
	if err != nil {
		return TransactionsBroadcast{}, err
	}
	msg, err := t.client.handleResponse(*res)
	if err != nil {
		return TransactionsBroadcast{}, err
	}
	buf := TransactionsBroadcastResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return TransactionsBroadcast{}, err
	}
	return TransactionsBroadcast{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}

// Info gets a transaction
func (t Transactions) Info(id string) (TransactionsInfo, error) {
	r, err := t.client.prepBody("transactions.info", Params{ID: id})
	if err != nil {
		return TransactionsInfo{}, err
	}
	res, err := t.client.send(contenttype, r)
	if err != nil {
		return TransactionsInfo{}, err
	}
	msg, err := t.client.handleResponse(*res)
	if err != nil {
		return TransactionsInfo{}, err
	}
	buf := TransactionsInfoResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return TransactionsInfo{}, err
	}
	return TransactionsInfo{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}

// Bip38Create creates a transaction from a BIP38 wallet
// amount is the amout in arktoshis, recipientID is the address of the recipient, bip38 is the bip38 of the sender,
// userID is the ID of the sender
func (t *Transactions) Bip38Create(amount string, recipientID string, bip38 string, userID string) (TransactionsBIP38Create, error) {
	r, err := t.client.prepBody("transactions.bip38.create", Params{RecipientID: recipientID, Amount: amount,
		Bip38: bip38, UserID: userID})
	if err != nil {
		return TransactionsBIP38Create{}, err
	}
	res, err := t.client.send(contenttype, r)
	if err != nil {
		return TransactionsBIP38Create{}, err
	}
	msg, err := t.client.handleResponse(*res)
	if err != nil {
		return TransactionsBIP38Create{}, err
	}
	buf := TransactionsBIP38CreateResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return TransactionsBIP38Create{}, err
	}
	return TransactionsBIP38Create{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}
