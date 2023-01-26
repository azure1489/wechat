package model

type CreateChatRoomReq struct {
	WxidList string `json:"wxidlist"`
}

type CreateChatRoomResult struct {
	Retstr string `json:"retstr"`
	Gid    string `json:"gid"`
}

type AddFriendToChatRoomReq struct {
	Gid      string `json:"gid"`
	WxidList string `json:"wxidlist"`
}

type AddFriendToChatRoomResult struct {
	AddFriendToChatRoom string `json:"AddFriendToChatRoom"`
}

type InviteFriendToChatRoomReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
}

type InviteFriendToChatRoomResult struct {
	InviteFriendToChatRoom string `json:"InviteFriendToChatRoom"`
}

type AgreeJoinGroupRoomReq struct {
	GroupAccessUrl string `json:"GroupAccessUrl"`
}

type AgreeJoinGroupResult struct {
	AgreeJoinGroup string `json:"AgreeJoinGroup"`
}

type ChatRoomEditMarkReq struct {
	Gid  string `json:"gid"`
	Mark string `json:"mark"`
}

type ChatRoomEditMarkResult struct {
	ChatRoomEditMark string `json:"ChatRoomEditMark"`
}

type ChatRoomEditNameReq struct {
	Gid   string `json:"gid"`
	Gname string `json:"gname"`
}

type ChatRoomEditNameResult struct {
	ChatRoomEditName string `json:"ChatRoomEditName"`
}

type ChatRoomEditAnnouncementReq struct {
	Gid          string `json:"gid"`
	OwnerWxid    string `json:"ownerwxid"`
	Announcement string `json:"announcement"`
}

type ChatRoomEditAnnouncementResult struct {
	ChatRoomEditAnnouncement string `json:"ChatRoomEditAnnouncement"`
}

type ChatRoomVoipReq struct {
	Gid      string `json:"gid"`
	WxidList string `json:"wxidlist"`
}

type ChatRoomVoipResult struct {
	ChatRoomVoip string `json:"ChatRoomVoip"`
}

type QueryChatRoomOwnerWxidReq struct {
	Gid string `json:"gid"`
}

type QueryChatRoomOwnerWxidResult struct {
	OwnerWxid string `json:"OwnerWxid"`
}

type GetChatroomMemberDetailInfoReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
}

type GetChatroomMemberDetailInfoResult struct {
	Type          string `json:"type"`
	Wxid          string `json:"wxid"`
	Nickname      string `json:"nickname"`
	Province      string `json:"province"`
	Area          string `json:"area"`
	Signinfo      string `json:"signinfo"`
	Wxaccount     string `json:"wxaccount"`
	Timelinebgurl string `json:"timelinebgurl"`
	Country       string `json:"country"`
	Headurl       string `json:"headurl"`
	V3            string `json:"v3"`
	Fromchatroom  string `json:"fromchatroom"`
	V4            string `json:"v4"`
}

type GetChatrooMmemberDetailReq struct {
	Gid string `json:"gid"`
}

type GetChatrooMmemberDetailResult struct {
	Type           string                                    `json:"type"`
	Gid            string                                    `json:"gid"`
	MemberCount    string                                    `json:"member_count"`
	MasterWxid     string                                    `json:"master_wxid"`
	MasterNickname string                                    `json:"master_nickname"`
	ChatroomName   string                                    `json:"chatroom_name"`
	Member         []GetChatrooMmemberDetailResultMemberItem `json:"member"`
}

type GetChatrooMmemberDetailResultMemberItem struct {
	Wxid          string `json:"wxid"`
	Nickname      string `json:"nickname"`
	UserHeadBig   string `json:"user_head_big"`
	UserHeadSmall string `json:"user_head_small"`
	UserFlag      string `json:"user_flag"`
	InviterWxid   string `json:"inviter_wxid"`
}

type QueryChatRoomMemberCountReq struct {
	Gid string `json:"gid"`
}

type QueryChatRoomMemberCountResult struct {
	Count string `json:"count"`
}

type BatchGetChatRoomMemberWxidReq struct {
	Gid string `json:"gid"`
}

type BatchGetChatRoomMemberWxidResult struct {
	Data []BatchGetChatRoomMemberWxidResultDataItem `json:"data"`
}

type BatchGetChatRoomMemberWxidResultDataItem struct {
	Wxid string `json:"wxid"`
}

type QueryChatRoomMemberNickNameReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
}

type QueryChatRoomMemberNickNameResult struct {
	Nickname string `json:"nickname"`
}
