package redis

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestSubPub(t *testing.T) {
	type User struct {
		Name   string
		Gender string
		Age    int
	}
	Init()
	alice := &User{"alice", "female", 18}
	bob := &User{"bob", "male", 20}
	frank := &User{"frank", "male", 20}
	aliceBytes, _ := json.Marshal(alice)
	bobBytes, _ := json.Marshal(bob)
	frankBytes, _ := json.Marshal(frank)

	ctx := context.Background()
	n,err := Client.Publish(ctx,"testChannel1:key1:key2:key3", string(aliceBytes)).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(n)
	n,err = Client.Publish(ctx,"testChannel2:key1:key2:key3", string(bobBytes)).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(n)
	n,err = Client.Publish(ctx,"testChannel2:key1:key2:key3", string(frankBytes)).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(n)

	ch := Client.Subscribe(ctx,"testChannel1:key1:key2:key3","testChannel2:key1:key2:key3").Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
