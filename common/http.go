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
)

type HttpClientConfig struct {
	Ip      string
	Port    string
	Url     string
	Secret  string
	Timeout time.Duration
}

const appId = "uEVq0SDj34HwVltpNKzdgxmK"

// http请求接口
type HttpClientService interface {
	// 执行Get请求
	DoGet(model string) ([]byte, error)
	// 执行Post请求
	DoPost(model string, postBody interface{}) ([]byte, error)

	JSON(body interface{}) []byte

	Reader(postBody interface{}) ([]byte, error)

	GetMD5Encode(data string) string
}

// http请求默认实现(json传参)
type HttpClientServiceImpl struct {
	config *HttpClientConfig
	// client *http.Client
}

func NewHttpClientService(ip, port, url, secret string, timeout time.Duration) HttpClientService {

	hc := &HttpClientConfig{
		Url:     url,
		Ip:      ip,
		Port:    port,
		Secret:  secret,
		Timeout: timeout,
	}

	return &HttpClientServiceImpl{
		config: hc,
		// client:       &http.Client{Timeout: timeout},
	}
}

type CommonResult struct {
	Success string `json:"success"`
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

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", w.config.Url, model), nil)
	if err != nil {
		return nil, err
	}

	if w.config.Secret != "" {

		jsonText := model

		signText := appId + jsonText + w.config.Secret
		//
		signToLowerCase := w.GetMD5Encode(signText)

		req.Header.Add("ip", w.config.Ip)
		req.Header.Add("port", w.config.Port)
		req.Header.Add("sign", signToLowerCase)
	}

	resp, err := (&http.Client{Timeout: w.config.Timeout}).Do(req)
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

	fmt.Println("body:", string(body))

	return body, nil
}

func (w *HttpClientServiceImpl) DoPost(model string, postBody interface{}) ([]byte, error) {

	var bodyReader io.Reader
	var err error

	jsonText := ""

	if postBody != nil {
		bs, err := json.Marshal(&postBody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bs)
		jsonText = string(bs)
	}

	url := fmt.Sprintf("%s%s", w.config.Url, model)
	// fmt.Println("url:", url)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, err
	}

	if w.config.Secret != "" {

		signText := appId + jsonText + w.config.Secret
		signToLowerCase := w.GetMD5Encode(signText)

		req.Header.Add("ip", w.config.Ip)
		req.Header.Add("port", w.config.Port)
		req.Header.Add("sign", signToLowerCase)
	}

	resp, err := (&http.Client{Timeout: w.config.Timeout}).Do(req)
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

	fmt.Println("body:", string(body))

	return body, nil
}
