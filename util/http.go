package util

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

type HttpClient struct {
	Ip      string
	Port    string
	Url     string
	Secret  string
	Timeout time.Duration
}

func NewClient(ip, port, url string, timeout time.Duration) (*HttpClient, error) {
	w := &HttpClient{
		Url:     url,
		Timeout: timeout,
		Ip:      ip,
		Port:    port,
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
func (c *HttpClient) GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func (w *HttpClient) DoPost(model string, postBody interface{}) ([]byte, error) {

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

	url := fmt.Sprintf("%s%s", w.Url, model)
	// fmt.Println("url:", url)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, err
	}

	appId := "uEVq0SDj34HwVltpNKzdgxmK"
	// appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"

	signText := appId + jsonText + w.Secret
	//
	signToLowerCase := w.GetMD5Encode(signText)

	req.Header.Add("ip", w.Ip)
	req.Header.Add("port", w.Port)
	req.Header.Add("sign", signToLowerCase)

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

	jsonText := model

	appId := "uEVq0SDj34HwVltpNKzdgxmK"
	// appSecret := "81fb6a51e232411c09575bb96bf71675980da0ac"

	signText := appId + jsonText + w.Secret
	//
	signToLowerCase := w.GetMD5Encode(signText)

	req.Header.Add("ip", w.Ip)
	req.Header.Add("port", w.Port)
	req.Header.Add("sign", signToLowerCase)

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
