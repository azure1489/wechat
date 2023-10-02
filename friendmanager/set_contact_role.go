package friendmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// wxid	是	string	对方微信ID
// role	是	string	权限值,1~6 1=仅聊天 2=聊天、朋友圈、微信运动等
type SetContactRoleReq struct {
	Wxid string `json:"wxid"` // 对方微信ID
	Role string `json:"role"` // 权限值,1~6 1=仅聊天 2=聊天、朋友圈、微信运动等
}

type SetContactRoleResult struct {
	SetContactRole string `json:"SetContactRole"`
}

// SetContactRole 设置好友备注 https://www.showdoc.com.cn/WeChatProject/9644721002325179
func (l *FriendManagerServiceImpl) SetContactRole(wxid string, role string) error {

	req := SetContactRoleReq{
		Wxid: wxid,
		Role: role,
	}

	resultBody, err := l.http.DoPost("/SetContactRole", req)
	if err != nil {
		return err
	}

	commonResult := SetContactRoleResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SetContactRole != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
