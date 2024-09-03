package api

import (
	"encoding/json"
	"fmt"
	util "handler/function/internal/util"

	"github.com/go-playground/validator"
	log "github.com/rs/zerolog/log"
)

type IBKROrderRequest struct {
	AcctId          string  `json:"acctId" validate:"required,min=8,max=10`
	Conid           int     `json:"conid" validate:"gte=0"`
	SecType         string  `json:"secType" validate:"required"`
	COID            string  `json:"cOID"`
	ParentId        string  `json:"parentId"`
	OrderType       string  `json:"orderType" validate:"required,oneof=MKT LMT STP STP_LIMIT"`
	ListingExchange string  `json:"listingExchange,omitempty" validate:"oneof=SMART AMEX NYSE CBOE ISE CHX ARCA ISLAND DRCTEDGE BEX BATS EDGEA CSFBALGO JE FFALGO BYX IEX FOXRIVER TPLUS1 NYSENAT PSX`
	OutsideRTH      bool    `json:"outsideRTH"`
	Price           float32 `json:"price,omitempty"`
	Side            string  `json:"side" validate:"required,oneof=SELL BUY"`
	Ticker          string  `json:"ticker"`
	Tif             string  `json:"tif" validate:"required,oneof=GTC DAY"`
	Referrer        string  `json:"referrer"`
	Quantity        float32 `json:"quantity" validate:"gte=0.0,lte=1000.0"`
	UseAdaptive     bool    `json:"useAdaptive"`
}

type IBKRPreviewOrderResponse struct {
	Amount      *IBKRPreviewOrderResponseAmount `json:"amount"`
	Equity      *IBKRPreviewOrderResponseEquity `json:"equity"`
	Initial     *IBKRPreviewOrderResponseEquity `json:"initial"`
	Maintenance *IBKRPreviewOrderResponseEquity `json:"maintenance"`
	Warn        string                          `json:"warn"`
	Error       string                          `json:"error"`
}

type IBKRPreviewOrderResponseAmount struct {
	Amount     string `json:"amount"`
	Commission string `json:"commision"`
	Total      string `json:"total"`
}

type IBKRPreviewOrderResponseEquity struct {
	Current string `json:"current"`
	Change  string `json:"change"`
	After   string `json:"after"`
}

type IBKROrderResponse struct {
	OrderId      string `json:"order_id"`
	LocalOrderId string `json:"local_order_id"`
	OrderStatus  string `json:"order_status"`
}

type IBKROrderResponse2 struct {
	Id      string   `json:"id"`
	Message []string `json:"message"`
}

func DefaultIBKROrderRequest() *IBKROrderRequest {
	return &IBKROrderRequest{
		ParentId:        "",
		OrderType:       "MKT",
		ListingExchange: "SMART",
		OutsideRTH:      false,
		Ticker:          "",
		Tif:             "DAY",
		Referrer:        "QuickTrade",
		Quantity:        0,
		UseAdaptive:     false,
	}
}

func ParseIBKROrderRequest(data []byte) (orderRequest *IBKROrderRequest, err error) {
	log.Trace().Msg("api.ParseIBKROrderRequest")
	orderRequest = &IBKROrderRequest{}
	err = json.Unmarshal(data, orderRequest)
	return
}

func ParseIBKRPreviewOrderResponse(data []byte) (previewOrderResponse *IBKRPreviewOrderResponse, err error) {
	log.Trace().Msg("api.ParseIBKRPreviewOrderResponse")
	previewOrderResponse = &IBKRPreviewOrderResponse{}
	err = json.Unmarshal(data, previewOrderResponse)
	return
}

func ParseIBKROrderResponse(data []byte) (orderResponse []*IBKROrderResponse, err error) {
	log.Trace().Msg("api.ParseIBKROrderResponse")
	orderResponse = make([]*IBKROrderResponse, 0)
	err = json.Unmarshal(data, &orderResponse)
	return
}

func ParseIBKROrderResponse2(data []byte) (orderResponse2 []*IBKROrderResponse2, err error) {
	log.Trace().Msg("api.ParseIBKROrderResponse2")
	orderResponse2 = make([]*IBKROrderResponse2, 0)
	err = json.Unmarshal(data, &orderResponse2)
	return
}

func (this *IBKROrderRequest) Validate() error {
	log.Trace().Msg("api.IBKROrderRequest.Validate")
	return validator.New().Struct(this)
}

func (this *IBKROrderRequest) ToData() ([]byte, error) {
	log.Trace().Msg("api.IBKROrderRequest.ToData")
	return json.Marshal(this)
}

