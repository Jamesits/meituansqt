package sqt

import (
	"context"
	"github.com/jamesits/meituansqt/pkg/sqtapi"
	"net/http"
)

type StaffBatchAddRequest struct {
	StaffInfos []StaffInfo `json:"staffInfos"` // 添加到美团企业版的人员信息
}

type StaffAddResult int

const (
	StaffAddResultSucceed                  StaffAddResult = 0
	StaffAddResultFailed                   StaffAddResult = 1
	StaffAddResultFailedToSaveCardInfoOnly StaffAddResult = 2
)

type StaffAddResultItem struct {
	Result      StaffAddResult `json:"result"`                // 是否成功添加 0-已成功 1-未成功 2-员工信息保存成功但证件信息保存失败）
	Msg         string         `json:"msg"`                   // 描述
	StaffId     int64          `json:"staffId,omitempty"`     // 美团企业版员工ID
	EntStaffNum string         `json:"entStaffNum,omitempty"` // 员工工号
	Phone       string         `json:"phone,omitempty"`       // 手机号
	Email       string         `json:"email,omitempty"`       // 邮箱
}

type StaffBatchAddResponse struct {
	StaffAddResultItems []StaffAddResultItem `json:"staffAddResultItems"` // 添加结果列表
}

// 批量添加员工接口
func (sqt *SQT) StaffBatchAdd(ctx context.Context, req *StaffBatchAddRequest) (*StaffBatchAddResponse, *http.Response, error) {
	return sqtapi.WrappedApi[StaffBatchAddRequest, StaffBatchAddResponse](http.MethodPost, "staff/batch/add", "staff.batch.add")(ctx, sqt.Client, &sqt.Config, req)
}
