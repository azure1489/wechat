package labelmanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type LabelManagerService interface {

	// GetContactLabelList 获取所有标签清单 get_contact_label_list.go
	GetContactLabelList() (*GetContactLabelListResult, error)

	// QueryDB 查询用户所属标签ID清单 query_db.go
	QueryDB(wxidList []string) (*QueryDBResult, error)

	// MoveToLabel 移动成员到指定的标签下 move_to_label.go
	MoveToLabel(wxidList []string, labelid string) error

	// CreateLabel 创建标签 create_label.go
	CreateLabel(labelname string) error

	// EditLabelName 编辑标签名称 edit_label_name.go
	EditLabelName(labelId string, labelName string) error

	// DeleteLabel 删除标签 delete_label.go
	DeleteLabel(labelId string) error
}

type LabelManagerServiceImpl struct {
	// config *wechat.WechatConfig
	http common.HttpClientService
}

func NewLabelManagerService(config *wechat.WechatConfig) LabelManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.PublicKeyPath, config.Timeout)

	return &LabelManagerServiceImpl{
		// config: config,
		http: httpClientService,
	}
}
