package model

type TimelineGetFristPageResult struct {
	Count string                               `json:"count"`
	Data  []TimelineGetFristPageResultDataItem `json:"data"`
}
type TimelineGetFristPageResultDataItem struct {
	ID         string `json:"id"`
	Msgtime    string `json:"msgtime"`
	Wxid       string `json:"wxid"`
	Nickname   string `json:"nickname"`
	Content    string `json:"content"`
	CommentInc string `json:"comment_inc"`
}

type TimelineGetNextPageReq struct {
	Id string `json:"id"`
}

type TimelineGetNextPageResult struct {
	Count string                              `json:"count"`
	Data  []TimelineGetNextPageResultDataItem `json:"data"`
}
type TimelineGetNextPageResultDataItem struct {
	ID         string `json:"id"`
	Msgtime    string `json:"msgtime"`
	Wxid       string `json:"wxid"`
	Nickname   string `json:"nickname"`
	Content    string `json:"content"`
	CommentInc string `json:"comment_inc"`
}

type GetFriendTimelineReq struct {
	Wxid string `json:"wxid"`
	Id   string `json:"id,omitempty"`
}

type GetFriendTimelineResult struct {
	Data []GetFriendTimelineResultDataItem `json:"data"`
}
type GetFriendTimelineResultDataItem struct {
	ID           string `json:"id"`
	Msgtime      string `json:"msgtime"`
	Wxid         string `json:"wxid"`
	Nickname     string `json:"nickname"`
	NicknameUTF8 string `json:"nickname_UTF8"`
	Content      string `json:"content"`
	CommentInc   string `json:"comment_inc"`
}
