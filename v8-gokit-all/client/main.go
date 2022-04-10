package client

/*
import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sirupsen/logrus"
	"go-kit-demo/v7-gokit-grpc/pb"
	v7 "go-kit-demo/v8-gokit-all/book"
	"go-kit-demo/v8-gokit-all/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		println("grpc.Dial err:", err)
		return
	}
	defer conn.Close()
	svr := NewGRPCClient(conn, logrus.Logger{})
	ack, err := svr.Login(context.Background(), &pb.Login{
		Account:  "hwholiday",
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ack)
}

func NewGRPCClient(conn *grpc.ClientConn, log logrus.Logger) v7.BookServiceClient {
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
			//UUID := uuid.NewV5(uuid.NewV4(), "req_uuid").String()
			//log.Debug("给请求添加uuid", zap.Any("UUID", UUID))
			//md.Set(v5_service.ContextReqUUid, UUID)
			ctx = metadata.NewOutgoingContext(context.Background(), *md)
			return ctx
		}),
	}
	var bookEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpctransport.NewClient(
			conn,
			"pb.User",
			"RpcUserLogin",
			RequestLogin,
			ResponseLogin,
			pb.LoginAck{},
			options...).Endpoint()
	}
	return service.Endpoints{
		LoginEndpoint: loginEndpoint,
	}

}
*/
