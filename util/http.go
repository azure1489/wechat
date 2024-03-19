package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/azure1489/wechat/util/httptools"
)

// 默认超时时间
const DefaultTimeout = 60 * time.Second

type HttpClient struct {
	Ip            string
	Port          string
	Url           string
	PublicKeyPath string
	Timeout       time.Duration
}

func NewClient(ip, port, url, publicKeyPath string, timeout time.Duration) (*HttpClient, error) {

	if timeout == 0 {
		timeout = DefaultTimeout
	}

	w := &HttpClient{
		Url:     url,
		Timeout: timeout,
		Ip:      ip,
		Port:    port,
		// Secret:  secret,
	}
	return w, nil
}

func (w *HttpClient) JSON(body interface{}) []byte {
	bs, _ := json.Marshal(&body)
	return bs
}

func (w *HttpClient) Reader(postBody interface{}) ([]byte, error) {
	bs, err := json.Marshal(&postBody)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

type CommonResult struct {
	Success string `json:"success"`
}

// GetMD5Encode 返回一个32位md5加密后的字符串
// func (c *HttpClient) GetMD5Encode(data string) string {
// 	h := md5.New()
// 	h.Write([]byte(data))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func (w *HttpClient) DoGet(model string) ([]byte, error) {

	var proxyBody httptools.ProxyBody

	var err error

	proxyBody.Domain = "http://" + w.Ip + ":" + w.Port
	proxyBody.UrlPath = model
	proxyBody.Method = "GET"

	sendRequestBody, err := httptools.GetSendRequestBody(proxyBody, w.PublicKeyPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, w.Url, bytes.NewBuffer(sendRequestBody))
	if err != nil {
		return nil, err
	}

	resp, err := (&http.Client{Timeout: w.Timeout}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("response status code=%d, body=%s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (w *HttpClient) DoPost(model string, postBody interface{}) ([]byte, error) {

	bodyStr := ""

	if postBody != nil {
		bs, err := json.Marshal(&postBody)
		if err != nil {
			return nil, err
		}
		bodyStr = string(bs)
		// jsonText = string(bs)
	}

	var proxyBody httptools.ProxyBody

	var err error

	proxyBody.Domain = "http://" + w.Ip + ":" + w.Port
	proxyBody.UrlPath = model
	proxyBody.Body = bodyStr
	proxyBody.Method = "POST"
	proxyBody.Headers = map[string][]string{
		"Content-Type": {"application/json"},
		// "Authorization": {"Bearer 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"},
	}

	sendRequestBody, err := httptools.GetSendRequestBody(proxyBody, w.PublicKeyPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, w.Url, bytes.NewBuffer(sendRequestBody))
	if err != nil {
		return nil, err
	}

	resp, err := (&http.Client{Timeout: w.Timeout}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("response status code=%d, body=%s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println("body:", string(body))

	return body, nil
}
