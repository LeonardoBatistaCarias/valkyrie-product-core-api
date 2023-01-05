package rest

import (
	"encoding/json"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"io/ioutil"
	"net/http"
)

type RestClient interface {
	Get(url string, returnType any) error
}

type restClient struct {
	log logger.Logger
}

func NewRestClient(log logger.Logger) *restClient {
	return &restClient{log: log}
}
func (r *restClient) Get(url string, responseType any) error {
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &responseType)
	if err != nil {
		r.log.WarnMsg("Error in json.Unmarshal", err)
		return err
	}

	return nil
}
