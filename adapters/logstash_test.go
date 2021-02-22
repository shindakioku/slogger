package adapters

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogstashAdapter_Execute(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	}))
	defer func() { testServer.Close() }()
	logstashAdapter := NewLogstash("http://localhost:8080", testServer.Client(), 1)
	logstashAdapter.messages = make(chan []byte, 100)

	cases := []struct {
		name    string
		message []byte
	}{
		{
			name:    "Write hello world",
			message: []byte("hello world"),
		},
	}

	for _, c := range cases {
		logstashAdapter.Execute(c.message)

		if !assert.Equal(t, c.message, <-logstashAdapter.messages) {
			t.Error(c.name + " don't passed")
		}
	}
}
