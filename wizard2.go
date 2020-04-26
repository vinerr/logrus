package logrus

import (
	"github.com/kr/pretty"
)

var (
	FieldsLogger = std
)

func SetFieldsLogger(loger *Logger) {
	loger.mu.Lock()
	defer loger.mu.Unlock()
	FieldsLogger = loger
}

func (f Fields) Tracef(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Tracef(format, args...)
}

func (f Fields) Debugf(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Debugf(format, args...)
}

func (f Fields) Infof(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Infof(format, args...)
}

func (f Fields) Printf(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Printf(format, args...)
}

func (f Fields) Warnf(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Warnf(format, args...)
}

func (f Fields) Warningf(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Warnf(format, args...)
}

func (f Fields) Errorf(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Errorf(format, args...)
}

func (f Fields) Fatalf(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Fatalf(format, args...)
}

func (f Fields) Panicf(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Panicf(format, args...)
}

func (f Fields) Trace(args ...interface{}) {
	FieldsLogger.WithFields(f).Trace(args...)
}

func (f Fields) Debug(args ...interface{}) {
	FieldsLogger.WithFields(f).Debug(args...)
}

func (f Fields) Info(args ...interface{}) {
	FieldsLogger.WithFields(f).Info(args...)
}

func (f Fields) Print(args ...interface{}) {
	FieldsLogger.WithFields(f).Info(args...)
}

func (f Fields) Warn(args ...interface{}) {
	FieldsLogger.WithFields(f).Warn(args...)
}

func (f Fields) Warning(args ...interface{}) {
	FieldsLogger.WithFields(f).Warn(args...)
}

func (f Fields) Error(args ...interface{}) {
	FieldsLogger.WithFields(f).Error(args...)
}

func (f Fields) Fatal(args ...interface{}) {
	FieldsLogger.WithFields(f).Fatal(args...)
}

func (f Fields) Panic(args ...interface{}) {
	FieldsLogger.WithFields(f).Panic(args...)
}

func (f Fields) Traceln(args ...interface{}) {
	FieldsLogger.WithFields(f).Traceln(args...)
}

func (f Fields) Debugln(args ...interface{}) {
	FieldsLogger.WithFields(f).Debugln(args...)
}

func (f Fields) Infoln(args ...interface{}) {
	FieldsLogger.WithFields(f).Infoln(args...)
}

func (f Fields) Println(args ...interface{}) {
	FieldsLogger.WithFields(f).Println(args...)
}

func (f Fields) Warnln(args ...interface{}) {
	FieldsLogger.WithFields(f).Warnln(args...)
}

func (f Fields) Warningln(args ...interface{}) {
	FieldsLogger.WithFields(f).Warnln(args...)
}

func (f Fields) Errorln(args ...interface{}) {
	FieldsLogger.WithFields(f).Errorln(args...)
}

func (f Fields) Fatalln(args ...interface{}) {
	FieldsLogger.WithFields(f).Fatalln(args...)
}

func (f Fields) Panicln(args ...interface{}) {
	FieldsLogger.WithFields(f).Panicln(args...)
}

// for fields to pretty printing for Go values
func (f Fields) Tracefp(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Tracef(format+" ==>%# v", pretty.Formatter(args))
}

func (f Fields) Debugfp(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Debugf(format+" ==>%# v", pretty.Formatter(args))
}

func (f Fields) Infofp(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Infof(format+" ==>%# v", pretty.Formatter(args))
}

func (f Fields) Printfp(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Printf(format+" ==>%# v", pretty.Formatter(args))
}

func (f Fields) Warnfp(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Warnf(format+" ==>%# v", pretty.Formatter(args))
}

func (f Fields) Warningfp(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Warnf(format+" ==>%# v", pretty.Formatter(args))
}

func (f Fields) Errorfp(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Errorf(format+" ==>%# v", pretty.Formatter(args))
}

func (f Fields) Fatalfp(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Fatalf(format+" ==>%# v", pretty.Formatter(args))
}

func (f Fields) Panicfp(format string, args ...interface{}) {
	FieldsLogger.WithFields(f).Panicf(format+" ==>%# v", pretty.Formatter(args))
}

// exported, for logrus to pretty printing for Go values
func Tracefp(format string, args ...interface{}) {
	FieldsLogger.Tracef(format+" ==>%# v", pretty.Formatter(args))
}

// Debugf logs a message at level Debug on the standard logger.
func Debugfp(format string, args ...interface{}) {
	FieldsLogger.Debugf(format+" ==>%# v", pretty.Formatter(args))
}

// Printf logs a message at level Info on the standard logger.
func Printfp(format string, args ...interface{}) {
	FieldsLogger.Printf(format+" ==>%# v", pretty.Formatter(args))
}

// Infof logs a message at level Info on the standard logger.
func Infofp(format string, args ...interface{}) {
	FieldsLogger.Infof(format+" ==>%# v", pretty.Formatter(args))
}

// Warnf logs a message at level Warn on the standard logger.
func Warnfp(format string, args ...interface{}) {
	FieldsLogger.Warnf(format+" ==>%# v", pretty.Formatter(args))
}

// Warningf logs a message at level Warn on the standard logger.
func Warningfp(format string, args ...interface{}) {
	FieldsLogger.Warnf(format+" ==>%# v", pretty.Formatter(args))
}

// Errorf logs a message at level Error on the standard logger.
func Errorfp(format string, args ...interface{}) {
	FieldsLogger.Errorf(format+" ==>%# v", pretty.Formatter(args))
}

// Panicf logs a message at level Panic on the standard logger.
func Panicfp(format string, args ...interface{}) {
	FieldsLogger.Panicf(format+" ==>%# v", pretty.Formatter(args))
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalfp(format string, args ...interface{}) {
	FieldsLogger.Fatalf(format+" ==>%# v", pretty.Formatter(args))
}

// for logger to pretty printing for Go values
func (logger *Logger) Tracefp(format string, args ...interface{}) {
	FieldsLogger.Trace(format+" ==>%# v", pretty.Formatter(args))
}

func (logger *Logger) Debugfp(format string, args ...interface{}) {
	FieldsLogger.Debug(format+" ==>%# v", pretty.Formatter(args))
}

func (logger *Logger) Infofp(format string, args ...interface{}) {
	FieldsLogger.Infof(format+" ==>%# v", pretty.Formatter(args))
}

func (logger *Logger) Printfp(format string, args ...interface{}) {
	FieldsLogger.Printf(format+" ==>%# v", pretty.Formatter(args))
}

func (logger *Logger) Warnfp(format string, args ...interface{}) {
	FieldsLogger.Warnf(format+" ==>%# v", pretty.Formatter(args))
}

func (logger *Logger) Warningfp(format string, args ...interface{}) {
	FieldsLogger.Warnf(format+" ==>%# v", pretty.Formatter(args))
}

func (logger *Logger) Errorfp(format string, args ...interface{}) {
	FieldsLogger.Errorf(format+" ==>%# v", pretty.Formatter(args))
}

func (logger *Logger) Fatalfp(format string, args ...interface{}) {
	FieldsLogger.Fatalf(format+" ==>%# v", pretty.Formatter(args))
}

func (logger *Logger) Panicfp(format string, args ...interface{}) {
	FieldsLogger.Panicf(format+" ==>%# v", pretty.Formatter(args))
}
