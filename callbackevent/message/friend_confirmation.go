package message

// FriendConfirmation 好友确认消息
type FriendConfirmation struct {
	V3           string `json:"v3"`
	V4           string `json:"v4"`
	Content      string `json:"content"`
	FromUserName string `json:"fromUserName"`
	FromNickName string `json:"fromNickName"`
	HeadImgURL   string `json:"headImgUrl"`
	Sex          string `json:"sex"` // 1: 男 2: 女
}

type BrandList struct {
	Count int    `xml:"count,attr"`
	Ver   string `xml:"ver,attr"`
}

type FriendConfirmationMsgXml struct {
	FromUserName      string    `xml:"fromusername,attr"`
	EncryptUserName   string    `xml:"encryptusername,attr"`
	FromNickName      string    `xml:"fromnickname,attr"`
	Content           string    `xml:"content,attr"`
	FullPy            string    `xml:"fullpy,attr"`
	ShortPy           string    `xml:"shortpy,attr"`
	ImageStatus       string    `xml:"imagestatus,attr"`
	Scene             string    `xml:"scene,attr"`
	Country           string    `xml:"country,attr"`
	Province          string    `xml:"province,attr"`
	City              string    `xml:"city,attr"`
	Sign              string    `xml:"sign,attr"`
	PerCard           string    `xml:"percard,attr"`
	Sex               string    `xml:"sex,attr"`
	Alias             string    `xml:"alias,attr"`
	Weibo             string    `xml:"weibo,attr"`
	AlbumFlag         string    `xml:"albumflag,attr"`
	AlbumStyle        string    `xml:"albumstyle,attr"`
	AlbumBgImgID      string    `xml:"albumbgimgid,attr"`
	SnsFlag           string    `xml:"snsflag,attr"`
	SnsBgImgID        string    `xml:"snsbgimgid,attr"`
	SnsBgObjectID     string    `xml:"snsbgobjectid,attr"`
	MHash             string    `xml:"mhash,attr"`
	MFullHash         string    `xml:"mfullhash,attr"`
	BigHeadImgURL     string    `xml:"bigheadimgurl,attr"`
	SmallHeadImgURL   string    `xml:"smallheadimgurl,attr"`
	Ticket            string    `xml:"ticket,attr"`
	OpCode            string    `xml:"opcode,attr"`
	GoogleContact     string    `xml:"googlecontact,attr"`
	QrTicket          string    `xml:"qrticket,attr"`
	ChatRoomUserName  string    `xml:"chatroomusername,attr"`
	SourceUserName    string    `xml:"sourceusername,attr"`
	SourceNickName    string    `xml:"sourcenickname,attr"`
	ShareCardUserName string    `xml:"sharecardusername,attr"`
	ShareCardNickName string    `xml:"sharecardnickname,attr"`
	CardVersion       string    `xml:"cardversion,attr"`
	ExtFlag           string    `xml:"extflag,attr"`
	BrandList         BrandList `xml:"brandlist"`
}
