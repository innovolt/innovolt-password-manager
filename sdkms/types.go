package sdkms

import (
//  "encoding/json"
)

type Secret struct {
	Name    string   `json:"name"`
	GroupId string   `json:"group_id"`
	KeyOps  []string `json:"key_ops"`
	ObjType string   `json:"obj_type"`
	Value   string   `json:"value"`
}
