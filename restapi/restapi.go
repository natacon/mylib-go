package restapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiRequest struct {
	Method    string
	Url       string
	BasicAuth BasicAuth
}

type BasicAuth struct {
	Username string
	Password string
}

// ベーシック認証を行う必要があるリクエストかを返す
// username,password両方を指定している場合にtrueを返す。
func (r *ApiRequest) needBasicAuth() bool {
	return r.BasicAuth.Username != "" && r.BasicAuth.Password != ""
}

func CallApi(apiRequest ApiRequest) ([]byte, error) {
	req, _ := http.NewRequest(apiRequest.Method, apiRequest.Url, nil)
	if apiRequest.needBasicAuth() {
		req.SetBasicAuth(apiRequest.BasicAuth.Username, apiRequest.BasicAuth.Password)
	}
	client := http.Client{}
	resp, e := client.Do(req)
	if e != nil {
		return nil, fmt.Errorf("[ERROR] Faild to apiRequest: %v", e)
	} else if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("[ERROR] Faild to apiRequest: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	//
	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, fmt.Errorf("[ERROR] Faild to read response: %d", e)
	}
	return bytes, nil
}
