package pili

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

var VERSION = "1.0.1"

func formatUserAgent() string {
	return fmt.Sprintf("pili-sdk-go/%s %s %s/%s", VERSION, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

// -----------------------------------------------------------------------------------------------------------

type RPC_Client struct {
	*http.Client
}

var DefaultClient = RPC_Client{http.DefaultClient}

func (r RPC_Client) Do(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("User-Agent", formatUserAgent())
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

// -----------------------------------------------------------------------------------------------------------

type Error struct {
	Err int    `json:"error"`
	Msg string `json:"message"`
}

type ErrorRet struct {
	Err     int              `json:"error"`
	Msg     string           `json:"message"`
	Details map[string]Error `json:"details"`
}

type ErrorInfo struct {
	ErrorRet
	Code int
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
