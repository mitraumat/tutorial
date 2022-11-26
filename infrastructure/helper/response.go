package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Option func(f *Response)

func WithData(data interface{}) Option {
	return func(r *Response) {
		r.Data = data
	}
}

func NewResponse(code int, status bool, message string, options ...Option) Response {
	r := Response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    nil,
	}

	for _, option := range options {
		option(&r)
	}
	return r
}

func ResponseJson(context *fiber.Ctx, response Response) error {
	return context.Status(response.Code).JSON(response)
}
