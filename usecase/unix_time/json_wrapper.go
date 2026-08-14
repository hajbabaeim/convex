package unix_time

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func (u *UnixTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		u.Time = time.Time{}
		return nil
	}

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case float64:
		u.Time = time.Unix(int64(val), 0).UTC()
		return nil
	case string:
		if val == "" {
			u.Time = time.Time{}
			return nil
		}
		// try numeric string
		if sec, err := strconv.ParseInt(val, 10, 64); err == nil {
			u.Time = time.Unix(sec, 0).UTC()
			return nil
		}
		// try RFC3339Nano
		if t, err := time.Parse(time.RFC3339Nano, val); err == nil {
			u.Time = t.UTC()
			return nil
		}
		return errors.New("invalid time format, use Unix timestamp or RFC3339")
	default:
		return errors.New("unsupported time format")
	}
}

func (u *UnixTime) MarshalJSON() ([]byte, error) {
	if u.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%d", u.Unix())), nil
}
