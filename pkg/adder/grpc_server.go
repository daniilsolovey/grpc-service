package adder

import (
	"context"

	"github.com/daniilsolovey/grpc-service/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
}

func (server *GRPCServer) Add(
	ctx context.Context, req *api.AddRequest,
) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

func NewAdderServer() *GRPCServer {
	return &GRPCServer{}
}
