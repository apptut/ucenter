package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
	"github.com/tabalt/gracehttp"
	"path"
	"runtime"
	"ucenter/app"
	"ucenter/app/routes"
	"ucenter/app/svc/logger"
)

func parseArgs() string {
	prjHome := flag.String("prjHome", "", "project home dir")
	flag.Parse()

	if *prjHome == "" {
		_, filePath, _, ok := runtime.Caller(0)
		if ok {
			return path.Dir(filePath)
		}

		fmt.Println(aurora.Red("\n😡 params error"))
		flag.PrintDefaults()
	}

	return *prjHome
}

func main() {
	// 解析参数
	prjHome := parseArgs()

	// 初始化配置
	err := app.Init(prjHome)
	if err != nil {
		app.Destruct()
		logger.Fatal(err)
	}

	// 配置log
	if app.IsProd() {
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
	}

	err = gracehttp.ListenAndServe(app.Config.Http.Listen, routes.New())
	if err != nil {
		app.Destruct()
		logger.Fatal(err)
	}
}
