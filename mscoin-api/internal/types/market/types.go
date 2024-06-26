package types

type (
	RateRequest struct {
		Unit string `path:"unit" json:"unit"`
		Ip   string `json:"ip,optional"`
	}

	RateResponse struct {
		Rate float64 `json:"rate"`
	}
)

type (
	MarketReq struct {
		Ip string `json:"ip,optional"`
	}
	CoinThumbResp struct {
		Symbol       string    `json:"symbol"`
		Open         float64   `json:"open"`
		High         float64   `json:"high"`
		Low          float64   `json:"low"`
		Close        float64   `json:"close"`
		Chg          float64   `json:"chg"`    //变化百分比
		Change       float64   `json:"change"` // 变化金额
		Volume       float64   `json:"volume"`
		Turnover     float64   `json:"turnover"`
		LastDayClose float64   `json:"lastDayClose"`
		UsdRate      float64   `json:"usdRate"`        // USDT汇率
		BaseUsdRate  float64   `json:"baseUsdRate"`    // 基础USDT汇率
		Zone         int       `json:"zone"`           // 交易区
		Trend        []float64 `json:"trend,optional"` //价格趋势
	}
)
