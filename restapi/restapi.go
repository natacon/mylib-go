package restapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ApiRequest struct {
	Method    string
	Url       string
	BasicAuth BasicAuth
	Body      string
}

type BasicAuth struct {
	Username string
	Password string
}

// ベーシック認証を行う必要があるリクエストであるかを返す
// username,password両方を指定している場合にtrueを返す。
func (r *ApiRequest) needBasicAuth() bool {
	return r.BasicAuth.Username != "" && r.BasicAuth.Password != ""
}

func CallApi(apiRequest ApiRequest) ([]byte, error) {
	log.Printf("Call api=%+v\n", apiRequest)
	req, _ := http.NewRequest(apiRequest.Method, apiRequest.Url, bytes.NewBufferString(apiRequest.Body))
	if apiRequest.needBasicAuth() {
		req.SetBasicAuth(apiRequest.BasicAuth.Username, apiRequest.BasicAuth.Password)
	}
	client := http.Client{}
	resp, e := client.Do(req)
	if e != nil {
		return nil, fmt.Errorf("[ERROR] Faild to request: %v", e)
	} else if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("[ERROR] Faild to request: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	//
	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, fmt.Errorf("[ERROR] Faild to read response: %d", e)
	}
	return bytes, nil
}
