package message

// Voice 语音消息
type Voice struct {
	VoiceLen  string `json:"voice_len"`
	VoiceData string `json:"voice_data"`
}

// GroupVoice 语音消息
type GroupVoice struct {
	Voice
	CommonGroupMsg
}
