package public

import "time"

const (
	TIME_FORMAT        = "2006-01-02 15:04:05.999999999 -0700 MST"
	TIME_FORMAT_SECOND = "2006-01-02 15:04:05"
	TIME_FORMAT_ZONE   = "2006-01-02 15:04:05 -0700 MST"
)

// token
const (
	TOKEN_KEY      = "Authorization"
	TOKEN_ISSUER   = "Oliodynamical"
	TOKEN_SECRET   = "*&IU^%YT$#TR7890"
	TOKEN_SURVIVAL = 24 * 7 // Hour
)

func FormatTime() string {
	return time.Now().Format(TIME_FORMAT)
}

// 时间单位
const (
	TIME_DAY    = "d" // 日
	TIME_HOUR   = "h" // 时
	TIME_MINUTE = "m" // 分
	TIME_SECOND = "s" // 秒
)
