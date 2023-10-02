package friendmanager

import (
	"encoding/json"
	"fmt"
)

//	{
//	    "v3": "v3_020b3826fd03010000000000519c389df92e05000000501ea9a3dba12f95f6b60a0536a1adb6a5fe833159e36adbba153c16649b8da81936d019e8f5d3d2b48692fbac566e5b2ff089a19cd6ccd47693fe51ef9fad11754dec3b282fcea4a80c9f5d5a@stranger",
//	    "v4": "v4_000b708f0b040000010000000000744aedde2a4a2b975f194365e3621000000050ded0b020927e3c97896a09d47e6e9ee6a92f3a8adc008ba7cea408ff1850bc5e4383b5bbe3d87d2d56d66982212e656f7363950a20ff12ba2aa8984c302de95ad574a40ed3254f0f88c18f613974095d9aef785cf7642414da31e215545775a432f05e5c14d58f2938a0dc5598eeaa079aba9518bf350931ca31e40256052b88e0e018ec29edf4c6167a465588c3021df446e556e70ab0c15811a7fae91e78b40db6b2fc42ffd73aac2e99f6c51f56896c6fb97321e29e24044d7b6303f8237f195a77364bb11c6de40412270fcc152b192bfd09c2dc5d699336c1a32a3157602d920b7f8a2ed925f9d105d69daa23122a65d76fec766b876a8fa33f6ed71bba78992e274fd4c5a4d4c75437ea2618b347359c3cd4322c72@stranger",
//	    "role": "0",
//	    "from": "17"
//	}
type VerifyFriendReq struct {
	V3   string `json:"v3"`
	V4   string `json:"v4"`
	Role string `json:"role"`
	From string `json:"from"`
}

type VerifyFriendResult struct {
	Success string `json:"success"`
}

// VerifyFriend 通过好友申请 https://www.showdoc.com.cn/WeChatProject/8996001801759077
func (l *FriendManagerServiceImpl) VerifyFriend(req *VerifyFriendReq) error {

	resultBody, err := l.http.DoPost("/VerifyFriend", req)
	if err != nil {
		return err
	}

	commonResult := VerifyFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
