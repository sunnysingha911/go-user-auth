package listeners

import (
	"log"
	"net"

	"github.com/sunnysingha911/user-service/gen/user-service/userpb"
	"github.com/sunnysingha911/user-service/grpc/handlers"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewUserGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {

	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatal("Failed to run: %v", err)
	}

	grpcServer := grpc.NewServer()

	loginService := &handlers.AuthServiceServer{}
	userpb.RegisterUserServiceServer(grpcServer, loginService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
