package api

import (
	"encoding/json"
	"errors"
	"fmt"
	util "handler/function/internal/util"

	log "github.com/rs/zerolog/log"
)

const (
	IBKR_API_PORTFOLIO_GETACCOUNTS = "/v1/portal/portfolio/accounts"
)

type IBKRAccountsGetAllRequest struct {
}

type IBKRAccountsGetAllResponse struct {
	Id           string `json:"id"`
	AccountId    string `json:"accountId"`
	AccountTitle string `json:"accountTitle"`
}

func ParseIBKRAccountsGetAllResponse(data []byte) (accounts []*IBKRAccountsGetAllResponse, err error) {
	log.Trace().Msg("api.ParseIBKRAccountsGetAllRequest")
	accounts = make([]*IBKRAccountsGetAllResponse, 0)
	err = json.Unmarshal(data, &accounts)
	return
}

func (this IBKRAccountsGetAllRequest) GetAtIndex(index uint8) (account *IBKRAccountsGetAllResponse, err error) {
	log.Trace().Msg("api.IBKRAccountsGetAllRequest.GetAtIndex")
	var responseData []byte
	var statusCode int
	if responseData, statusCode, err = doGet(IBKR_API_PORTFOLIO_GETACCOUNTS); err == nil {
		if err = util.RxToErr(statusCode); err == nil {
			var accounts []*IBKRAccountsGetAllResponse = nil
			if accounts, err = ParseIBKRAccountsGetAllResponse(responseData); err == nil {
				if len(accounts) != 0 {
					i := int(index)
					if len(accounts) > i {
						account = accounts[i]
					} else {
						err = errors.New(fmt.Sprintf("Index outside of bounds: %d", index))
					}
				} else {
					err = errors.New("No accounts were retrieved")
				}
			}
		}
	}
	return
}
