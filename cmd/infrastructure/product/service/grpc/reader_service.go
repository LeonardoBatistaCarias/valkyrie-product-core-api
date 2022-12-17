package grpc

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	reader_service2 "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/grpc/pb/reader_service"
	uuid "github.com/satori/go.uuid"
)

type ReaderService struct {
	rc reader_service2.ProductReaderServiceClient
}

func NewReaderService(rc reader_service2.ProductReaderServiceClient) *ReaderService {
	return &ReaderService{rc: rc}
}

func (s *ReaderService) GetProductByID(productID uuid.UUID) (*product.Product, error) {
	cmd := &reader_service2.GetProductByIDReq{ProductID: productID.String()}

	res, err := s.rc.GetProductByID(context.Background(), cmd)
	if err != nil {
		return nil, err
	}

	return product.ProductFromGrpcMessage(*res.GetProduct()), nil
}
