# 请求支付宝接口封装类
### 目的在于在应用中更加快速、简便的使用支付宝接口以及新增接口。使得整个支付宝的接口使用全部面向对象化。

### 安装
```
go get -u github.com/reaburoa/alipaySDK
```

### 使用
在使用中只需要初始化相关支付宝接口类，即刻快速在应用中使用支付宝收款等一系列操作。接口返回数据可以根据需要获得不同类型的数据。

```go
aliClient := alipay.NewClient(
        "app_id",
        "https://openapi.alipay.com/gateway.do",
        "MIIEvwIBADA", // 私钥
        "MIIBIjANB", // 支付宝公钥
        "RSA2"
)

// 预下单
reqData := map[string]interface{}{
    "out_trade_no": "test_go_test_order0002uuy2",
    "subject": "test_go_test_order0002",
    "total_amount": "0.01",
    "buyer_id" :"*****",
}
requestObj := &request.AlipayTradeCreate{}
requestObj.SetBizContent(reqData)
requestObj.SetNotifyUrl("http://www.baidu.com")
by, err := aliClient.Execute(requestObj, "POST", "auth_token", "")
if err != nil {
    fmt.Println(err.Error())
}
```
