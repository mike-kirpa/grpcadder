package adder

import (
	"context"
	"grpcadder/pkg/api"
)

type GRPCServer struct{}

func (g *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponce, error) {
	return &api.AddResponce{Result: req.GetX() + req.GetY()}, nil
}
