package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"job-center/internal/config"
	"log"
	"sync"
	"time"
)

type KafkaData struct {
	Group string
	Topic string
	Key   []byte
	Data  []byte
}

type KafkaClient struct {
	w         *kafka.Writer
	readers   map[group]*kafka.Reader
	writeChan chan KafkaData
	c         *config.KafkaConfig
	closed    bool
	mutex     sync.Mutex
	Consumers map[group]Consumer
}

type group string

type Consumer interface {
	Consume(data []byte) error
	Topic() string
	Group() string
}

func NewKafkaClient(c *config.KafkaConfig) *KafkaClient {
	return &KafkaClient{
		c:         c,
		readers:   make(map[group]*kafka.Reader, 16),
		Consumers: make(map[group]Consumer, 16),
	}
}

func (k *KafkaClient) RegisterConsumer(c Consumer) {
	k.Consumers[group(c.Group())] = c
}

func (k *KafkaClient) StartWrite() {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(k.c.Addr),
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		RequiredAcks:           kafka.RequireAll, // ack模式
		Async:                  true,
	}
	k.w = w
	k.writeChan = make(chan KafkaData, k.c.WriteCap)
	go k.sendKafka()
}

func (w *KafkaClient) Send(data KafkaData) {
	defer func() {
		if err := recover(); err != nil {
			logx.Info("kafka:写通道关闭")
			w.closed = true
		}
	}()
	w.writeChan <- data
	w.closed = false
}

// 关闭writer和reader
func (w *KafkaClient) Close() {
	if w.w != nil {
		w.w.Close()
		w.mutex.Lock()
		defer w.mutex.Unlock()
		if !w.closed {
			close(w.writeChan)
			w.closed = true
		}
	}

	for _, r := range w.readers {
		r.Close()
	}
}

func (w *KafkaClient) sendKafka() {
	for {
		select {
		case data := <-w.writeChan:
			messages := []kafka.Message{
				{
					Topic: data.Topic,
					Key:   data.Key,
					Value: data.Data,
				},
			}
			var err error
			const retries = 3
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			success := false
			for i := 0; i < retries; i++ {
				// attempt to create topic prior to publishing the message
				err = w.w.WriteMessages(ctx, messages...)
				if err == nil {
					success = true
					break
				}
				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					time.Sleep(time.Millisecond * 250)
					success = false
					continue
				}
				if err != nil {
					success = false
					log.Printf("kafka send writemessage err %s \n", err.Error())
				}
			}
			if !success {
				//重新放进去等待消费
				w.Send(data)
			}
		}
	}

}

// 配置reader，并启动线程监听消息队列的消息
func (k *KafkaClient) StartRead() {
	//构建并启动所有reader
	for g, c := range k.Consumers {
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{k.c.Addr},
			GroupID:  string(g),
			Topic:    c.Topic(),
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		})
		k.readers[g] = r
		go k.readMsg(r)
	}
}

func (k *KafkaClient) readMsg(reader *kafka.Reader) {

	for {
		//若没有消息则一直阻塞
		ctx := context.Background()
		m, err := reader.FetchMessage(ctx)
		if err != nil {
			logx.Error("kafka:拉取消息异常" + err.Error())
			continue
		}

		k.Consumer(reader, &m)
	}
}

func (k *KafkaClient) Consumer(reader *kafka.Reader, m *kafka.Message) {
	ctx := context.Background()
	g := string(reader.Config().GroupID)
	data := &KafkaData{
		Group: g,
		Key:   m.Key,
		Data:  m.Value,
	}
	c, ok := k.Consumers[group(g)]
	if !ok {
		logx.Info("kafka: 消费者" + data.Group + "未注册")
		return
	}
	err := c.Consume(data.Data)
	if err = reader.CommitMessages(ctx, *m); err != nil {
		logx.Info("kafka:消息提交失败：" + err.Error())
	}
	fmt.Println("kafka:消息提交成功：" + string(m.Value))
}
