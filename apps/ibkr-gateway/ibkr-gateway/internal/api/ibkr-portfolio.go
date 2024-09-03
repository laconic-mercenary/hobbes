package api

import (
	"encoding/json"

	log "github.com/rs/zerolog/log"
)

type IBKRPortfolioGetPositionsResponse struct {
	Entries []*IBKRPortfolioGetPositionsResponseEntry `json:"ACCTID"`
}

type IBKRPortfolioGetPositionsResponseEntry struct {
	AccountId                  string   `json:"acctId"`
	Conid                      int      `json:"conid"`
	ContractDescription        string   `json:"contractDesc"`
	AssetClass                 string   `json:"assetClass"`
	Position                   int      `json:"position"`
	MarketPrice                float32  `json:"mktPrice"`
	MarketValue                float32  `json:"mktValue"`
	Currency                   string   `json:"currency"`
	AverageCost                float32  `json:"avgCost"`
	AveragePrice               float32  `json:"avgPrice"`
	RealizedProfitAndLoss      float32  `json:"realizedPnl"`
	UnrealizedProfitAndLoss    float32  `json:"unrealizedPnl"`
	Exchanges                  string   `json:"exchs"`
	Expiry                     string   `json:"expiry"`
	PutOrCall                  string   `json:"putOrCall"`
	Multiplier                 float32  `json:"multiplier"`
	Strike                     float32  `json:"strike"`
	ExerciseStyle              string   `json:"exerciseStyle"`
	UndConid                   int      `json:"undConid"`
	ContractExchangeMap        []string `json:"conExchMap"`
	BaseMarketValue            float32  `json:"baseMktValue"`
	BaseMarketPrice            float32  `json:"baseMktPrice"`
	BaseAverageCost            float32  `json:"baseAvgCost"`
	BaseAveragePrice           float32  `json:"baseAvgPrice"`
	BaseRealizedProfitAndLoss  float32  `json:"baseRealizedPnl"`
	BaseUnealizedProfitAndLoss float32  `json:"baseUnrealizedPnl"`
	Name                       string   `json:"name"`
	LastTradingDay             string   `json:"lastTradingDay"`
	Group                      string   `json:"group"`
	Sector                     string   `json:"sector"`
	SectorGroup                string   `json:"sectorGroup"`
	Ticker                     string   `json:"ticker"`
	UndComp                    string   `json:"undComp"`
	UndSym                     string   `json:"undSym"`
	FullName                   string   `json:"fullName"`
	PageSize                   int      `json:"pageSize"`
	Model                      string   `json:"model"`
}

func ParseIBKRPortfolioGetPositionsResponse(data []byte) (getPositionsResponse *IBKRPortfolioGetPositionsResponse, err error) {
	log.Trace().Msg("api.ParseIBKRPortfolioGetPositionsResponse")
	getPositionsResponse = &IBKRPortfolioGetPositionsResponse{}
	err = json.Unmarshal(data, &getPositionsResponse)
	return
}
