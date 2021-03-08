package logging

import (
	"strings"
	"time"
)

// Level defines the level type
type Level uint

const (
	// Error level
	Error Level = 0

	// Warn level
	Warn Level = 1

	// Info level
	Info Level = 2

	// Debug level
	Debug Level = 3

	// Trace level
	Trace Level = 4
)

// ExtraParametersFormatCallback defines the format of the string conversion of the parameter in K struct format
type ExtraParametersFormatCallback func([]K) string

// CustomLogFormatData defines input parameters used to format a log message
type CustomLogFormatData struct {
	Level         Level
	Message       string
	CurrentTime   time.Time
	CorrelationID string
	ExtraParams   []K
}

// CustomLogFormatCallback defines the format of the final message
type CustomLogFormatCallback func(CustomLogFormatData) string

// CustomTimeCallback defines the time value written to the final message
type CustomTimeCallback func() (string, time.Time)

// Defines new type for the list of message items.
type PartOrder uint

const (
	// Print the level information of the final message, if requested
	PartOrderLevel = 0x0

	// Print the timestamp information of the final message, if requested
	PartOrderTimestamp = 0x1

	// Print the level correlation id of the final message, if requested
	PartOrderCorrelationId = 0x2

	// Print the level information of the final message. Always printed
	PartOrderMessage = 0x3

	// Print the level information of the final message, only if there is one extra parameter
	PartOrderExtra = 0x4
)

// LoggerConfig is the configuration of the logger behavior
type Config struct {
	ExtraParametersFormat ExtraParametersFormatCallback
	CustomLogFormat       CustomLogFormatCallback
	CustomTime            CustomTimeCallback
	Level                 Level
	TimeFormat            string
	ExtraParametersPrefix string
	Order                 []PartOrder

	SkipPrintLevel           bool
	SkipPrintCorrelationID   bool
	SkipPrintTimestamp       bool
	SkipExtraParameterPrefix bool
}

// Logger abstract the selected log backend
type Logger interface {
	Info(ctx Context, msg string, params ...interface{})
	Debug(ctx Context, msg string, params ...interface{})
	Trace(ctx Context, msg string, params ...interface{})
	Error(ctx Context, msg string, params ...interface{})
	Warn(ctx Context, msg string, params ...interface{})

	BeginMethod(ctx Context)
	BeginMethodParams(ctx Context, format string, params ...interface{})
	EndMethod(ctx Context)
	EndMethodParams(ctx Context, format string, params ...interface{})
}

// LoggerFactory defines the behavior of the factory used to generate logging from a template at runtime
type Factory interface {
	Create() Logger
	CreateFromConfig(Config) Logger
}

// Atol converts a string into a logger level. Utility function
func Atol(level string) Level {
	var logLevel Level
	switch strings.ToLower(level) {
	case "error":
		logLevel = Error
	case "info":
		logLevel = Info
	case "warn":
		logLevel = Warn
	case "debug":
		logLevel = Debug
	case "trace":
		logLevel = Trace
	default:
		logLevel = Info
	}

	return logLevel
}
