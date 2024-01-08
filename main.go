package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	docs "shift-scheduler-service/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"shift-scheduler-service/config"
	"shift-scheduler-service/internal/handlers"
	"shift-scheduler-service/pkg/db/postgres"
	"shift-scheduler-service/pkg/db/redis"
	"shift-scheduler-service/pkg/logger"
)

// Path: ShiftSchedulerService/main.go
// @Title ShiftSchedulerService API
// @Description:
// @Version 1.0.0
// @Schemes http https
// @BasePath /api-shifts

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

var isConfigSuccess = false

// var equals string = strings.Repeat("=", 50)

const (
	// APP_NAME = "localhost:9097/shift-scheduler-service/"
	APP_NAME = "shift-scheduler-service"
)

func main() {
	mode := config.C.App.Mode
	port := config.C.App.Port

	router := gin.New()
	router.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter))
	inAppCache := redis.NewInAppCacheStore(time.Minute)
	cacheConn, cacheContext := redis.NewRedisCacheConnection(config.C.Cache.Url)
	dbConn := postgres.NewPostgresDB(config.C.DB.Url)

	jaegerCfgInstance := jaegercfg.Configuration{
		ServiceName: config.C.Jaeger.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           config.C.Jaeger.LogSpans,
			LocalAgentHostPort: config.C.Jaeger.Host,
		},
	}

	tracer, closer, err := jaegerCfgInstance.NewTracer(
		jaegercfg.Logger(jaegerlog.StdLogger),
		jaegercfg.Metrics(metrics.NullFactory),
	)

	if err != nil {
		log.Fatal("cannot create tracer", err)
	}

	// create application service
	shiftsvc := handlers.NewShiftService(
		inAppCache,
		cacheConn,
		cacheContext,
		dbConn,
	)

	// check env and set gin mode
	setApplicationMode(mode, router)
	shiftsvc.InitRouter(router)

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	logger.CLogger.Info("Opentracing connected")

	// run application
	logger.CLogger.Info("INIT: Application " + APP_NAME + " started on port " + port)
	logger.CLogger.Info(router.Run(":" + port))
}

// Initialize Application
func init() {
	isConfigSuccess = configureApplication()
	if !isConfigSuccess {
		logger.CLogger.Error("INIT: Application configuration failed.")
		os.Exit(1)
	} else {
		logger.CLogger.Info("INIT: Application configuration success.")
	}
}

// Configure Application with config file
func configureApplication() bool {
	dir, err := os.Getwd()
	if err != nil {
		logger.CLogger.Error("INIT: Cannot get current working directory os.Getwd()")
		return false
	} else {
		config.ReadConfig(dir)
		return true
	}
}

// Set Application Mode
func setApplicationMode(md string, router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	if md == "prod" || md == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		router.Use(gin.Logger())
		// logger time gonna be Istanbul
		gin.SetMode(gin.DebugMode)
	}

	// check env and set swagger
	if !(md == "prod" || md == "production") {
		docs.SwaggerInfo.BasePath = handlers.API_PREFIX
		router.GET(handlers.API_PREFIX+"/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
