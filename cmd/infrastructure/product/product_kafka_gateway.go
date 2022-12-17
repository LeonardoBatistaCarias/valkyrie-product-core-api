package product

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	kafkaMessage "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/grpc/pb/kafka"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/grpc/pb/model"
	producer "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/kafka"
	"github.com/golang/protobuf/proto"
	uuid "github.com/satori/go.uuid"
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

func (g *ProductKafkaGateway) CreateProduct(ctx context.Context, p product.Product) error {
	msg := newProductKafkaMessage(p)

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

func (g *ProductKafkaGateway) DeleteProductByID(ctx context.Context, productID uuid.UUID) error {
	msg := &kafkaMessage.DeactivateProductByIDKafkaMessage{ProductID: productID.String()}

	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return g.producer.PublishMessage(ctx, kafka.Message{
		Topic: g.cfg.KafkaTopics.ProductDelete.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	})
}

func (g *ProductKafkaGateway) DeactivateProductByID(ctx context.Context, productID uuid.UUID) error {
	msg := &kafkaMessage.DeactivateProductByIDKafkaMessage{ProductID: productID.String()}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return g.producer.PublishMessage(ctx, kafka.Message{
		Topic: g.cfg.KafkaTopics.ProductDeactivate.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	})
}

func (g *ProductKafkaGateway) UpdateProductByID(ctx context.Context, p product.Product) error {
	msg := newProductKafkaMessage(p)

	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return g.producer.PublishMessage(ctx, kafka.Message{
		Topic: g.cfg.KafkaTopics.ProductUpdate.TopicName,
		Value: msgBytes,
		Time:  time.Now().UTC(),
	})
}

func newProductKafkaMessage(p product.Product) *model.Product {
	//pbImages := make([]*protoProduct.ProductImage, len(p.ProductImages))

	//for i, image := range p.ProductImages {
	//	pi := &protoProduct.ProductImage{
	//		Name:      image.Name,
	//		ProductID: image.ProductID.String(),
	//	}
	//	pbImages[i] = pi
	//}

	return &model.Product{
		ProductID:     p.ProductID.String(),
		Name:          p.Name,
		Description:   p.Description,
		Brand:         int32(p.Brand),
		Price:         p.Price,
		Quantity:      p.Quantity,
		CategoryID:    p.CategoryID.String(),
		ProductImages: nil,
		Active:        p.Active,
		CreatedAt:     p.CreatedAt.String(),
		UpdatedAt:     "",
		DeletedAt:     "",
	}
}
