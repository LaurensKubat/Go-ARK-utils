package jsonrpc

import "encoding/json"

// Wallets is the service supporting all wallet related api calls in https://docs.ark.io/api/json-rpc/wallets.html
type Wallets Service

// Create creates a wallet with passphrase as the passphrase to be used for the wallet
func (w *Wallets) Create(passphrase string) (WalletsCreate, error) {
	r, err := w.client.prepBody("wallets.create", Params{Passphrase: passphrase})
	if err != nil {
		return WalletsCreate{}, err
	}
	res, err := w.client.send(contenttype, r)
	if err != nil {
		return WalletsCreate{}, err
	}
	msg, err := w.client.handleResponse(*res)
	if err != nil {
		return WalletsCreate{}, err
	}
	buf := WalletsCreateResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return WalletsCreate{}, err
	}
	return WalletsCreate{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}

// Info gets a wallet, address is the address of the wallet to be retrieved
func (w *Wallets) Info(address string) (WalletsInfo, error) {
	r, err := w.client.prepBody("wallets.info", Params{Address: address})
	if err != nil {
		return WalletsInfo{}, err
	}
	res, err := w.client.send(contenttype, r)
	if err != nil {
		return WalletsInfo{}, err
	}
	msg, err := w.client.handleResponse(*res)
	if err != nil {
		return WalletsInfo{}, err
	}
	buf := WalletsInfoResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return WalletsInfo{}, err
	}
	return WalletsInfo{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}

// Transactions lists all transactions of a wallet. address is the address of the wallet. offset is the offset of
// transactions to be retrieved
func (w Wallets) Transactions(address string, offset int) (WalletsTransactions, error) {
	r, err := w.client.prepBody("wallets.transactions", Params{Address: address, Offset: offset})
	if err != nil {
		return WalletsTransactions{}, err
	}
	res, err := w.client.send(contenttype, r)
	if err != nil {
		return WalletsTransactions{}, err
	}
	msg, err := w.client.handleResponse(*res)
	if err != nil {
		return WalletsTransactions{}, err
	}
	buf := WalletsTransactionsResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return WalletsTransactions{}, err
	}
	return WalletsTransactions{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}

// Bip38Info gets a bip38 wallet, userID is the identifier of the wallet to be retrieved
func (w *Wallets) Bip38Info(userID string) (WalletsBIP38Info, error) {
	r, err := w.client.prepBody("wallets.bip38.info", Params{UserID: userID})
	if err != nil {
		return WalletsBIP38Info{}, err
	}
	res, err := w.client.send(contenttype, r)
	if err != nil {
		return WalletsBIP38Info{}, err
	}
	msg, err := w.client.handleResponse(*res)
	if err != nil {
		return WalletsBIP38Info{}, err
	}
	buf := WalletsBIP38InfoResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return WalletsBIP38Info{}, err
	}
	return WalletsBIP38Info{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}

// Bip38Create creates a bip38 wallet. bip38 is the bip38 of the wallet to be retrieved. userID is the identifier of
// the wallet to be retrieved
func (w Wallets) Bip38Create(bip38 string, userID string) (WalletsBIP38Create, error) {
	r, err := w.client.prepBody("wallets.bip38.info", Params{Bip38: bip38, UserID: userID})
	if err != nil {
		return WalletsBIP38Create{}, err
	}
	res, err := w.client.send(contenttype, r)
	if err != nil {
		return WalletsBIP38Create{}, err
	}
	msg, err := w.client.handleResponse(*res)
	if err != nil {
		return WalletsBIP38Create{}, err
	}
	buf := WalletsBIP38CreateResult{}
	err = json.Unmarshal(msg.Result, &buf)
	if err != nil {
		return WalletsBIP38Create{}, err
	}
	return WalletsBIP38Create{
		Jsonrpc: msg.JsonRPC,
		ID:      msg.ID,
		Result:  buf,
	}, nil
}
