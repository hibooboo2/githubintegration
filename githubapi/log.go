package githubapi

import (
	stdlog "log"
	"os"

	"github.com/fatih/color"
)

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
	// Error ...
	Error
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

func (l *logger) Debugln(v ...interface{}) {
	if l.Level <= Debug {
		l.debug.Println(v...)
	}
}

func (l *logger) Debugf(format string, v ...interface{}) {
	if l.Level <= Debug {
		l.debug.Printf(format, v)
	}
}

func (l *logger) Debug(v ...interface{}) {
	if l.Level <= Debug {
		l.debug.Print(v)
	}
}

func (l *logger) Warnln(v ...interface{}) {
	if l.Level <= Warn {
		l.warn.Println(v...)
	}
}

func (l *logger) Warnf(format string, v ...interface{}) {
	if l.Level <= Warn {
		l.warn.Printf(format, v)
	}
}

func (l *logger) Warn(v ...interface{}) {
	if l.Level <= Warn {
		l.warn.Print(v)
	}
}

func (l *logger) Infoln(v ...interface{}) {
	if l.Level <= Info {
		l.info.Println(v...)
	}
}

func (l *logger) Infof(format string, v ...interface{}) {
	if l.Level <= Info {
		l.info.Printf(format, v)
	}
}

func (l *logger) Info(v ...interface{}) {
	if l.Level <= Info {
		l.info.Print(v)
	}
}

func (l *logger) Fatalln(v ...interface{}) {
	if l.Level <= Fatal {
		l.fatal.Fatalln(v...)
	}
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	if l.Level <= Fatal {
		l.fatal.Fatalf(format, v)
	}
}

func (l *logger) Fatal(v ...interface{}) {
	if l.Level <= Fatal {
		l.fatal.Fatal(v)
	}
}

func (l *logger) Errorln(v ...interface{}) {
	if l.Level <= Error {
		l.err.Println(v...)
	}
}

func (l *logger) Errorf(format string, v ...interface{}) {
	if l.Level <= Error {
		l.err.Printf(format, v)
	}
}

func (l *logger) Error(v ...interface{}) {
	if l.Level <= Error {
		l.err.Print(v)
	}
}
func newLogger() *logger {
	flag := 0
	aLogger := &logger{
		std:   stdlog.New(os.Stdout, color.New(color.FgCyan).SprintFunc()("Std: "), flag),
		info:  stdlog.New(os.Stdout, color.New(color.FgHiGreen).SprintFunc()("Info: "), flag),
		debug: stdlog.New(os.Stdout, color.New(color.FgBlue).SprintFunc()("Debug: "), flag),
		warn:  stdlog.New(os.Stdout, color.New(color.FgHiYellow).SprintFunc()("Warn: "), flag),
		err:   stdlog.New(os.Stderr, color.New(color.FgRed).SprintFunc()("Error: "), flag),
		fatal: stdlog.New(os.Stderr, color.New(color.FgHiRed).SprintFunc()("Fatal: "), flag),
		Level: Info,
	}
	return aLogger
}
