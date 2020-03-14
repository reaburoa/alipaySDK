package alipay

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/reaburoa/elec-signature/signature"
    "io/ioutil"
    "net/http"
    "net/url"
    "reflect"
    "regexp"
    "sort"
    "strconv"
    "strings"
    "time"
)

var (
    responseFix = "_response"
    requestCharset = "UTF-8"
    format = "json"
    errResponse = "error_response"
    signTag = "sign"
    signTypeRSA2 = "RSA2"
)

type AliPayClient struct {
    AppId                 string
    GateWay               string
    Format                string
    RsaPrivateKeyFilePath string
    RsaPrivateKey         string
    AliPublicKeyFilePath  string
    AliPublicKey          string
    RequestCharset        string
    SignType              string
    EncryptKey            string
    EncryptType           string
    Client                *http.Client
}

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

func (r *Response) GetResponse(req requestKernel, client *AliPayClient) (map[string]interface{}, error) {
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

func NewClient(appId, gateWay, privateKey, aliPublicKey, signType string) *AliPayClient {
    return &AliPayClient{
        AppId: appId,
        GateWay: gateWay,
        Format: format,
        RsaPrivateKey: privateKey,
        AliPublicKey: aliPublicKey,
        RequestCharset: requestCharset,
        SignType: signType,
        Client: http.DefaultClient,
    }
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

func (a *AliPayClient) sortContentByKeys(data map[string]interface{}) []string {
    keys := []string{}
    for k, _ := range data {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    return keys
}

func (a *AliPayClient) number2String(number interface{}) string {
    kStr := reflect.TypeOf(number).Kind().String()
    switch kStr {
    case "int64":
        number = strconv.FormatInt(number.(int64), 10)
    case "int32":
        number = strconv.FormatInt(number.(int64), 10)
    case "int":
        number = strconv.Itoa(number.(int))
    case "float64":
        number = strconv.FormatFloat(number.(float64), 'f', -1, 64)
    case "float32":
        number = strconv.FormatFloat(number.(float64), 'f', -1, 64)
    case "string":
    default:
        number = ""
    }
    
    return number.(string)
}

func (a *AliPayClient) genSignContent(data map[string]interface{}) string {
    sortedKeys := a.sortContentByKeys(data)
    toSignData := []string{}
    for _, key := range sortedKeys {
        value := data[key]
        if value == nil || strings.Trim(value.(string), "") == "" {
            continue
        }
        toSignData = append(toSignData, fmt.Sprintf("%s=%v", key, value))
    }
    return strings.Join(toSignData, "&")
}

func (a *AliPayClient) genSign(data, signType string) string {
    priKey := []byte(signature.FastFormatPrivateKey(a.RsaPrivateKey))
    var sign string
    if signType == signTypeRSA2 {
        sign, _ = signature.SignSha256WithRsa(data, priKey)
    } else {
        sign, _ = signature.SignSha1WithRsa(data, priKey)
    }
    
    return sign
}

func (a *AliPayClient) checkSign(data, sign, signType string) bool {
    pubKey := []byte(signature.FastFormatPublicKey(a.AliPublicKey))
    var err error
    if signType == signTypeRSA2 {
        err = signature.VerifySignSha256WithRsa(data, sign, pubKey)
    } else {
        err = signature.VerifySignSha1WithRsa(data, sign, pubKey)
    }
    return err == nil
}

func (a *AliPayClient) CheckNotifySign(notifyData map[string]interface{}) bool {
    if notifyData == nil || len(notifyData) == 0 {
        return false
    }
    sign := notifyData["sign"]
    signType := notifyData["sign_type"]
    delete(notifyData, "sign")
    delete(notifyData, "sign_type")
    toVerifyData := a.genSignContent(notifyData)
    verifyStr, _ := url.QueryUnescape(toVerifyData)
    return a.checkSign(verifyStr, sign.(string), signType.(string))
}

func (a *AliPayClient) formatUrlValue(data map[string]interface{}) url.Values {
    var formData = make(url.Values)
    for key, val := range data {
        val = a.number2String(val)
        formData.Set(key, val.(string))
    }
    
    return formData
}

func (a *AliPayClient) methodNameToResponseName(req requestKernel) string {
    method := req.GetApiMethod()
    respStr := strings.Replace(method, ".", "_", -1)
    return respStr + responseFix
}

func (a *AliPayClient) setHeader(req *http.Request) {
    req.Header.Set("content-type", "application/x-www-form-urlencoded;charset="+a.RequestCharset)
}

func (a *AliPayClient) genReqData(req requestKernel, authToken string) map[string]interface{} {
    commonReq := CommonRequest{
        AppId:        a.AppId,
        Method:       req.GetApiMethod(),
        Format:       a.Format,
        Charset:      a.RequestCharset,
        SignType:     a.SignType,
        Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
        Version:      req.GetApiVersion(),
        NotifyUrl:    req.GetNotifyUrl(),
        AppAuthToken: authToken,
        BizContent:   req.GetBizContent(),
    }
    clientMap := commonReq.toMap()
    commonReq.Sign = a.genSign(a.genSignContent(clientMap), a.SignType)
    clientMap["sign"] = commonReq.Sign
    
    return clientMap
}

func (a *AliPayClient) Execute(req requestKernel, method, authToken, appAuthToken string) (Response, error) {
    formData := a.formatUrlValue(a.genReqData(req, authToken))
    buf := strings.NewReader(formData.Encode())
    reqes, err := http.NewRequest(method, a.GateWay, buf)
    if err != nil {
        return "", err
    }
    a.setHeader(reqes)
    resp, err := a.Client.Do(reqes)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    parsedBody, err := a.parseBody(body, req)
    if err != nil {
        return "", err
    }
    if parsedBody["sign"] != "" {
        checkRet := a.checkSign(parsedBody["sign_data"], parsedBody["sign"], a.SignType)
        if checkRet != true {
            return "", errors.New("Check Sign Error")
        }
    }
    return Response(body), nil
}

func (a *AliPayClient) parseBody(body []byte, req requestKernel) (map[string]string, error) {
    bodyStr := string(body)
    responseReg := a.methodNameToResponseName(req)
    if strings.Index(bodyStr, errResponse) > -1 {
        responseReg = errResponse
    }
    mapResp := make(map[string]interface{})
    err := json.Unmarshal(body, &mapResp)
    if err != nil {
        return nil, err
    }
    reg, sign := "", ""
    if strings.Index(bodyStr, signTag) == -1 {
        reg = "{\"" + responseReg + `":\s?{(.*)}`
    } else {
        reg = "{\"" + responseReg + `":\s?{(.*)},`
        sign = mapResp["sign"].(string)
    }
    re, err := regexp.Compile(reg)
    if err != nil {
        return nil, err
    }
    toVerifyStr := re.FindString(bodyStr)
    start := len("{\""+responseReg+"\":")
    end := len(toVerifyStr) - 1
    return map[string]string{
        "sign_data": strings.Trim(string(toVerifyStr[start:end]), ""),
        "sign": sign,
    }, nil
}
