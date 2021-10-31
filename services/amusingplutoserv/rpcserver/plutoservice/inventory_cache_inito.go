package plutoservice

import (
	"amusingx.fit/amusingx/protos/amusingplutoservice/plutoservice"
	"amusingx.fit/amusingx/services/amusingplutoserv/rpcserver/promotionalbook"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func (s *PlutoService) InventoryCacheInit(ctx context.Context, req *plutoservice.InventoryCacheInitRequest) (*plutoservice.InventoryCacheInitResponse, error) {
	var promotionalObj PromotionalObj

	if req == nil {
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "req is nil")
	}

	switch req.GetObj() {
	case plutoservice.InventoryCacheInitRequest_Book:
		promotionalObj = promotionalbook.NewPromotionalBook(-1)
	case plutoservice.InventoryCacheInitRequest_Mac:
		return &plutoservice.InventoryCacheInitResponse{Succeed: false, Message: "unsupported type"}, nil
	default:
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "unexpected obj type")
	}

	err := promotionalObj.InventoryCacheInit(ctx)
	if err != nil {
		return nil, err
	}

	return &plutoservice.InventoryCacheInitResponse{Succeed: true}, nil
}
