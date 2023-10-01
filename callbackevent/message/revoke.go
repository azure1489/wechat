package message

// Revoke 撤回消息

// "revoke_msg":"<msg><emoji fromusername=\"azure1489\" tousername=\"wxid_pckegem76ewv12\" type=\"2\" idbuffer=\"media:0_0\" md5=\"9122f8a77147be1c45c57ae0fe86db7c\" len=\"73241\" productid=\"\" androidmd5=\"9122f8a77147be1c45c57ae0fe86db7c\" androidlen=\"73241\" s60v3md5=\"9122f8a77147be1c45c57ae0fe86db7c\" s60v3len=\"73241\" s60v5md5=\"9122f8a77147be1c45c57ae0fe86db7c\" s60v5len=\"73241\" cdnurl=\"http://wxapp.tc.qq.com/262/20304/stodownload?m=9122f8a77147be1c45c57ae0fe86db7c&amp;filekey=30350201010421301f020201060402534804109122f8a77147be1c45c57ae0fe86db7c0203011e19040d00000004627466730000000132&amp;hy=SH&amp;storeid=263195e3300010da1000000000000010600004f505348014c596096857a969&amp;bizid=1023\" designerid=\"\" thumburl=\"\" encrypturl=\"http://wxapp.tc.qq.com/262/20304/stodownload?m=64a8f0e76814adf6984166696db0df8f&amp;filekey",
// "fromtype":"1",
type Revoke struct {
	RevokeMsg  string `json:"revoke_msg"`
	Session    string `json:"session"`
	MsgId      string `json:"msgid"`
	NewMsgId   string `json:"newmsgid"`
	ReplaceMsg string `json:"replacemsg"`
}

// GroupRevoke 群撤回消息
type GroupRevoke struct {
	Revoke
	CommonGroupMsg
}
