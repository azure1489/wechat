package message

// Text 文本消息
type Text struct {
	CommonMsg
	ToCommonMsg
	MsgSource string `json:"msgsource"` // 消息源内容
}

// 群文本消息
type GroupText struct {
	CommonMsg
	ToCommonMsg
	MsgSource string `json:"msgsource"` // 消息源内容
	CommonGroupMsg
}
