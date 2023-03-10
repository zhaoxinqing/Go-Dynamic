package logger

import (
	"io"
	"log"
	"os"
)

const (
	LOG_FILE = "./log/backend_finance.log" // log file
)

func WriteLogToFile() {
	logFile, err := os.OpenFile(LOG_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	defer func() {
		logFile.Close()
	}()

	// 组合一下即可，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[agent-purchase]")
	log.SetFlags(log.LstdFlags | log.Llongfile | log.LUTC)
}
