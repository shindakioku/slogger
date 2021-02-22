package adapters

import (
	"bytes"
	"log"
	"net/http"
)

// Logstash adapter (https://www.elastic.co/logstash)
// Adapter works by goroutines so it will be not block's your common lifecycle.
// You must use `New` function for create instance of the adapter
// adapters.NewLogstash("http://localhost:8080", &http.Client{}, 3)
type LogstashAdapter struct {
	url         string
	httpClient  *http.Client
	messages    chan []byte
	writeErrors chan []byte
}

func (a LogstashAdapter) Execute(message []byte) {
	a.messages <- message
}

func (a LogstashAdapter) waiting() {
	for {
		select {
		case message := <-a.messages:
			req, err := http.NewRequest(http.MethodPut, a.url, bytes.NewBuffer(message))
			if err != nil {
				a.writeErrors <- []byte(err.Error())

				continue
			}

			req.Header.Set("Content-Type", "application/json; charset=utf-8")
			if _, err = a.httpClient.Do(req); err != nil {
				a.writeErrors <- []byte(err.Error())
			}
		case message := <-a.writeErrors:
			// TODO: Write to file
			log.Println(string(message))
		}
	}
}

func NewLogstash(url string, httpClient *http.Client, chanLength uint8) *LogstashAdapter {
	adapter := &LogstashAdapter{
		url:         url,
		messages:    make(chan []byte, chanLength),
		writeErrors: make(chan []byte, chanLength),
		httpClient:  httpClient,
	}

	go adapter.waiting()

	return adapter
}
