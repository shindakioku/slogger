### SLogger

Very simple logger. All what you need: create/use adapter for logging. You can use many adapters in the same time.
Adapter it's a simple (or not) handler for your logging actions. It will be call when you pass the message for log.

### Install

```go
go get -u https: //github.com/shindakioku/slogger
```

### How to use

```go
logger := slogger.New()
logger := slogger.New(adapters.StdoutAdapter{}, MyAdapter{}) // You can pass adapters on the init stage
logger.AddAdapter(MyNewCoolAdapter{})

logger.Write([]byte("Hello World!"))
logger.WriteString("Hello World!")
```

### How to create custom adapter?

```go
type MyAdapter struct {}

func (a MyAdapter) Execute(message []byte) {
    fmt.Println(message)
}
```
So it's all. Then you can use this.
```go
slogger.New(MyAdapter{}).WriteString("Hello World!")
```

Of course you can do anything you want with your adapter, for example:
```go
type MyAdapter struct {
	messages chan []byte
}
func InitMyAdapter() MyAdapter {
	// Here you can setup mysql or elasticsearch...or write loop for :) Anything.
	MyAdapter{make(chan []byte, 1)}
	go adapter.readFromChannel()
	
	return adapter
}

func (a MyAdapter) Execute(message []byte) {
	a.messages <- message
}

func (a MyAdapter) readFromChannel() {
	for {
		log.Println(<-a.messages)
    }
}
```