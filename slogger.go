package slogger

type SLogger struct {
	Adapters []Adapter
}

// Adapters for execute logging messages
// slogger.New().AddAdapters(SomeAdapter{})
func (l *SLogger) AddAdapters(adapters ...Adapter) *SLogger {
	l.Adapters = append(l.Adapters, adapters...)

	return l
}

// Send message to the all adapters
// slogger.New().Write([]byte("Hello World!"))
func (l SLogger) Write(message []byte) {
	l.sendToAdapters(message)
}

// Send message to the all adapters
// slogger.New().WriteString("Hello World!")
func (l SLogger) WriteString(message string) {
	l.sendToAdapters([]byte(message))
}

func (l SLogger) sendToAdapters(message []byte) {
	for _, adapter := range l.Adapters {
		adapter.Execute(message)
	}
}

func New(adapters ...Adapter) *SLogger {
	if adapters == nil {
		adapters = []Adapter{}
	}

	return &SLogger{Adapters: adapters}
}
