package repo

import (
	"poller-app/internal/base"
	"poller-app/internal/domain"
	"poller-app/internal/service/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateURIWithParams(t *testing.T) {
	apiKey, baseCurrency, baseDomain := "test_api_key", "INR", "www.test.com/"
	t.Run("success: get correct URI name", func(t *testing.T) {
		params := domain.RequestParams{BaseCurrency: &baseCurrency, ApiKey: &apiKey}
		uri := createURIWithParams(baseDomain, params)
		assert.Equal(t, "www.test.com/test_api_key/latest/INR", uri)
	})
	t.Run("failure: returns only base domain", func(t *testing.T) {
		params := domain.RequestParams{}
		uri := createURIWithParams(baseDomain, params)
		assert.NotEqual(t, "www.test.com/test_api_key/latest/INR", uri)
	})
}

func TestGetURI(t *testing.T) {
	logger := log.NewLogger()
	t.Run("success: api returns success", func(t *testing.T) {
		baseCurrency, apiKey := "INR", "7ab0e18899eb2e26880830bf"
		params := domain.RequestParams{BaseCurrency: &baseCurrency, ApiKey: &apiKey}
		res, err := GetURI(base.ApiURI, params, logger)
		assert.Nil(t, err)
		assert.Equal(t, "success", *res.Result)
	})
	t.Run("failure: wrong API key", func(t *testing.T) {
		baseCurrency, apiKey := "INR", "test_api_key"
		params := domain.RequestParams{BaseCurrency: &baseCurrency, ApiKey: &apiKey}
		_, err := GetURI(base.ApiURI, params, logger)
		// assert
		assert.NotNil(t, err)
	})
	t.Run("failure: wrong type of base currency", func(t *testing.T) {
		baseCurrency, apiKey := "asdas", "7ab0e18899eb2e26880830bf"
		params := domain.RequestParams{BaseCurrency: &baseCurrency, ApiKey: &apiKey}
		_, err := GetURI(base.ApiURI, params, logger)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "unsupported-code")
	})
}
