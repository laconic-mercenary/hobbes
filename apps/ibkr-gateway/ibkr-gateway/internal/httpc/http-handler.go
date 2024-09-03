package httpc

import (
	api "handler/function/internal/api"
	util "handler/function/internal/util"

	zerolog "github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	pkgerrors "github.com/rs/zerolog/pkgerrors"

	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

const (
	PARAM_TYPE_KEY = "type"
)

var (
	isValidRequestType = regexp.MustCompile(`^[A-Za-z][A-Za-z-]{1,40}[A-Za-z]$`).MatchString
)

func HandleRequest(body []byte, method string, queryString string, headers http.Header) (response string, statusCode int, responseHeaders http.Header, err error) {
	initLogging()
	log.Trace().Msg("httpc.HandleRequest")
	requestId := util.RandomRequestId()
	responseHeaders = getHeaders()
	log.Info().
		Str("context", "httpc.HandleRequest").
		Str("method", method).
		Str("queryString", queryString).
		Str("requestBody", string(body)).
		Float64("requestId", requestId).
		Msg("request received")
	if response, statusCode, err = multiplexor(body, method, queryString); err != nil {
		log.Error().Str("context", "httpc.HandleRequest").Err(err).Msg("error from multiplexor")
	}
	log.Info().
		Str("context", "httpc.HandleRequest").
		Int("statusCode", statusCode).
		Str("response", response).
		Float64("requestId", requestId).
		Msg("request handling complete")
	return
}

func initLogging() {
	log.Trace().Msg("httpc.initLogging")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimestampFieldName = "timestamp"
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func multiplexor(body []byte, method, queryString string) (response string, statusCode int, err error) {
	log.Trace().Msg("httpc.multiplexor")
	if method == http.MethodOptions {
		response = "OK"
		statusCode = http.StatusOK
	} else {
		var requestType string
		statusCode = http.StatusBadRequest
		if requestType, err = getRequestType(queryString); err == nil {
			switch requestType {
			case "health-check":
				response, statusCode = handleHealthCheck()
			case "auth-check":
				response, statusCode, err = handleAuthCheck()
			case "portfolio-get-accounts":
				response, statusCode, err = handlePortfolioGetAccounts()
			case "portfolio-get-positions":
				response, statusCode, err = handlePortfolioGetPositions(body)
			case "orders-preview":
				response, statusCode, err = handleOrderPreview(body)
			case "orders-place":
				response, statusCode, err = handleOrderPlace(body)
			case "orders-list":
				response, statusCode, err = handleOrdersList()
			case "market-data":
				response, statusCode, err = handleMarketDataGet(body)
			default:
				err = errors.New(fmt.Sprintf("unknown requestType '%s'", requestType))
			}
		}
	}
	return
}

func getRequestType(queryString string) (requestType string, err error) {
	log.Trace().Msg("httpc.getRequestType")
	var parameters url.Values
	if parameters, err = url.ParseQuery(queryString); err == nil {
		if requestTypes, exists := parameters[PARAM_TYPE_KEY]; exists {
			if err = validateRequestType(requestTypes); err == nil {
				requestType = requestTypes[0]
			}
		} else {
			log.Warn().
				Str("context", "http.multiplexor").
				Msg("'type' parameter not found - assuming health-check is requested")
			requestType = "health-check"
		}
	}
	return
}

func validateRequestType(requestTypes []string) (err error) {
	log.Trace().Msg("httpc.validateRequestType")
	if len(requestTypes) != 1 {
		err = errors.New("invalid requestType - only 1 value is allowed")
	} else {
		if valid := isValidRequestType(requestTypes[0]); !valid {
			err = errors.New(fmt.Sprintf("invalid requestType - only letters and hyphens are allowed - got '%s'", requestTypes[0]))
		}
	}
	return
}

func handleHealthCheck() (response string, statusCode int) {
	log.Trace().Msg("httpc.handleHealthCheck")
	if healthy := api.IsHealthy(); healthy {
		statusCode = http.StatusOK
		response = "OK"
	} else {
		statusCode = http.StatusServiceUnavailable
		response = "Unhealthy"
		log.Warn().Str("result", "Unhealthy").Msg("healthcheck")
	}
	return
}

func handleAuthCheck() (response string, statusCode int, err error) {
	log.Trace().Msg("httpc.handleAuthCheck")
	if response, err = api.AuthCheck(); err == nil {
		statusCode = http.StatusOK
		if response == "login-required" {
			statusCode = http.StatusUnauthorized
		}
	} else {
		statusCode = http.StatusInternalServerError
		response = "Err"
	}
	return
}

func handlePortfolioGetAccounts() (response string, statusCode int, err error) {
	log.Trace().Msg("httpc.handlePortfolioGetAccounts")
	if response, err = api.PortfolioGetAccounts(); err == nil {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
	}
	return
}

func handlePortfolioGetPositions(body []byte) (response string, statusCode int, err error) {
	log.Trace().Msg("httpc.handlePortfolioGetPositions")
	if response, err = api.PortfolioGetPositions(body); err == nil {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
	}
	return
}

func handleOrderPreview(body []byte) (response string, statusCode int, err error) {
	log.Trace().Msg("httpc.handleOrderPreview")
	if response, err = api.OrderPreview(body); err == nil {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
		response = "Err"
	}
	return
}

func handleOrderPlace(body []byte) (response string, statusCode int, err error) {
	log.Trace().Msg("httpc.handleOrderPlace")
	if err = api.OrderPlace(body); err == nil {
		statusCode = http.StatusOK
		response = "OK"
	} else {
		statusCode = http.StatusInternalServerError
		response = "Err"
	}
	return
}

func handleOrdersList() (response string, statusCode int, err error) {
	log.Trace().Msg("httpc.handleOrdersList")
	if response, err = api.OrdersList(); err == nil {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
		response = "Err"
	}
	return
}

func handleMarketDataGet(body []byte) (response string, statusCode int, err error) {
	log.Trace().Msg("httpc.handleMarketDataGet")
	if response, err = api.MarketDataGet(body); err == nil {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
		response = "Err"
	}
	return
}

func getHeaders() (headers http.Header) {
	headers = make(map[string][]string)
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	headers.Add("Access-Control-Allow-Credentials", "true")
	headers.Add("Access-Control-Expose-Headers", "Content-Length, Content-Type, api_key, Authorization")
	headers.Add("Access-Control-Allow-Headers", "Content-Length, Content-Type, api_key, Authorization")
	return headers
}
