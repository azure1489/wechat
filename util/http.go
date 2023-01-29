package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	Url     string
	Timeout time.Duration
}

func NewClient(url string, timeout time.Duration) (*HttpClient, error) {
	w := &HttpClient{
		Url:     url,
		Timeout: timeout,
	}
	return w, nil
}

func (w *HttpClient) JSON(body interface{}) []byte {
	bs, _ := json.Marshal(&body)
	return bs
}

func (w *HttpClient) Reader(postBody interface{}) (io.Reader, error) {
	bs, err := json.Marshal(&postBody)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(bs), nil
}

type CommonResult struct {
	Success string `json:"success"`
}

func (w *HttpClient) DoPost(model string, postBody interface{}) ([]byte, error) {

	var bodyReader io.Reader
	var err error

	if postBody != nil {
		bodyReader, err = w.Reader(postBody)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", w.Url, model), bodyReader)
	if err != nil {
		return nil, err
	}
	resp, err := (&http.Client{Timeout: w.Timeout}).Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("response status code=%d, body=%s", resp.StatusCode, string(body))
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (w *HttpClient) DoGet(model string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", w.Url, model), nil)
	if err != nil {
		return nil, err
	}
	resp, err := (&http.Client{Timeout: w.Timeout}).Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("response status code=%d, body=%s", resp.StatusCode, string(body))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
