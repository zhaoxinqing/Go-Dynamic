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
func HttpResult(c *gin.Context, data interface{}, err *ErrResponseMsg) {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code,
			"message": err.Message,
			"data":    data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "ok",
			"data":    data,
		})
	}
	return

}
