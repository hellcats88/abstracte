package base

import (
	"fmt"
	"runtime"
	"time"

	"github.com/hellcats88/abstracte/lib/logging"
)

func defaultExtraParameterFormat(extras []logging.K) string {
	return fmt.Sprintf("%+v", extras)
}

type WriteMessage func(level logging.LoggerLevel, msg string)

type LoggerBase struct {
	config               logging.Config
	writeMessageCallback WriteMessage
}

func NewLoggerBase(config logging.Config, writeMessageCallback WriteMessage) logging.Logger {
	log := LoggerBase{
		config:               config,
		writeMessageCallback: writeMessageCallback,
	}

	if len(log.config.Order) == 0 {
		log.config.Order = append(log.config.Order, logging.Level, logging.Timestamp, logging.CorrelationId, logging.Message, logging.Extra)
	}

	if log.config.ExtraParametersFormat == nil {
		log.config.ExtraParametersFormat = defaultExtraParameterFormat
	}

	if !log.config.SkipExtraParameterPrefix {
		if log.config.ExtraParametersPrefix == "" {
			log.config.ExtraParametersPrefix = "extra-parameters="
		}
	}

	if log.config.CustomTime == nil {
		log.config.CustomTime = func() (string, time.Time) {
			now := time.Now()
			if log.config.TimeFormat != "" {
				return now.Format(log.config.TimeFormat), now
			} else {
				return now.String(), now
			}
		}
	}

	return log
}

func (cns LoggerBase) composeMessage(level string, ctx logging.Context, message string) string {
	msg := ""

	for _, part := range cns.config.Order {
		switch part {
		case logging.Level:
			if !cns.config.SkipPrintLevel {
				msg += level + ":"
			}

		case logging.Timestamp:
			if !cns.config.SkipPrintTimestamp {
				str, _ := cns.config.CustomTime()
				msg += "[" + str + "]:"
			}

		case logging.CorrelationId:
			if !cns.config.SkipPrintCorrelationID {
				msg += "[" + ctx.CorrID() + "]:"
			}

		case logging.Message:
			msg += message

		case logging.Extra:
			if len(ctx.GetExtras()) > 0 {
				if msg != "" {
					msg += ":" + cns.config.ExtraParametersPrefix + cns.config.ExtraParametersFormat(ctx.GetExtras())
				} else {
					msg += cns.config.ExtraParametersPrefix + cns.config.ExtraParametersFormat(ctx.GetExtras())
				}
			}

		default:
			break
		}
	}

	return msg[:len(msg)-1]
}

func (cns LoggerBase) _print(ctx logging.Context, referenceLevel logging.LoggerLevel, levelText string, msg string, params ...interface{}) {
	if referenceLevel <= cns.config.Level {
		if cns.config.CustomLogFormat != nil {
			_, now := cns.config.CustomTime()
			cns.writeMessageCallback(referenceLevel, cns.config.CustomLogFormat(logging.CustomLogFormatData{
				Level:         referenceLevel,
				CorrelationID: ctx.CorrID(),
				CurrentTime:   now,
				Message:       fmt.Sprintf(msg, params...),
				ExtraParams:   ctx.GetExtras(),
			}))
		} else {
			cns.writeMessageCallback(referenceLevel, cns.composeMessage(levelText, ctx, fmt.Sprintf(msg, params...)))
		}
	}
}

func (cns LoggerBase) Debug(ctx logging.Context, msg string, params ...interface{}) {
	cns._print(ctx, logging.Debug, "DEBUG", msg, params...)
}

func (cns LoggerBase) Trace(ctx logging.Context, msg string, params ...interface{}) {
	cns._print(ctx, logging.Trace, "TRACE", msg, params...)
}

func (cns LoggerBase) Error(ctx logging.Context, msg string, params ...interface{}) {
	cns._print(ctx, logging.Error, "ERROR", msg, params...)
}

func (cns LoggerBase) Info(ctx logging.Context, msg string, params ...interface{}) {
	cns._print(ctx, logging.Info, "INFO", msg, params...)
}

func (cns LoggerBase) Warn(ctx logging.Context, msg string, params ...interface{}) {
	cns._print(ctx, logging.Warn, "WARN", msg, params...)
}

func (cns LoggerBase) BeginMethod(ctx logging.Context) {
	if logging.Debug <= cns.config.Level {
		fpcs := make([]uintptr, 1)
		runtime.Callers(2, fpcs)
		caller := runtime.FuncForPC(fpcs[0] - 1)
		cns.Debug(ctx, "Begin "+caller.Name())
	}
}

func (cns LoggerBase) BeginMethodParams(ctx logging.Context, format string, params ...interface{}) {
	if logging.Debug <= cns.config.Level {
		fpcs := make([]uintptr, 1)
		runtime.Callers(2, fpcs)
		caller := runtime.FuncForPC(fpcs[0] - 1)
		cns.Debug(ctx, "Begin "+caller.Name()+" "+format, params...)
	}
}

func (cns LoggerBase) EndMethod(ctx logging.Context) {
	if logging.Debug <= cns.config.Level {
		fpcs := make([]uintptr, 1)
		runtime.Callers(2, fpcs)
		caller := runtime.FuncForPC(fpcs[0] - 1)
		cns.Debug(ctx, "End "+caller.Name())
	}
}

func (cns LoggerBase) EndMethodParams(ctx logging.Context, format string, params ...interface{}) {
	if logging.Debug <= cns.config.Level {
		fpcs := make([]uintptr, 1)
		runtime.Callers(2, fpcs)
		caller := runtime.FuncForPC(fpcs[0] - 1)
		cns.Debug(ctx, "End "+caller.Name()+" "+format, params...)
	}
}
