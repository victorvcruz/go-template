package httputils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type BaseResponse struct {
	Msg string `json:"msg"`
}

func JSON(response *fasthttp.Response, body any, status int) {
	b, _ := json.Marshal(body)
	response.Header.Set(fasthttp.HeaderContentType, "application/json")
	response.SetBody(b)
	response.SetStatusCode(status)
}
