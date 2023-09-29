package message

// Quote 引用回复消息
type Quote struct {
	CommonMsg
	ToCommonMsg
	QuoteMsgId string `json:"quote_msg_id"` // 引用的消息id
	QuoteMsg   string `json:"quote_msg"`    // 引用的消息内容
	ReplyMsg   string `json:"reply_msg"`    // 回复的消息内容
}

// GroupQuote 群引用回复消息
type GroupQuote struct {
	CommonMsg
	ToCommonMsg
	QuoteMsgId string `json:"quote_msg_id"` // 引用的消息id
	QuoteMsg   string `json:"quote_msg"`    // 引用的消息内容
	ReplyMsg   string `json:"reply_msg"`    // 回复的消息内容
	CommonGroupMsg
}
