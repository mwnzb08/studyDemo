package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"strconv"
	"sync"
	"time"
)

func main() {
	g := new(sync.WaitGroup)
	i:=0
	now:=time.Now()
	client := CreateClient()
	defer client.Close()
	// 存储10000个信息到pulsar 500*200
	for i < 500 {
		g.Add(1)
		i++
		go SendMsg(client,"my-topic", "hello id is " + strconv.Itoa(i), g)
	}
	g.Wait()
	fmt.Println(i * 200, time.Since(now))

}

func CreateClient() pulsar.Client {
	client, _ := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://192.168.149.130:6650",
		//OperationTimeout: time.Second * 20,
		//ConnectionTimeout: time.Second * 20,
	})
	return client
}
func SendMsg(client pulsar.Client, topic string, msg string, g *sync.WaitGroup) {
	defer g.Done()
	// 创建发送者
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
		//MaxPendingMessages: 1000,
		//SendTimeout: time.Second * 1,
	})
	defer producer.Close()
	// 创建一个发送者200个信息
	for i:=0; i < 200 ; i++ {
		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(msg),
		})
		if err != nil {
			fmt.Println("fail", err)
		}
	}
	if err != nil {
		fmt.Println("fail", err)
	}
}
