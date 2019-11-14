package rabbitmq

import (
	"encoding/json"
	"github.com/natacon/mylib-go/restapi"
)

// RabbitMQの api/queuesに問い合わせて、その結果をQueuesにして返却する
// TODO: メソッド名称をGet以外に変更する。
func getQueues() (Queues, error) {
	preUnmarshalQueues, err := restapi.CallApi(restapi.ApiRequest{
		Method: "GET",
		Url:    "http://192.168.56.101:15672/api/queues",
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
