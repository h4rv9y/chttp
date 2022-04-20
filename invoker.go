package chttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type result struct {
	status int
	err    error
}

func (r result) Status() int {
	return r.status
}

func (r result) Error() error {
	return r.err
}

func (r *request) setup() (req *http.Request, err error) {
	switch r.method {
	case "GET":
		url := fmt.Sprintf("%s?%s", r.url, r.params.Encode())
		req, err = http.NewRequest("GET", url, strings.NewReader(""))
		if err != nil {
			return
		}
		req.Header = r.meta
	case "POST", "PUT", "DELETE":
		contentType := r.meta["Content-Type"]
		if len(contentType) == 0 || strings.Compare(contentType[0], "application/json") == 0 {
			payload, err := json.Marshal(r.params)
			if err != nil {
				return nil, err
			}
			req, err = http.NewRequest(r.method, r.url, bytes.NewBuffer(payload))
			if err != nil {
				return nil, err
			}
			req.Header = r.meta
		} else if strings.Compare(contentType[0], "application/x-www-form-urlencoded") == 0 {
			req, err = http.NewRequest(r.method, r.url, strings.NewReader(r.params.Encode()))
			if err != nil {
				return nil, err
			}
			req.Header = r.meta
		}

	}
	return
}

func (r *request) invoker() (*http.Response, error) {
	req, err := r.setup()
	if err != nil {
		return nil, err
	}
	return r.c.httpClient.Do(req)
}

func (r *request) Send() result {
	res, err := r.invoker()
	if err != nil {
		return result{err: err}
	}
	defer res.Body.Close()
	return result{
		status: res.StatusCode,
	}
}

type stringResult struct {
	result
	str string
}

func (r stringResult) String() string {
	return r.str
}

func (r *request) ToString() stringResult {
	res, err := r.invoker()
	if err != nil {
		return stringResult{result: result{err: err}}
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return stringResult{result: result{err: err}}
	}
	return stringResult{
		result: result{
			status: res.StatusCode,
		},
		str: string(data),
	}
}

func (r *request) ToStruct(v interface{}) result {
	res, err := r.invoker()
	if err != nil {
		return result{err: err}
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result{err: err}
	}
	if err = json.Unmarshal(data, v); err != nil {
		return result{err: err}
	}
	return result{
		status: res.StatusCode,
	}
}
