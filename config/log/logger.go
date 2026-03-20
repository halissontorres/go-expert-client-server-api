package log

import (
	"io"
	"log"
	"os"
	"strings"
)

// Logger fornece logging estruturado por níveis, com suporte a debug, info, warning e error.
type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	fatal   *log.Logger
	writer  io.Writer
}

// NewLogger retorna uma nova instância de Logger com níveis de log pré-configurados e um writer de saída.
func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)

	basePrefix := ""
	if p != "" {
		basePrefix = strings.ToUpper(p) + "-"
	}
	flags := log.Ldate | log.Ltime

	return &Logger{
		debug:   log.New(writer, basePrefix+"DEBUG: ", flags),
		info:    log.New(writer, basePrefix+"INFO: ", flags),
		warning: log.New(writer, basePrefix+"WARNING: ", flags),
		err:     log.New(writer, basePrefix+"ERROR: ", flags),
		fatal:   log.New(writer, basePrefix+"FATAL: ", flags),
		writer:  writer,
	}
}

// Registra uma mensagem formatada em nível debug usando o logger interno de depuração.
func (l *Logger) Debug(format string, v ...any) {
	l.debug.Printf(format, v...)
}

// Registra uma mensagem formatada em nível info usando o logger interno de depuração.
func (l *Logger) Info(format string, v ...any) {
	l.info.Printf(format, v...)
}

// Registra uma mensagem formatada em nível warning usando o logger interno de depuração.
func (l *Logger) Warn(format string, v ...any) {
	l.warning.Printf(format, v...)
}

// Registra uma mensagem formatada em nível error usando o logger interno de depuração.
func (l *Logger) Error(format string, v ...any) {
	l.err.Printf(format, v...)
}

func (l *Logger) Fatal(format string, v ...any) {
	l.err.Fatalf(format, v...)
}
