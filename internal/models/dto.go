package models

import (
	"encoding/json"
)

type Order struct {
	OrderUid string          `json:"order_uid"`
	Data     json.RawMessage `json:"data"`
}

func (o *Order) MarshalJSON() ([]byte, error) {
	var dataMap map[string]interface{}

	if err := json.Unmarshal(o.Data, &dataMap); err != nil {
		return nil, err
	}

	dataMap["order_uid"] = o.OrderUid

	return json.Marshal(dataMap)
}
