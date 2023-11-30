package message

// Text 文本消息
type Text struct {
	MsgSource string `json:"msgsource"` // 消息源内容
}

// 群文本消息
// type GroupText struct {
// 	Text
// 	CommonGroupMsg
// }
