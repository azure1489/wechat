package chatroommanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type ChatRoomManagerService interface {
	// CreateChatRoom 创建群聊 create_chat_room.go https://www.showdoc.com.cn/WeChatProject/9019609319630104
	CreateChatRoom(wxidList []string) (string, error)
	// AddFriendToChatRoom 添加好友进群(40人以下) add_friend_to_chat_room.go https://www.showdoc.com.cn/WeChatProject/9020971162414036
	AddFriendToChatRoom(gid string, wxidList []string) error
	// InviteFriendToChatRoom 邀请好友进群(40人以上) invite_friend_to_chat_room.go https://www.showdoc.com.cn/WeChatProject/9021493538921140
	InviteFriendToChatRoom(wxid string, gid string) error
	// AgreeJoinGroup 同意加入群聊 agree_join_group.go https://www.showdoc.com.cn/WeChatProject/9032882631445023
	AgreeJoinGroup(groupAccessUrl string) error
	// ChatRoomEditMark 修改群聊备注 chat_room_edit_mark.go https://www.showdoc.com.cn/WeChatProject/8994766712410900
	ChatRoomEditMark(gid string, mark string) error
	// ChatRoomEditName 修改群聊名称 chat_room_edit_name.go https://www.showdoc.com.cn/WeChatProject/8993745484621751
	ChatRoomEditName(gid string, gname string) error
	// ChatRoomEditAnnouncement 修改群聊公告（群主身份才可以使用） chat_room_edit_announcement.go https://www.showdoc.com.cn/WeChatProject/8994689342748667
	ChatRoomEditAnnouncement(gid string, ownerWxid string, announcement string) error
	// ChatRoomVoip 群聊多人语音 chat_room_voip.go https://www.showdoc.com.cn/WeChatProject/9019559882256250
	ChatRoomVoip(gid string, wxidList []string) error
	// QueryChatRoomOwnerWxid 查询群主微信ID query_chat_room_owner_wxid.go https://www.showdoc.com.cn/WeChatProject/9019577323953766
	QueryChatRoomOwnerWxid(gid string) (string, error)
	// GetChatroomMemberDetailInfo 网络获取群成员详细信息,可以获取群成员微信号 get_chatroom_member_detail_info.go https://www.showdoc.com.cn/WeChatProject/9158592965427244
	GetChatroomMemberDetailInfo(gid string, wxid string) (*GetChatroomMemberDetailInfoResult, error)
	// GetChatrooMmemberDetail 批量获取群成员详细邀请信息,可以获取群成员是由谁邀请进来的 get_chatroo_mmember_detail.go https://www.showdoc.com.cn/WeChatProject/9745666070753468
	GetChatrooMmemberDetail(gid string) (*GetChatrooMmemberDetailResult, error)
	// QueryChatRoomMemberCount 本地查询群成员总数(如果查询不到则需要先点一下群聊或者调用网络获取好友或群详细信息) query_chat_room_member_count.go https://www.showdoc.com.cn/WeChatProject/9019566491411017
	QueryChatRoomMemberCount(gid string) (int, error)
	// BatchGetChatRoomMemberWxid 批量获取群成员微信ID batch_get_chat_room_member_wxid.go https://www.showdoc.com.cn/WeChatProject/9021207508542836
	BatchGetChatRoomMemberWxid(gid string) ([]string, error)
	// QueryChatRoomMemberNickName 查询群成员昵称 query_chat_room_member_nick_name.go https://www.showdoc.com.cn/WeChatProject/9019565404784998
	QueryChatRoomMemberNickName(gid string, wxid string) (string, error)
	// ChatRoomMemberBatchDelete 批量删除群成员 chat_room_member_batch_delete.go https://www.showdoc.com.cn/WeChatProject/8993612184500147
	ChatRoomMemberBatchDelete(gid string, wxidList []string) error
	// ChangeChatRoomNicknameBySelf 修改自己的群昵称 change_chat_room_nickname_by_self.go https://www.showdoc.com.cn/WeChatProject/9019580625786730
	ChangeChatRoomNicknameBySelf(gid string, wxid string, nickname string) error
	// QuitChatRoom 退出群聊 quit_chat_room.go https://www.showdoc.com.cn/WeChatProject/9002651226705549
	QuitChatRoom(gid string) error
	// DisplayChatRoomMemberNickName 显示群成员昵称 display_chat_room_member_nick_name.go https://www.showdoc.com.cn/WeChatProject/9063349491429947
	DisplayChatRoomMemberNickName(gid string) error
	// TurnOffDisplayChatRoomMemberNickName 关闭显示群成员昵称 turn_off_display_chat_room_member_nick_name.go https://www.showdoc.com.cn/WeChatProject/9063350057038457
	TurnOffDisplayChatRoomMemberNickName(gid string) error
	// CollapseChat 群聊折叠 collapse_chat.go https://www.showdoc.com.cn/WeChatProject/9063366035200767
	CollapseChat(gid string) error
	// UncollapseChat 取消群聊折叠 uncollapse_chat.go https://www.showdoc.com.cn/WeChatProject/9063366996691400
	UncollapseChat(gid string) error
	// SaveToContact 保存到通讯录 save_to_contact.go https://www.showdoc.com.cn/WeChatProject/9078597995448349
	SaveToContact(gidorwxid string) error
	// RemoveToContact 从通讯录删除 remove_to_contact.go https://www.showdoc.com.cn/WeChatProject/9078598797893485
	RemoveToContact(gidorwxid string) error
	// ChatRoomInviteCfg 群聊邀请确认 chat_room_invite_cfg.go https://www.showdoc.com.cn/WeChatProject/9859658923614486
	ChatRoomInviteCfg(gid string, opt string) (string, error)
	// AddChatRoomAdmin 添加群管理员 add_chat_room_admin.go https://www.showdoc.com.cn/WeChatProject/10518529776146702
	AddChatRoomAdmin(gid string, wxid string) error
	// DelChatRoomAdmin 删除群管理员 del_chat_room_admin.go https://www.showdoc.com.cn/WeChatProject/10518531857265007
	DelChatRoomAdmin(gid string, wxid string) error
	// TransferChatRoomOwner 转让群主 transfer_chat_room_owner.go https://www.showdoc.com.cn/WeChatProject/10518532504405659
	TransferChatRoomOwner(gid string, wxid string) error
}

type ChatRoomManagerServiceImpl struct {
	// config *wechat.WechatConfig
	http common.HttpClientService
}

func NewChatRoomManagerService(config *wechat.WechatConfig) ChatRoomManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.Secret, config.Timeout)

	return &ChatRoomManagerServiceImpl{
		// config: config,
		http: httpClientService,
	}
}
