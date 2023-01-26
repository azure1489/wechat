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

type AddNewFriendReq struct {
	Wxid string `json:"wxid"`
	Desc string `json:"desc"`
	Type string `json:"type"`
	Role string `json:"role"`
}

type AddNewFriendResult struct {
	AddNewFriend string `json:"AddNewFriend"`
}

type DeleteFriendReq struct {
	Wxid string `json:"wxid"`
}

type DeleteFriendResult struct {
	DeleteFriend string `json:"DeleteFriend"`
}

type BlackFriendReq struct {
	Wxid string `json:"wxid"`
}

type BlackFriendResult struct {
	BlackFriend string `json:"BlackFriend"`
}

type VerifyFriendReq struct {
	V3   string `json:"v3"`
	V4   string `json:"v4"`
	Role string `json:"role"`
	From string `json:"from"`
}

type VerifyFriendResult struct {
	VerifyFriend string `json:"VerifyFriend"`
}

type VoiceCallFriendReq struct {
	Wxid string `json:"wxid"`
}

type VoiceCallFriendResult struct {
	VoiceCallFriend string `json:"VoiceCallFriend"`
}

type VideoCallFriendReq struct {
	Wxid string `json:"wxid"`
}

type VideoCallFriendResult struct {
	VideoCallFriend string `json:"VideoCallFriend"`
}

type SetStarTagReq struct {
	Wxid string `json:"wxid"`
}

type SetStarTagResult struct {
	SetStarTag string `json:"SetStarTag"`
}

type UndoStarTagReq struct {
	Wxid string `json:"wxid"`
}

type UndoStarTagResult struct {
	UndoStarTag string `json:"UndoStarTag"`
}

type SetContactRoleReq struct {
	Wxid string `json:"wxid"`
	Role string `json:"role"` // 权限值,1~6 1=仅聊天 2=聊天、朋友圈、微信运动等
}

type SetContactRoleResult struct {
	SetContactRole string `json:"SetContactRole"`
}
