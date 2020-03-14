package request

import "encoding/json"

type Requester interface {
    SetBizContent(data map[string]interface{})
    GetBizContent() string
    GetApiMethod() string
    GetApiVersion() string
    GetNotifyUrl() string
    SetNotifyUrl(url string)
}

func JsonEncode(data map[string]interface{}) string {
    b, err := json.Marshal(data)
    if err != nil {
        panic("Json Encode Error With:" + err.Error())
    }
    
    return string(b)
}
