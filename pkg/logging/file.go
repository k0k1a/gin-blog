package logging

import (
	"fmt"
	"github.com/k0k1a/go-gin-example/pkg/setting"
	"log"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filepath string) *os.File {
	_, err := os.Stat(filepath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission:%v", err)
	}
	handle, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Fail to OpenFile:%v", err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	if err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm); err != nil {
		panic(err)
	}

}
