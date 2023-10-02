package contactmanager

import "encoding/json"

//	{
//		"wxid": "themid",
//		"v3": "v3_.............",
//		"nickname": "电信-小张",
//		"name": "我爱中国",
//		"headurl": "http://....."
//		}
type GetCurrentChatObjectInfoResult struct {
	Wxid     string `json:"wxid"`
	V3       string `json:"v3"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Headurl  string `json:"headurl"`
}

// 获取当前聊天对象信息 https://www.showdoc.com.cn/WeChatProject/9019586393938696
func (c *ContactManagerServiceImpl) GetCurrentChatObjectInfo() (*GetCurrentChatObjectInfoResult, error) {

	resultBody, err := c.http.DoPost("/GetCurrentChatObjectInfo", nil)
	if err != nil {
		return nil, err
	}

	commonResult := GetCurrentChatObjectInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
