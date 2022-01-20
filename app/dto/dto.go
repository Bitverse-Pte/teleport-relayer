package dto

const (
	Success    int = 0
	BadRequest int = -1 //bind error
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ReqUpgradeHeight struct {
	Height uint64 `json:"height" form:"height"`
}
