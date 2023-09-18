package sqt

import (
	"context"
	"github.com/jamesits/meituansqt/pkg/sqtapi"
	"net/http"
)

type StaffBatchDeleteRequest struct {
	StaffIdType      int      `json:"staffIdType"`
	StaffIdentifiers []string `json:"staffIdentifiers"`
}

type StaffDeleteResult int

const (
	StaffDeleteResultSucceed StaffDeleteResult = 0
	StaffDeleteResultFailed  StaffDeleteResult = 1
)

type StaffDeleteResultItem struct {
	Result          StaffDeleteResult `json:"result"`
	Msg             string            `json:"msg"`
	StaffIdentifier string            `json:"staffIdentifier,omitempty"`
}

type StaffBatchDeleteResponse struct {
	StaffQueryResultItems []StaffDeleteResultItem `json:"StaffDeleteResultItems"`
}

// 批量删除员工接口
func (sqt *SQT) StaffBatchDelete(ctx context.Context, req *StaffBatchDeleteRequest) (*StaffBatchDeleteResponse, *http.Response, error) {
	return sqtapi.WrappedApi[StaffBatchDeleteRequest, StaffBatchDeleteResponse](http.MethodPost, "staff/batch/delete", "staff.batch.delete")(ctx, sqt.Client, &sqt.Config, req)
}
