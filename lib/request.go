package lib

import (
	"github.com/gin-gonic/gin"
)

const (
	TOKEN_KEY = "Authorization" // Token key
)

// PageByQuery ...
type PageQueryReq struct {
	Page     uint64 `form:"page" binding:"min=1"`
	PageSize uint64 `form:"page_size" binding:"min=1"`
}

// GetPageIndex 获取页码
func GetPage(c *gin.Context) uint64 {
	pageStr := c.Query("page")
	return StrToUint64(pageStr)
}

// GetPageSize 获取每页记录数
func GetPageSize(c *gin.Context) uint64 {
	pageSizeStr := c.Query("page_size")
	return StrToUint64(pageSizeStr)
}

// GetPagesByQuery ...
func GetPagesByQuery(c *gin.Context) (*PageQueryReq, error) {
	var pages PageQueryReq
	err := c.ShouldBind(&pages)
	return &pages, err
}

// GetToken 从header获取token
func GetToken(c *gin.Context) string {
	return c.GetHeader(TOKEN_KEY)
}

// GetID
func GetIDFromQuery(c *gin.Context) uint64 {
	return StrToUint64(c.Query("id"))
}

// GetQueryParam ...
func GetQueryParam(c *gin.Context, queryKey string) string {
	return c.Query(queryKey)
}
