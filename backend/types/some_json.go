package types

type SomeJsonQuery struct {
	UserID int `form:"user_id" binding:"required"`
}

type SomeHandlerResp struct {
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Married   bool    `json:"married"`
	RingRatio float64 `json:"ring_ration"`
}
