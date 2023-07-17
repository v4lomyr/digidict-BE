package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/v4lomyr/digidict-BE/internal/config"
	"github.com/v4lomyr/digidict-BE/internal/constants"
	"github.com/v4lomyr/digidict-BE/internal/helper"

	"github.com/sirupsen/logrus"
)

type logrusEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	log *logrus.Logger
}

// setting up logrus as our logger
func SetLogrusLogger(cfg *config.Config) {
	log := logrus.New()

	// let's say that in production, we want to output our log into a file called "log.txt"
	// otherwise, it will only be output to standard output
	if cfg.App.Environment == constants.AppEnvironmentProduction {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})

		f, err := os.OpenFile(filepath.Join(helper.GetRootPath(), "log", "log.txt"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("ERROR OPENING LOG FILE")
		}
		log.SetOutput(f)
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
			ForceColors:     true,
		})
		log.SetReportCaller(true)
		log.SetOutput(os.Stdout)
	}

	logrusLogger := &logrusLogger{log: log}

	SetLogger(logrusLogger)
}

func (l *logrusLogger) GetWriter() io.Writer {
	return l.log.Out
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.log.Fatalf(format, args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *logrusLogger) WithFields(args map[string]interface{}) Logger {
	return &logrusEntry{
		entry: l.log.WithFields(args),
	}
}

func (l *logrusLogger) Printf(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

func (l *logrusEntry) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *logrusEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logrusEntry) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logrusEntry) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *logrusEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusEntry) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *logrusEntry) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *logrusEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logrusEntry) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *logrusEntry) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *logrusEntry) WithFields(args map[string]interface{}) Logger {
	return &logrusEntry{
		entry: l.entry.WithFields(args),
	}
}

func (l *logrusEntry) GetWriter() io.Writer {
	return l.entry.Logger.Out
}

func (l *logrusEntry) Printf(format string, args ...interface{}) {
	l.entry.Printf(format, args...)
}
