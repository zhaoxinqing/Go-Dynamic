package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PageDataRes ...
type PageResponseData struct {
	Meta  PageMata    `json:"meta"`
	Items interface{} `json:"items"`
}

// Pages ...
type PageMata struct {
	PageNo   uint64 `json:"page_no"`
	PageSize uint64 `json:"page_size"`
	Total    uint64 `json:"total"`
}

type ErrResponseMsg struct {
	Code    int64       `json:"code"`
	Message interface{} `json:"message"`
}

// HttpResult ...
func HttpResult(c *gin.Context, data interface{}, err *ProtocolError) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": err.ErrorCode, "message": err.ErrorString, "data": data})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 0, "message": "ok", "data": data})
	}
}
