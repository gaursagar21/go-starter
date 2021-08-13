package grpc

import (
	"context"
	"fmt"
	"github.com/gaursagarMT/starter/pb/yelo"
	"github.com/gaursagarMT/starter/src/config"
	"github.com/gaursagarMT/starter/src/storage"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

func StartGRPCServer(conf config.Config) error {

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				ServerRequestInterceptor(conf.ApplicationName),
				ServerHeaderInterceptor(),
			),
		),
	)
	ctx := context.Background()

	//Storage, Dispatch, etc.
	mysqlStorageImpl, _ := storage.GetMySQLStorage(ctx, conf.DatabaseConfig.MySQLConfig)

	yeloServerImpl := GetYeloServer(mysqlStorageImpl)
	yelo.RegisterYeloServiceServer(grpcServer, yeloServerImpl)

	hostPort := fmt.Sprintf("0.0.0.0:%s", "5000")
	lis, _ := net.Listen("tcp", hostPort)

	_ = grpcServer.Serve(lis)
	return nil
}
