package category

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/category"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/category/model"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/rest"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
)

type CategoryRestGateway struct {
	cfg *config.Config
	log logger.Logger
	rc  rest.RestClient
}

func NewCategoryRestGateway(cfg *config.Config, log logger.Logger, rc rest.RestClient) *CategoryRestGateway {
	return &CategoryRestGateway{
		cfg: cfg,
		log: log,
		rc:  rc,
	}
}

func (g *CategoryRestGateway) FindCategoryByID(ctx context.Context, categoryID string) (*category.Category, error) {
	url := fmt.Sprintf("%s/%s", g.cfg.Rest.CategoryServicePath, categoryID)

	var response model.CategoryResponse
	if err := g.rc.Get(url, response); err != nil {
		g.log.WarnMsg("There was an error in rest call to Category API", err)
		return nil, err
	}

	if response.ID == "" {
		return nil, fmt.Errorf("the category with ID: %s doesn't exist", categoryID)
	}

	return response.ToModel(), nil
}
