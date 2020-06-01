package request

type AlipayTradeCreate struct {
    BizContent string
    NotifyUrl string
}

// biz_content 中的buyer_id需要传递
func (t *AlipayTradeCreate) SetBizContent(data map[string]interface{}) {
    t.BizContent = JsonEncode(data)
}

func (t *AlipayTradeCreate) GetBizContent() string {
    return t.BizContent
}

func (t *AlipayTradeCreate) GetApiMethod() string {
    return "alipay.trade.create"
}

func (t *AlipayTradeCreate) GetApiVersion() string {
    return "1.0"
}

func (t *AlipayTradeCreate) SetNotifyUrl(str string) {
    t.NotifyUrl = str
}

func (t *AlipayTradeCreate) GetNotifyUrl() string {
    return t.NotifyUrl
}