package repo

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"poller-app/internal/domain"
	"time"
)

func GetURI(baseURI string, param domain.RequestParams, logger *log.Logger) (*domain.Response, error) {
	uriStr := createURIWithParams(baseURI, param)
	resp, err := http.Get(uriStr)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	//Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	// Unmarshal the response to response struct
	response := new(domain.Response)
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	if *response.Result == "error" {
		if response.ErrorType != nil {
			logger.Println(err)
			return nil, errors.New(*response.ErrorType)
		}
		logger.Println("unknow error in calling API")
		return nil, errors.New("unknow error in calling API")
	}
	// adding a timestamp for the response call being saved
	response.CreatedTime = time.Now()
	return response, nil
}

func createURIWithParams(baseURI string, param domain.RequestParams) string {
	if param.BaseCurrency != nil && param.ApiKey != nil {
		return baseURI + *param.ApiKey + "/latest/" + *param.BaseCurrency
	} else {
		return baseURI
	}
}
