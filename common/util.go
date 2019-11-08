package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallApi(method, url string) ([]byte, error) {
	req, _ := http.NewRequest(method, url, nil)
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
