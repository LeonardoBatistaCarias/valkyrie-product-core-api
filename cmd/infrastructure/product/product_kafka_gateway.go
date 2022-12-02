package product

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/config"
	producer "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/kafka"
	protoProduct "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/proto"
	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
	"time"
)

type ProductKafkaGateway struct {
	producer producer.Producer
	cfg      *config.Config
}

func NewProductKafkaGateway(cfg *config.Config, producer producer.Producer) *ProductKafkaGateway {
	return &ProductKafkaGateway{
		cfg:      cfg,
		producer: producer,
	}
}

func (g *ProductKafkaGateway) Create(ctx context.Context, p product.Product) error {
	msg := newProductCreateKafkaMessage(p)

	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return g.producer.PublishMessage(ctx, kafka.Message{
		Topic: g.cfg.KafkaTopics.ProductCreate.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	})
}

func newProductCreateKafkaMessage(p product.Product) *protoProduct.Product {
	pbImages := make([]*protoProduct.ProductImage, len(p.ProductImages))

	for i, image := range p.ProductImages {
		pi := &protoProduct.ProductImage{
			Name:      image.Name,
			ProductID: image.ProductID.String(),
		}
		pbImages[i] = pi
	}

	return &protoProduct.Product{
		ProductID:     p.ProductID.String(),
		Name:          p.Name,
		Description:   p.Description,
		Brand:         int32(p.Brand),
		Price:         p.Price,
		Quantity:      p.Quantity,
		CategoryID:    p.CategoryID.String(),
		ProductImages: pbImages,
		Active:        p.Active,
		CreatedAt:     p.CreatedAt.String(),
		UpdatedAt:     "",
		DeletedAt:     "",
	}
}
