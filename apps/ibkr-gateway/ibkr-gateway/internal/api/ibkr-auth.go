package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	util "handler/function/internal/util"

	log "github.com/rs/zerolog/log"
)

const (
	IBKR_API_AUTH_STATUS  = "/v1/portal/iserver/auth/status"
	IBKR_API_AUTH_TICKLE  = "/v1/portal/tickle"
	IBKR_API_SSO_VALIDATE = "/v1/portal/sso/validate"
	IBKR_API_AUTH_REAUTH  = "/v1/portal/iserver/reauthenticate"
)

var (
	emptyRequestBody = []byte{}
)

type IBKRAuthStatusRequest struct {
	client *util.HttpClient `json:"-"`
}

type IBKRAuthStatusResponse struct {
	Authenticated bool     `json:"authenticated"`
	Connected     bool     `json:"connected"`
	Competing     bool     `json:"competing"`
	Fail          string   `json:"fail"`
	Message       string   `json:"message"`
	Prompts       []string `json:"prompts"`
}

func ParseIBKRAuthStatusResponse(data []byte) (asr *IBKRAuthStatusResponse, err error) {
	log.Trace().Msg("api.ParseIBKRAuthStatusResponse")
	asr = &IBKRAuthStatusResponse{}
	err = json.Unmarshal(data, asr)
	return
}

func (this IBKRAuthStatusRequest) Submit() (authStatusResponse *IBKRAuthStatusResponse, loginRequired bool, err error) {
	log.Trace().Msg("api.IBKRAuthStatusRequest.Submit")
	checkClient(this.client, "api.IBKRAuthStatusRequest.Submit")
	loginRequired = false
	var response *http.Response = nil
	if response, err = this.client.Post(IBKR_API_AUTH_STATUS, emptyRequestBody); err == nil {
		defer response.Body.Close()
		var responseData []byte = []byte{}
		if response.StatusCode == http.StatusOK {
			if responseData, err = ioutil.ReadAll(response.Body); err == nil {
				authStatusResponse, err = ParseIBKRAuthStatusResponse(responseData)
			}
		} else if response.StatusCode == http.StatusUnauthorized {
			loginRequired = true
		} else {
			if responseData, err = ioutil.ReadAll(response.Body); err == nil {
				err = util.RxToErr(response.StatusCode)
			}
		}
		if e := log.Debug(); e.Enabled() {
			log.Debug().Str("context", "api.IBKRAuthStatusRequest.Submit").
				Str("responseData", string(responseData)).
				Int("responseStatusCode", response.StatusCode).
				Str("requestBody", "(empty)").
				Msg("Finished")
		}
	}
	return
}

func (this *IBKRAuthStatusResponse) ToData() ([]byte, error) {
	log.Trace().Msg("api.IBKRAuthStatusResponse.ToData")
	return json.Marshal(this)
}

type IBKRAuthTickleRequest struct {
	client *util.HttpClient `json:"-"`
}

type IBKRAuthTickleResponse struct {
	Session    string                                `json:"session"`
	SSOExpires int                                   `json:"ssoExpires"`
	Collission bool                                  `json:"collission"`
	UserID     int                                   `json:"userId"`
	IServer    *IBKRAuthTickleResponseISserverStatus `json:"iserver"`
}

type IBKRAuthTickleResponseISserverStatus struct {
	AuthStatus *IBKRAuthTickleResponseAuthStatus `json:"authStatus"`
}

type IBKRAuthTickleResponseAuthStatus struct {
	Authenticated bool   `json:"authenticated"`
	Competing     bool   `json:"competing"`
	Connected     bool   `json:"connected"`
	Message       string `json:"message"`
	Mac           string `json:"MAC"`
}

func (this IBKRAuthTickleRequest) Submit() (tickleResponse *IBKRAuthTickleResponse, err error) {
	log.Trace().Msg("api.IBKRAuthTickleRequest.Submit")
	checkClient(this.client, "api.IBKRAuthTickleRequest.Submit")
	var response *http.Response = nil
	if response, err = this.client.Post(IBKR_API_AUTH_TICKLE, emptyRequestBody); err == nil {
		defer response.Body.Close()
		if err = util.RxToErr(response.StatusCode); err == nil {
			if responseData, err := ioutil.ReadAll(response.Body); err == nil {
				tickleResponse, err = ParseIBKRAuthTickleResponse(responseData)
			}
		}
	}
	return
}

