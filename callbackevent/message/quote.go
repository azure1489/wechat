package message

import "encoding/xml"

// Quote 引用回复消息
type Quote struct {
	CommonMsg
	ToCommonMsg
	QuoteMsgId   string `json:"quote_msg_id"`   // 引用的消息id
	QuoteMsg     string `json:"quote_msg"`      // 引用的消息内容
	QuoteMsgType string `json:"quote_msg_type"` // 引用的消息类型
	ReplyMsg     string `json:"reply_msg"`      // 回复的消息内容
	MsgSource    string `json:"msgsource"`      // 消息源内容
}

// GroupQuote 群引用回复消息
type GroupQuote struct {
	CommonMsg
	ToCommonMsg
	QuoteMsgId   string `json:"quote_msg_id"`   // 引用的消息id
	QuoteMsg     string `json:"quote_msg"`      // 引用的消息内容
	QuoteMsgType string `json:"quote_msg_type"` // 引用的消息类型
	ReplyMsg     string `json:"reply_msg"`      // 回复的消息内容
	MsgSource    string `json:"msgsource"`      // 消息源内容
	CommonGroupMsg
}

// "<?xml version=\"1.0\"?>
// <msg>
//
//	<appmsg appid=\"\" sdkver=\"0\">
//	<title>?????</title>
//	<type>57</type>
//	<appattach><cdnthumbaeskey /><aeskey></aeskey></appattach>
//	<refermsg>
//		<type>47</type>
//		<svrid>5810306839995267584</svrid>
//		<fromusr>44095400957@chatroom</fromusr>
//		<chatusr>wxid_vz893djs61rs22</chatusr>
//		<displayname>???</displayname>
//		<content>&lt;msg&gt;&lt;emoji fromusername=\"wxid_vz893djs61rs22\" tousername=\"44095400957@chatroom\" type=\"1\" idbuffer=\"media:0_0\" md5=\"297b75f565038f9f24d151d8a2245843\" len=\"70242\" productid=\"\" androidmd5=\"297b75f565038f9f24d151d8a2245843\" androidlen=\"70242\" s60v3md5=\"297b75f565038f9f24d151d8a2245843\" s60v3len=\"70242\" s60v5md5=\"297b75f565038f9f24d151d8a2245843\" s60v5len=\"70242\" cdnurl=\"http://vweixinf.tc.qq.com/110/20401/stodownload?m=297b75f565038f9f24d151d8a2245843&amp;amp;filekey=30440201010430302e02016e04025348042032393762373566353635303338663966323464313531643861323234353834330203011262040d00000004627466730000000131&amp;amp;hy=SH&amp;amp;storeid=323032313039303232323032343730303030356134303463633538373564613536626234306230303030303036653031303034666231&amp;amp;ef=1&amp;amp;bizid=1022\" designerid=\"\" thumburl=\"\" encrypturl=\"http://vweixinf.tc.qq.com/110/20402/stodownload?m=9561208bb4d681d69ffd2bfd26e837b1&amp;amp;filekey=30440201010430302e02016e04025348042039353631323038626234643638316436396666643262666432366538333762310203011270040d00000004627466730000000131&amp;amp;hy=SH&amp;amp;storeid=323032313039303232323032343730303032393762333463633538373564613536626234306230303030303036653032303034666232&amp;amp;ef=2&amp;amp;bizid=1022\" aeskey=\"7f1376ee272742868c2a6a71586260a4\" externurl=\"http://vweixinf.tc.qq.com/110/20403/stodownload?m=e3cd198e05959ed11354e2c9f930845c&amp;amp;filekey=3043020101042f302d02016e040253480420653363643139386530353935396564313133353465326339663933303834356302021260040d00000004627466730000000131&amp;amp;hy=SH&amp;amp;storeid=323032313039303232323032343730303034353161643463633538373564613536626234306230303030303036653033303034666233&amp;amp;ef=3&amp;amp;bizid=1022\" externmd5=\"9f4cfb3983b31cb8a77e043cc3fba0dc\" width=\"222\" height=\"226\" tpurl=\"\" tpauthkey=\"\" attachedtext=\"\" attachedtextcolor=\"\" lensid=\"\" emojiattr=\"\" linkid=\"\" desc=\"\"&gt;&lt;/emoji&gt;&lt;/msg&gt;</content>
//		<msgsource>&lt;msgsource&gt;&lt;sequence_id&gt;895254454&lt;/sequence_id&gt;\n\t&lt;silence&gt;0&lt;/silence&gt;\n\t&lt;membercount&gt;3&lt;/membercount&gt;\n\t&lt;signature&gt;v1_oF8Jvd3F&lt;/signature&gt;\n\t&lt;tmp_node&gt;\n\t\t&lt;publisher-id&gt;&lt;/publisher-id&gt;\n\t&lt;/tmp_node&gt;\n&lt;/msgsource&gt;\n</msgsource>
//		<createtime>1695809590</createtime>
//	</refermsg>
//	</appmsg>
//	<fromusername>azure1489</fromusername>
//	<scene>0</scene>
//	<appinfo><version>1</version><appname /></appinfo>
//	<commenturl />
//
// </msg>"
// type AppMsgXml struct {
// 	Msg struct {
// 		AppMsg struct {
// 			Title    string `xml:"title"`
// 			Type     string `xml:"type"`
// 			ReferMsg struct {
// 				Type        string `xml:"type"`
// 				Svrid       string `xml:"svrid"`
// 				Fromusr     string `xml:"fromusr"`
// 				Chatusr     string `xml:"chatusr"`
// 				DisplayName string `xml:"displayname"`
// 				Content     string `xml:"content"`
// 				MsgSource   string `xml:"msgsource"`
// 				CreateTime  string `xml:"createtime"`
// 			} `xml:"refermsg"`
// 			FromUserName string `xml:"fromusername"`
// 		} `xml:"appmsg"`
// 	} `xml:"msg"`
// }

type AppMsgXml struct {
	XMLName xml.Name `xml:"msg"`
	AppMsg  struct {
		Title     string `xml:"title"`
		Type      string `xml:"type"`
		AppAttach struct {
			CdnThumbAesKey string `xml:"cdnthumbaeskey"`
			AesKey         string `xml:"aeskey"`
		} `xml:"appattach"`
		ReferMsg struct {
			Type        string `xml:"type"`
			Svrid       string `xml:"svrid"`
			Fromusr     string `xml:"fromusr"`
			Chatusr     string `xml:"chatusr"`
			DisplayName string `xml:"displayname"`
			Content     string `xml:"content"`
			MsgSource   string `xml:"msgsource"`
			CreateTime  string `xml:"createtime"`
		} `xml:"refermsg"`
		FromUserName string `xml:"fromusername"`
	} `xml:"appmsg"`
}

// type QuoteXmlMsg struct {
// 	AppMsg QuoteXmlAppMsg `xml:"appmsg"`
// }

// type QuoteXmlAppMsg struct {
// 	Title        string                 `xml:"title"`
// 	Type         string                 `xml:"type"`
// 	ReferMsg     QuoteXmlAppMsgReferMsg `xml:"refermsg"`
// 	FromUserName string                 `xml:"fromusername"`
// }

// type QuoteXmlAppMsgReferMsg struct {
// 	Type        string `xml:"type"`
// 	Svrid       string `xml:"svrid"`
// 	Fromusr     string `xml:"fromusr"`
// 	Chatusr     string `xml:"chatusr"`
// 	Displayname string `xml:"displayname"`
// 	Content     string `xml:"content"`
// 	Msgsource   string `xml:"msgsource"`
// 	Createtime  string `xml:"createtime"`
// }
