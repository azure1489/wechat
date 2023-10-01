package message

import "encoding/xml"

// "msg":"
// "
// <sysmsg type=\"revokemsg\">
//
//	<revokemsg>
//	<session>azure1489</session>
//	<msgid>1799034216</msgid>
//	<newmsgid>4300207189504548072</newmsgid>
//	<replacemsg><![CDATA[\"?????\" ???????]]></replacemsg>
//	</revokemsg>
//
// </sysmsg>
// ",
// "msgsvrid":"4138569175",
//
//	"revoke_msg":"?",
type SysMsgXml struct {
	XMLName   xml.Name `xml:"sysmsg"`
	Type      string   `xml:"type,attr"`
	RevokeMsg struct {
		Session    string `xml:"session"`
		MsgId      string `xml:"msgid"`
		NewMsgId   string `xml:"newmsgid"`
		ReplaceMsg string `xml:"replacemsg"`
	} `xml:"revokemsg"`
}
