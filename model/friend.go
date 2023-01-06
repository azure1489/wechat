package model

type QueryBodyInfoReq struct {
	SerachWhat string `json:"serachwhat"`
}

type QueryBodyInfoResult struct {
	Retstatus   string `json:"retstatus"` // 查询状态 Everything is OK=查询成功 User do not exist= 不存在该用户
	Tel         string `json:"tel"`       //手机号码
	Nickname    string `json:"nickname"`
	Wxid        string `json:"wxid"`
	Sex         string `json:"sex"`
	Country     string `json:"country"`
	Province    string `json:"province"`
	Region      string `json:"region"`
	Signature   string `json:"signature"`
	Bigavatar   string `json:"bigavatar"`
	Smallavatar string `json:"smallavatar"`
	V3          string `json:"v3"`
	V4          string `json:"v4"`
}

type EditFriendMarkReq struct {
	Wxid string `json:"wxid"`
	Mark string `json:"mark"`
}

type EditFriendMarkResult struct {
	EditFriendMark string `json:"EditFriendMark"`
}
