package router

import (
	"database/sql"
	"myapp/config"
	"myapp/controller"
	"myapp/cors"
	"myapp/repository"
	"myapp/usecase"

	"github.com/Abhi-singh-karuna/my_Liberary/baselogger"
	"github.com/Abhi-singh-karuna/my_Liberary/cachehandler"
	email "github.com/Abhi-singh-karuna/my_Liberary/email"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewRouter(db *sql.DB,cfg *config.Config, validate *validator.Validate, logger *baselogger.BaseLogger, emailService *email.EmailService, redisClient cachehandler.CacheHandler) *gin.Engine {
	r := gin.Default()

	// Apply CORS middleware
	r.Use(cors.SetupCORS())

	userRepo := repository.NewRepository(db, logger, cfg, redisClient)
	userUseCase := usecase.NewUserInteractor(userRepo, logger, cfg, emailService)
	userController := controller.NewUserController(userUseCase, validate, logger, cfg, redisClient)

	// Only Auth-Check Middleware

	// Auth and verified  Middleware
	verifiedUser := r.Group("/")

	// Public routes
	verifiedUser.POST("/password", userController.AddPassword)

	verifiedUser.POST("/count", userController.CountVisitWebsite)

	verifiedUser.GET("/biodata", userController.GetBioDataTrackerInfo)

	verifiedUser.POST("/weekly", userController.GetWeeklyData)

	verifiedUser.POST("/buffer", userController.GetPageBufferPercentages)

	// subscribe 

	verifiedUser.POST("/subscribe", userController.Subscribe)

	// get all subscribers
	verifiedUser.POST("/subscribers", userController.GetAllSubscribers)

	verifiedUser.POST("/dashboard-data", userController.CalculatePercentageChange)

	// sp_GetCountsWithPercentage

	verifiedUser.POST("/count-data", userController.GetCountsWithPercentage)

	return r
}
