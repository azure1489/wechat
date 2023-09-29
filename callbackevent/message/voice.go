package message

// Voice 语音消息
type Voice struct {
	CommonMsg
	ToCommonMsg
	VoiceLen string `json:"voice_len"`
	VoiceHex string `json:"voice_hex"`
}

// GroupVoice 语音消息
type GroupVoice struct {
	CommonMsg
	ToCommonMsg
	VoiceLen string `json:"voice_len"`
	VoiceHex string `json:"voice_hex"`
	CommonGroupMsg
}
