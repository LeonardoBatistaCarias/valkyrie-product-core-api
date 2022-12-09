package grpc

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/proto/pb"
	uuid "github.com/satori/go.uuid"
)

type ReaderService struct {
	rc pb.ProductReaderServiceClient
}

func NewReaderService(rc pb.ProductReaderServiceClient) *ReaderService {
	return &ReaderService{rc: rc}
}

func (s *ReaderService) GetProductByID(productID uuid.UUID) (*product.Product, error) {
	cmd := pb.GetProductByIDReq{ProductID: productID.String()}

	res, err := s.rc.GetProductByID(context.Background(), &cmd)
	if err != nil {
		return nil, err
	}

	return product.ProductFromGrpcMessage(*res.GetProduct()), nil
}
