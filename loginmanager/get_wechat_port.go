package loginmanager

import (
	"encoding/json"
	"net/url"
	"strconv"
)

//	{
//	    "total_num": "2",
//	    "List": [
//	        {
//	            "Index": 1,
//	            "ProcessName": "WeChat.exe",
//	            "PID": 40588,
//	            "Par": "StartPort=30001&Proxy_IP=183.162.226.236&Proxy_Port=31722"
//	        },
//	        {
//	            "Index": 2,
//	            "ProcessName": "WeChat.exe",
//	            "PID": 820,
//	            "Par": "StartPort=30002"
//	        }
//	    ]

// GetWeChatPort 获取微信端口号
func (l *LoginManagerServiceImpl) GetWeChatPort() ([]string, error) {

	resultBody, err := l.http.DoPost("/Get_WeChatProcessNumber", nil)
	if err != nil {
		return nil, err
	}

	// log.Println("resultBody:", string(resultBody))

	commonResult := GetWeChatProcessNumberResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	if commonResult.TotalNum == "0" || len(commonResult.List) == 0 {
		return []string{}, nil
	}

	urlStr := l.config.Url

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	newUrl := u.Scheme + "://" + u.Hostname() + ":" + u.Port() + "/process-ports"
	// log.Println("newUrl:", newUrl)

	pids := make([]string, 0)
	for i := range commonResult.List {
		pids = append(pids, strconv.Itoa(commonResult.List[i].PID))
	}

	ports, err := l.GetProcessPorts(newUrl, l.config.Timeout, pids, l.config.PublicKeyPath)
	if err != nil {
		return nil, err
	}

	return ports, nil
}
