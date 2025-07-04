package response

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InvalidArgumentError(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}

func InternalError(msg string) error {
	return status.Error(codes.Internal, msg)
}
