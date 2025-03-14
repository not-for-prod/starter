package starter

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// Service is interface u should implement to use starter
type Service interface {
	Name() string
	Swagger() []byte
	RegisterGRPC(s *grpc.Server)
	RegisterGateway(
		ctx context.Context,
		mux *runtime.ServeMux,
	)
}
