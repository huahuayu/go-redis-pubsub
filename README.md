## Problem

go-redis pub sub not working, no message write to the queue.

**step1:** create three user instance 

```go
	alice := &User{"alice", "female", 18}
	bob := &User{"bob", "male", 20}
	frank := &User{"frank", "male", 20}
```

**step2:** put alice in channel: testChannel1:key1:key2:key3

```go
	n,err := Client.Publish(ctx,"testChannel1:key1:key2:key3", string(aliceBytes)).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(n)
```

**step3:** put bob & frank in channel: testChannel2:key1:key2:key3

**step4:** listen to both channels.

```go
	ch := Client.Subscribe(ctx,"testChannel1:key1:key2:key3","testChannel2:key1:key2:key3").Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
```

the output:

```text
time="2021-01-14T16:24:48+08:00" level=info msg=0
time="2021-01-14T16:24:48+08:00" level=info msg=0
time="2021-01-14T16:24:48+08:00" level=info msg=0
```

no message write to the queue, thus no message can read from it.


## Reproduce

clone this example

```bash
git clone https://github.com/huahuayu/go-redis-pubsub.git
```

run test in `redis_test.go`

```bash
cd go-redis-pubsub/redis
go test
```



