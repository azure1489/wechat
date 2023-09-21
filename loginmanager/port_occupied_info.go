package loginmanager

import (
	"encoding/json"
)

type PortOccupiedInfoReq struct {
	CheckPort string `json:"CheckPort"`
}

type PortOccupiedInfoResult struct {
	Occupied string `json:"Occupied"`
}

// Get_PortOccupiedInfo 获取进程端口占用信息 https://www.showdoc.com.cn/WeChatProject/10128598529687136
func (l *LoginManagerServiceImpl) GetPortOccupiedInfo(req *PortOccupiedInfoReq) (string, error) {

	resultBody, err := l.http.DoPost("/Get_PortOccupiedInfo", req)
	if err != nil {
		return "", err
	}

	commonResult := PortOccupiedInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.Occupied, nil
}
