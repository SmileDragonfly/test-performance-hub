package logger

import (
	"encoding/json"
	"errors"
	"os"
)

type ILogger interface {
	Debug(a ...any)
	Info(a ...any)
	Warning(a ...any)
	Error(a ...any)
	Debugf(template string, a ...any)
	Infof(template string, a ...any)
	Warningf(template string, a ...any)
	Errorf(template string, a ...any)
	DebugFunc(a ...any)
	InfoFunc(a ...any)
	WarningFunc(a ...any)
	ErrorFunc(a ...any)
	DebugFuncf(template string, a ...any)
	InfoFuncf(template string, a ...any)
	WarningFuncf(template string, a ...any)
	ErrorFuncf(template string, a ...any)
}

var instance ILogger

type LoggerConfig struct {
	Type       string
	FileName   string
	MaxSize    int // MB
	MaxBackups int
	MaxAge     int
	Level      int // Debug/Info/Warning/Error = -1/0/1/2
}

const (
	ZAP_LOGGER string = "zap"
)

func NewLogger(sPath string) error {
	// Init logger
	byteCfg, err := os.ReadFile(sPath)
	if err != nil {
		panic(err)
	}
	var logCfg LoggerConfig
	err = json.Unmarshal(byteCfg, &logCfg)
	if err != nil {
		panic(err)
	}
	switch logCfg.Type {
	case ZAP_LOGGER:
		logger, err := NewZapLogger(logCfg)
		if err != nil {
			return err
		}
		instance = logger
		break
	default:
		return errors.New("Invalid logger type")
	}
	instance.Info("==================================================")
	instance.Info("Start logger succesfully")
	return nil
}

func Debug(a ...any) {
	instance.Debug(a...)
}
func Info(a ...any) {
	instance.Info(a...)
}
func Warning(a ...any) {
	instance.Warning(a...)
}
func Error(a ...any) {
	instance.Error(a...)
}
func Debugf(template string, a ...any) {
	instance.Debugf(template, a...)
}
func Infof(template string, a ...any) {
	instance.Infof(template, a...)
}
func Warningf(template string, a ...any) {
	instance.Warningf(template, a...)
}
func Errorf(template string, a ...any) {
	instance.Errorf(template, a...)
}

func DebugFunc(a ...any) {
	instance.DebugFunc(a...)
}
func InfoFunc(a ...any) {
	instance.InfoFunc(a...)
}
func WarningFunc(a ...any) {
	instance.WarningFunc(a...)
}
func ErrorFunc(a ...any) {
	instance.ErrorFunc(a...)
}
func DebugFuncf(template string, a ...any) {
	instance.DebugFuncf(template, a...)
}
func InfoFuncf(template string, a ...any) {
	instance.InfoFuncf(template, a...)
}
func WarningFuncf(template string, a ...any) {
	instance.WarningFuncf(template, a...)
}
func ErrorFuncf(template string, a ...any) {
	instance.ErrorFuncf(template, a...)
}
