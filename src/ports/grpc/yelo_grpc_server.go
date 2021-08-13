package grpc

import (
	"context"
	"fmt"
	"github.com/gaursagarMT/starter/pb/yelo"
	"github.com/gaursagarMT/starter/src/env"
	"google.golang.org/grpc/metadata"
)

func GetYeloServer(appEnv *env.Environment) YeloServer {
	return YeloServer{}
}

type YeloServer struct {
	appEnv env.Environment
}

func (YeloServer) AddNewPlace(ctx context.Context, request *yelo.AddNewPlaceRequest) (*yelo.AddNewPlaceResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md, ok)
	md, ok = metadata.FromOutgoingContext(ctx)
	fmt.Println(md, ok)
	return &yelo.AddNewPlaceResponse{
		PlaceId: 0,
	}, nil
}

func (YeloServer) ListPlaces(ctx context.Context, request *yelo.ListPlacesRequest) (*yelo.ListPlacesResponse, error) {
	return &yelo.ListPlacesResponse{
		Places: nil,
		Page:   nil,
	}, nil
}
