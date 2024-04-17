package loginmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/azure1489/wechat/util/httptools"
)

// GetProcessPorts 获取进程端口
func (l *LoginManagerServiceImpl) GetProcessPorts(url string, timeout time.Duration, process []string, publicKeyPath string) ([]string, error) {

	var processPortsReq httptools.ProcessPortsReq

	var err error

	processPortsReq.Process = process

	sendRequestBody, err := httptools.GetSendRequestBody(processPortsReq, publicKeyPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(sendRequestBody))
	if err != nil {
		return nil, err
	}

	resp, err := (&http.Client{Timeout: timeout}).Do(req)
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

	res := httptools.ProcessPortsRes{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(body))

	return res.Ports, nil
}
