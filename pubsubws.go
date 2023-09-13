package main

import "bufio"

type PubSub struct {
	Subscribers map[string][]*bufio.ReadWriter
}

func InitiatePubSubFeed() *PubSub{
	init := make(map[string][] *bufio.ReadWriter)
	return &PubSub{Subscribers: init} 
}

func(pb *PubSub) Subscribe(topic string, wr *bufio.ReadWriter) {
	pb.Subscribers[topic] = append(pb.Subscribers[topic], wr)
}

func(pb *PubSub) Publish(topic string,message string){
	if len(pb.Subscribers) == 0{
		return
	}
	
}

// Need to figure out how I should implement the closing of a connection. 
// If connection is closed and it's inside of a list, what happens to that closed connection
func(pb *PubSub) UnSub(){

}