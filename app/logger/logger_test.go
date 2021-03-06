package logger_test

import (
	"testing"

	"github.com/ngocphuongnb/tetua/app/logger"
	"github.com/ngocphuongnb/tetua/app/mock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mockLogger := mock.CreateLogger()
	logger.New(mockLogger)
	assert.Equal(t, mockLogger, logger.Get())

	logger.Debug("Debug message")
	assert.Equal(t, mockLogger.Last().Params[0], "Debug message")
	logger.Info("Info message")
	assert.Equal(t, mockLogger.Last().Params[0], "Info message")
	logger.Warn("Warn message")
	assert.Equal(t, mockLogger.Last().Params[0], "Warn message")
	logger.Error("Error message")
	assert.Equal(t, mockLogger.Last().Params[0], "Error message")
	logger.DPanic("DPanic message")
	assert.Equal(t, mockLogger.Last().Params[0], "DPanic message")
	logger.Panic("Panic message")
	assert.Equal(t, mockLogger.Last().Params[0], "Panic message")
	logger.Fatal("Fatal message")
	assert.Equal(t, mockLogger.Last().Params[0], "Fatal message")
}
