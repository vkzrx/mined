package vm

import (
	pb "github.com/vkzrx/mined/vm/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MinecraftClient struct {
	Service    pb.MinecraftServiceClient
	connection *grpc.ClientConn
}

// Creates a new gRPC MinecraftClient.
// It establishes a connection to the provided address.
//
// Parameters:
//
//	addr (string): The address to connect to e.g. "localhost:5001"
func NewMinecraftClient(addr string) (*MinecraftClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	mc := &MinecraftClient{
		Service:    pb.NewMinecraftServiceClient(conn),
		connection: conn,
	}
	return mc, err
}

func (c *MinecraftClient) Close() error {
	return c.connection.Close()
}
