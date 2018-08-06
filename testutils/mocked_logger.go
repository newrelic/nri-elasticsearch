package testutils

import "github.com/stretchr/testify/mock"

// MockedLogger is a mocked logger
type MockedLogger struct {
	mock.Mock
}

// Debugf debug format
func (l *MockedLogger) Debugf(format string, args ...interface{}) {
	args = append([]interface{}{format}, args...)
	l.Called(args...)
}

// Infof info format
func (l *MockedLogger) Infof(format string, args ...interface{}) {
	args = append([]interface{}{format}, args...)
	l.Called(args...)
}

// Errorf error format
func (l *MockedLogger) Errorf(format string, args ...interface{}) {
	args = append([]interface{}{format}, args...)
	l.Called(args...)
}

// Warnf warning format
func (l *MockedLogger) Warnf(format string, args ...interface{}) {
	args = append([]interface{}{format}, args...)
	l.Called(args...)
}
