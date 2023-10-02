package contactmanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type ContactManagerService interface {

	// 初始化通讯录
	InitContact() (*InitContactResult, error)

	// 获取好友和群清单
	GetFriendAndChatRoomList(string) (*GetFriendAndChatRoomListResult, error)

	// 批量获取联系人简明信息
	BatchGetContactBriefInfo([]string) (*BatchGetContactBriefInfoReult, error)

	// 获取头像信息
	GetHeadIMG(string) (*GetHeadIMGResult, error)

	// GetFriendOrChatroomDetailInfo 网络获取好友或群详细信息,好友则返回好友信息，群聊则返回群成员数量。成员微信id，成员昵称 get_friend_or_chatroom_detail_info
	GetFriendOrChatroomDetailInfo(string) (*GetFriendOrChatroomDetailInfoResult, error)

	// GetCurrentChatObjectInfo 获取当前聊天对象信息 get_current_chat_object_info
	GetCurrentChatObjectInfo() (*GetCurrentChatObjectInfoResult, error)

	// SwitchCurrentChatObject 切换当前聊天对象 switch_current_chat_object
	SwitchCurrentChatObject(wxidorgid string) error

	// TurnOnDoNotDisturb 开启消息免打扰 turn_on_do_not_disturb
	TurnOnDoNotDisturb(gidorwxid string) error

	// TurnOffDoNotDisturb 关闭消息免打扰 turn_off_do_not_disturb
	TurnOffDoNotDisturb(gidorwxid string) error

	// StickyChat 置顶聊天 sticky_chat
	StickyChat(gidorwxid string) error

	// UnpinChat 取消置顶聊天 unpin_chat
	UnpinChat(gidorwxid string) error

	// SaveToContact 保存到通讯录 save_to_contact
	SaveToContact(gidorwxid string) error

	// RemoveToContact 从通讯录删除 remove_to_contact
	RemoveToContact(gidorwxid string) error

	// MarkAsReadSession 标记为已读 mark_as_read_session
	MarkAsReadSession(gidorwxid string) error

	// MarkAsNoReadSession 标记为未读 mark_as_no_read_session
	MarkAsNoReadSession(gidorwxid string) error

	// BeforeTransfer 查询好友实名 before_transfer
	BeforeTransfer(wxid string) (string, error)
}

type ContactManagerServiceImpl struct {
	// config *wechat.WechatConfig
	http common.HttpClientService
}

func NewContactManagerService(config *wechat.WechatConfig) ContactManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.Secret, config.Timeout)

	return &ContactManagerServiceImpl{
		// config: config,
		http: httpClientService,
	}
}
