---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meituansqt_staff Resource - meituansqt"
subcategory: ""
description: |-
  Staff
---

# meituansqt_staff (Resource)

Staff

## Example Usage

```terraform
resource "meituansqt_staff" "user1" {
  name  = "姓名"
  phone = "13800000000"
  email = "user@example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) 姓名

### Optional

- `email` (String) 邮箱
- `ent_staff_num` (String) 工号
- `phone` (String) 手机号

### Read-Only

- `id` (Number) 用户 ID
