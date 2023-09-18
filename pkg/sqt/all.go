package sqt

import (
	"context"
	"github.com/jamesits/meituansqt/pkg/sqtapi"
	"net/http"
)

type AdminDivisionInfo struct {
	AdCode       string `json:"adCode"`                 // 行政区划编号
	ChineseName  string `json:"chineseName"`            // 行政区划中文名称
	ParentAdCode int    `json:"parentAdCode,omitempty"` // 父级行政区划编号
	Level        int    `json:"level"`                  // 区划层级
}

func (sqt *SQT) AllAdminDivisionQuery(ctx context.Context) (*[]AdminDivisionInfo, *http.Response, error) {
	return sqtapi.WrappedApi[any, []AdminDivisionInfo](http.MethodPost, "tmc/trip/query/allAdminDivision", "all.admin.division.query")(ctx, sqt.Client, &sqt.Config, nil)
}