func (this *IBKRAuthTickleResponse) ToData() ([]byte, error) {
	log.Trace().Msg("api.IBKRAuthTickleResponse.ToData")
	return json.Marshal(this)
}

func (this *IBKRAuthTickleResponse) AboutToExpire() bool {
	log.Trace().Msg("api.IBKRAuthTickleResponse.AboutToExpire")
	return this.SSOExpires < this.getMaxTimeBeforeExpire()
}

func (this *IBKRAuthTickleResponse) getMaxTimeBeforeExpire() int {
	config := util.Env(util.ENV_MAX_PREEXPIRY_TIME)
	maxPreExpiryTime, err := strconv.Atoi(config)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return maxPreExpiryTime
}

func ParseIBKRAuthTickleResponse(data []byte) (asr *IBKRAuthTickleResponse, err error) {
	log.Trace().Msg("api.ParseIBKRAuthTickleResponse")
	asr = &IBKRAuthTickleResponse{}
	err = json.Unmarshal(data, asr)
	return
}

type IBKRAuthValidateSSORequest struct {
	client *util.HttpClient `json:"-"`
}

type IBKRAuthValidateSSOResponse struct {
	LoginType int    `json:"LOGIN_TYPE"`
	UserName  string `json:"USER_NAME"`
	UserID    int    `json:"USER_ID"`
	Expire    int    `json:"expire"`
	Result    bool   `json:"RESULT"`
	AuthTime  int    `json:"AUTH_TIME"`
}

func (this IBKRAuthValidateSSORequest) Submit() (ssoResponse *IBKRAuthValidateSSOResponse, loginRequired bool, err error) {
	log.Trace().Msg("api.IBKRAuthValidateSSORequest.Submit")
	checkClient(this.client, "api.IBKRAuthValidateSSORequest.Submit")
	loginRequired = false
	var response *http.Response = nil
	if response, err = this.client.Get(IBKR_API_SSO_VALIDATE); err == nil {
		defer response.Body.Close()
		if response.StatusCode == http.StatusUnauthorized {
			loginRequired = true
		} else {
			var responseData []byte
			if responseData, err = ioutil.ReadAll(response.Body); err == nil {
				ssoResponse, err = ParseIBKRAuthValidateSSOResponse(responseData)
			}
		}
	}
	return
}

func (this *IBKRAuthValidateSSOResponse) ToData() ([]byte, error) {
	log.Trace().Msg("api.IBKRAuthValidateSSOResponse.ToData")
	return json.Marshal(this)
}

func ParseIBKRAuthValidateSSOResponse(data []byte) (asr *IBKRAuthValidateSSOResponse, err error) {
	log.Trace().Msg("api.ParseIBKRAuthValidateSSOResponse")
	asr = &IBKRAuthValidateSSOResponse{}
	err = json.Unmarshal(data, asr)
	return
}

type IBKRAuthReauthenticateRequest struct {
	client *util.HttpClient `json:"-"`
}

func (this IBKRAuthReauthenticateRequest) Submit() (authStatusResponse *IBKRAuthStatusResponse, err error) {
	log.Trace().Msg("api.IBKRAuthReauthenticateRequest.Submit")
	checkClient(this.client, "api.IBKRAuthReauthenticateRequest.Submit")
	var response *http.Response = nil
	if response, err = this.client.Post(IBKR_API_AUTH_REAUTH, emptyRequestBody); err == nil {
		defer response.Body.Close()
		if err = util.RxToErr(response.StatusCode); err == nil {
			var responseData []byte
			if responseData, err = ioutil.ReadAll(response.Body); err == nil {
				authStatusResponse, err = ParseIBKRAuthStatusResponse(responseData)
			}
		}
	}
	return
}

func checkClient(client *util.HttpClient, context string) {
	if client == nil {
		log.Fatal().Str("context", context).Msg("client not initialized")
	}
}
