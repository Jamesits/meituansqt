# meituansqt

[美团企业版（原商企通）API](https://h5.dianping.com/app/bep-docs/sky-doc/) 纯 Golang 实现。

## 功能 

- `pkg/sqtobfuscate`：美团企业版 API 参数 ~~签名~~ ~~加密~~ 混淆算法（[官方文档](https://h5.dianping.com/app/bep-docs/sky-doc/api.html#_1-4-%E7%AD%BE%E5%90%8D%E6%96%B9%E6%B3%95)）
- `pkg/sqtapi`：基于泛型的通用 API 包装函数，自动处理传入传出类型转换和错误码解析成 Error
- `pkg/sqt`：部分 API 实现，包括参数结构体定义

## 示例

```go
package main

import (
	"context"
	"fmt"
	"github.com/jamesits/meituansqt/pkg/sqt"
)

const (
	entId     = 114514
	accessKey = "XXXXXXXXXXXX-TK"
	secretKey = "XXXXXXXXXXXXXXXXXXXXXXX"
)

func main() {
	s := sqt.NewProduction(entId, accessKey, secretKey)
	
	resp, _, err := s.AllAdminDivisionQuery(context.Background())
	if err != nil {
		panic(err)
    }
	fmt.Printf("%v", resp)
	
	resp, _, err = s.StaffBatchQuery(context.Background(), &StaffBatchQueryRequest{
		StaffIdType:      0,
		StaffIdentifiers: []string{},
	})
	if err != nil {
		panic(err)
    }
	fmt.Printf("%v", resp)
}
```
