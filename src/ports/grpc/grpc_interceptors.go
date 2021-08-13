package grpc

import (
	"context"
	"fmt"
	"github.com/gaursagarMT/starter/src/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func GenerateRequestId(appName string, namespace string) string {
	return fmt.Sprintf("%s-%s-%s-%d", namespace, appName, helper.GetHostName(), time.Now().Unix())
}

func ServerRequestInterceptor(appName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		calleeRequestId := GenerateRequestId(appName, "governance")

		// For external use
		ctx = metadata.AppendToOutgoingContext(ctx, "RequestId", calleeRequestId)

		// For internal use
		ctx = context.WithValue(ctx, "RequestId", calleeRequestId)

		return handler(ctx, req)
	}
}

var requiredMetaFields = []string{
	"authorizer",
}

func ServerHeaderInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		// check if required fields are present in request metadata and set those in context
		metaData, _ := metadata.FromIncomingContext(ctx)

		for _, field := range requiredMetaFields {
			fieldVal := metaData.Get(field)
			if len(fieldVal) == 0 || helper.IsEmpty(fieldVal[0]) {
				return nil, status.Errorf(codes.InvalidArgument, "Required field %s is missing in request metadata", field)
			}
			ctx = context.WithValue(ctx, field, fieldVal[0])
		}

		// return handler
		return handler(ctx, req)
	}
}
