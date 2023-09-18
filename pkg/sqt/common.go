package sqt

import (
	"github.com/jamesits/meituansqt/pkg/sqtapi"
	"net/http"
)

const (
	DefaultVersion         = "1.0"
	ApiEndpointProduction  = "https://bep-openapi.meituan.com/api/sqt/openapi"
	ApiEndpointDevelopment = "https://inf-openapi.apigw.test.meituan.com/api/sqt/openapi"
)

type SQT struct {
	sqtapi.Config
	Client *http.Client
}

func NewProduction(enterpriseId int64, accessKey string, secretKey string) *SQT {
	return &SQT{
		Config: sqtapi.Config{
			Version:      DefaultVersion,
			ApiEndpoint:  ApiEndpointProduction,
			EnterpriseID: enterpriseId,
			AccessKey:    accessKey,
			SecretKey:    secretKey,
		},
	}
}

func NewDevelopment(enterpriseId int64, accessKey string, secretKey string) *SQT {
	return &SQT{
		Config: sqtapi.Config{
			Version:      DefaultVersion,
			ApiEndpoint:  ApiEndpointDevelopment,
			EnterpriseID: enterpriseId,
			AccessKey:    accessKey,
			SecretKey:    secretKey,
		},
	}
}
