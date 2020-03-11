package request

import "github.com/reaburoa/alipaySDK/alipay"

type AliPayTradeRefund struct {
    BizContent string
    NotifyUrl string
}

func (t AliPayTradeRefund) SetBizContent(data map[string]interface{}) {
    t.BizContent = alipay.JsonEncode(data)
}

func (t AliPayTradeRefund) GetBizContent() string {
    return t.BizContent
}

func (t AliPayTradeRefund) GetApiMethod() string {
    return "alipay.trade.refund"
}

func (t AliPayTradeRefund) GetApiVersion() string {
    return "1.0"
}

func (t AliPayTradeRefund) SetNotifyUrl(str string) {
    t.NotifyUrl = str
}

func (t AliPayTradeRefund) GetNotifyUrl() string {
    return t.NotifyUrl
}
