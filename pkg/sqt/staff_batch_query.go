package sqt

import (
	"context"
	"github.com/jamesits/meituansqt/pkg/sqtapi"
	"net/http"
)

type StaffIdType int32

const (
	StaffIdTypeEmail       StaffIdType = 20
	StaffIdTypeEntStaffNum StaffIdType = 30
	StaffIdTypeStaffId     StaffIdType = 40
	StaffIdTypePhone       StaffIdType = 50
)

type StaffBatchQueryRequest struct {
	StaffIdType      StaffIdType `json:"staffIdType"`
	StaffIdentifiers []string    `json:"staffIdentifiers"`
}

type StaffBatchQueryResultItem struct {
	StaffId       int32  `json:"staffId"`
	EntStaffNum   string `json:"entStaffNum,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email,omitempty"`
	BindStatus    int32  `json:"bindStatus,omitempty"`
	ParentStaffId int32  `json:"parentStaffId,omitempty"`
	Level         string `json:"level,omitempty"`
	UserId        string `json:"userId,omitempty"`
}

type StaffBatchQueryResponse struct {
	StaffQueryResultItems []StaffBatchQueryResultItem `json:"staffQueryResultItems"`
}

// 批量查询员工信息接口
func (sqt *SQT) StaffBatchQuery(ctx context.Context, req *StaffBatchQueryRequest) (*StaffBatchQueryResponse, *http.Response, error) {
	return sqtapi.WrappedApi[StaffBatchQueryRequest, StaffBatchQueryResponse](http.MethodPost, "staff/batch/query", "staff.batch.query")(ctx, sqt.Client, &sqt.Config, req)
}
