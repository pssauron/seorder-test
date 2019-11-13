//============================================================
// 描述: 
// 作者: Simon
// 日期: 2019/11/13 3:17 下午
//
//============================================================

package main

import (
	"flag"
	"github.com/labstack/echo"
	"github.com/pssauron/log4go"
	"github.com/seorder-test/cfgs"
	"github.com/seorder-test/routes"
	"github.com/seorder-test/storages"
	"runtime"
)

func main() {

	// 设置 运行cpu 数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 解析 命令行 输入的参数    seorder.exe -l ./log.xml -c ./cfg.yaml
	flag.Parse()

	//初始化日志

	initLog()

	//序列化配置文件
	initCfg()

	// 序列化数据对象操作数据库
	storages.InitStorage(cfgs.Cfg.MssqlConf)

	e  := echo.New()

	routes.InitHandlers(e)

	if err := e.Start(":9999");err != nil {
		log4go.Error(err)
		panic("http api 服务器启动异常")
	}

}


// 初始化日志
func initLog() {
	var logPath string
	flag.StringVar(&logPath,"l","./log.xml","日志配置文件地址")

	log4go.LoadConfiguration(logPath)

}

func initCfg() {
	var cfgPath string
	flag.StringVar(&cfgPath,"c","./cfg.yaml","基础配置文件地址")

	cfgs.InitCfg(cfgPath)
}




