package sqtapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type BaseRequest struct {
	Method       string `json:"method"`
	Timestamp    int64  `json:"ts"`
	EnterpriseID int64  `json:"entId"`
}

type Request[T any] struct {
	BaseRequest
	Data *T
}

func (r *Request[T]) MarshalJSON() ([]byte, error) {
	if r.Data == nil {
		return json.Marshal(r.BaseRequest)
	}

	b, err := json.Marshal(r.BaseRequest)
	if err != nil {
		return nil, fmt.Errorf("sqt: %w", err)
	}

	d, err := json.Marshal(r.Data)
	if err != nil {
		return nil, fmt.Errorf("sqt: %w", err)
	}

	// hack to flatten the struct
	return []byte(fmt.Sprintf("{%s,%s}", bytes.Trim(b, "{}"), bytes.Trim(d, "{}"))), nil
}
