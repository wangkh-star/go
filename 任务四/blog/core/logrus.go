package core

import (
	"fmt"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据日志级别设置颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	// 获取调用者的文件信息和行号
	var fileInfo string
	if entry.HasCaller() {
		file := entry.Caller.File
		line := entry.Caller.Line
		// 缩短文件路径，只显示最后两部分
		if parts := strings.Split(file, "/"); len(parts) > 2 {
			file = strings.Join(parts[len(parts)-2:], "/")
		}
		fileInfo = fmt.Sprintf(" \033[%dm(%s:%d)\033[0m", gray, file, line)
	}

	// 格式化时间
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	// 构建日志消息
	levelStr := fmt.Sprintf("\033[%dm[%s]\033[0m", levelColor, strings.ToUpper(entry.Level.String()))
	msg := fmt.Sprintf("%s %s %s", timestamp, levelStr, entry.Message)

	// 如果有字段，添加到日志中
	if len(entry.Data) > 0 {
		msg += " " + fmt.Sprint(entry.Data)
	}

	// 添加文件信息
	if fileInfo != "" {
		msg += fileInfo
	}

	return []byte(msg + "\n"), nil
}

var once sync.Once

// InitGlobalLogger 初始化全局日志配置
func InitGlobalLogger(level logrus.Level) {
	once.Do(func() {
		// 设置日志级别
		logrus.SetLevel(level)

		// 启用调用者信息（显示文件和行号）
		logrus.SetReportCaller(true)

		// 只设置自定义格式化器，不要重复设置
		logrus.SetFormatter(&LogFormatter{})
	})
}

// 安全的快捷函数
func Debug(args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		logrus.Debug(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		logrus.Debugf(format, args...)
	}
}

func Info(args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.InfoLevel) {
		logrus.Info(args...)
	}
}

func Infof(format string, args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.InfoLevel) {
		logrus.Infof(format, args...)
	}
}

func Warn(args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.WarnLevel) {
		logrus.Warn(args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.WarnLevel) {
		logrus.Warnf(format, args...)
	}
}

func Error(args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.ErrorLevel) {
		logrus.Error(args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.ErrorLevel) {
		logrus.Errorf(format, args...)
	}
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// WithField 添加字段到日志
func WithField(key string, value interface{}) *logrus.Entry {
	return logrus.WithField(key, value)
}

// WithFields 添加多个字段到日志
func WithFields(fields map[string]interface{}) *logrus.Entry {
	return logrus.WithFields(fields)
}

// GetLogger 获取一个带有模块名的logger
func GetLogger(module string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"module": module,
	})
}
