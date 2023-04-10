package logic

import (
	"backend-go/public"
)

type SomeJsonQuery struct {
	UserID int `form:"user_id" binding:"required,gte=1,lte=1000"`
}
type SomeHandlerResp struct {
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Married   bool    `json:"married"`
	RingRatio float64 `json:"ring_ratio"`
}

func SomeJson(*SomeJsonQuery) (*SomeHandlerResp, *public.ProtocolError) {
	return &SomeHandlerResp{Name: "Dynamic", Age: 1, Married: false, RingRatio: float64(-2) / float64(15)}, nil
}
