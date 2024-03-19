package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/azure1489/wechat/util/httptools"
)

type HttpClientConfig struct {
	Ip            string
	Port          string
	Url           string
	PublicKeyPath string
	Timeout       time.Duration
}

// http请求接口
type HttpClientService interface {
	// 执行Get请求
	DoGet(model string) ([]byte, error)
	// 执行Post请求
	DoPost(model string, postBody interface{}) ([]byte, error)

	JSON(body interface{}) []byte

	Reader(postBody interface{}) ([]byte, error)

	// GetMD5Encode(data string) string
}

// http请求默认实现(json传参)
type HttpClientServiceImpl struct {
	config *HttpClientConfig
	// client *http.Client
}

func NewHttpClientService(ip, port, url, publicKeyPath string, timeout time.Duration) HttpClientService {

	hc := &HttpClientConfig{
		Url:           url,
		Ip:            ip,
		Port:          port,
		PublicKeyPath: publicKeyPath,
		Timeout:       timeout,
	}

	return &HttpClientServiceImpl{
		config: hc,
		// client:       &http.Client{Timeout: timeout},
	}
}

type CommonResult struct {
	Code string `json:"code"`
}

func (w *HttpClientServiceImpl) JSON(body interface{}) []byte {
	bs, _ := json.Marshal(&body)
	return bs
}

func (w *HttpClientServiceImpl) Reader(postBody interface{}) ([]byte, error) {
	bs, err := json.Marshal(&postBody)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// GetMD5Encode 返回一个32位md5加密后的字符串
func (c *HttpClientServiceImpl) GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func (w *HttpClientServiceImpl) DoGet(model string) ([]byte, error) {

	var proxyBody httptools.ProxyBody

	var err error

	proxyBody.Domain = "http://" + w.config.Ip + ":" + w.config.Port
	proxyBody.UrlPath = model
	proxyBody.Method = "GET"

	sendRequestBody, err := httptools.GetSendRequestBody(proxyBody, w.config.PublicKeyPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, w.config.Url, bytes.NewBuffer(sendRequestBody))
	if err != nil {
		return nil, err
	}

	resp, err := (&http.Client{Timeout: w.config.Timeout}).Do(req)
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

func (w *HttpClientServiceImpl) DoPost(model string, postBody interface{}) ([]byte, error) {

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

	proxyBody.Domain = "http://" + w.config.Ip + ":" + w.config.Port
	proxyBody.UrlPath = model
	proxyBody.Body = bodyStr
	proxyBody.Method = "POST"
	proxyBody.Headers = map[string][]string{
		"Content-Type": {"application/json"},
		// "Authorization": {"Bearer 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"},
	}

	sendRequestBody, err := httptools.GetSendRequestBody(proxyBody, w.config.PublicKeyPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, w.config.Url, bytes.NewBuffer(sendRequestBody))
	if err != nil {
		return nil, err
	}

	resp, err := (&http.Client{Timeout: w.config.Timeout}).Do(req)
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
