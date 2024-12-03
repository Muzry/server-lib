package server

const (
	SuccessCode = 0
)

type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type List struct {
	Total int64 `json:"total"`
	Items any   `json:"items"`
}

func SuccessResponse(data any) Response {
	return Response{
		Code: SuccessCode,
		Data: data,
	}
}

func ListResponse(total int64, data any) Response {
	return Response{
		Code: SuccessCode,
		Data: List{
			Total: total,
			Items: data,
		},
	}
}
