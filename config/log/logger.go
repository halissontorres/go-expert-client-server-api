package log

import (
	"io"
	"log"
	"os"
)

// Logger fornece logging estruturado por níveis, com suporte a debug, info, warning e error.
type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

// NewLogger retorna uma nova instância de Logger com níveis de log pré-configurados e um writer de saída.
func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
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
