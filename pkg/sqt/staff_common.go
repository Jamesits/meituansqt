package sqt

type Gender int

const (
	GenderUnknown Gender = 0
	GenderMale    Gender = 1
	GenderFemale  Gender = 2
)

type StaffStatus int

const (
	StaffStatusEnabled  StaffStatus = 0
	StaffStatusDisabled StaffStatus = 3
)

type StaffInfo struct {
	StaffId       int64       `json:"staffId,omitempty"`       // 美团企业版员工ID
	Name          string      `json:"name"`                    // 姓名(不能为空)
	Gender        Gender      `json:"gender,omitempty"`        // 性别（0:未知；1:男；2:女）默认为"0:未知")
	Phone         string      `json:"phone,omitempty"`         // 手机号，且需要企业内唯一。当手机号为企业唯一标识时，此字段必传（注：手机号不包含国家代码）
	EntStaffNum   string      `json:"entStaffNum,omitempty"`   // 工号，且需要企业内唯一。当工号为企业唯一标识时，此字段必传
	Email         string      `json:"email,omitempty"`         // 邮箱，且需要企业内唯一。当邮箱为企业唯一标识时，此字段必传
	TaxNumber     string      `json:"taxNumber,omitempty"`     // 发票税号（需提前在企业后台添加发票信息）
	City          string      `json:"city,omitempty"`          // 员工所属城市
	CityId        string      `json:"cityId,omitempty"`        // 城市id（国标）
	Level         string      `json:"level,omitempty"`         // 职级
	ParentStaffId int         `json:"parentStaffId,omitempty"` // 上级美团企业版员工ID
	NotifyStaffId int         `json:"notifyStaffId,omitempty"` // 消费通知接收人美团企业版员工ID
	UserId        string      `json:"userId,omitempty"`        // 企业员工在第三方平台（如钉钉、企微等）上的唯一标识，如有传值，需要企业内唯一，当需要从钉钉或企微生态平台单点登录美团企业版时，此字段必传
	StaffStatus   StaffStatus `json:"staffStatus,omitempty"`   // 员工在职状态(0:在职；3:停用)
	CardInfos     []CardInfo  `json:"cardInfos,omitempty"`     // 员工证件信息列表
}

type CardType int

const (
	CardTypePrcId CardType = 0
	CardTypeVisa  CardType = 1
	CardTypeOther CardType = 2
)

type CardInfo struct {
	CardType    CardType `json:"cardType"`              // 证件类型，0：身份证，1：护照，2：其它
	CardName    string   `json:"cardName"`              // 证件姓名
	FirstName   string   `json:"firstName,omitempty"`   // 名（英文名）
	MiddleName  string   `json:"middleName,omitempty"`  // 中间名（英文名）
	LastName    string   `json:"lastName,omitempty"`    // 姓（英文名）
	CardNum     string   `json:"cardNum"`               // 证件号码
	Nationality string   `json:"nationality,omitempty"` // 国籍（国家标准二字码）
	Birthday    string   `json:"birthday,omitempty"`    // 出生年月日（pattern:yyyy-MM-dd）
	Sex         Gender   `json:"sex,omitempty"`         // 性别，1：男，2：女
}
