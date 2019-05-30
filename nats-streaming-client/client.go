package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/go-nats-streaming"
	"github.com/nats-io/go-nats"
)


func main() {
	var clusterID string
	var clientID string
	var showTime bool
	var startSeq uint64
	var startDelta string
	var deliverAll bool
	var deliverLast bool
	var durable string
	var qgroup string
	var unsubscribe bool
	var URL string


	clientID = "whale-penguin"
	clusterID = "yonghui"
	// Connect to a server
	nc, _ := nats.Connect("nats://10.216.155.34:4222,nats://10.216.155.35:4222,nats://10.216.155.36:4222")

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Simple Sync Subscriber
	sub, err := nc.SubscribeSync("foo")
	m, err := sub.NextMsg(timeout)

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("foo", ch)
	msg := <- ch

	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Requests
	msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)

	// Replies
	nc.Subscribe("help", func(m *Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}
