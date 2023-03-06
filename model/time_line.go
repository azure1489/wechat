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

type SendTimelineReq struct {
	SendPyqXml string `json:"send_pyq_xml"`
}

type SendTimelineResult struct {
	SendTimeline string `json:"SendTimeline"`
}

type TimelineLikeReq struct {
	ID string `json:"id"`
}

type TimelineLikeResult struct {
	TimelineLike string `json:"TimelineLike"`
}

type TimelineUndoLikeReq struct {
	ID string `json:"id"`
}

type TimelineUndoLikeResult struct {
	TimelineUndoLike string `json:"TimelineUndoLike"`
}

type TimelineCommentReq struct {
	ID  string `json:"id"`
	Msg string `json:"msg"`
}

type TimelineCommentResult struct {
	TimelineComment string `json:"TimelineComment"`
}

type TimelineDeleteCommentReq struct {
	Index string `json:"index"`
	ID    string `json:"id"`
}

type TimelineDeleteCommentResult struct {
	TimelineDeleteComment string `json:"TimelineDeleteComment"`
}

type SwitchTimelineCommentReq struct {
	Option string `json:"option"`
}

type SwitchTimelineCommentResult struct {
	SwitchTimelineComment string `json:"SwitchTimelineComment"`
}

type TimelineUploadPicReq struct {
	PicPath string `json:"PicPath"`
}

type TimelineUploadPicResult struct {
	PicURLSmall string `json:"pic_url_small"`
	PicURLBig   string `json:"pic_url_big"`
}
