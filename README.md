# kafka-go
Learning how to publish and consume messages in kafka with go


### How to execute

First, connect to the app container. Ex:

```sh
docker exec -it kafka-go_app_1 bash
```

#### Producer

Run the producer (it will publish a message)
```
go run producer/main.go
```

#### Consumer

Run the consumer (it will consume the messages)
```
go run consumer/main.go
```