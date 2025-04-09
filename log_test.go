package main

import "testing"

func TestLogging(t *testing.T) {
	log := NewLogger()
	log.EnableDebug = true

	log.Debug("This is a debug message")
	log.Log("This is a log message")
	log.Info("This is a info message")
	log.Warn("This is a warning")
	log.Error("This is a error message")
}
