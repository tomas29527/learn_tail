package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"strings"
)

var ConfigMap map[string]map[string]interface{} = make(map[string]map[string]interface{}, 20)

func init() {
	configer, err := config.NewConfig("ini", "./resource/app.ini")
	if err != nil {
		fmt.Println("open ini file is err:", err)
	}
	level := strings.ToLower(configer.String("log::level"))
	fileName := configer.String("log::logPath")

	logconf := make(map[string]interface{})
	if level == "debug" {
		logconf["level"] = logs.LevelDebug
	} else if level == "info" {
		logconf["level"] = logs.LevelInformational
	}
	logconf["filename"] = fileName
	ConfigMap["logconf"] = logconf
}
