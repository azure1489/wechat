package message

// Video 视频消息
type Video struct {
	VideoPath string `json:"video_path"`
	Info      string `json:"info"`
}

// GroupVideo 视频消息
// type GroupVideo struct {
// 	Video
// 	CommonGroupMsg
// }
