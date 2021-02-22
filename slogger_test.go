package slogger

import (
	"github.com/shindakioku/slogger/adapters"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CustomAdapter struct {
}

func (a CustomAdapter) Execute(message []byte) {
}

func TestAddAdapters(t *testing.T) {
	logger := New()
	cases := []struct {
		name   string
		before func(logger2 *SLogger)
		assert func(logger2 *SLogger) bool
	}{
		{
			name: "Adapters must be empty",
			assert: func(logger2 *SLogger) bool {
				return assert.Equal(t, logger2.Adapters, []Adapter{})
			},
		},
		{
			name: "Must be init with stdout adapter",
			assert: func(_ *SLogger) bool {
				stdoutAdapter := adapters.StdoutAdapter{}

				return assert.Equal(t, New(stdoutAdapter).Adapters, []Adapter{stdoutAdapter})
			},
		},
		{
			name: "Must be init with stdout adapter and logstash",
			assert: func(_ *SLogger) bool {
				stdoutAdapter := adapters.StdoutAdapter{}
				logstashAdapter := adapters.LogstashAdapter{}

				return assert.Equal(t, New(stdoutAdapter, logstashAdapter).Adapters, []Adapter{stdoutAdapter, logstashAdapter})
			},
		},
		{
			name: "Must add custom adapter",
			assert: func(logger2 *SLogger) bool {
				customAdapter := &CustomAdapter{}
				logger2.AddAdapters(customAdapter)

				return assert.Equal(t, logger2.Adapters, []Adapter{customAdapter})
			},
		},
	}

	for _, c := range cases {
		copiedLogger := logger

		if c.before != nil {
			c.before(copiedLogger)
		}

		if !assert.True(t, c.assert(copiedLogger)) {
			t.Error(c.name + " don't passed")
		}
	}
}
