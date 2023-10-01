package message

// Quote 引用回复消息
type Quote struct {
	QuoteMsgId   string `json:"quote_msg_id"`   // 引用的消息id
	QuoteMsg     string `json:"quote_msg"`      // 引用的消息内容
	QuoteMsgType string `json:"quote_msg_type"` // 引用的消息类型
	ReplyMsg     string `json:"reply_msg"`      // 回复的消息内容
	MsgSource    string `json:"msgsource"`      // 消息源内容
}

// GroupQuote 群引用回复消息
type GroupQuote struct {
	Quote
	CommonGroupMsg
}
