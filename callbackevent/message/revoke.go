package message

// Revoke 撤回消息
// "msg":"<sysmsg type=\"functionmsg\">
// <functionmsg>
// <cgi>/cgi-bin/micromsg-bin/pullfunctionmsg</cgi>
// <cmdid>614</cmdid>
// <functionmsgid>MMKANYIKAN_REDDOT_1695881100</functionmsgid>
// <retryinterval>150</retryinterval>
// <retrycount>3</retrycount>
// <custombuff>YImkAXIKMTY5NTg4MTEwMHoGd2VpeGluigEcTU1LQU5ZSUtBTl9SRURET1RfMTY5NTg4MTEwMKoBWEp0WlRZSGF1ZVhQZDhLY2kxY1VaZW1NVDFTUnJDZHROWnNLU1dhanp4U01ZaG5MUEh4V3Izd0FTL0xrcVp1NlNRM293Ni9WYy9mcGw1bWwyRzg2Tk9RPT0=</custombuff>
// <businessid>21001</businessid>
// <actiontime>1695880809</actiontime>
// <functionmsgversion>2</functionmsgversion>
// </functionmsg>
// </sysmsg>",
//
//	"revoke_msg":"<msg><emoji fromusername=\"azure1489\" tousername=\"wxid_pckegem76ewv12\" type=\"2\" idbuffer=\"media:0_0\" md5=\"9122f8a77147be1c45c57ae0fe86db7c\" len=\"73241\" productid=\"\" androidmd5=\"9122f8a77147be1c45c57ae0fe86db7c\" androidlen=\"73241\" s60v3md5=\"9122f8a77147be1c45c57ae0fe86db7c\" s60v3len=\"73241\" s60v5md5=\"9122f8a77147be1c45c57ae0fe86db7c\" s60v5len=\"73241\" cdnurl=\"http://wxapp.tc.qq.com/262/20304/stodownload?m=9122f8a77147be1c45c57ae0fe86db7c&amp;filekey=30350201010421301f020201060402534804109122f8a77147be1c45c57ae0fe86db7c0203011e19040d00000004627466730000000132&amp;hy=SH&amp;storeid=263195e3300010da1000000000000010600004f505348014c596096857a969&amp;bizid=1023\" designerid=\"\" thumburl=\"\" encrypturl=\"http://wxapp.tc.qq.com/262/20304/stodownload?m=64a8f0e76814adf6984166696db0df8f&amp;filekey",
//	"fromtype":"1",
type Revoke struct {
	CommonMsg
	ToCommonMsg
	RevokeMsg string `json:"revoke_msg"`
}

// GroupRevoke 群撤回消息
type GroupRevoke struct {
	CommonMsg
	ToCommonMsg
	RevokeMsg string `json:"revoke_msg"`
	CommonGroupMsg
}
