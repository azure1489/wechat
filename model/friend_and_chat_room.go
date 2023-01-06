package model

type FriendAndChatRoomResult struct {
	Friends  []FriendItem   `json:"friend"`
	Chatroom []ChatroomItem `json:"chatroom"`
	Gh       []GhItem       `json:"gh"`
}

type FriendItem struct {
	Index    string `json:"index"`
	Wxid     string `json:"wxid"`
	Account  string `json:"account"`
	Markname string `json:"markname"`
	Nickname string `json:"nickname"`
	Headimg  string `json:"headimg"`
	V3       string `json:"v3"`
}

type ChatroomItem struct {
	Index    string `json:"index"`
	Gid      string `json:"gid"`
	Markname string `json:"markname"`
	Gname    string `json:"gname"`
	V3       string `json:"v3"`
}

type GhItem struct {
	Index    string `json:"index"`
	Wxid     string `json:"wxid"`
	Account  string `json:"account"`
	Markname string `json:"markname"`
	Nickname string `json:"nickname"`
	Headimg  string `json:"headimg"`
	V3       string `json:"v3"`
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
