package githubapi

import "log"

type logger struct {
	debug log.Logger
	info  log.Logger
	std   log.Logger
	warn  log.Logger
	err   log.Logger
	fatal log.Logger
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
