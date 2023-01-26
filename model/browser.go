package model

type OpenInlineBrowserReq struct {
	URL    string `json:"url"`
	ToWxid string `json:"to_wxid,omitempty"`
}

type OpenInlineBrowserResult struct {
	OpenInlineBrowser string `json:"OpenInlineBrowser"`
}
