package tracing

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type JaegerLoggerAdapter struct {
	logger *log.Entry
}

func NewLogger(l *log.Entry) *JaegerLoggerAdapter {
	return &JaegerLoggerAdapter{
		logger: l,
	}
}

func (l *JaegerLoggerAdapter) Error(msg string) {
	l.logger.Error(msg)
}

func (l *JaegerLoggerAdapter) Infof(msg string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(msg, args...))
}
