package client

import (
	"goim/network"
	"goim/rpc"
	"golang.org/x/net/context"
	"errors"
	"sync"
)

var NotConnectedError = errors.New("Client is not connected")

type Client struct {
	sync.Mutex
	con *network.Connection
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Connect(ctx context.Context, address string, cert []byte, username string, password string) error {
	c.Lock()
	defer c.Unlock()

	if c.con != nil {
		c.con.Close()
	}

	connection, err := network.Connect(ctx, address, nil)
	if err != nil {
		return nil
	}

	c.con = connection

	login := &rpc.LoginRequest{Username: &username, Password: &password}
	if err := connection.Write(login.Serialize()); err != nil {
		return err
	}

	buf, err := connection.Read()
	if err != nil {
		return err
	}

	loginResponse := rpc.LoginResponse{}
	if err := loginResponse.Deserialize(buf); err != nil {
		return err
	}

	if loginResponse.Status != rpc.SuccessStatus {
		return errors.New("got unexpected status from server")
	}

	return nil
}

func (c *Client) Disconnect() {
	c.Lock()
	if c.con != nil {
		c.con.Close()
	}
	c.con = nil
	c.Unlock()
}

func (c *Client) GetKeys() (res []*string, err error, status int32) {
	request := &rpc.GetKeysRequest{}
	response := &rpc.GetKeysResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Arr, nil, response.Status
}

func (c *Client) GetStr(key *string) (res *string, expires int64, err error, status int32) {
	request := &rpc.GetStrRequest{Key: key}
	response := &rpc.GetStrResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Str, response.Expires, nil, response.Status
}

func (c *Client) SetStr(key *string, str *string, ttl int64) (expires int64, err error, status int32) {
	request := &rpc.SetStrRequest{Key: key, Str: str, TTL: ttl}
	response := &rpc.SetStrResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Expires, nil, response.Status
}

func (c *Client) GetArr(key *string) (res []*string, expires int64, err error, status int32) {
	request := &rpc.GetArrRequest{Key: key}
	response := &rpc.GetArrResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Arr, response.Expires, nil, response.Status
}

func (c *Client) SetArr(key *string, arr []*string, ttl int64) (expires int64, err error, status int32) {
	request := &rpc.SetArrRequest{Key: key, Arr: arr, TTL: ttl}
	response := &rpc.SetArrResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Expires, nil, response.Status
}

func (c *Client) GetArrItem(key *string, index int32) (res *string, expires int64, err error, status int32) {
	request := &rpc.GetArrItemRequest{Key: key, Index: index}
	response := &rpc.GetArrItemResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Str, response.Expires, nil, response.Status
}

func (c *Client) SetArrItem(key *string, index int32, str *string, ttl int64) (expires int64, err error, status int32) {
	request := &rpc.SetArrItemRequest{Key: key, Index: index, Str: str, TTL: ttl}
	response := &rpc.SetArrItemResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Expires, nil, response.Status
}

func (c *Client) GetDict(key *string) (res map[string]*string, expires int64, err error, status int32) {
	request := &rpc.GetDictRequest{Key: key}
	response := &rpc.GetDictResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Dict, response.Expires, nil, response.Status
}

func (c *Client) SetDict(key *string, dict map[string]*string, ttl int64) (expires int64, err error, status int32) {
	request := &rpc.SetDictRequest{Key: key, Dict: dict, TTL: ttl}
	response := &rpc.SetDictResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Expires, nil, response.Status
}

func (c *Client) GetDictItem(key *string, subkey *string) (res *string, expires int64, err error, status int32) {
	request := &rpc.GetDictItemRequest{Key: key, Subkey: subkey}
	response := &rpc.GetDictItemResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Str, response.Expires, nil, response.Status
}

func (c *Client) SetDictItem(key *string, subkey *string, str *string, ttl int64) (expires int64, err error, status int32) {
	request := &rpc.SetDictItemRequest{Key: key, Subkey: subkey, Str: str, TTL: ttl}
	response := &rpc.SetDictItemResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return response.Expires, nil, response.Status
}

func (c *Client) RemoveEntry(key *string) (err error, status int32) {
	request := &rpc.RemoveEntryRequest{Key: key}
	response := &rpc.RemoveEntryResponse{}

	if err = c.doRequest(request, response); err != nil {
		return
	}

	return nil, response.Status
}

func (c *Client) doRequest(request rpc.Binarizer, response rpc.Binarizer) error {
	// as long as this implementation does not support request multiplexing only one request per time can be done
	c.Lock()
	defer c.Unlock()

	con := c.con
	if con == nil {
		return NotConnectedError
	}

	con.Write(request.Serialize())

	buf, err := con.Read()
	if err != nil {
		return err
	}

	if err := response.Deserialize(buf); err != nil {
		return err
	}

	return nil
}
