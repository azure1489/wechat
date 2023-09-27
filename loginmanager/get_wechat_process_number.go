package loginmanager

import "encoding/json"

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
//	}
type GetWeChatProcessNumberResult struct {
	TotalNum string `json:"total_num"`
	List     []struct {
		Index        int    `json:"Index"`
		ProcessName  string `json:"ProcessName"`
		PID          int    `json:"PID"`
		Par          string `json:"Par"`
		IsLogin      string `json:"IsLogin"`
		IsWeChatLive string `json:"IsWeChatLive"`
	} `json:"List"`
}

// GetWeChatProcessNumber 获取微信进程总数 https://www.showdoc.com.cn/WeChatProject/9794632555004152
func (l *LoginManagerServiceImpl) GetWeChatProcessNumber() (*GetWeChatProcessNumberResult, error) {

	resultBody, err := l.http.DoPost("/Get_WeChatProcessNumber", nil)
	if err != nil {
		return nil, err
	}

	commonResult := GetWeChatProcessNumberResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
