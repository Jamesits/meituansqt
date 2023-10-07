# meituansqt

[美团企业版（原商企通）API](https://h5.dianping.com/app/bep-docs/sky-doc/) 纯 Golang 实现。

## 功能

- `pkg/sqtobfuscate`：美团企业版 API 参数 ~~签名~~ ~~加密~~
  混淆算法（[官方文档](https://h5.dianping.com/app/bep-docs/sky-doc/api.html#_1-4-%E7%AD%BE%E5%90%8D%E6%96%B9%E6%B3%95)）
- `pkg/sqtapi`：基于泛型的通用 API 包装函数，自动处理传入传出类型转换和错误码解析成 Error
- `pkg/sqt`：部分 API 实现，包括参数结构体定义
- `cmd/terraform-provider-meituansqt`：一个 Terraform provider

## 示例

### 作为 Terraform Provider

本地编译最新的开发版本：
```shell
goreleaser build --clean --snapshot
```

配置 Terraform 加载本地编译的 provider：
```shell
cat > .terraformrc <<EOF
provider_installation {
    dev_overrides {
        "registry.terraform.io/jamesits/meituansqt" = "path/to/meituansqt/dist/terraform-provider-meituansqt_windows_amd64_v1"
    }
    direct {}
}
EOF

export TF_CLI_CONFIG_FILE="./.terraformrc"
```

最简单的 Terraform 例子：
```hcl2
terraform {
    required_providers {
      meituansqt = {
        source = "registry.terraform.io/jamesits/meituansqt"
      }
    }
}

provider "meituansqt" {
    ent_id = 114514
    access_key = "XXXXXXXXXXXXX-TK"
    secret_key = ""
}

resource "meituansqt_staff" "huanjie_zhu" {
    name = "姓名"
    phone = "13800000000"
    email = "user@example.com"
}
```

运行：
```shell
terraform apply
```

### 作为 Golang 库

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

	resp, _, err = s.StaffBatchQuery(context.Background(), &sqt.StaffBatchQueryRequest{
		StaffIdType:      0,
		StaffIdentifiers: []string{},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", resp)
}
```
