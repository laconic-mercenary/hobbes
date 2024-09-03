package api

import (
	"errors"
	"fmt"
	"math/rand"

	log "github.com/rs/zerolog/log"

	util "handler/function/internal/util"
)

func AuthCheck() (response string, err error) {
	log.Trace().Msg("api.AuthCheck")
	var httpClient util.HttpClient
	var authStatusResponse *IBKRAuthStatusResponse = nil
	var loginRequired bool = false
	if httpClient, err = util.NewClient(ibkrAPIBase()); err == nil {
		if _, loginRequired, err = (IBKRAuthStatusRequest{client: &httpClient}).Submit(); err == nil && !loginRequired {
			log.Debug().Str("context", "api.AuthCheck").Msg("Sending Tickle...")
			var tickleResponse *IBKRAuthTickleResponse = nil
			if tickleResponse, err = (IBKRAuthTickleRequest{client: &httpClient}).Submit(); err == nil {
				var tickleResponseData []byte
				if tickleResponseData, err = tickleResponse.ToData(); err == nil {
					log.Info().
						Str("context", "api.AuthCheck").
						Str("responseData", util.ToString(tickleResponseData)).
						Msg("IBKRAuthTickleResponse response data")
					log.Debug().Str("context", "api.AuthCheck").Msg("Sending SSOValidate...")
					var ssoValidateResponse *IBKRAuthValidateSSOResponse = nil
					if ssoValidateResponse, loginRequired, err = (IBKRAuthValidateSSORequest{client: &httpClient}).Submit(); err == nil && !loginRequired {
						var ssoValidateResponseData []byte
						if ssoValidateResponseData, err = ssoValidateResponse.ToData(); err == nil {
							log.Info().
								Str("context", "api.AuthCheck").
								Str("responseData", util.ToString(ssoValidateResponseData)).
								Msg("IBKRAuthValidateSSOResponse response data")
							if tickleResponse.AboutToExpire() {
								log.Warn().Str("context", "api.AuthCheck").Msg("ssoSession about to expire, sending reauthentication request...")
								if authStatusResponse, err = (IBKRAuthReauthenticateRequest{client: &httpClient}).Submit(); err == nil {
									var authStatusResponseData []byte
									if authStatusResponseData, err = authStatusResponse.ToData(); err == nil {
										log.Info().
											Str("context", "api.AuthCheck").
											Str("responseData", util.ToString(authStatusResponseData)).
											Msg("IBKRAuthTickleResponse(Reauthenticate) response data")
										response = "OK"
									}
								}
							} else {
								response = "OK"
							}
						}
					}
				}
			}
		}
		if loginRequired {
			response = "login-required"
			log.Warn().
				Str("context", "api.AuthCheck").
				Str("code", "LOGIN_REQUIRED").
				Msg("Login to the portal is required")
		}
	}
	return
}

func OrdersList() (response string, err error) {
	log.Trace().Msg("api.OrdersList")
	var responseData []byte
	var statusCode int
	responseData, statusCode, err = doGet(IBKR_API_ORDERS_LIST)
	log.Info().
		Str("context", "api.OrdersList").
		Str("target", IBKR_API_ORDERS_LIST).
		Str("responseFromServer", string(responseData)).
		Int("responseStatusCode", statusCode).
		Msg("OrdersList - response from server")
	if err = util.RxToErr2(statusCode, err); err == nil {
		var orderListResponse *IBKROrderListResponse = nil
		if orderListResponse, err = ParseIBKROrderListResponse(responseData); err == nil {
			responseData, err = orderListResponse.ToData()
			response = util.ToString(responseData)
		}
	}
	return
}

func OrderPlace(data []byte) (err error) {
	log.Trace().Msg("api.OrderPlace")
	var clientOrderRequest *OrderRequest = nil
	if clientOrderRequest, err = ParseOrderRequest(data); err == nil {
		log.Debug().
			Str("context", "api.OrderPlace").
			Msg("successfully parsed OrderRequest")
		if err = clientOrderRequest.Validate(); err == nil {
			log.Debug().
				Str("context", "api.OrderPlace").
				Msg("successfully validated OrderRequest")
			contractSearchRequest := DefaultIBKRContractSearchRequest()
			contractSearchRequest.Symbol = clientOrderRequest.Symbol
			var contractSearchResponseList []*IBKRContractSearchResponse = nil
			if contractSearchResponseList, err = contractSearchRequest.Perform(); err == nil {
				if contractSearchResponse, found := getStock(contractSearchResponseList); found {
					log.Debug().
						Str("context", "api.OrderPlace").
						Msg("IBKRContractSearchRequest succeeded")
					var account *IBKRAccountsGetAllResponse = nil
					if account, err = (IBKRAccountsGetAllRequest{}).GetAtIndex(getMyAccountIndex()); err == nil {
						orderRequest := DefaultIBKROrderRequest()
						orderRequest.AcctId = account.AccountId
						orderRequest.Conid = contractSearchResponse.Conid
						orderRequest.SecType = fmt.Sprintf("%d:STK", contractSearchResponse.Conid)
						orderRequest.COID = fmt.Sprintf("hobbes-%s-%d-%d", contractSearchResponse.Symbol, contractSearchResponse.Conid, rand.Uint32())
						orderRequest.Side = clientOrderRequest.BuyOrSell
						orderRequest.Quantity = clientOrderRequest.Quantity
						_, err = orderRequest.Submit()
					}
				} else {
					err = errors.New(fmt.Sprintf("No STOCK ('STK') is available for symbol: '%s'", clientOrderRequest.Symbol))
				}
			}
		}
	}
	return
}

