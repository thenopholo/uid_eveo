package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "\033[32mDEBUG:\033[0m ", logger.Flags()),
		info:    log.New(writer, "\033[34mINFO:\033[0m ", logger.Flags()),
		warning: log.New(writer, "\033[33mWARNING:\033[0m ", logger.Flags()),
		err:     log.New(writer, "\033[31mERROR:\033[0m ", logger.Flags()),
		writer:  writer,
}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Debugf(f string, v ...interface{}) {
	l.debug.Printf(f, v...)
}
func (l *Logger) Infof(f string, v ...interface{}) {
	l.info.Printf(f, v...)
}
func (l *Logger) Warningf(f string, v ...interface{}) {
	l.warning.Printf(f, v...)
}
func (l *Logger) Errorf(f string, v ...interface{}) {
	l.err.Printf(f, v...)
}