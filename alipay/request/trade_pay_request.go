package request

import (
    "github.com/reaburoa/alipaySDK/alipay"
)

type TradePayRequest struct {
    BizContent string
    NotifyUrl string
}

func (t TradePayRequest) SetBizContent(data map[string]interface{}) {
    t.BizContent = alipay.JsonEncode(data)
}

func (t TradePayRequest) GetBizContent() string {
    return t.BizContent
}

func (t TradePayRequest) GetApiMethod() string {
    return "alipay.trade.pay"
}

func (t TradePayRequest) GetApiVersion() string {
    return "1.0"
}

func (t TradePayRequest) SetNotifyUrl(str string) {
    t.NotifyUrl = str
}

func (t TradePayRequest) GetNotifyUrl() string {
    return t.NotifyUrl
}

