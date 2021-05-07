package logger

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogHandler interface{
	LogInfo(s string)
	LogError(s string)
	LogFatal(s string)
}

type Zaploger struct {
	Log *zap.SugaredLogger
}

func NewZapLoger() *Zaploger {

	writerSyncer := getLogWriter()
	encoder := getEncoder()		
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	l := zap.New(core, zap.AddCaller())

	return &Zaploger{
		Log: l.Sugar(),
	}
}

func (z *Zaploger) LogInfo(warning string) {
	z.Log.Info(warning)
}

func (z *Zaploger) LogError(err string) {
	z.Log.Error(err)
}

func (z *Zaploger) LogFatal(fatal string) {
	z.Log.Fatal(fatal)
}

func getLogWriter() zapcore.WriteSyncer {
	cwd := os.Getenv("WorkStation")
	filename := filepath.Join(cwd, "test.log")
		lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
