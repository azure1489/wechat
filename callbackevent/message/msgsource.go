package message

import "encoding/xml"

// <msgsource>
//
//	<atuserlist>wxid_vz893djs61rs22</atuserlist>
//	<bizflag>0</bizflag>
//	<silence>0</silence>
//	<membercount>3</membercount>
//	<signature>v1_dH8Qqmsj</signature>
//	<tmp_node>
//		<publisher-id></publisher-id>
//	</tmp_node>
//
// </msgsource>
type MsgSourceXml struct {
	XMLName     xml.Name `xml:"msgsource"`
	AtUserList  string   `xml:"atuserlist"`
	BizFlag     int      `xml:"bizflag"`
	Silence     int      `xml:"silence"`
	MemberCount int      `xml:"membercount"`
	Signature   string   `xml:"signature"`
	TmpNode     struct {
		PublisherID string `xml:"publisher-id"`
	} `xml:"tmp_node"`
}
