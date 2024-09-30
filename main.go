package main

import (
	"blogGo/conf"
	_ "blogGo/docs"
	"blogGo/src/model"
	"blogGo/src/utils"
	"blogGo/src/views"
	"context"
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"
	"time"
)

func setMode() {
	ginConf := conf.CFG.GinConf

	if ginConf.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

var ctx = context.Background()

func main() {
	defer func() {
		if err := recover(); err != nil {
			utils.Logging.Error(err)
		}
	}()

	conf.InitConfig()

	r := gin.Default()
	logger := utils.GetLogger()
	r.Use(ginlogrus.Logger(logger), gin.Recovery())
	setMode()
	// 注册 Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	model.InitDB()
	utils.InitRedis(ctx)

	sqlDB, _ := model.DB.DB()
	// 显式闭包 关闭数据库
	defer func(sqlDB *sql.DB) {
		if err := sqlDB.Close(); err != nil {
			panic(err)
		}
	}(sqlDB)
	// 显式闭包 关闭redis
	defer func(RC *redis.Client) {
		if err := RC.Close(); err != nil {
			panic(err)
		}
	}(utils.RC)

	router := r.Group("/api/v1")
	views.AuthViewRoutes(router)
	views.BlogRouter(router)

	// 跨域问题
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://github.com", "https://shinestar.fun"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	_ = r.Run()
}
