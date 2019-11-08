package rabbitmq

import (
	"errors"
)

type Queue struct {
	Arguments struct {
		XDeadLetterExchange   string `json:"x-dead-letter-exchange"`
		XDeadLetterRoutingKey string `json:"x-dead-letter-routing-key"`
	} `json:"arguments,omitempty"`
	AutoDelete         bool `json:"auto_delete"`
	BackingQueueStatus struct {
		AvgAckEgressRate  float64       `json:"avg_ack_egress_rate"`
		AvgAckIngressRate float64       `json:"avg_ack_ingress_rate"`
		AvgEgressRate     float64       `json:"avg_egress_rate"`
		AvgIngressRate    float64       `json:"avg_ingress_rate"`
		Delta             []interface{} `json:"delta"`
		Len               int           `json:"len"`
		Mode              string        `json:"mode"`
		NextSeqID         int           `json:"next_seq_id"`
		Q1                int           `json:"q1"`
		Q2                int           `json:"q2"`
		Q3                int           `json:"q3"`
		Q4                int           `json:"q4"`
		TargetRAMCount    string        `json:"target_ram_count"`
	} `json:"backing_queue_status"`
	ConsumerUtilisation       interface{}   `json:"consumer_utilisation"`
	Consumers                 int           `json:"consumers"`
	Durable                   bool          `json:"durable"`
	EffectivePolicyDefinition []interface{} `json:"effective_policy_definition"`
	Exclusive                 bool          `json:"exclusive"`
	ExclusiveConsumerTag      interface{}   `json:"exclusive_consumer_tag"`
	GarbageCollection         struct {
		FullsweepAfter  int `json:"fullsweep_after"`
		MaxHeapSize     int `json:"max_heap_size"`
		MinBinVheapSize int `json:"min_bin_vheap_size"`
		MinHeapSize     int `json:"min_heap_size"`
		MinorGcs        int `json:"minor_gcs"`
	} `json:"garbage_collection"`
	HeadMessageTimestamp       interface{} `json:"head_message_timestamp"`
	IdleSince                  string      `json:"idle_since"`
	Memory                     int         `json:"memory"`
	MessageBytes               int         `json:"message_bytes"`
	MessageBytesPagedOut       int         `json:"message_bytes_paged_out"`
	MessageBytesPersistent     int         `json:"message_bytes_persistent"`
	MessageBytesRAM            int         `json:"message_bytes_ram"`
	MessageBytesReady          int         `json:"message_bytes_ready"`
	MessageBytesUnacknowledged int         `json:"message_bytes_unacknowledged"`
	MessageStats               struct {
		Ack        int `json:"ack"`
		AckDetails struct {
			Rate int `json:"rate"`
		} `json:"ack_details"`
		Deliver        int `json:"deliver"`
		DeliverDetails struct {
			Rate int `json:"rate"`
		} `json:"deliver_details"`
		DeliverGet        int `json:"deliver_get"`
		DeliverGetDetails struct {
			Rate int `json:"rate"`
		} `json:"deliver_get_details"`
		DeliverNoAck        int `json:"deliver_no_ack"`
		DeliverNoAckDetails struct {
			Rate int `json:"rate"`
		} `json:"deliver_no_ack_details"`
		Get        int `json:"get"`
		GetDetails struct {
			Rate int `json:"rate"`
		} `json:"get_details"`
		GetNoAck        int `json:"get_no_ack"`
		GetNoAckDetails struct {
			Rate int `json:"rate"`
		} `json:"get_no_ack_details"`
		Publish        int `json:"publish"`
		PublishDetails struct {
			Rate int `json:"rate"`
		} `json:"publish_details"`
		Redeliver        int `json:"redeliver"`
		RedeliverDetails struct {
			Rate int `json:"rate"`
		} `json:"redeliver_details"`
	} `json:"message_stats,omitempty"`
	Messages        int `json:"messages"`
	MessagesDetails struct {
		Rate int `json:"rate"`
	} `json:"messages_details"`
	MessagesPagedOut     int `json:"messages_paged_out"`
	MessagesPersistent   int `json:"messages_persistent"`
	MessagesRAM          int `json:"messages_ram"`
	MessagesReady        int `json:"messages_ready"`
	MessagesReadyDetails struct {
		Rate int `json:"rate"`
	} `json:"messages_ready_details"`
	MessagesReadyRAM              int `json:"messages_ready_ram"`
	MessagesUnacknowledged        int `json:"messages_unacknowledged"`
	MessagesUnacknowledgedDetails struct {
		Rate int `json:"rate"`
	} `json:"messages_unacknowledged_details"`
	MessagesUnacknowledgedRAM int         `json:"messages_unacknowledged_ram"`
	Name                      string      `json:"name"`
	Node                      string      `json:"node"`
	OperatorPolicy            interface{} `json:"operator_policy"`
	Policy                    interface{} `json:"policy"`
	RecoverableSlaves         interface{} `json:"recoverable_slaves"`
	Reductions                int         `json:"reductions"`
	ReductionsDetails         struct {
		Rate int `json:"rate"`
	} `json:"reductions_details"`
	State string `json:"state"`
	Vhost string `json:"vhost"`
}

type Queues []Queue

func (queues Queues) SelectByQueueName(queueName string) (Queue, error) {
	for _, q := range queues {
		if q.Name == queueName {
			return q, nil
		}
	}
	return Queue{}, ErrorNotFound
}

var ErrorNotFound = errors.New("Queue not found")
