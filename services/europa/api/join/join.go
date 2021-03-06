package join

import (
	"amusingx.fit/amusingx/apistruct/europa"
	"amusingx.fit/amusingx/services/europa/app/joinapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandleJoin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	params, err := getAndValidParams(r)
	if err != nil {
		logger.Errorf("params is invalid: %s", err.Error())

		rest.FailJsonResponse(w, aerror.Code.CreateUserError, err.Message())
		return
	}

	u, err := joinapp.CreateUser(ctx, params)
	if err != nil {
		logger.Errorf("create user failed, params: %s, err: %s", logger.ToJson(params), err.Error())
		rest.FailJsonResponse(w, aerror.Code.CreateUserError, err.Message())
		return
	}

	resp := europa.JoinResponse{
		Nickname: u.Nickname,
		Phone:    u.Phone,
	}

	rest.SucceedJsonResponse(w, resp)
}

func getAndValidParams(r *http.Request) (*europa.JoinRequest, aerror.Error) {
	joinRequest := europa.JoinRequest{}

	err := httputil.DecodeJsonBody(r, &joinRequest)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.CParamsError, "decode request body failed")
	}

	xErr := joinRequest.Valid()
	if xErr != nil {
		return nil, xErr
	}

	return &joinRequest, nil
}
