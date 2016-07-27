package githubapi

import stdlog "log"

type logger struct {
	debug *stdlog.Logger
	info  *stdlog.Logger
	std   *stdlog.Logger
	warn  *stdlog.Logger
	err   *stdlog.Logger
	fatal *stdlog.Logger
	Level logLevel
}

type logLevel int

const (
	// Debug ...
	Debug logLevel = iota
	// Info ...
	Info
	// Std ...
	Std
	// Warn ...
	Warn
	// Err ...
	Err
	// Fatal ...
	Fatal
)

func (l *logger) Println(v ...interface{}) {
	if l.Level <= Std {
		l.std.Println(v)
	}
}

func (l *logger) Printf(format string, v ...interface{}) {
	if l.Level <= Std {
		l.std.Printf(format, v)
	}
}

func (l *logger) Print(v ...interface{}) {
	if l.Level <= Std {
		l.std.Print(v)
	}
}

func newLogger() *logger {
	aLogger := &logger{
		std: &stdlog.Logger{},
	}
	return aLogger
}
