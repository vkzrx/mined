package vm

import (
	"context"
	"net"

	"github.com/coreos/go-systemd/dbus"
	pb "github.com/vkzrx/mined/vm/proto"
	"google.golang.org/grpc"
)

type MinecraftService struct {
	pb.UnsafeMinecraftServiceServer
	dbus *dbus.Conn
}

func NewMinecraftService(ctx context.Context) (*MinecraftService, error) {
	dbus, err := dbus.NewWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return &MinecraftService{dbus: dbus}, nil
}

func (s *MinecraftService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	_, err := s.dbus.StartUnitContext(ctx, "minecraft.service", "replace", nil)
	if err != nil {
		return nil, err
	}
	return &pb.StartResponse{Message: "success"}, nil
}

type MinecraftServiceServer struct {
	port    string
	server  *grpc.Server
	service *MinecraftService
}

func NewMinecraftServiceServer(port string) (*MinecraftServiceServer, error) {
	ctx := context.Background()
	service, err := NewMinecraftService(ctx)
	if err != nil {
		return nil, err
	}
	server := &MinecraftServiceServer{
		server:  grpc.NewServer(),
		port:    port,
		service: service,
	}
	return server, nil
}

func (s *MinecraftServiceServer) Launch() error {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}

	pb.RegisterMinecraftServiceServer(s.server, s.service)

	if err := s.server.Serve(listener); err != nil {
		return err
	}

	return nil
}
