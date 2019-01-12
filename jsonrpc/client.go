package jsonrpc

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	jsonrpcversion = "2.0"
	contenttype    = "application/json"
)

type Service struct {
	client *Client
}

type Client struct {
	client *http.Client
	url    *url.URL

	Transactions *Transactions
	Blocks       *Blocks
	Wallets      *Wallets
}

type request struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      string `json:"id"`
	Params  Params `json:"params, omitempty"`
}

type Params struct {
	ID          string   `json:"id, omitempty"`
	Offset      int      `json:"offset, omitempty"`
	Txs         []string `json:"transactions, omitempty"`
	RecipientID string   `json:"recipientId, omitempty"`
	Amount      string   `json:"amount, omitempty"`
	Passphrase  string   `json:"passphrase, omitempty"`
	Bip38       string   `json:"bip38, omitempty"`
	UserID      string   `json:"userId, omitempty"`
	Address     string   `json:"address, omitempty"`
}

func NewClient(rawurl string) (*Client, error) {
	URL, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	c := Client{
		client: http.DefaultClient,
		url:    URL,
	}
	c.Wallets = &Wallets{client: &c}
	c.Blocks = &Blocks{client: &c}
	c.Transactions = &Transactions{client: &c}
	return &c, nil
}

func (c Client) send(contentType string, msg io.Reader) (*http.Response, error) {
	res, err := c.client.Post(c.url.String(), contentType, msg)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// prepBody preps the body of the post request. the nested json params contains different tags depending on the parameter
// in the message. Thus that part of the body is made by marshalling a map instead of marshalling a struct.
func (c Client) prepBody(method string, params Params) (io.Reader, error) {
	body := request{
		JsonRPC: jsonrpcversion,
		Method:  method,
		ID:      "",
		Params:  params,
	}
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(buf)
	return r, nil
}

type Response struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      string          `json:"id"`
	Result  json.RawMessage `json:"result, omitempty"`
	Error   json.RawMessage `json:"error, omitempty"`
}

// handleResponse handles all responses by partially unmarshalling the http response and checking if an error is present
// in the response struct. If not, the Response struct is returned, if an error is present in the response struct.
// the struct itself is returned for more potential error handling plus an error
func (c Client) handleResponse(response http.Response) (Response, error) {
	// Unmarshall the response
	res := Response{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return Response{}, err
	}
	// check if the error field in the response is nil, else unmarshall that field further and return the contents as
	// an error
	if res.Error == nil {
		return res, nil
	}
	errp := ErrorP{}
	err = json.Unmarshal(res.Error, errp)
	if err != nil {
		return Response{}, err
	}
	errmsg := ErrorMessage{
		Jsonrpc: res.JsonRPC,
		ID:      res.ID,
		Error:   errp,
	}
	return res, errors.New("id" + errmsg.ID + "errored with: " + errmsg.Error.Message + " (" +
		errmsg.Error.Code + ")")
}
