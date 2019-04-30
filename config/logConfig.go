package config

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

var Logger *logs.BeeLogger

func Logconf() {
	logconf := ConfigMap["logconf"]
	bytes, err := json.Marshal(logconf)
	if err != nil {
		fmt.Println("json is err:", err)
		return
	}
	Logger = logs.NewLogger(10000)
	Logger.SetLogger(logs.AdapterFile, string(bytes))
	Logger.EnableFuncCallDepth(true)

}
