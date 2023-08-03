package http

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var ErrTimeout = errors.New("Timeout")

var defaultClient = &http.Client{
	Timeout: time.Second * 40,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

// Request model of http client request
type Request struct {
	Method         string
	URL            string
	Headers        map[string]string
	PostData       []byte
	RespData       []byte
	RespStatusCode int
	Username       string
	Password       string
}

// Send - sends a request;
func (r *Request) Send(ctx context.Context, client *http.Client) (retry bool, err error) {

	if client == nil {
		client = defaultClient
	}

	var bodyReader io.Reader
	if len(r.PostData) > 0 {
		bodyReader = bytes.NewReader(r.PostData)
	}

	// prepare request;
	httpReqs, err := http.NewRequestWithContext(ctx, r.Method, r.URL, bodyReader)
	if err != nil {
		err = errors.New("http.NewRequest error: " + err.Error())
		return
	}
	if len(r.Username) != 0 && len(r.Password) != 0 {
		httpReqs.SetBasicAuth(r.Username, r.Password)
	}
	httpReqs.Header.Set("Content-Type", "application/json; charset=utf-8")
	httpReqs.Header.Set("Accept", "application/json")
	if len(r.Headers) != 0 {
		for key, value := range r.Headers {
			httpReqs.Header.Set(key, value)
		}
	}

	// send request;
	httpResp, err := client.Do(httpReqs)
	if err != nil {
		retry = true
		if err2, ok := err.(net.Error); ok && err2.Timeout() {
			err = ErrTimeout
		} else {
			err = errors.New("client.Do error: " + err.Error())
		}
		return
	}
	defer httpResp.Body.Close()
	r.RespStatusCode = httpResp.StatusCode

	if httpResp.Body == nil {
		err = errors.New("empty_http_response")
		return
	}

	// read response;
	r.RespData, err = ioutil.ReadAll(httpResp.Body)
	if err != nil {
		retry = true
		err = errors.New("ioutil.ReadAll error: " + err.Error())
		return
	}

	return
}

func (r *Request) JSONMarshal(v interface{}) (err error) {
	r.PostData, err = json.Marshal(v)
	return
}

func (r *Request) JSONUnmarshal(v interface{}) (err error) {
	err = json.Unmarshal(r.RespData, v)
	return
}
