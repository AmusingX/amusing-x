package logout

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"amusingx.fit/amusingx/services/europa/session"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandlerLogOut(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	sid, ok := session.GetSessionID(r)
	if !ok && sid == "" {
		rest.SucceedJsonResponse(w, ganymedeservice.LogoutResponse{Logout: true})
		return
	}

	resp, err := ganymede.RPCClient.Client.LogOut(ctx, &ganymedeservice.LogoutRequest{SessionID: sid})
	if err != nil {
		logger.Errorf("log out failed: %s", err)
		rest.FailJsonResponse(w, aerror.Code.SUnexpectedErr, "logout failed")
		return
	}

	session.DeleteSession(w)

	rest.SucceedJsonResponse(w, resp)
}
