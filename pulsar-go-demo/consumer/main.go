package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"sync"
	"time"
)

func main() {
	var w = new(sync.WaitGroup)
	i:=0
	now:=time.Now()
	client := CreateClient()
	defer client.Close()
	// 从pulsar 获取10000个数据100 * 100
	// 接收的吞吐量平均3000/s
	for i < 100 {
		i++
		w.Add(1)
		go GetMsg(client,"my-topic","my-bc2", w)
	}
	w.Wait()
	defer func() {
		fmt.Println(i * 100, time.Since(now))
	}()


}
// 创建客户端
func CreateClient() pulsar.Client {
	client, _ := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://192.168.149.130:6650",
		//OperationTimeout: time.Second * 20,
		//ConnectionTimeout: time.Second * 20,
	})
	return client
}
// 获取消息
func GetMsg(client pulsar.Client, topic string,sub string, group *sync.WaitGroup)  {
	defer group.Done()
	subscribe, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: sub,
		Type:             pulsar.Failover,
	})

	defer subscribe.Close()
	// 一次穿创建,拿100个消息，减少创建接收者操作
	for i:=0; i < 100 ;i++ {
		// 接收消息
	receive, err := subscribe.Receive(context.Background())
	if err != nil  {
		log.Fatal(err)
	}
	subscribe.Ack(receive) // 特别备注，如果不加这个确认消息会一直拿缓存，如果缓存没有就去监听最新的发送，完全就是监听的模式，即使本来就有消息没被消费。
	fmt.Printf("Received message msgId: %#v -- content: '%s'\n", receive.ID(),receive.Payload())
	}
	//if err := subscribe.Unsubscribe(); err != nil {
	//	log.Fatal(err)
	//}
	if err != nil {
		log.Fatal(err)
	}

}