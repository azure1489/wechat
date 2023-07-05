package model

type FriendAndChatRoomResult struct {
	Friends  []FriendItem   `json:"friend"`
	Chatroom []ChatroomItem `json:"chatroom"`
	Gh       []GhItem       `json:"gh"`
}

type FriendItem struct {
	Index     string `json:"index"`
	Wxid      string `json:"wxid"`     // 好友微信ID
	Account   string `json:"account"`  // 好友微信号
	Markname  string `json:"markname"` // 好友备注名称
	Nickname  string `json:"nickname"` // 好友昵称
	Headimg   string `json:"headimg"`  // 好友头像URL地址
	V3        string `json:"v3"`
	Sex       string `json:"sex"`       // 性别 0=未设置 1=男 2=女
	Starrole  string `json:"starrole"`  // 星标 65/67=星标 1/3=未星标
	Dontseeit string `json:"dontseeit"` // 1=开启了不让他看
	Dontseeme string `json:"dontseeme"` // 1=开启了不看他 128/129=仅聊天
	Lag       string `json:"lag"`       // 所属标签清单，多开会用逗号隔开
}

type ChatroomItem struct {
	Index    string `json:"index"`
	Gid      string `json:"gid"`      // 群ID
	Markname string `json:"markname"` // 群备注名称
	Gname    string `json:"gname"`    // 群名称
	V3       string `json:"v3"`       // V3数据
}

type GhItem struct {
	Index    string `json:"index"`
	Wxid     string `json:"wxid"`     // 公众号微信ID
	Account  string `json:"account"`  // 微信号
	Markname string `json:"markname"` // 备注名称
	Nickname string `json:"nickname"` // 昵称
	Headimg  string `json:"headimg"`  // 头像URL地址
	V3       string `json:"v3"`       // V3数据

}

type StickyChatReq struct {
	Gidorwxid string `json:"gidorwxid"`
}

type StickyChatResult struct {
	StickyChat string `json:"StickyChat"`
}

type UnpinChatResult struct {
	UnpinChat string `json:"UnpinChat"`
}

type BatchGetContactBriefInfoReq struct {
	WxidList string `json:"wxidlist"`
}

type BatchGetContactBriefInfoResult struct {
	Count string             `json:"count"`
	Info  []ContactBriefInfo `json:"info"`
}
type ContactBriefInfo struct {
	Wxid      string `json:"wxid"`
	Nickname  string `json:"nickname"`
	Markname  string `json:"markname"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Bighead   string `json:"bighead"`
	Smallhead string `json:"smallhead"`
}

type GetFriendOrChatroomDetailInfoReq struct {
	WxidOrGid string `json:"wxidorgid"`
}

type GetFriendOrChatroomDetailInfoResult struct {
	Type string `json:"type"`
	V3   string `json:"v3"`
	GetFriendOrChatroomDetailInfoResultFriendItem
	GetFriendOrChatroomDetailInfoResultChatroomItem
}

type GetFriendOrChatroomDetailInfoResultFriendItem struct {
	Wxid          string `json:"wxid,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
	Province      string `json:"province,omitempty"`
	Area          string `json:"area,omitempty"`
	Signinfo      string `json:"signinfo,omitempty"`
	Wxaccount     string `json:"wxaccount,omitempty"`
	Timelinebgurl string `json:"timelinebgurl,omitempty"`
	Country       string `json:"country,omitempty"`
	Headurl       string `json:"headurl,omitempty"`
	Fromchatroom  string `json:"fromchatroom,omitempty"`
	V4            string `json:"v4,omitempty"`
}

type GetFriendOrChatroomDetailInfoResultChatroomItem struct {
	Gid             string                                                  `json:"gid,omitempty"`
	ChatroomMaster  string                                                  `json:"chatroom_master,omitempty"`
	ChatroomHeadURL string                                                  `json:"chatroom_head_url,omitempty"`
	MemberCount     string                                                  `json:"member_count,omitempty"`
	Member          []GetFriendOrChatroomDetailInfoResultChatroomItemMember `json:"member,omitempty"`
}
type GetFriendOrChatroomDetailInfoResultChatroomItemMember struct {
	Wxid     string `json:"wxid,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Fromsouc string `json:"fromsouc,omitempty"`
}

type MarkAsReadSessionReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type MarkAsReadSessionResult struct {
	MarkAsReadSession string `json:"MarkAsReadSession"`
}

type InitContactResult struct {
	InitContact string                       `json:"InitContact"`
	Batch       []InitContactResultBatchItem `json:"batch"`
}

type InitContactResultBatchItem struct {
	Number string `json:"number"`
	List   string `json:"list"`
}

type GetCurrentChatObjectInfoResult struct {
	Wxid     string `json:"wxid"`
	V3       string `json:"v3"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Headurl  string `json:"headurl"`
}

type SwitchCurrentChatObjectReq struct {
	WxidOrGid string `json:"wxidorgid"`
}

type SwitchCurrentChatObjectResult struct {
	SwitchCurrentChatObject string `json:"SwitchCurrentChatObject"`
}

type TurnOnDoNotDisturbReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type TurnOnDoNotDisturbResult struct {
	TurnOnDoNotDisturb string `json:"TurnOnDoNotDisturb"`
}

type TurnOffDoNotDisturbReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type TurnOffDoNotDisturbResult struct {
	TurnOffDoNotDisturb string `json:"TurnOffDoNotDisturb"`
}

type SaveToContactReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type SaveToContactResult struct {
	SaveToContact string `json:"SaveToContact"`
}

type RemoveToContactReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type RemoveToContactResult struct {
	RemoveToContact string `json:"RemoveToContact"`
}

type MarkAsNoReadSessionReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type MarkAsNoReadSessionResult struct {
	MarkAsNoReadSession string `json:"MarkAsNoReadSession"`
}
