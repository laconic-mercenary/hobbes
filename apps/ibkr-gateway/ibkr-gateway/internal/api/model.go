package api

import (
	"encoding/json"

	"github.com/go-playground/validator"
	log "github.com/rs/zerolog/log"
)

type OrderRequest struct {
	Symbol    string  `json:"symbol" validate:"required,min=1,max=5"`
	BuyOrSell string  `json:"buyOrSell" validate:"required,oneof=BUY SELL"`
	Quantity  float32 `json:"quantity" validate:"required,gte=1.0,lte=2000.0"`
}

func ParseOrderRequest(data []byte) (req *OrderRequest, err error) {
	log.Trace().Msg("api.ParseOrderRequest")
	req = &OrderRequest{}
	err = json.Unmarshal(data, req)
	return
}

func (this *OrderRequest) Validate() error {
	return validator.New().Struct(this)
}

func (this *OrderRequest) ToData() ([]byte, error) {
	return json.Marshal(this)
}

type PortfolioGetPositionsRequest struct {
	Symbol string `json:"symbol" validate:"required,min=1,max=10"`
}

func ParsePortfolioGetPositionsRequest(data []byte) (getPositionsRequest *PortfolioGetPositionsRequest, err error) {
	log.Trace().Msg("api.ParsePortfolioGetPositionsRequest")
	getPositionsRequest = &PortfolioGetPositionsRequest{}
	err = json.Unmarshal(data, &getPositionsRequest)
	return
}

func (this *PortfolioGetPositionsRequest) Validate() error {
	log.Trace().Msg("api.PortfolioGetPositionsRequest.Validate")
	return validator.New().Struct(this)
}

type MarketDataGetRequest struct {
	Symbol string `json:"symbol" validate:"required,alphanum,min=1,max=10"`
	Bar    string `json:"bar" validate:"required,oneof=1min 2min 3min 5min 10min 15min 30min 1h 2h 3h 4h 8h 1d 1w 1m"`
	Period string `json:"period" validate:"required"`
}

func ParseMarketDataGetRequest(data []byte) (req *MarketDataGetRequest, err error) {
	log.Trace().Msg("api.ParseMarketDataGetRequest")
	req = &MarketDataGetRequest{}
	err = json.Unmarshal(data, req)
	return
}

func (this *MarketDataGetRequest) Validate() error {
	log.Trace().Msg("api.MarketDataGetRequest.Validate")
	return validator.New().Struct(this)
}

func (this *MarketDataGetRequest) ToData() ([]byte, error) {
	log.Trace().Msg("api.MarketDataGetRequest.ToData")
	return json.Marshal(this)
}
