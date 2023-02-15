package lib

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// TimeToString ...
func TimeToString(time time.Time) string {
	if time.IsZero() {
		return ""
	}
	return time.Format(FormatTime)
}

// StringToTime ...
func StringToTime(timeStr string) time.Time {
	time, _ := time.ParseInLocation(FormatTime, timeStr, time.UTC)
	return time
}

// StringToUnix ...
func StringToUnix(timeStr string) int64 {
	time, _ := time.ParseInLocation(FormatTime, timeStr, time.UTC)
	return time.Unix()
}

// TimestampUTC ... 获取当前UTC时间戳
func GetUTCTimestamp() int64 {
	return time.Now().UTC().Unix()
}

// GetUTCFormatTime ... 获取格式化时间
func GetUTCFormatTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

var FormatTime = "2006-01-02 15:04:05"

type Time time.Time

// MarshalJSON ...
func (t Time) MarshalJSON() ([]byte, error) {

	b := make([]byte, 0, len(FormatTime)+2)
	b = append(b, '"')

	b = time.Time(t).AppendFormat(b, FormatTime)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON ...
func (t *Time) UnmarshalJSON(data []byte) (err error) {

	now, err := time.ParseInLocation(`"`+FormatTime+`"`, string(data), time.Local)
	*t = Time(now)

	return
}

// Scan ...
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)

	if ok {
		*t = Time(value)
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Value ...
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time

	var ti = time.Time(t)

	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return ti, nil
}

// UnixTemp ...
func (t Time) UnixTemp() int64 {
	var ti = time.Time(t)
	return ti.Unix()
}

// FormatMD ...
func (t Time) FormatMD() string {

	if time.Time(t).IsZero() {
		return ""
	}

	return time.Time(t).Format("01/02")
}

// String ...
func (t Time) String() string {

	return time.Time(t).Format(FormatTime)

}

// StringNotNull ...
func (t Time) StringNotNull() string {

	if time.Time(t).IsZero() {
		return ""
	}

	return time.Time(t).Format(FormatTime)
}

// StringUTC ...
func (t Time) StringUTC() string {

	return time.Time(t).UTC().Format(FormatTime)

}

func Timex() {
	fmt.Println(time.Now().UTC().Truncate(24 * time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().UTC().Truncate(24 * time.Hour).Add(24 * time.Hour).Format("2006-01-02 15:04:05"))
}
