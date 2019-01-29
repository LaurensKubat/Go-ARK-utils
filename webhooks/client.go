package webhooks

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Client is the client for the ARK webhooks api as described in https://docs.ark.io/api/webhooks
type Client struct {
	client *http.Client
	url    url.URL
}

func (c *Client) get(URL string) ([]byte, error) {
	res, err := c.client.Get(URL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func (c *Client) post(contentType string, body io.Reader) ([]byte, error) {
	res, err := c.client.Post(c.url.String(), contentType, body)
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	return resBody, err
}

func (c *Client) put(URL string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, URL, body)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	return res, err
}

func (c *Client) delete(URL string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, URL, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	return res, err
}

func (c *Client) List(page int, limit int) (List, error) {
	URL := c.url
	if page != 0 || limit != 0 {
		if page != 0 {
			URL.RawQuery = url.QueryEscape(strconv.Itoa(page))
		}
		if limit != 0 {
			URL.RawQuery = url.QueryEscape(strconv.Itoa(limit))
		}
	}
	res, err := c.get(c.url.String())
	if err != nil {
		return List{}, err
	}
	list := List{}
	err = json.Unmarshal(res, &list)
	return list, err
}

func (c *Client) Retrieve(id string) (Retrieve, error) {
	res, err := c.get(c.url.String() + "/" + id)
	if err != nil {
		return Retrieve{}, nil
	}
	ret := Retrieve{}
	err = json.Unmarshal(res, &ret)
	return ret, err
}

func (c *Client) Create(event string, target string, enabled string, conditions []Condition) (Create, error) {
	req := CreateBody{Event: event, Target: target, Enabled: enabled, Conditions: conditions}
	body, err := json.Marshal(req)
	if err != nil {
		return Create{}, nil
	}
	readbody := bytes.NewReader(body)
	res, err := c.post("application/json", readbody)
	cr := Create{}
	err = json.Unmarshal(res, &cr)
	return cr, err
}

func (c *Client) Update(id string, event string, target string, enabled string, conditions []Condition) error {
	req := CreateBody{Event: event, Target: target, Enabled: enabled, Conditions: conditions}
	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	readbody := bytes.NewReader(body)
	_, err = c.put(c.url.String()+"/"+id, readbody)
	return err
}

func (c *Client) Delete(id string) error {
	_, err := c.delete(c.url.String() + "/" + id)
	return err
}
