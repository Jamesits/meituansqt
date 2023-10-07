package sqt

import (
	"context"
	"github.com/jamesits/meituansqt/pkg/sqtapi"
	"net/http"
)

type StaffBatchUpdateRequest struct {
	StaffInfos []StaffInfo `json:"staffInfos"` // 要更新的人员信息
}

type StaffUpdateResult int

const (
	StaffUpdateResultSucceed                  StaffUpdateResult = 0
	StaffUpdateResultFailed                   StaffUpdateResult = 1
	StaffUpdateResultFailedToSaveCardInfoOnly StaffUpdateResult = 2
)

type StaffUpdateResultItem struct {
	Result      StaffUpdateResult `json:"result"`                // 是否成功更新 0-已成功 1-未成功 2-员工信息保存成功但证件信息保存失败
	Msg         string            `json:"msg"`                   // 描述
	StaffId     int64             `json:"staffId,omitempty"`     // 美团企业版员工ID
	EntStaffNum string            `json:"entStaffNum,omitempty"` // 员工工号
	Phone       string            `json:"phone,omitempty"`       // 手机号
	Email       string            `json:"email,omitempty"`       // 邮箱
}

type StaffBatchUpdateResponse struct {
	StaffBatchUpdateResultItems []StaffUpdateResultItem `json:"staffUpdateResultItems"` // 更新结果列表
}

// 批量更新员工接口
func (sqt *SQT) StaffBatchUpdate(ctx context.Context, req *StaffBatchUpdateRequest) (*StaffBatchUpdateResponse, *http.Response, error) {
	return sqtapi.WrappedApi[StaffBatchUpdateRequest, StaffBatchUpdateResponse](http.MethodPost, "staff/batch/update", "staff.batch.update")(ctx, sqt.Client, &sqt.Config, req)
}
