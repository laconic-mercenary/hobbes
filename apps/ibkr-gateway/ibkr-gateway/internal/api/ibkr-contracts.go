package api

import (
	"encoding/json"
	util "handler/function/internal/util"

	"github.com/go-playground/validator"

	log "github.com/rs/zerolog/log"
)

type IBKRContractSearchRequest struct {
	Symbol  string `json:"symbol" validate:"required,min=2,max=10"`
	Name    bool   `json:"name"`
	SecType string `json:"secType" validate:"required,oneof=STK"`
}

type IBKRContractSearchResponse struct {
	Conid         int                                  `json:"conid"`
	CompanyName   string                               `json:"companyName"`
	CompanyHeader string                               `json:"companyHeader"`
	Symbol        string                               `json:"symbol"`
	Description   string                               `json:"description"`
	Opt           string                               `json:"opt"`
	War           string                               `json:"war"`
	Sections      []IBKRContractSearchResponseSecurity `json:"sections"`
}

type IBKRContractSearchResponseSecurity struct {
	SecType string `json:"secType"`
}

func DefaultIBKRContractSearchRequest() *IBKRContractSearchRequest {
	return &IBKRContractSearchRequest{
		Symbol:  "---fill-this-in---",
		Name:    false,
		SecType: SECURITY_TYPE,
	}
}

func ParseIBKRContractSearchRequest(data []byte) (contractSearch *IBKRContractSearchRequest, err error) {
	log.Trace().Msg("api.ParseIBKRContractSearchRequest")
	contractSearch = &IBKRContractSearchRequest{}
	err = json.Unmarshal(data, contractSearch)
	return
}

func ParseIBKRContractSearchResponse(data []byte) (contractSearch []*IBKRContractSearchResponse, err error) {
	log.Trace().Msg("api.ParseIBKRContractSearchResponse")
	contractSearch = make([]*IBKRContractSearchResponse, 0)
	err = json.Unmarshal(data, &contractSearch)
	return
}

func (this *IBKRContractSearchRequest) Validate() error {
	log.Trace().Msg("api.IBKRContractSearchRequest.Validate")
	return validator.New().Struct(this)
}

func (this *IBKRContractSearchRequest) ToData() ([]byte, error) {
	log.Trace().Msg("api.IBKRContractSearchRequest.ToData")
	return json.Marshal(this)
}

func (this *IBKRContractSearchRequest) Perform() (contractSearchResponse []*IBKRContractSearchResponse, err error) {
	log.Trace().Msg("api.IBKRContractSearchRequest.Perform")
	if err = this.Validate(); err == nil {
		log.Debug().Str("context", "api.IBKRContractSearchRequest.Perform").Msg("validation succeeded")
		var responseData []byte
		var statusCode int
		var postData []byte
		if postData, err = this.ToData(); err == nil {
			if responseData, statusCode, err = doPost(IBKR_API_CONTRACTS_SEARCH, postData); err == nil {
				if e := log.Debug(); e.Enabled() {
					log.Debug().
						Str("context", "api.IBKRContractSearchRequest.Perform").
						Int("statusCode", statusCode).
						Str("response", string(responseData)).
						Msg("response received")
				}
				if err = util.RxToErr(statusCode); err == nil {
					contractSearchResponse, err = ParseIBKRContractSearchResponse(responseData)
				}
			}
		}
	}
	return
}

func getStock(contractSearchResponseList []*IBKRContractSearchResponse) (*IBKRContractSearchResponse, bool) {
	// makes sure that STOCK is available for purchase - as opposed to BONDS or OPTIONS
	log.Trace().Msg("api.getStock")
	log.Debug().Msg("checking if STK is available...")
	for i := 0; i < len(contractSearchResponseList); i++ {
		for j := 0; j < len(contractSearchResponseList[i].Sections); j++ {
			if contractSearchResponseList[i].Sections[j].SecType == SECURITY_TYPE {
				return contractSearchResponseList[i], true
			}
		}
	}
	return nil, false
}
