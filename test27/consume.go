package main

import (
	"github.com/sevenNt/rocketmq"
	"time"
	"fmt"
)

func main() {

	group := "dev-VodHotClacSrcData"
	topic := "canal_vod_collect__video_collected_count_live"
	var timeSleep = 60 * time.Second

	conf := &rocketmq.Config{
		Namesrv:   "rocketmq-cs.zxbike.top:32075",
		ClientIp:"192.168.60.94",
		InstanceName: "DEFAULT",
	}
    buf :=[]string{"1","2","3"}
	//consumer, err := rocketmq.NewDefaultConsumer("KUBE_TOPIC_CONSUMER", conf)
	//if err != nil {
	//	panic(err)
	//}
	////rocketmq.NewDefaultProducer("KUBE_TOPIC_CONSUMER",conf)
	//consumer.Subscribe("kube_topic", "")
	consumer, err := rocketmq.NewDefaultConsumer(group, conf)
	if err != nil {
		return
	}
	consumer.Subscribe(topic, "")
	consumer.RegisterMessageListener(
		func(msgs []*rocketmq.MessageExt) error {
			for i, msg := range msgs {
				fmt.Println("msg=", i, msg.Topic, msg.Flag, msg.Properties, string(msg.Body))
			}
			buf = append(buf,"test" )
			fmt.Println("####3######",buf)
			fmt.Println("Consume success!")
			return nil
		})
	consumer.Start()
	fmt.Println("######1####",buf)
	time.Sleep(timeSleep)
	fmt.Println("#######2###",buf)
}
