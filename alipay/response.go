package alipay

import (
    "encoding/json"
    "errors"
    "github.com/reaburoa/alipaySDK/alipay/request"
)

type Response string

func (r *Response) ToMap() (map[string]interface{}, error) {
    if *r == "" {
        return nil, errors.New("Response Is Empty")
    }
    var mapResp = make(map[string]interface{})
    err := json.Unmarshal([]byte(*r), &mapResp)
    if err != nil {
        return nil, err
    }
    
    return mapResp, nil
}

func (r *Response) GetResponse(req request.Requester, client *AliPayClient) (map[string]interface{}, error) {
    mapResp, err := r.ToMap()
    if err != nil {
        return nil, err
    }
    respKey := client.methodNameToResponseName(req)
    if value, ok := mapResp[respKey]; ok {
        return value.(map[string]interface{}), nil
    } else {
        return mapResp[errResponse].(map[string]interface{}), nil
    }
}
