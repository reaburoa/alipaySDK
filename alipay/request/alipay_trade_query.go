package request

import "github.com/reaburoa/alipaySDK/alipay"

type AliPayTradeQuery struct {
    BizContent string
}

func (t AliPayTradeQuery) SetBizContent(data map[string]interface{}) {
    t.BizContent = alipay.JsonEncode(data)
}

func (t AliPayTradeQuery) GetBizContent() string {
    return t.BizContent
}

func (t AliPayTradeQuery) GetApiMethod() string {
    return "alipay.trade.query"
}

func (t AliPayTradeQuery) GetApiVersion() string {
    return "1.0"
}

func (t AliPayTradeQuery) SetNotifyUrl(str string) {

}

func (t AliPayTradeQuery) GetNotifyUrl() string {
    return ""
}