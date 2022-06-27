package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Logger struct {
	Logger *logrus.Logger
	file   *os.File
}

func InitializeLogger(serviceName string, typeMessage string) *Logger {
	logger := &Logger{}
	if err := CreateLogDir(filepath.FromSlash("../logfiles/" + serviceName)); err != nil {
		logrus.Fatalf("Failed to create directory for log files | %v\n", err)
	}

	file := filepath.FromSlash("./logfiles/" + serviceName + "/" + typeMessage + ".log")

	rotatingLogs, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   file,
		MaxSize:    100, //MB
		MaxBackups: 50,
		MaxAge:     1, //DAYS
		Level:      logrus.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%lvl% | %time% | %msg% \n",
		},
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize rotating hook | %v\n", err)
	}
	logger.Logger = logrus.New()
	logger.Logger.AddHook(rotatingLogs)
	logger.Logger.SetReportCaller(true)
	logger.Logger.SetOutput(os.Stdout)
	logger.Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return frame.Function, fmt.Sprintf("%s:%d", FormatFilePath(frame.File), frame.Line)
		}})
	logger.Logger.SetLevel(logrus.InfoLevel)

	return &Logger{Logger: logger.Logger, file: logger.file}
}

func FormatFilePath(file string) string {
	arr := strings.Split(filepath.ToSlash(file), "/")
	return arr[len(arr)-1]
}

func CreateLogDir(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return os.MkdirAll(filePath, os.ModeDir|0777)
	}
	return nil
}

func (l *Logger) CloseLogger() {
	l.file.Close()
}

func (l *Logger) ErrorMessage(message string) {
	l.Logger.Error(message)
}

func (l *Logger) InfoMessage(message string) {
	l.Logger.Info(message)
}

func (l *Logger) FatalMessage(message string) {
	l.Logger.Fatal(message)
}
