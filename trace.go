package aliyuniot

// Logger is the fundamental interface for all log operations. Log creates a
// log event from keyvals, a variadic sequence of alternating keys and values.
// Implementations must be safe for concurrent use by multiple goroutines. In
// particular, any implementation of Logger that appends to keyvals or
// modifies or retains any of its elements must make a copy first.
// from github.com/go-kit/kit/log
type Logger interface {
	Log(keyvals ...interface{}) error
}

// NOOPLogger implements the logger that does not perform any operation
// by default. This allows us to efficiently discard the unwanted messages.
type NOOPLogger struct{}

// Log do nothing
func (l *NOOPLogger) Log(keyvals ...interface{}) error { return nil }

// Internal levels of library output that are initialised to not print
// anything but can be overridden by programmer
var (
	Error    Logger = &NOOPLogger{}
	Critical Logger = &NOOPLogger{}
	Warning  Logger = &NOOPLogger{}
	Debug    Logger = &NOOPLogger{}
)
