package testutils

// TestLogger a logger that logs to a generic format function, used with testing.T.Logf
type TestLogger struct {
	F func(format string, args ...interface{})
}

// Debugf debug format
func (l *TestLogger) Debugf(format string, args ...interface{}) {
	l.F("DEBUG: "+format, args...)
}

// Infof info format
func (l *TestLogger) Infof(format string, args ...interface{}) {
	l.F("INFO : "+format, args...)
}

// Errorf error format
func (l *TestLogger) Errorf(format string, args ...interface{}) {
	l.F("ERROR: "+format, args...)
}

// Warnf warning format
func (l *TestLogger) Warnf(format string, args ...interface{}) {
	l.F("WARN: "+format, args...)
}
