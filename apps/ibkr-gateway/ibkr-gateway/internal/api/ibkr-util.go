package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	util "handler/function/internal/util"

	log "github.com/rs/zerolog/log"
)

const (
	IBKR_API_PORTFOLIO_GETPOSITIONS = "/v1/portal/portfolio/positions/%d"
	IBKR_API_CONTRACTS_SEARCH       = "/v1/portal/iserver/secdef/search"
	IBKR_API_ORDERS_PREVIEW         = "/v1/portal/iserver/account/%s/order/whatif"
	IBKR_API_ORDERS_PLACE           = "/v1/portal/iserver/account/%s/order"
	IBKR_API_ORDERS_REPLY           = "/v1/portal/iserver/reply/%s"
	IBKR_API_ORDERS_LIST            = "/v1/portal/iserver/account/orders"
	IBKR_API_MARKETDATA             = "/v1/portal/iserver/marketdata/history"
	EXCHANGE_AUTHORITY              = "NYSE"
	SECURITY_TYPE                   = "STK" // STOCK
)

type IBKRRequest interface {
	Validate() error
	ToData() ([]byte, error)
}

func ibkrAPIBase() string {
	return util.Env(util.ENV_IBKR_API_URL)
}

func ibkrAPITarget(target string) string {
	if prefixed := strings.HasPrefix(target, "/"); !prefixed {
		target = "/" + target
	}
	return target
}

func doGet(target string) (responseData []byte, statusCode int, err error) {
	log.Trace().Msg("api.doGet")
	statusCode = -1
	responseData = []byte{}
	var client util.HttpClient
	if client, err = util.NewClient(ibkrAPIBase()); err == nil {
		log.Debug().Str("context", "api.doGet").Msg("created client")
		qualifiedTarget := ibkrAPITarget(target)
		var response *http.Response = nil
		if response, err = client.Get(qualifiedTarget); err == nil {
			log.Debug().Str("context", "api.doGet").Msg("request succeeded")
			defer response.Body.Close()
			statusCode = response.StatusCode
			responseData, err = ioutil.ReadAll(response.Body)
			if e := log.Debug(); e.Enabled() {
				log.Debug().
					Str("context", "api.doGet").
					Int64("responseContentLength", response.ContentLength).
					Str("response", string(responseData)).
					Int("serverStatusCode", response.StatusCode).
					Str("target", qualifiedTarget).
					Msg("response details")
			}
		}
	}
	return
}

func doPost(target string, body []byte) (responseData []byte, statusCode int, err error) {
	log.Trace().Msg("api.doPost")
	statusCode = -1
	responseData = []byte{}
	var client util.HttpClient
	if client, err = util.NewClient(ibkrAPIBase()); err == nil {
		log.Debug().Str("context", "api.doPost").Msg("created client")
		qualifiedTarget := ibkrAPITarget(target)
		var response *http.Response = nil
		if response, err = client.Post(qualifiedTarget, body); err == nil {
			log.Debug().Str("context", "api.doPost").Msg("request succeeded")
			defer response.Body.Close()
			statusCode = response.StatusCode
			responseData, err = ioutil.ReadAll(response.Body)
			if e := log.Debug(); e.Enabled() {
				log.Debug().
					Str("context", "api.doPost").
					Int64("responseContentLength", response.ContentLength).
					Str("response", string(responseData)).
					Int("serverStatusCode", response.StatusCode).
					Str("target", qualifiedTarget).
					Msg("response details")
			}
		}
	}
	return
}

func getMyAccountIndex() uint8 {
	accountIndexStr := util.Env(util.ENV_MY_ACCOUNT_INDEX)
	var result uint64
	var err error
	if result, err = strconv.ParseUint(accountIndexStr, 10, 8); err != nil {
		log.Fatal().Str("context", "api.getMyAccountIndex").
			Msg(fmt.Sprintf("Invalid index format: %s", accountIndexStr))
	}
	return uint8(result)
}
