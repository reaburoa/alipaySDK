package request

import "github.com/reaburoa/alipaySDK/alipay"

type AliPayTradeFastPayRefundQuery struct {
    BizContent string
}

func (t AliPayTradeFastPayRefundQuery) SetBizContent(data map[string]interface{}) {
    t.BizContent = alipay.JsonEncode(data)
}

func (t AliPayTradeFastPayRefundQuery) GetBizContent() string {
    return t.BizContent
}

func (t AliPayTradeFastPayRefundQuery) GetApiMethod() string {
    return "alipay.trade.fastpay.refund.query"
}

func (t AliPayTradeFastPayRefundQuery) GetApiVersion() string {
    return "1.0"
}

func (t AliPayTradeFastPayRefundQuery) SetNotifyUrl(str string) {

}

func (t AliPayTradeFastPayRefundQuery) GetNotifyUrl() string {
    return ""
}