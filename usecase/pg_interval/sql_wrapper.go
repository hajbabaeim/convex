package pg_interval

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (d PGInterval) Value() (driver.Value, error) {
	dur := time.Duration(d)
	if dur == 0 {
		return "0", nil
	}
	hours := dur / time.Hour
	dur %= time.Hour
	minutes := dur / time.Minute
	dur %= time.Minute
	seconds := dur / time.Second

	parts := make([]string, 0)
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d hours", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d minutes", minutes))
	}
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%d seconds", seconds))
	}
	if len(parts) == 0 {
		return "0 seconds", nil
	}
	return strings.Join(parts, " "), nil
}

func (d *PGInterval) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid type for PGInterval: %T", value)
	}

	// HH:MM:SS
	if matched, _ := regexp.MatchString(`^\d+:\d{2}:\d{2}$`, str); matched {
		parts := strings.Split(str, ":")
		h, _ := strconv.Atoi(parts[0])
		m, _ := strconv.Atoi(parts[1])
		s, _ := strconv.Atoi(parts[2])
		total := time.Duration(h)*time.Hour +
			time.Duration(m)*time.Minute +
			time.Duration(s)*time.Second
		*d = PGInterval(total)
		return nil
	}

	// unit-based: "2 hours 30 minutes"
	fields := strings.Fields(str)
	var total time.Duration
	for i := 0; i < len(fields); i += 2 {
		if i+1 >= len(fields) {
			return fmt.Errorf("invalid interval format: %s", str)
		}
		qty, unit := fields[i], fields[i+1]
		val, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid quantity in interval: %s", qty)
		}
		switch {
		case strings.HasPrefix(unit, "hour"):
			total += time.Duration(val) * time.Hour
		case strings.HasPrefix(unit, "minute"):
			total += time.Duration(val) * time.Minute
		case strings.HasPrefix(unit, "second"):
			total += time.Duration(val) * time.Second
		default:
			return fmt.Errorf("unrecognized interval unit: %s", unit)
		}
	}
	*d = PGInterval(total)
	return nil
}

func (PGInterval) GormDataType() string {
	return "interval"
}
