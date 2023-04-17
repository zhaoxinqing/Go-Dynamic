package timex

import "time"

// GetUTCFormatTime ... 获取格式化时间
func GetUTCFormatTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}
