package types

type SomeJsonQuery struct {
	UserID int `form:"user_id" binding:"required,gte=1,lte=1000"`
}

type SomeHandlerResp struct {
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Married   bool    `json:"married"`
	RingRatio float64 `json:"ring_ratio"`
}
