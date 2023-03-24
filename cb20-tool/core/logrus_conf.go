package core_lib

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func InitLogrus() {
	// var log = logrus.New()
	log := &logrus.Logger{
		Out:          os.Stdout,
		Level:        logrus.TraceLevel,
		ReportCaller: true,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}
	// log.Formatter = new(logrus.TextFormatter)                     //default
	// log.Formatter.(*logrus.TextFormatter).DisableColors = false // remove colors
	// log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	// log.Level = logrus.TraceLevel
	// log.SetReportCaller(true)

	logFile := "build/log/outputlog.txt"

	logf, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile: " + logFile)
		panic(err)
	}

	defer logf.Close()
	mw := io.MultiWriter(os.Stdout, logf)
	log.SetOutput(mw)
	// log.Out = os.Stdout
	log.Info("Init Logrus successfully")
}
