package grpc

import (
	"context"
	"fmt"
	"github.com/gaursagarMT/starter/pb/yelo"
	"github.com/gaursagarMT/starter/src/storage"
	"google.golang.org/grpc/metadata"
)

func GetYeloServer(iStorage storage.IStorage) YeloServer {
	return YeloServer{
		storage: iStorage,
	}
}

type YeloServer struct {
	storage storage.IStorage
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