func (this *IBKROrderRequest) DryRun() (previewOrderResponse *IBKRPreviewOrderResponse, err error) {
	log.Trace().Msg("api.IBKROrderRequest.DryRun")
	if err = this.Validate(); err == nil {
		var responseData []byte
		var statusCode int
		var postData []byte
		if postData, err = this.ToData(); err == nil {
			if responseData, statusCode, err = doPost(fmt.Sprintf(IBKR_API_ORDERS_PREVIEW, this.AcctId), postData); err == nil {
				if e := log.Debug(); e.Enabled() {
					log.Debug().
						Str("context", "api.IBKROrderRequest.DryRun").
						Int("statusCode", statusCode).
						Str("response", string(responseData)).
						Msg("response received")
				}
				if err = util.RxToErr(statusCode); err == nil {
					previewOrderResponse, err = ParseIBKRPreviewOrderResponse(responseData)
				}
			}
		}
	}
	return
}

func (this *IBKROrderRequest) Submit() (orderResponse []*IBKROrderResponse, err error) {
	log.Trace().Msg("api.IBKROrderRequest.Submit")
	if err = this.Validate(); err == nil {
		var postData []byte
		if postData, err = this.ToData(); err == nil {
			var responseData []byte
			var statusCode int
			if responseData, statusCode, err = doPost(fmt.Sprintf(IBKR_API_ORDERS_PLACE, this.AcctId), postData); err == nil {
				if e := log.Debug(); e.Enabled() {
					log.Debug().
						Str("context", "api.IBKROrderRequest.Submit").
						Int("statusCode", statusCode).
						Str("response", string(responseData)).
						Msg("response received")
				}
				if err = util.RxToErr(statusCode); err == nil {
					if orderResponse, err = ParseIBKROrderResponse(responseData); err != nil {
						log.Warn().
							Str("context", "api.IBKROrderRequest.Submit").
							Msg("order not submitted -- response may requiree replies")
						var questionResponse []*IBKROrderResponse2 = nil
						if questionResponse, err = ParseIBKROrderResponse2(responseData); err == nil {
							log.Debug().
								Str("context", "api.IBKROrderRequest.Submit").
								Msg("replying to questions")
							answerRequest := IBKROrderReplyRequest{
								Confirmed: true,
								ReplyId:   questionResponse[0].Id,
							}
							orderResponse, err = answerRequest.Submit()
						}
					}
				}
			}
		}
	}
	return
}

func (this *IBKROrderResponse) ToData() ([]byte, error) {
	log.Trace().Msg("api.IBKROrderResponse.ToData")
	return json.Marshal(this)
}

func (this *IBKRPreviewOrderResponse) ToData() ([]byte, error) {
	log.Trace().Msg("api.IBKRPreviewOrderResponse.ToData")
	return json.Marshal(this)
}

type IBKROrderReplyRequest struct {
	Confirmed bool   `json:"confirmed"`
	ReplyId   string `json:"-"`
}

func (this *IBKROrderReplyRequest) ToData() (data []byte, err error) {
	log.Trace().Msg("api.IBKROrderReplyRequest.ToData")
	return json.Marshal(this)
}

func (this *IBKROrderReplyRequest) Submit() (orderResponse []*IBKROrderResponse, err error) {
	log.Trace().Msg("api.IBKROrderReplyRequest.Submit")
	var postData []byte
	if postData, err = this.ToData(); err == nil {
		var responseData []byte
		var statusCode int
		if responseData, statusCode, err = doPost(fmt.Sprintf(IBKR_API_ORDERS_REPLY, this.ReplyId), postData); err == nil {
			if e := log.Debug(); e.Enabled() {
				log.Debug().
					Str("context", "api.IBKROrderReplyRequest.Submit").
					Int("statusCode", statusCode).
					Str("response", string(responseData)).
					Msg("response received")
			}
			orderResponse, err = ParseIBKROrderResponse(responseData)
		}
	}
	return
}

type IBKROrderListResponse struct {
	Orders []*IBKROrderListResponseOrder `json:"orders"`
}

type IBKROrderListResponseOrder struct {
	Account           string  `json:"acct"`
	Conid             int     `json:"conid"`
	OrderDescription  string  `json:"orderDesc"`
	Description1      string  `json:"description1"`
	Ticker            string  `json:"ticker"`
	SecurityType      string  `json:"secType"`
	ListingExchange   string  `json:"listingExchange"`
	RemainingQuantity float32 `json:"remainingQuantity"`
	FilledQuantity    float32 `json:"filledQuantity"`
	CompanyName       string  `json:"companyName"`
	Status            string  `json:"status"`
	OriginalOrderType string  `json:"origOrderType"`
	Side              string  `json:"side"`
	Price             float32 `json:"price"`
	BGColor           string  `json:"bgColor"`
	FGColor           string  `json:"fgColor"`
	OrderId           int     `json:"orderId"`
	ParentId          int     `json:"parentId"`
}

func ParseIBKROrderListResponse(data []byte) (orderListResponse *IBKROrderListResponse, err error) {
	log.Trace().Msg("api.ParseIBKROrderListResponse")
	orderListResponse = &IBKROrderListResponse{}
	err = json.Unmarshal(data, &orderListResponse)
	return
}

func (this *IBKROrderListResponse) ToData() (data []byte, err error) {
	log.Trace().Msg("api.IBKROrderListResponse.ToData")
	return json.Marshal(this)
}
