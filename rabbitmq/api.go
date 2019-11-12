package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/natacon/mylib-go/restapi"
	"io/ioutil"
	"net/http"
)

// RabbitMQの restapi/queuesに問い合わせて、その結果をQueuesにして返却する
// TODO: メソッド名称をGet以外に変更する。
func getQueues() (Queues, error) {
	preUnmarshalQueues, err := restapi.CallApi(restapi.ApiRequest{
		Method: "GET",
		Url:    "http://192.168.56.101:15672/restapi/queues",
		BasicAuth: restapi.BasicAuth{
			Username: "guest",
			Password: "guest",
		},
	})
	if err != nil {
		return nil, err
	}
	var queues Queues
	err = json.Unmarshal(preUnmarshalQueues, &queues)
	return queues, err
}

func FindByQueueName(queueName string) (Queue, error) {
	queues, _ := getQueues()
	return queues.SelectByQueueName(queueName)
}

func callRabbitMQApi(method, url string) ([]byte, error) {
	req, _ := http.NewRequest(method, url, nil)
	req.SetBasicAuth("guest", "guest")
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
