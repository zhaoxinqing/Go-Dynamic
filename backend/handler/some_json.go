package handler

import (
	"demo/backend/logic"
	"demo/backend/types"
	"demo/public"

	"github.com/gin-gonic/gin"
)

func SomeJson(c *gin.Context) {
	query := types.SomeJsonQuery{}

	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	if errA := c.ShouldBindQuery(&query); errA != nil {
		public.HttpResult(c, errA.Error(), public.ErrBadQuery)
		return
	}

	v, err := logic.SomeJson(&query)
	public.HttpResult(c, v, err)
}
