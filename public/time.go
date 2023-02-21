package public

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
	return time.Format(TIME_FORMAT_SECOND)
}

// StringToTime ...
func StringToTime(timeStr string) time.Time {
	time, _ := time.ParseInLocation(TIME_FORMAT_SECOND, timeStr, time.UTC)
	return time
}

// StringToUnix ...
func StringToUnix(timeStr string) int64 {
	time, _ := time.ParseInLocation(TIME_FORMAT_SECOND, timeStr, time.UTC)
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

type Time time.Time

// MarshalJSON ...
func (t Time) MarshalJSON() ([]byte, error) {

	b := make([]byte, 0, len(TIME_FORMAT_SECOND)+2)
	b = append(b, '"')

	b = time.Time(t).AppendFormat(b, TIME_FORMAT_SECOND)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON ...
func (t *Time) UnmarshalJSON(data []byte) (err error) {

	now, err := time.ParseInLocation(`"`+TIME_FORMAT_SECOND+`"`, string(data), time.Local)
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

	return time.Time(t).Format(TIME_FORMAT_SECOND)

}

// StringNotNull ...
func (t Time) StringNotNull() string {

	if time.Time(t).IsZero() {
		return ""
	}

	return time.Time(t).Format(TIME_FORMAT_SECOND)
}

// StringUTC ...
func (t Time) StringUTC() string {

	return time.Time(t).UTC().Format(TIME_FORMAT_SECOND)

}

func Timex() {
	fmt.Println(time.Now().UTC().Truncate(24 * time.Hour).Format(TIME_FORMAT_SECOND))
	fmt.Println(time.Now().UTC().Truncate(24 * time.Hour).Add(24 * time.Hour).Format(TIME_FORMAT_SECOND))
}
