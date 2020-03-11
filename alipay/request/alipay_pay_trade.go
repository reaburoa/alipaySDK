package request

import (
    "github.com/reaburoa/alipaySDK/alipay"
)

type AliPayTradePay struct {
    BizContent string
    NotifyUrl string
}

func (t AliPayTradePay) SetBizContent(data map[string]interface{}) {
    t.BizContent = alipay.JsonEncode(data)
}

func (t AliPayTradePay) GetBizContent() string {
    return t.BizContent
}

func (t AliPayTradePay) GetApiMethod() string {
    return "alipay.trade.pay"
}

func (t AliPayTradePay) GetApiVersion() string {
    return "1.0"
}

func (t AliPayTradePay) SetNotifyUrl(str string) {
    t.NotifyUrl = str
}

func (t AliPayTradePay) GetNotifyUrl() string {
    return t.NotifyUrl
}

