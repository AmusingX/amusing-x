package ganymede

import (
	"amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"amusingx.fit/amusingx/rpcclient"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/grpc"
)

var (
	Conn   *grpc.ClientConn
	Client ganymedeservice.GanymedeServiceClient
)

func InitClient(addr string) aerror.Error {
	if Conn != nil {
		return nil
	}

	var err aerror.Error
	if len(addr) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.SUnexpectedErr, "addr is blank")
	}

	Conn, err = rpcclient.InitRPCClientConn(addr)
	if err != nil {
		return err
	}

	Client = ganymedeservice.NewGanymedeServiceClient(Conn)

	return nil
}

func CloseClient(addr string) {
	if Conn != nil {
		Client = nil
		Conn.Close()
	}
}
