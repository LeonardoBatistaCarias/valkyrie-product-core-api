package grpc

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/proto/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/constants"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func NewReaderServiceClient(ctx context.Context, cfg *config.Config) (pb.ProductReaderServiceClient, error) {
	readerServiceConn, err := newReaderServiceConn(ctx, cfg)
	if err != nil {
		return nil, err
	}
	//defer readerServiceConn.Close()

	return pb.NewProductReaderServiceClient(readerServiceConn), nil
}

func newReaderServiceConn(ctx context.Context, cfg *config.Config) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(constants.BACKOFF_LINEAR)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
		grpc_retry.WithMax(constants.BACKOFF_RETRIES),
	}

	readerServiceConn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("localhost:%s", cfg.Grpc.ReaderServicePort),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.DialContext")
	}

	return readerServiceConn, nil
}
