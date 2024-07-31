package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go-template/cmd/api/httputils"
	"go-template/internal/user"
	"strconv"
)

type User struct {
	service user.Service
}

// GetUser godoc
// @Summary Show a user
// @Description get user by ID
// @Tags user
// @Accept json
// @Produce json
// @Header 200 {object} user.Model
// @Failure 400 {object} httputils.BaseResponse
// @Failure 404 {object} httputils.BaseResponse
// @Failure 500 {object} httputils.BaseResponse
// @Router /user/{id} [get]
func (u *User) GetUser(ctx *fasthttp.RequestCtx) {
	idStr, ok := ctx.UserValue("id").(string)
	if idStr == "" || !ok {
		httputils.JSON(&ctx.Response, &httputils.BaseResponse{Msg: "id not provide in url path"}, fasthttp.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		httputils.JSON(&ctx.Response, &httputils.BaseResponse{Msg: "id must be a number"}, fasthttp.StatusBadRequest)
		return
	}

	resp, err := u.service.GetByID(ctx, id)
	if err != nil {
		switch {
		default:
			httputils.JSON(&ctx.Response, err, fasthttp.StatusInternalServerError)
		}
		return
	}

	httputils.JSON(&ctx.Response, resp, fasthttp.StatusOK)
}

// InsertUser godoc
// @Summary Insert a user
// @Description insert user
// @Tags user
// @Accept json
// @Produce json
// @Param user body user.Model true "user"
// @Header 200 {object} user.Model
// @Failure 400 {object} httputils.BaseResponse
// @Failure 500 {object} httputils.BaseResponse
// @Router /user [post]
func (u *User) InsertUser(ctx *fasthttp.RequestCtx) {
	var user user.Model
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		httputils.JSON(&ctx.Response, &httputils.BaseResponse{Msg: err.Error()}, fasthttp.StatusBadRequest)
		return
	}

	err := u.service.Insert(ctx, &user)
	if err != nil {
		switch {
		default:
			httputils.JSON(&ctx.Response, &httputils.BaseResponse{Msg: err.Error()}, fasthttp.StatusInternalServerError)
		}
		return
	}

	httputils.JSON(&ctx.Response, user, fasthttp.StatusOK)
}

func NewUser(service user.Service) *User {
	return &User{
		service: service,
	}
}
