package jsonrpc

import "time"

// structs contains all the response structs used for the json rpc api, the responses are divided up into two parts
// this is done because the api returns a message containing either an error or a response and by using multiple structs
// unmarshalling

// ErrorMessage is the message
type ErrorMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Error   ErrorP `json:"error"`
}

type ErrorP struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type BlocksInfo struct {
	Jsonrpc string           `json:"jsonrpc"`
	ID      string           `json:"id"`
	Result  BlocksInfoResult `json:"result"`
}

type BlocksInfoResult struct {
	ID       string `json:"id"`
	Version  int    `json:"version"`
	Height   int    `json:"height"`
	Previous string `json:"previous"`
	Forged   struct {
		Reward int `json:"reward"`
		Fee    int `json:"fee"`
		Total  int `json:"total"`
	} `json:"forged"`
	Payload struct {
		Hash   string `json:"hash"`
		Length int    `json:"length"`
	} `json:"payload"`
	Generator struct {
		Username  string `json:"username"`
		Address   string `json:"address"`
		PublicKey string `json:"publicKey"`
	} `json:"generator"`
	Signature    string `json:"signature"`
	Transactions int    `json:"Transactions"`
	Timestamp    struct {
		Epoch int       `json:"epoch"`
		Unix  int       `json:"unix"`
		Human time.Time `json:"human"`
	} `json:"timestamp"`
}

type BlocksLatest struct {
	Jsonrpc string             `json:"jsonrpc"`
	ID      string             `json:"id"`
	Result  BlocksLatestResult `json:"result"`
}

type BlocksLatestResult struct {
	ID       string `json:"id"`
	Version  int    `json:"version"`
	Height   int    `json:"height"`
	Previous string `json:"previous"`
	Forged   struct {
		Reward int `json:"reward"`
		Fee    int `json:"fee"`
		Total  int `json:"total"`
	} `json:"forged"`
	Payload struct {
		Hash   string `json:"hash"`
		Length int    `json:"length"`
	} `json:"payload"`
	Generator struct {
		Username  string `json:"username"`
		Address   string `json:"address"`
		PublicKey string `json:"publicKey"`
	} `json:"generator"`
	Signature    string `json:"signature"`
	Transactions int    `json:"Transactions"`
	Timestamp    struct {
		Epoch int       `json:"epoch"`
		Unix  int       `json:"unix"`
		Human time.Time `json:"human"`
	} `json:"timestamp"`
}

type BlocksTransactions struct {
	Jsonrpc string                   `json:"jsonrpc"`
	ID      string                   `json:"id"`
	Result  BlocksTransactionsResult `json:"result"`
}

type BlocksTransactionsResult struct {
	Count int `json:"count"`
	Data  []struct {
		ID            string `json:"id"`
		BlockID       string `json:"blockId"`
		Type          int    `json:"type"`
		Amount        int64  `json:"amount"`
		Fee           int    `json:"fee"`
		Sender        string `json:"sender"`
		Recipient     string `json:"recipient"`
		Signature     string `json:"signature"`
		Confirmations int    `json:"confirmations"`
		Timestamp     struct {
			Epoch int       `json:"epoch"`
			Unix  int       `json:"unix"`
			Human time.Time `json:"human"`
		} `json:"timestamp"`
	} `json:"data"`
}

type TransactionsInfo struct {
	Jsonrpc string                 `json:"jsonrpc"`
	ID      string                 `json:"id"`
	Result  TransactionsInfoResult `json:"result"`
}

type TransactionsInfoResult struct {
	ID            string `json:"id"`
	BlockID       string `json:"blockId"`
	Type          int    `json:"type"`
	Amount        int    `json:"amount"`
	Fee           int    `json:"fee"`
	Sender        string `json:"sender"`
	Recipient     string `json:"recipient"`
	Signature     string `json:"signature"`
	Confirmations int    `json:"confirmations"`
	Timestamp     struct {
		Epoch int       `json:"epoch"`
		Unix  int       `json:"unix"`
		Human time.Time `json:"human"`
	} `json:"timestamp"`
}

