package main

import (
	"context"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)
func main(){
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		// logをdebugとtrace両方falseにする
		watermill.NewStdLogger(false, false),
		)
	message, err := pubSub.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}
	go process(message)
}

func process(messages <- chan *message.Message){
	for msg := range messages {
		log.Printf("received message: %s, payload %s", msg.UUID, string(msg.Payload))
		// messageを受け取り処理したことを知らせる
		msg.Ack()
	}
}