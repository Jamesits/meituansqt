package sqt

import (
	"context"
	"github.com/jamesits/meituansqt/pkg/sqtapi"
	"net/http"
)

type StaffBatchQueryRequest struct {
	StaffIdType      int      `json:"staffIdType"`
	StaffIdentifiers []string `json:"staffIdentifiers"`
}

type StaffBatchQueryResultItem struct {
	StaffId       int    `json:"staffId"`
	EntStaffNum   string `json:"entStaffNum,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email,omitempty"`
	BindStatus    int    `json:"bindStatus,omitempty"`
	ParentStaffId int    `json:"parentStaffId,omitempty"`
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
