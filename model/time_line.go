package model

type TimelineGetFristPageResult struct {
	Havenewmsg string                               `json:"havenewmsg"`
	Data       []TimelineGetFristPageResultDataItem `json:"data"`
}
type TimelineGetFristPageResultDataItem struct {
	MsgidFull10 string `json:"msgid_full_10"`
	MsgidFull16 string `json:"msgid_full_16"`
	Msgid1      string `json:"msgid_1"`
	Msgid2      string `json:"msgid_2"`
	Msgtime     string `json:"msgtime"`
	Wxid        string `json:"wxid"`
	Nickname    string `json:"nickname"`
	Content     string `json:"content"`
}

type GetFriendTimelineReq struct {
	Wxid string `json:"wxid"`
	Id   string `json:"id,omitempty"`
}

type GetFriendTimelineResult struct {
	Data []GetFriendTimelineResultDataItem `json:"data"`
}
type GetFriendTimelineResultDataItem struct {
	MsgidFull10 string `json:"msgid_full_10"`
	MsgidFull16 string `json:"msgid_full_16"`
	Msgid1      string `json:"msgid_1"`
	Msgid2      string `json:"msgid_2"`
	Msgtime     string `json:"msgtime"`
	Wxid        string `json:"wxid"`
	Nickname    string `json:"nickname"`
	Content     string `json:"content"`
}
