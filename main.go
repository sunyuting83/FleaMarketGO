package main

import (
	"FleaMarket/router"
	"FleaMarket/utils"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	OS := runtime.GOOS
	CurrentPath, _ := utils.GetCurrentPath()

	confYaml, err := utils.CheckConfig(OS, CurrentPath)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Duration(10) * time.Second)
		os.Exit(0)
	}

	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)
	app := router.InitRouter(confYaml.SECRET_KEY)

	app.Run(strings.Join([]string{":", confYaml.Port}, ""))
}
