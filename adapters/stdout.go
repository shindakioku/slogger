package adapters

import (
	"fmt"
)

// Adapter for pass messages to stdout
type StdoutAdapter struct {
}

func (a StdoutAdapter) Execute(message []byte) {
	fmt.Println(message)
}