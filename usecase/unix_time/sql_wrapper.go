package unix_time

import (
	"database/sql/driver"
	"fmt"
	"time"
)

func (u *UnixTime) Scan(value interface{}) error {
	if value == nil {
		u.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		u.Time = v.UTC()
	case []byte:
		return u.UnmarshalJSON(v)
	case string:
		return u.UnmarshalJSON([]byte(v))
	default:
		return fmt.Errorf("unsupported type for UnixTime: %T", value)
	}
	return nil
}

func (u *UnixTime) Value() (driver.Value, error) {
	if u.Time.IsZero() {
		return nil, nil
	}
	return u.Time.UTC(), nil
}

func (UnixTime) GormDataType() string {
	return "timestamptz"
}
