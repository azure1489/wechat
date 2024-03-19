package friendmanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type FriendManagerService interface {
	// AddNewFriend 添加好友 add_new_friend.go
	AddNewFriend(req AddNewFriendReq) (*AddNewFriendResult, error)
	// DeleteFriend 删除好友 delete_friend.go
	DeleteFriend(wxid string) error
	// BlackFriend 拉黑好友 black_friend.go
	BlackFriend(wxid string) error
	// VerifyFriend 同意好友申请 verify_friend.go
	VerifyFriend(req *VerifyFriendReq) error
	// EditFriendMark 编辑好友备注 edit_friend_mark.go
	EditFriendMark(wxid string, mark string) error
	// VoiceCallFriend 好友语音聊天 voice_call_friend.go
	VoiceCallFriend(wxid string) error
	// VideoCallFriend 好友视频聊天 video_call_friend.go
	VideoCallFriend(wxid string) error
	// QueryBodyInfo 网络查询陌生人信息 query_body_info.go
	QueryBodyInfo(serachwhat string) (*QueryBodyInfoResult, error)
	// SetStarTag 设置星标好友 set_star_tag.go
	SetStarTag(wxid string) error
	// UndoStarTag 取消星标好友 undo_star_tag.go
	UndoStarTag(wxid string) error
	// SetContactRole 设置朋友权限 set_contact_role.go
	SetContactRole(wxid string, role string) error
}

type FriendManagerServiceImpl struct {
	// config *wechat.WechatConfig
	http common.HttpClientService
}

func NewFriendManagerService(config *wechat.WechatConfig) FriendManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.PublicKeyPath, config.Timeout)

	return &FriendManagerServiceImpl{
		// config: config,
		http: httpClientService,
	}
}
