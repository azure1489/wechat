package message

// Gif gif消息
type Gif struct {
	GifPath string `json:"gif_path"`
	Desc    string `json:"desc"`
}

// GroupGif gif消息
// type GroupGif struct {
// 	Gif
// 	CommonGroupMsg
// }

type EmojiMsg struct {
	Emoji Emoji `xml:"emoji"`
}

type Emoji struct {
	FromUserName      string  `xml:"fromusername,attr"`
	ToUserName        string  `xml:"tousername,attr"`
	Type              string  `xml:"type,attr"`
	IDBuffer          string  `xml:"idbuffer,attr"`
	MD5               string  `xml:"md5,attr"`
	Len               string  `xml:"len,attr"`
	ProductID         string  `xml:"productid,attr"`
	AndroidMD5        string  `xml:"androidmd5,attr"`
	AndroidLen        string  `xml:"androidlen,attr"`
	S60v3MD5          string  `xml:"s60v3md5,attr"`
	S60v3Len          string  `xml:"s60v3len,attr"`
	S60v5MD5          string  `xml:"s60v5md5,attr"`
	S60v5Len          string  `xml:"s60v5len,attr"`
	CdnURL            string  `xml:"cdnurl,attr"`
	DesignerID        string  `xml:"designerid,attr"`
	ThumbURL          string  `xml:"thumburl,attr"`
	EncryptURL        string  `xml:"encrypturl,attr"`
	AesKey            string  `xml:"aeskey,attr"`
	ExternURL         string  `xml:"externurl,attr"`
	ExternMD5         string  `xml:"externmd5,attr"`
	Width             string  `xml:"width,attr"`
	Height            string  `xml:"height,attr"`
	TPURL             string  `xml:"tpurl,attr"`
	TPAuthKey         string  `xml:"tpauthkey,attr"`
	AttachedText      string  `xml:"attachedtext,attr"`
	AttachedTextColor string  `xml:"attachedtextcolor,attr"`
	LensID            string  `xml:"lensid,attr"`
	EmojiAttr         string  `xml:"emojiattr,attr"`
	LinkID            string  `xml:"linkid,attr"`
	Desc              string  `xml:"desc,attr"`
	GameExt           GameExt `xml:"gameext"`
}

type GameExt struct {
	Type    string `xml:"type,attr"`
	Content string `xml:"content,attr"`
}
