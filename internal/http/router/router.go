package router

import (
	"io"
	api "laboratory/api/login"
	middleware "laboratory/internal/http/middleware"
	"laboratory/log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	workdir, _ := os.Getwd()
	logfile, err := os.Create(workdir + "/log/logs/gin_http.log")
	if err != nil {
		log.SugarLogger.Errorf("配置ginhttp日志失败:%v", err)
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(logfile)
	r := gin.Default()
	r.Use(middleware.Cors())

	api1 := r.Group("/api")

	api1.GET("/auth", api.GetAuthCode)
	api1.POST("/register/teacher", api.RegisterTeacher)
	api1.POST("/register/student", api.RegisterStudent)
	api1.POST("/login", api.LoginByPwd)
	api1.POST("/login2", api.LoginByAuthCode)

	la := api1.Group("/identify")

	la.Use(middleware.Auth())

	return r
}