type TransactionsBroadcast struct {
	Jsonrpc string                      `json:"jsonrpc"`
	ID      string                      `json:"id"`
	Result  TransactionsBroadcastResult `json:"result"`
}

type TransactionsBroadcastResult struct {
	ID              string `json:"id"`
	Signature       string `json:"signature"`
	Timestamp       int    `json:"timestamp"`
	Type            int    `json:"type"`
	Fee             int    `json:"fee"`
	SenderPublicKey string `json:"senderPublicKey"`
	Amount          int    `json:"amount"`
	RecipientID     string `json:"recipientId"`
}

type TransactionsCreate struct {
	Jsonrpc string                   `json:"jsonrpc"`
	ID      string                   `json:"id"`
	Result  TransactionsCreateResult `json:"result"`
}

type TransactionsCreateResult struct {
	ID              string `json:"id"`
	Signature       string `json:"signature"`
	Timestamp       int    `json:"timestamp"`
	Type            int    `json:"type"`
	Fee             int    `json:"fee"`
	SenderPublicKey string `json:"senderPublicKey"`
	Amount          int    `json:"amount"`
	RecipientID     string `json:"recipientId"`
}

type TransactionsBIP38Create struct {
	Jsonrpc string                        `json:"jsonrpc"`
	ID      string                        `json:"id"`
	Result  TransactionsBIP38CreateResult `json:"result"`
}

type TransactionsBIP38CreateResult struct {
	ID              string `json:"id"`
	Signature       string `json:"signature"`
	Timestamp       int    `json:"timestamp"`
	Type            int    `json:"type"`
	Fee             int    `json:"fee"`
	SenderPublicKey string `json:"senderPublicKey"`
	Amount          int    `json:"amount"`
	RecipientID     string `json:"recipientId"`
}

type WalletsInfo struct {
	Jsonrpc string            `json:"jsonrpc"`
	ID      string            `json:"id"`
	Result  WalletsInfoResult `json:"result"`
}

type WalletsInfoResult struct {
	Address         string      `json:"address"`
	PublicKey       string      `json:"publicKey"`
	SecondPublicKey interface{} `json:"secondPublicKey"`
	Balance         int64       `json:"balance"`
	IsDelegate      bool        `json:"isDelegate"`
}

type WalletsTransactions struct {
	Jsonrpc string                    `json:"jsonrpc"`
	ID      string                    `json:"id"`
	Result  WalletsTransactionsResult `json:"result"`
}

type WalletsTransactionsResult struct {
	Count int `json:"count"`
	Data  []struct {
		ID            string `json:"id"`
		BlockID       string `json:"blockId"`
		Type          int    `json:"type"`
		Amount        int64  `json:"amount"`
		Fee           int    `json:"fee"`
		Sender        string `json:"sender"`
		Recipient     string `json:"recipient"`
		Signature     string `json:"signature"`
		Confirmations int    `json:"confirmations"`
		Timestamp     struct {
			Epoch int       `json:"epoch"`
			Unix  int       `json:"unix"`
			Human time.Time `json:"human"`
		} `json:"timestamp"`
	} `json:"data"`
}

type WalletsCreate struct {
	Jsonrpc string              `json:"jsonrpc"`
	ID      string              `json:"id"`
	Result  WalletsCreateResult `json:"result"`
}

type WalletsCreateResult struct {
	PublicKey string `json:"publicKey"`
	Address   string `json:"address"`
}

type WalletsBIP38Info struct {
	Jsonrpc string                 `json:"jsonrpc"`
	ID      string                 `json:"id"`
	Result  WalletsBIP38InfoResult `json:"result"`
}

type WalletsBIP38InfoResult struct {
	Wif string `json:"wif"`
}

type WalletsBIP38Create struct {
	Jsonrpc string                   `json:"jsonrpc"`
	ID      string                   `json:"id"`
	Result  WalletsBIP38CreateResult `json:"result"`
}

type WalletsBIP38CreateResult struct {
	PublicKey string `json:"publicKey"`
	Address   string `json:"address"`
	Wif       string `json:"wif"`
}
