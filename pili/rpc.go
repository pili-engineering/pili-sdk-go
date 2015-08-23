package pili

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type RPC_Client struct {
	*http.Client
}

func (r RPC_Client) Do(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("User-Agent", UserAgent())
	resp, err = r.Client.Do(req)
	if err != nil {
		return
	}
	return
}

func (r RPC_Client) RequestWith(
	method string,
	url string,
	bodyType string,
	body io.Reader,
	bodyLength int) (resp *http.Response, err error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", bodyType)
	req.ContentLength = int64(bodyLength)
	return r.Do(req)
}

func (r RPC_Client) Post(url string, data interface{}) (resp *http.Response, err error) {
	msg, err := json.Marshal(data)
	if err != nil {
		return
	}
	return r.RequestWith("POST", url, "application/json", bytes.NewReader(msg), len(msg))
}

func (r RPC_Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	return r.Do(req)
}

func (r RPC_Client) Del(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}
	return r.Do(req)
}

func (r RPC_Client) PostCall(ret interface{}, url string, params interface{}) (err error) {
	resp, err := r.Post(url, params)
	if err != nil {
		return err
	}
	return callRet(ret, resp)
}

func (r RPC_Client) GetCall(ret interface{}, url string) (err error) {
	resp, err := r.Get(url)
	if err != nil {
		return err
	}
	return callRet(ret, resp)
}

func (r RPC_Client) DelCall(ret interface{}, url string) (err error) {
	resp, err := r.Del(url)
	if err != nil {
		return err
	}
	return callRet(ret, resp)
}
