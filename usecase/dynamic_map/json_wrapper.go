package dynamic_map

import (
	"encoding/json"
)

func (d DynamicMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(d)
}

func (d *DynamicMap) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	var value map[string]interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if len(value) == 0 {
		return nil
	}
	*d = value
	return nil
}
