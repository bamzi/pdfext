// Package common contains common properties used by the subpackages.
package common

import (
	_a "fmt"
	_e "io"
	_b "os"
	_cb "path/filepath"
	_ad "runtime"
	_g "time"
)

const _bde = 10

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct {
	LogLevel LogLevel
	Output   _e.Writer
}

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_db ConsoleLogger) IsLogLevel(level LogLevel) bool { return _db.LogLevel >= level }

// Error does nothing for dummy logger.
func (DummyLogger) Error(format string, args ...interface{}) {}

// Info logs info message.
func (_gd ConsoleLogger) Info(format string, args ...interface{}) {
	if _gd.LogLevel >= LogLevelInfo {
		_ebg := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_gd.output(_b.Stdout, _ebg, format, args...)
	}
}

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger(logLevel LogLevel, writer _e.Writer) *WriterLogger {
	_ff := WriterLogger{Output: writer, LogLevel: logLevel}
	return &_ff
}

// Warning logs warning message.
func (_ce ConsoleLogger) Warning(format string, args ...interface{}) {
	if _ce.LogLevel >= LogLevelWarning {
		_cg := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_ce.output(_b.Stdout, _cg, format, args...)
	}
}
func _dbd(_fec _e.Writer, _gab string, _dfb string, _bgff ...interface{}) {
	_, _faf, _dg, _dcf := _ad.Caller(3)
	if !_dcf {
		_faf = "\u003f\u003f\u003f"
		_dg = 0
	} else {
		_faf = _cb.Base(_faf)
	}
	_abe := _a.Sprintf("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ", _gab, _faf, _dg) + _dfb + "\u000a"
	_a.Fprintf(_fec, _abe, _bgff...)
}

// Debug logs debug message.
func (_gge WriterLogger) Debug(format string, args ...interface{}) {
	if _gge.LogLevel >= LogLevelDebug {
		_gbe := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_gge.logToWriter(_gge.Output, _gbe, format, args...)
	}
}

// Error logs error message.
func (_baf WriterLogger) Error(format string, args ...interface{}) {
	if _baf.LogLevel >= LogLevelError {
		_bgf := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_baf.logToWriter(_baf.Output, _bgf, format, args...)
	}
}

const _dcff = "\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034"

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{ LogLevel LogLevel }

func (_gda WriterLogger) logToWriter(_cad _e.Writer, _aa string, _bfg string, _ga ...interface{}) {
	_dbd(_cad, _aa, _bfg, _ga)
}

// Error logs error message.
func (_bg ConsoleLogger) Error(format string, args ...interface{}) {
	if _bg.LogLevel >= LogLevelError {
		_eb := "\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020"
		_bg.output(_b.Stdout, _eb, format, args...)
	}
}

// Info does nothing for dummy logger.
func (DummyLogger) Info(format string, args ...interface{}) {}

// Notice logs notice message.
func (_fb ConsoleLogger) Notice(format string, args ...interface{}) {
	if _fb.LogLevel >= LogLevelNotice {
		_fa := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_fb.output(_b.Stdout, _fa, format, args...)
	}
}

// Logger is the interface used for logging in the  package.
type Logger interface {
	Error(_d string, _ec ...interface{})
	Warning(_ee string, _ge ...interface{})
	Notice(_gb string, _f ...interface{})
	Info(_bb string, _dc ...interface{})
	Debug(_ba string, _ed ...interface{})
	Trace(_fd string, _bc ...interface{})
	IsLogLevel(_df LogLevel) bool
}

// Warning logs warning message.
func (_fed WriterLogger) Warning(format string, args ...interface{}) {
	if _fed.LogLevel >= LogLevelWarning {
		_ca := "\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020"
		_fed.logToWriter(_fed.Output, _ca, format, args...)
	}
}

const _cgc = 15

// SetLogger sets 'logger' to be used by the library.
func SetLogger(logger Logger) { Log = logger }

var Log Logger = DummyLogger{}

// Info logs info message.
func (_bcc WriterLogger) Info(format string, args ...interface{}) {
	if _bcc.LogLevel >= LogLevelInfo {
		_cf := "\u005bI\u004e\u0046\u004f\u005d\u0020"
		_bcc.logToWriter(_bcc.Output, _cf, format, args...)
	}
}

// Trace logs trace message.
func (_cge WriterLogger) Trace(format string, args ...interface{}) {
	if _cge.LogLevel >= LogLevelTrace {
		_ebgg := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_cge.logToWriter(_cge.Output, _ebgg, format, args...)
	}
}

const _aaf = 30

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat(t _g.Time) string { return t.Format(_dcff) + "\u0020\u0055\u0054\u0043" }

const (
	LogLevelTrace   LogLevel = 5
	LogLevelDebug   LogLevel = 4
	LogLevelInfo    LogLevel = 3
	LogLevelNotice  LogLevel = 2
	LogLevelWarning LogLevel = 1
	LogLevelError   LogLevel = 0
)

// Warning does nothing for dummy logger.
func (DummyLogger) Warning(format string, args ...interface{}) {}

const _cc = 21

// IsLogLevel returns true from dummy logger.
func (DummyLogger) IsLogLevel(level LogLevel) bool { return true }

// Trace does nothing for dummy logger.
func (DummyLogger) Trace(format string, args ...interface{}) {}

// Notice does nothing for dummy logger.
func (DummyLogger) Notice(format string, args ...interface{}) {}

const _fbg = 2024

// Notice logs notice message.
func (_dd WriterLogger) Notice(format string, args ...interface{}) {
	if _dd.LogLevel >= LogLevelNotice {
		_geb := "\u005bN\u004f\u0054\u0049\u0043\u0045\u005d "
		_dd.logToWriter(_dd.Output, _geb, format, args...)
	}
}

// NewConsoleLogger creates new console logger.
func NewConsoleLogger(logLevel LogLevel) *ConsoleLogger { return &ConsoleLogger{LogLevel: logLevel} }
func (_feg ConsoleLogger) output(_bd _e.Writer, _bba string, _bf string, _ab ...interface{}) {
	_dbd(_bd, _bba, _bf, _ab...)
}

// LogLevel is the verbosity level for logging.
type LogLevel int

// DummyLogger does nothing.
type DummyLogger struct{}

// Debug does nothing for dummy logger.
func (DummyLogger) Debug(format string, args ...interface{}) {}

const Version = "\u0033\u002e\u0036\u0033\u002e\u0030"

var ReleasedAt = _g.Date(_fbg, _bde, _cc, _cgc, _aaf, 0, 0, _g.UTC)

// Trace logs trace message.
func (_ef ConsoleLogger) Trace(format string, args ...interface{}) {
	if _ef.LogLevel >= LogLevelTrace {
		_fe := "\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020"
		_ef.output(_b.Stdout, _fe, format, args...)
	}
}

// Debug logs debug message.
func (_gg ConsoleLogger) Debug(format string, args ...interface{}) {
	if _gg.LogLevel >= LogLevelDebug {
		_de := "\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020"
		_gg.output(_b.Stdout, _de, format, args...)
	}
}

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_dcg WriterLogger) IsLogLevel(level LogLevel) bool { return _dcg.LogLevel >= level }
