package httptools

type ProxyBody struct {
	Domain  string              `json:"domain"`
	UrlPath string              `json:"urlPath"`
	Method  string              `json:"method"`
	Body    string              `json:"body"`
	Headers map[string][]string `json:"headers"`
}

type ProxyReq struct {
	EncryptedAESKey string `json:"encryptedAESKey"`
	EncryptedData   string `json:"encryptedData"`
}
