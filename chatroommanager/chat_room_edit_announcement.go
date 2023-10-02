package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type ChatRoomEditAnnouncementReq struct {
	Gid          string `json:"gid"`
	OwnerWxid    string `json:"ownerwxid"`
	Announcement string `json:"announcement"`
}

type ChatRoomEditAnnouncementResult struct {
	Success string `json:"success"`
}

// ChatRoomEditAnnouncement 修改群聊公告（群主身份才可以使用） https://www.showdoc.com.cn/WeChatProject/8994689342748667
func (l *ChatRoomManagerServiceImpl) ChatRoomEditAnnouncement(gid string, ownerWxid string, announcement string) error {

	req := ChatRoomEditAnnouncementReq{
		Gid:          gid,
		OwnerWxid:    ownerWxid,
		Announcement: announcement,
	}

	resultBody, err := l.http.DoPost("/ChatRoomEditAnnouncement", req)
	if err != nil {
		return err
	}

	commonResult := ChatRoomEditAnnouncementResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
