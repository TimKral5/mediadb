package utils_test

import (
	"mediadb/utils"
	"testing"
)

func TestNewMongoConnection(t *testing.T) {
	log := utils.NewLogger()
	log.DebugEnabled = true
	log.Info("This is an info message")
	log.Log("This is a log message")
	log.Debug("This is a debug message")
	log.Warn("This is a warning")
	log.Error("This is an error")
}
