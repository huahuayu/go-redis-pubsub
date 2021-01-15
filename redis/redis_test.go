package redis

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

func publish(ctx context.Context, channel string, message interface{}) {
	// Publish a message.
	err := Client.Publish(ctx, channel, message).Err()
	if err != nil {
		panic(err)
	}
}
func TestPubSub(t *testing.T) {
	Init()
	type User struct {
		Name   string
		Gender string
		Age    int
	}
	alice := &User{"alice", "female", 18}
	bob := &User{"bob", "male", 20}
	frank := &User{"frank", "male", 20}
	aliceBytes, _ := json.Marshal(alice)
	bobBytes, _ := json.Marshal(bob)
	frankBytes, _ := json.Marshal(frank)
	ctx := context.TODO()
	pubsub := Client.Subscribe(ctx, "testChannel1:key1:key2:key3", "testChannel2:key1:key2:key3")
	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}
	publish(ctx,  "testChannel1:key1:key2:key3", string(aliceBytes))
	publish(ctx,  "testChannel2:key1:key2:key3", string(bobBytes))
	publish(ctx,  "testChannel2:key1:key2:key3", string(frankBytes))
	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})
	// Consume messages.
	for msg := range pubsub.Channel() {
		t.Log(msg.Channel, msg.Payload)
	}
}

func TestSub(t *testing.T){
	Init()
	ctx := context.TODO()
	pubsub := Client.Subscribe(ctx, "testChannel1:key1:key2:key3", "testChannel2:key1:key2:key3")
	// Consume messages.
	for msg := range pubsub.Channel() {
		t.Log(msg.Channel, msg.Payload)
	}
}

func TestPub(t *testing.T){
	Init()
	type User struct {
		Name   string
		Gender string
		Age    int
	}
	alice := &User{"alice", "female", 18}
	bob := &User{"bob", "male", 20}
	frank := &User{"frank", "male", 20}
	aliceBytes, _ := json.Marshal(alice)
	bobBytes, _ := json.Marshal(bob)
	frankBytes, _ := json.Marshal(frank)
	ctx := context.TODO()
	publish(ctx,  "testChannel1:key1:key2:key3", string(aliceBytes))
	publish(ctx,  "testChannel2:key1:key2:key3", string(bobBytes))
	publish(ctx,  "testChannel2:key1:key2:key3", string(frankBytes))
}
