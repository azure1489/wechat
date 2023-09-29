package message

// Gif gif消息
type Gif struct {
	CommonMsg
	ToCommonMsg
	GifPath string `json:"gif_path"`
}

// GroupGif gif消息
type GroupGif struct {
	CommonMsg
	ToCommonMsg
	GifPath string `json:"gif_path"`
	CommonGroupMsg
}
