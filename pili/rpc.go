package pili

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func NewRPC(creds *Credentials) *RPC {
	t := NewTransport(creds, nil)
	tc := http.Client{Transport: t}
	return &RPC{&tc}
}

type RPC struct {
	*http.Client
}

func (r RPC) Do(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("User-Agent", UserAgent())
	resp, err = r.Client.Do(req)
	if err != nil {
		return
	}
	return
}

func (r RPC) RequestWith(
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

func (r RPC) Post(url string, data interface{}) (resp *http.Response, err error) {
	msg, err := json.Marshal(data)
	if err != nil {
		return
	}
	return r.RequestWith("POST", url, "application/json", bytes.NewReader(msg), len(msg))
}

func (r RPC) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	return r.Do(req)
}

func (r RPC) Del(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}
	return r.Do(req)
}

func (r RPC) PostCall(ret interface{}, url string, params interface{}) (err error) {
	resp, err := r.Post(url, params)
	if err != nil {
		return err
	}
	return callRet(ret, resp)
}

func (r RPC) GetCall(ret interface{}, url string) (err error) {
	resp, err := r.Get(url)
	if err != nil {
		return err
	}
	return callRet(ret, resp)
}

func (r RPC) DelCall(ret interface{}, url string) (err error) {
	resp, err := r.Del(url)
	if err != nil {
		return err
	}
	return callRet(ret, resp)
}

type ErrorInfo struct {
	Message string           `json:"message"`
	ErrCode int              `json:"error"`
	Details map[string]error `json:"details,omitempty"`
	Code    int              `json:"code"`
}

func (r *ErrorInfo) Error() string {
	msg, _ := json.Marshal(r)
	return string(msg)
}

func ResponseError(resp *http.Response) (err error) {

	e := &ErrorInfo{
		Code: resp.StatusCode,
	}

	if resp.StatusCode > 299 {
		if resp.ContentLength != 0 {
			if ct, ok := resp.Header["Content-Type"]; ok && ct[0] == "application/json" {
				json.NewDecoder(resp.Body).Decode(&e)
			}
		}
	}

	return e
}

func callRet(ret interface{}, resp *http.Response) (err error) {

	defer resp.Body.Close()

	if resp.StatusCode/100 == 2 {
		if ret != nil && resp.ContentLength != 0 {
			err = json.NewDecoder(resp.Body).Decode(ret)
			if err != nil {
				return
			}
		}
		return
	}
	return ResponseError(resp)
}