func OrderPreview(data []byte) (response string, err error) {
	log.Trace().Msg("api.OrderPreview")
	var clientOrderRequest *OrderRequest = nil
	if clientOrderRequest, err = ParseOrderRequest(data); err == nil {
		log.Debug().
			Str("context", "api.OrderPreview").
			Msg("successfully parsed OrderRequest")
		if err = clientOrderRequest.Validate(); err == nil {
			log.Debug().
				Str("context", "api.OrderPreview").
				Msg("successfully validated OrderRequest")
			contractSearchRequest := DefaultIBKRContractSearchRequest()
			contractSearchRequest.Symbol = clientOrderRequest.Symbol
			var contractSearchResponseList []*IBKRContractSearchResponse = nil
			if contractSearchResponseList, err = contractSearchRequest.Perform(); err == nil {
				if contractSearchResponse, found := getStock(contractSearchResponseList); found {
					log.Debug().
						Str("context", "api.OrderPreview").
						Msg("IBKRContractSearchRequest succeeded")
					var account *IBKRAccountsGetAllResponse = nil
					if account, err = (IBKRAccountsGetAllRequest{}).GetAtIndex(getMyAccountIndex()); err == nil {
						orderRequest := DefaultIBKROrderRequest()
						orderRequest.AcctId = account.AccountId
						orderRequest.Conid = contractSearchResponse.Conid
						orderRequest.SecType = fmt.Sprintf("%d:STK", contractSearchResponse.Conid)
						orderRequest.COID = fmt.Sprintf("hobbes-%s-%d-%d", contractSearchResponse.Symbol, contractSearchResponse.Conid, rand.Uint32())
						orderRequest.Side = clientOrderRequest.BuyOrSell
						orderRequest.Quantity = clientOrderRequest.Quantity
						var previewOrderResponse *IBKRPreviewOrderResponse
						if previewOrderResponse, err = orderRequest.DryRun(); err == nil {
							var responseBytes []byte
							responseBytes, err = previewOrderResponse.ToData()
							response = string(responseBytes)
						}
					}
				} else {
					err = errors.New(fmt.Sprintf("No STOCK ('STK') is available for symbol: '%s'", clientOrderRequest.Symbol))
				}
			}
		}
	}
	return
}

func PortfolioGetAccounts() (response string, err error) {
	log.Trace().Msg("api.PortfolioGetAccounts")
	var responseData []byte
	var statusCode int
	responseData, statusCode, err = doGet(IBKR_API_PORTFOLIO_GETACCOUNTS)
	if err = util.RxToErr2(statusCode, err); err == nil {
		response = util.ToString(responseData)
	}
	return
}

func PortfolioGetPositions(body []byte) (response string, err error) {
	log.Trace().Msg("api.PortfolioGetPositions")
	var getPositionsRequest *PortfolioGetPositionsRequest = nil
	if getPositionsRequest, err = ParsePortfolioGetPositionsRequest(body); err == nil {
		if err = getPositionsRequest.Validate(); err == nil {
			contractSearchRequest := DefaultIBKRContractSearchRequest()
			contractSearchRequest.Symbol = getPositionsRequest.Symbol
			var contractSearchResponseList []*IBKRContractSearchResponse = nil
			if contractSearchResponseList, err = contractSearchRequest.Perform(); err == nil {
				if contractSearchResponse, found := getStock(contractSearchResponseList); found {
					var responseData []byte
					var statusCode int
					target := fmt.Sprintf(IBKR_API_PORTFOLIO_GETPOSITIONS, contractSearchResponse.Conid)
					if responseData, statusCode, err = doGet(target); err == nil {
						log.Info().
							Str("context", "api.PortfolioGetPositions").
							Str("target", target).
							Str("responseFromServer", string(responseData)).
							Int("responseStatusCode", statusCode).
							Msg("PortfolioGetPositions - response from server")
						if err = util.RxToErr(statusCode); err == nil {
							response = string(responseData)
						}
					}
				}
			}
		}
	}
	return
}

func MarketDataGet(body []byte) (response string, err error) {
	log.Trace().Msg("api.MarketDataGet")
	var marketDataRequest *MarketDataGetRequest = nil
	if marketDataRequest, err = ParseMarketDataGetRequest(body); err == nil {
		if err = marketDataRequest.Validate(); err == nil {
			log.Debug().
				Str("context", "api.MarketDataGet").
				Msg("MarketDataGetRequest vaidation succeeded")
			contractSearchRequest := DefaultIBKRContractSearchRequest()
			contractSearchRequest.Symbol = marketDataRequest.Symbol
			var contractSearchResponseList []*IBKRContractSearchResponse = nil
			if contractSearchResponseList, err = contractSearchRequest.Perform(); err == nil {
				if contractSearchResponse, found := getStock(contractSearchResponseList); found {
					log.Debug().
						Str("context", "api.MarketDataGet").
						Msg("IBKRContractSearchRequest succeeded")
					var ibkrMarketDataResponse *IBKRMarketDataHistoryResponse = nil
					ibkrMarketDataRequest := IBKRMarketDataHistoryRequest{}
					ibkrMarketDataRequest.ContractId = contractSearchResponse.Conid
					ibkrMarketDataRequest.Bar = marketDataRequest.Bar
					ibkrMarketDataRequest.Period = marketDataRequest.Period
					ibkrMarketDataRequest.Exchange = EXCHANGE_AUTHORITY
					if ibkrMarketDataResponse, err = ibkrMarketDataRequest.Perform(); err == nil {
						log.Debug().
							Str("context", "api.MarketDataGet").
							Msg("IBKRMarketDataHistoryRequest succeeded")
						var responseBytes []byte = nil
						responseBytes, err = ibkrMarketDataResponse.ToData()
						response = util.ToString(responseBytes)
					}
				}
			}
		}
	}
	return
}
