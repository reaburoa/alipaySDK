package alipay

import "reflect"

type CommonRequest struct {
    AppId        string `json:"app_id"`
    Method       string `json:"method"`
    Format       string `json:"format"`
    Charset      string `json:"charset"`
    SignType     string `json:"sign_type"`
    Sign         string `json:"sign"`
    Timestamp    string `json:"timestamp"`
    Version      string `json:"version"`
    NotifyUrl    string `json:"notify_url"`
    AppAuthToken string `json:"app_auth_token"`
    BizContent   string `json:"biz_content"`
}

func (r *CommonRequest) toMap() map[string]interface{} {
    m := make(map[string]interface{})
    elemValues := reflect.ValueOf(r).Elem()
    elemTypes := elemValues.Type()
    for i := 0; i < elemTypes.NumField(); i ++ {
        m[elemTypes.Field(i).Tag.Get("json")] = elemValues.Field(i).Interface()
    }
    return m
}
