package api

import (
	"encoding/json"
	"fmt"
	util "handler/function/internal/util"
	"regexp"

	"github.com/go-playground/validator"
	log "github.com/rs/zerolog/log"
)

const (
	PERIOD_REGEX = "([12][0-9]|30|[1-9])min|([1-8])h|([0-9]|[1-9][0-9]|[1-9][0-9][0-9]|1000)d|([0-9]|[1-9][0-9]|[1-7][0-9][0-9])w|([0-9]|[1-9][0-9]|1[0-8][0-9])m|([1-9]|1[1-5])y"
)

type IBKRMarketDataHistoryRequest struct {
	ContractId int    `validate:"required"`
	Bar        string `validate:"required,oneof=1min 2min 3min 5min 10min 15min 30min 1h 2h 3h 4h 8h 1d 1w 1m"`
	Period     string `validate:"required,periodCheck"`
	Exchange   string `validate:"required,oneof=NYSE"`
}

type IBKRMarketDataHistoryResponse struct {
	Symbol                       string                              `json:"symbol"`
	Text                         string                              `json:"text"`
	CompanyName                  string                              `json:"companyName"`
	PriceFactor                  int                                 `json:"priceFactor"`
	StartTime                    string                              `json:"startTime"`
	High                         string                              `json:"high"`
	Low                          string                              `json:"low"`
	TimePeriod                   string                              `json:"timePeriod"`
	BarLengthSeconds             int                                 `json:"barLength"`
	MarketDataAvailability       string                              `json:"mdAvailability"`
	MarketDataDelayMills         int                                 `json:"mktDataDelay"`
	IsOutsideRegularTradingHours bool                                `json:"outsideRth"`
	TradingDayDurationSeconds    int                                 `json:"tradingDayDuration"`
	VolumeFactor                 int                                 `json:"volumeFactor"`
	PriceDisplayRule             int                                 `json:"priceDisplayRule"`
	PriceDisplayValue            string                              `json:"priceDisplayValue"`
	IsNegativeCapable            bool                                `json:"negativeCapable"`
	MessageVersion               int                                 `json:"messageVersion"`
	Points                       int                                 `json:"points"`
	TravelTime                   int                                 `json:"travelTime"`
	Data                         []IBKRMarketDataHistoryResponseData `json:"data"`
}

type IBKRMarketDataHistoryResponseData struct {
	OpenPrice     float32 `json:"o"`
	ClosePrice    float32 `json:"c"`
	HighPrice     float32 `json:"h"`
	LowPrice      float32 `json:"l"`
	Volume        float32 `json:"v"`
	UnixTimestamp float32 `json:"t"`
}

func (this *IBKRMarketDataHistoryRequest) Validate() error {
	log.Trace().Msg("api.IBKRMarketDataHistoryRequest.Validate")
	validation := validator.New()
	validation.RegisterValidation("periodCheck", this.validatePeriod)
	return validation.Struct(this)
}

func (this *IBKRMarketDataHistoryRequest) ToData() ([]byte, error) {
	log.Trace().Msg("api.IBKRMarketDataHistoryRequest.ToData")
	return json.Marshal(this)
}

func (this *IBKRMarketDataHistoryResponse) ToData() ([]byte, error) {
	log.Trace().Msg("api.IBKRMarketDataHistoryResponse.ToData")
	return json.Marshal(this)
}

func (this *IBKRMarketDataHistoryRequest) Perform() (marketDataResponse *IBKRMarketDataHistoryResponse, err error) {
	log.Trace().Msg("api.IBKRMarketDataHistoryRequest.Perform")
	if err = this.Validate(); err == nil {
		log.Debug().Str("context", "api.IBKRMarketDataHistoryRequest.Perform").Msg("validation succeeded")
		var responseData []byte
		var statusCode int
		if responseData, statusCode, err = doGet(this.toGet()); err == nil {
			if err = util.RxToErr(statusCode); err == nil {
				marketDataResponse, err = ParseIBKRMarketDataHistoryResponse(responseData)
			}
		}
	}
	return
}

func (this *IBKRMarketDataHistoryRequest) toGet() string {
	return fmt.Sprintf(
		"%s?conid=%d&exchange=%s&period=%s&bar=%s&outsideRth=false",
		IBKR_API_MARKETDATA,
		this.ContractId,
		this.Exchange,
		this.Period,
		this.Bar,
	)
}

func (this *IBKRMarketDataHistoryRequest) validatePeriod(fl validator.FieldLevel) bool {
	var err error
	var matched bool
	if matched, err = regexp.MatchString(PERIOD_REGEX, fl.Field().String()); err != nil {
		log.Fatal().Str("context", "api.IBKRMarketDataHistoryRequest.Validate").Msg("periodCheck failed")
		panic(err)
	}
	return matched
}

func ParseIBKRMarketDataHistoryResponse(data []byte) (marketDataRx *IBKRMarketDataHistoryResponse, err error) {
	log.Trace().Msg("api.ParseIBKRMarketDataHistoryResponse")
	marketDataRx = &IBKRMarketDataHistoryResponse{}
	err = json.Unmarshal(data, marketDataRx)
	return
}
