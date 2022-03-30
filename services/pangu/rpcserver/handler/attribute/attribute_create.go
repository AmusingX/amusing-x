package attribute

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/charonclient"
	"amusingx.fit/amusingx/services/comm"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
)

func HandlerAttributeCreate(ctx context.Context, req *panguservice.AttributeCreateRequest) *response.CommResponse {
	logger.Errorf("req: %s", logger.ToJson(req))
	resp, e := charonclient.Client.AttributeCreate(ctx, req)
	var err aerror.Error
	if e != nil {
		err = aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "unexpected error")
	}

	return comm.Response(resp, err)
}
