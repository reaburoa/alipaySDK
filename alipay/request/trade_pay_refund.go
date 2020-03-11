package request

import "alipaySDK/alipay"

type TradePayRefund struct {
    BizContent string
    NotifyUrl string
}

func (t TradePayRefund) SetBizContent(data map[string]interface{}) {
    t.BizContent = alipay.JsonEncode(data)
}

func (t TradePayRefund) GetBizContent() string {
    return t.BizContent
}

func (t TradePayRefund) GetApiMethod() string {
    return "alipay.trade.refund"
}

func (t TradePayRefund) GetApiVersion() string {
    return "1.0"
}

func (t TradePayRefund) SetNotifyUrl(str string) {
    t.NotifyUrl = str
}

func (t TradePayRefund) GetNotifyUrl() string {
    return t.NotifyUrl
}
