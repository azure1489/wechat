package message

// Image 图片消息
type Image struct {
	CommonMsg
	ToCommonMsg
	Info      string  `json:"info"`
	ImgPath   string  `json:"img_path"`
	ImgLen    float64 `json:"img_len"`
	ImgBase64 string  `json:"img_base64"`
}

// GroupImage 图片消息
type GroupImage struct {
	CommonMsg
	ToCommonMsg
	Info      string  `json:"info"`
	ImgPath   string  `json:"img_path"`
	ImgLen    float64 `json:"img_len"`
	ImgBase64 string  `json:"img_base64"`
	CommonGroupMsg
}
