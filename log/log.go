package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger
var Dir = "./log/"

func init() {
	GenerateLogger()
}

func SetLogDir(dir string) {
	Dir = dir
	GenerateLogger()
}

func GenerateLogger() {
	timeFormat := "2006-01-02 15:04:05"
	zerolog.TimeFieldFormat = timeFormat

	// create log dir
	err := os.MkdirAll(Dir, os.ModePerm)
	if err != nil {
		fmt.Println("Mkdir failed, err:", err)
		return
	}

	fileName := Dir + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	consoleWriter.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s;", i)
	}
	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)
	Logger = zerolog.New(multi).With().Timestamp().Logger()
}
