package server

import (
	"net"

	"github.com/ProtonFundationGroup/seventies/v2/log"
	"github.com/ProtonFundationGroup/seventies/v2/server/pb-go/api"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func New() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()
	// Register your gRPC service implementation here
	// example: pb.RegisterYourServiceServer(s.grpcServer, &yourService{})
	serviceMgr := newServiceManager()
	api.RegisterAPIServiceServer(s.grpcServer, serviceMgr)

	log.PrintGreen("Starting gRPC server on port 8080")
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() error {
	log.PrintGreen("server has been stopped.")
	s.grpcServer.Stop()
	return nil
}
