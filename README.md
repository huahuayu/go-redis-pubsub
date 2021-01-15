# go redis pubsub example

## get started

clone this example

```bash
git clone https://github.com/huahuayu/go-redis-pubsub.git
```

run test in `redis_test.go`

outoput:

```text
    redis_test.go:46: testChannel1:key1:key2:key3 {"Name":"alice","Gender":"female","Age":18}
    redis_test.go:46: testChannel2:key1:key2:key3 {"Name":"bob","Gender":"male","Age":20}
    redis_test.go:46: testChannel2:key1:key2:key3 {"Name":"frank","Gender":"male","Age":20}
```


