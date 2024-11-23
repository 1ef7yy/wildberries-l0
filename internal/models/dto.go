package models

import "encoding/json"

type Schema struct {
	OrderUid string `json:"order_uid"`
	Data     json.RawMessage
}
