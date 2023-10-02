package friendmanager

import "encoding/json"

type QueryBodyInfoReq struct {
	SerachWhat string `json:"serachwhat"`
}

type QueryBodyInfoResult struct {
	RetStatus   string `json:"retstatus"` // 查询状态 Everything is OK=查询成功 User do not exist= 不存在该用户
	Tel         string `json:"tel"`
	NickName    string `json:"nickname"`
	WxId        string `json:"wxid"`
	Sex         string `json:"sex"`
	Country     string `json:"country"`
	Province    string `json:"province"`
	Region      string `json:"region"`
	Signature   string `json:"signature"`
	BigAvatar   string `json:"bigavatar"`
	SmallAvatar string `json:"smallavatar"`
	V3          string `json:"v3"`
	V4          string `json:"v4"`
}

// QueryBodyInfo 查询好友信息 https://www.showdoc.com.cn/WeChatProject/8982367120882654
func (l *FriendManagerServiceImpl) QueryBodyInfo(serachWhat string) (*QueryBodyInfoResult, error) {

	req := QueryBodyInfoReq{
		SerachWhat: serachWhat,
	}

	resultBody, err := l.http.DoPost("/QueryBodyInfo", req)
	if err != nil {
		return nil, err
	}

	result := QueryBodyInfoResult{}
	err = json.Unmarshal(resultBody, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
