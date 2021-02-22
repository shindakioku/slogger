package slogger

type Adapter interface {
	Execute(message []byte)
}
