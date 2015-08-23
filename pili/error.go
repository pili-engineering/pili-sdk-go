package pili

import (
	"encoding/json"
	"net/http"
)

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
