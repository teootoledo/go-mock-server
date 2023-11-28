package logs_test

import (
	"context"
	"errors"
	"mock-server/internal/logs"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func createLogger(prefix string) (*logs.Logger, *test.Hook) {
	mock, hook := test.NewNullLogger()

	l := logs.New(prefix)
	l.WithLogger(mock)

	return &l, hook
}

func createContext() context.Context {
	return context.WithValue(context.Background(), "request_id", "12341234")
}

func TestLogger_Error(t *testing.T) {
	logger, hook := createLogger("my-component")

	logger.Error(createContext(), "error in component", errors.New("error in api call"))

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "[my-component] error in component - Error: error in api call [request_id:12341234]", hook.LastEntry().Message)
}

func TestLogger_ErrorWithoutContext(t *testing.T) {
	logger, hook := createLogger("my-component")

	logger.ErrorWithoutContext("error log", errors.New("error in api call"))

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "error log - Error: error in api call", hook.LastEntry().Message)
}

func TestLogger_Errorf(t *testing.T) {
	logger, hook := createLogger("my-component")

	logger.Errorf(createContext(), "error in component: %s", errors.New("error in api call"), "myComp")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "[my-component] error in component: [myComp] - Error: error in api call [request_id:12341234]", hook.LastEntry().Message)
}

func TestLogger_Info(t *testing.T) {
	logger, hook := createLogger("my-component")

	logger.Info(createContext(), "info log")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "[my-component] info log [request_id:12341234]", hook.LastEntry().Message)
}

func TestLogger_InfoWithoutContext(t *testing.T) {
	logger, hook := createLogger("my-component")

	logger.InfoWithoutContext("info log")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "info log", hook.LastEntry().Message)
}

func TestLogger_Infof(t *testing.T) {
	logger, hook := createLogger("my-component")

	logger.Infof(createContext(), "info log: %s", "myComp")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "[my-component] info log: [myComp] [request_id:12341234]", hook.LastEntry().Message)
}

func TestLogger_Warn(t *testing.T) {
	logger, hook := createLogger("my-component")

	logger.Warn(createContext(), "Warn log")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
	assert.Equal(t, "[my-component] Warn log [request_id:12341234]", hook.LastEntry().Message)
}

func TestLogger_Warnf(t *testing.T) {
	logger, hook := createLogger("my-component")

	logger.Warnf(createContext(), "Warn log: %s", "myComp")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
	assert.Equal(t, "[my-component] Warn log: [myComp] [request_id:12341234]", hook.LastEntry().Message)
}

func TestLogger_Debug(t *testing.T) {
	mock, hook := test.NewNullLogger()
	mock.SetLevel(logrus.DebugLevel)

	l := logs.New("my-component")
	l.WithLogger(mock)

	l.Debug(createContext(), "Debug log")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.DebugLevel, hook.LastEntry().Level)
	assert.Equal(t, "[my-component] Debug log [request_id:12341234]", hook.LastEntry().Message)
}

func TestLogger_Debugf(t *testing.T) {
	mock, hook := test.NewNullLogger()
	mock.SetLevel(logrus.DebugLevel)

	l := logs.New("my-component")
	l.WithLogger(mock)

	l.Debugf(createContext(), "Debug log: %s", "myComp")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.DebugLevel, hook.LastEntry().Level)
	assert.Equal(t, "[my-component] Debug log: [myComp] [request_id:12341234]", hook.LastEntry().Message)
}
