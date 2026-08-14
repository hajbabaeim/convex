package dynamic_map

import (
	"database/sql/driver"
	"encoding/json"
)

func (d DynamicMap) Value() (driver.Value, error) {
	if d == nil || len(d) == 0 {
		return []byte(""), nil
	}
	return json.Marshal(d)
}

func (d *DynamicMap) Scan(value any) error {
	if value == nil {
		return nil
	}
	var data map[string]interface{}
	err := json.Unmarshal(value.([]byte), &data)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	*d = data
	return nil
}

func (DynamicMap) GormDataType() string {
	return "jsonb"
}
