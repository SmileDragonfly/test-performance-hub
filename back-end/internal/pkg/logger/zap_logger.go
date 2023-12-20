package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strings"
	"time"
)

type ZapLogger struct {
	config LoggerConfig
	logger *zap.SugaredLogger
}

func addCaller(a ...any) []any {
	pc, _, _, _ := runtime.Caller(3)
	funcName := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	sRet := fmt.Sprintf("%s:", funcName[len(funcName)-1])
	arr := append([]any{sRet}, a...)
	return arr
}

func addCallerf(template string) string {
	pc, _, _, _ := runtime.Caller(3)
	funcName := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	sRet := fmt.Sprintf("%s:", funcName[len(funcName)-1])
	sRet = sRet + " " + template
	return sRet
}

func (z ZapLogger) DebugFunc(a ...any) {
	z.Debug(addCaller(a...))
}

func (z ZapLogger) InfoFunc(a ...any) {
	z.Info(addCaller(a...))
}

func (z ZapLogger) WarningFunc(a ...any) {
	z.Warning(addCaller(a...))
}

func (z ZapLogger) ErrorFunc(a ...any) {
	z.Error(addCaller(a...))
}

func (z ZapLogger) DebugFuncf(template string, a ...any) {
	z.Debugf(addCallerf(template), a...)
}

func (z ZapLogger) InfoFuncf(template string, a ...any) {
	z.Infof(addCallerf(template), a...)
}

func (z ZapLogger) WarningFuncf(template string, a ...any) {
	z.Warningf(addCallerf(template), a...)
}

func (z ZapLogger) ErrorFuncf(template string, a ...any) {
	z.Errorf(addCallerf(template), a...)
}

func NewZapLogger(config LoggerConfig) (*ZapLogger, error) {
	// lumberjack.Instance is already safe for concurrent use, so we don't need to
	// lock it.
	encoderCfg := zapcore.EncoderConfig{ //Cấu hình logging, sẽ không có stacktracekey
		MessageKey:   "message",
		TimeKey:      "time",
		LevelKey:     "level",
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder, //Lấy dòng code bắt đầu log
		EncodeLevel:  CustomLevelEncoder,        //Format cách hiển thị level log
		EncodeTime:   SyslogTimeEncoder,         //Format hiển thị thời điểm log
	}

	// Create a writer that logs to the console (os.Stdout)
	consoleDebugging := zapcore.Lock(os.Stdout)
	// Create a writer that logs to a file using lumberjack logger
	logRotation := &lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, // days
	}
	//rotator, err := rotatelogs.New(
	//	config.FileName,
	//	rotatelogs.WithMaxAge(60*24*time.Hour),
	//	rotatelogs.WithRotationTime(time.Hour))
	//if err != nil {
	//	panic(err)
	//}
	fileLogger := zapcore.AddSync(logRotation)
	// Combine both writers into a MultiWriteSyncer
	writeSyncer := zapcore.NewMultiWriteSyncer(consoleDebugging, zapcore.AddSync(fileLogger))
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		writeSyncer,
		zapcore.Level(config.Level),
	)
	logger := zap.New(core)
	// Schedule the first log rotation at midnight
	scheduleLogRotationMidnight(logRotation)
	return &ZapLogger{config: config, logger: logger.Sugar()}, nil
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func scheduleLogRotationMidnight(logRotate *lumberjack.Logger) {
	now := time.Now()
	nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	timeUntilMidnight := nextMidnight.Sub(now)
	// Schedule log rotation at midnight
	go func() {
		timer := time.NewTimer(timeUntilMidnight)
		defer timer.Stop()

		for {
			<-timer.C
			// Rotate logs at midnight
			err := logRotate.Rotate()
			if err != nil {
				fmt.Println("Error rotating logs:", err)
			}
			// Reschedule for the next midnight
			now = time.Now()
			nextMidnight = time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			timeUntilMidnight = nextMidnight.Sub(now)
			timer.Reset(timeUntilMidnight)
		}
	}()
}

func (z ZapLogger) Debug(a ...any) {
	z.logger.Debug(a...)
}

func (z ZapLogger) Info(a ...any) {
	z.logger.Info(a...)
}

func (z ZapLogger) Warning(a ...any) {
	z.logger.Warn(a...)
}
func (z ZapLogger) Error(a ...any) {
	z.logger.Error(a...)
}

func (z ZapLogger) Debugf(template string, a ...any) {
	z.logger.Debugf(template, a...)
}

func (z ZapLogger) Infof(template string, a ...any) {
	z.logger.Infof(template, a...)
}

func (z ZapLogger) Warningf(template string, a ...any) {
	z.logger.Warnf(template, a...)
}
func (z ZapLogger) Errorf(template string, a ...any) {
	z.logger.Errorf(template, a...)
}
