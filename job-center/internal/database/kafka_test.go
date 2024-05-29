package database

import (
	"coin-common/queue"
	"fmt"
	"job-center/internal/config"
	"strconv"
	"sync"
	"testing"
)

// 模拟一个消费者函数
type KafkaConsumer struct {
	topic string
	group string
	wg    *sync.WaitGroup
}

func NewConsumer(group string, topic string, wg *sync.WaitGroup) *KafkaConsumer {
	return &KafkaConsumer{
		topic: topic,
		group: group,
		wg:    wg,
	}
}
func (k *KafkaConsumer) Consume(data []byte) error {
	fmt.Printf("group:%s, topic %s 消费消息："+string(data)+"\n", k.group, k.topic)
	k.wg.Done()
	return nil
}

func (k *KafkaConsumer) Topic() string {
	return k.topic
}

func (k KafkaConsumer) Group() string {
	return k.group
}

func TestKafkaClient_Send(t *testing.T) {
	// 创建一个 Kafka 客户端实例
	k := queue.NewKafkaClient(&config.KafkaConfig{Addr: "localhost:9092", WriteCap: 128, ReadCap: 128})
	defer k.Close()
	wg := &sync.WaitGroup{}
	wg.Add(100)
	// 注册一个消费者函数
	c1 := NewConsumer("group0", "topic0", wg)
	c2 := NewConsumer("group1", "topic1", wg)
	c3 := NewConsumer("group2", "topic2", wg)
	c4 := NewConsumer("group3", "topic3", wg)
	k.RegisterConsumer(c1)
	k.RegisterConsumer(c2)
	k.RegisterConsumer(c3)
	k.RegisterConsumer(c4)

	// 启动写入和读取
	k.StartWrite()
	k.StartRead()

	// 发送消息
	for i := 0; i <= 99; i++ {
		index := i % 4
		k.Send(queue.KafkaData{Group: "group" + strconv.Itoa(index), Topic: "topic" + strconv.Itoa(index), Key: []byte("key1"), Data: []byte("value" + strconv.Itoa(i))})
	}
	wg.Wait()
}

//func TestKafkaClient_Read(t *testing.T) {
//	// 创建一个 Kafka 客户端实例
//	k := NewKafkaClient(KafkaConfig{Addr: "localhost:9092"})
//	defer k.Close()
//
//	// 注册一个消费者函数
//
//	// 启动写入和读取
//	k.StartWrite()
//	k.StartRead()
//
//	// 发送一条消息
//	k.Send(KafkaData{Group: "group1", Topic: "topic1", Key: []byte("key1"), Data: []byte("value1")})
//
//	// 等待一段时间以确保消息已被发送
//	time.Sleep(time.Second * 2)
//
//	// 读取消息
//	select {
//	case data := <-k.readChan:
//		if data.Group != "group1" || data.Topic != "topic1" || string(data.Key) != "key1" || string(data.Data) != "value1" {
//			t.Errorf("Expected message with group 'group1', topic 'topic1', key 'key1', and data 'value1', but got: %+v", data)
//		}
//	case <-time.After(time.Second * 5):
//		t.Error("Timeout waiting for message")
//	}
//}
//
//func TestKafkaClient_Close(t *testing.T) {
//	// 创建一个 Kafka 客户端实例
//	k := NewKafkaClient(KafkaConfig{Addr: "localhost:9092"})
//	defer k.Close()
//
//	// 注册一个消费者函数
//	k.RegisterConsumer("group1", mockConsumer("group1"))
//
//	// 启动写入和读取
//	k.StartWrite()
//	k.StartRead()
//
//	// 发送一条消息
//	k.Send(KafkaData{Group: "group1", Topic: "topic1", Key: []byte("key1"), Data: []byte("value1")})
//
//	// 等待一段时间以确保消息已被发送
//	time.Sleep(time.Second * 2)
//
//	// 关闭客户端
//	k.Close()
//
//	// 尝试发送消息
//	select {
//	case k.writeChan <- KafkaData{Group: "group1", Topic: "topic1", Key: []byte("key1"), Data: []byte("value1")}:
//		t.Error("Expected to fail sending message after client closed")
//	default:
//		//
//	}
//}
