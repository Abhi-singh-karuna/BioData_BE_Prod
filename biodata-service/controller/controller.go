package controller

import (
	"fmt"
	"myapp/config"
	"myapp/model"
	"myapp/usecase"
	"net/http"

	"github.com/Abhi-singh-karuna/my_Liberary/baselogger" // Import the logger package
	"github.com/Abhi-singh-karuna/my_Liberary/cachehandler"
	"github.com/gin-gonic/gin"

	_http "github.com/Abhi-singh-karuna/my_Liberary/http"
	httpErrors "github.com/Abhi-singh-karuna/my_Liberary/http/errors"

	"github.com/go-playground/validator/v10"
)

type Controller struct {
	userUseCase usecase.UserUseCase
	validator   *validator.Validate
	logger      *baselogger.BaseLogger
	cfg         *config.Config
	redisClient cachehandler.CacheHandler
}

func NewUserController(userUseCase usecase.UserUseCase, validator *validator.Validate, logger *baselogger.BaseLogger, cfg *config.Config, redisClient cachehandler.CacheHandler) *Controller {
	return &Controller{userUseCase, validator, logger, cfg, redisClient}
}

func ParseMessage(errorMessage string) model.ErrorMessage {
	return model.ErrorMessage{
		Message: errorMessage,
	}
}

const HeaderId = "HeaderId"

// Middleware Controller
// Validate User Is Verified
func (ctrl *Controller) ValidateUserVerified(userId string) (bool, *model.User, error) {
	return ctrl.userUseCase.ValidateUserVerified(userId)
}

// Add New Password
func (c *Controller) AddPassword(ctx *gin.Context) {

	headerId, _ := ctx.Get(HeaderId)
	headerIdValue, _ := headerId.(model.HeaderId)

	fmt.Println("headerIdValue.userid -------", headerIdValue.USER_ID)

	var pInfo = &model.PasswordReq{}
	pInfo.User_Id = headerIdValue.USER_ID

	if err := _http.ReadRequest(ctx, pInfo); err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(http.StatusBadRequest, ParseMessage(err.Error()))
		return
	}

	createdUser, err := c.userUseCase.AddPassword(pInfo)
	if err != nil {
		_http.LogResponseError(ctx, err)

		ctx.JSON(httpErrors.ErrorResponse(err))
		return
	}
	c.logger.Infof("Password created: %v", pInfo)

	ctx.JSON(http.StatusOK, createdUser)
}

func (c *Controller) CountVisitWebsite(ctx *gin.Context) {

	var count model.BioDataCount

	if err := _http.ReadRequest(ctx, &count); err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(http.StatusBadRequest, ParseMessage(err.Error()))
		return
	}

	go c.userUseCase.CountVisitWebsite(&count)

	ctx.Status(http.StatusNoContent)
}

// GetBioDataTrackerInfo
func (c *Controller) GetBioDataTrackerInfo(ctx *gin.Context) {

	info, err := c.userUseCase.GetBioDataTrackerInfo()
	if err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(httpErrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, info)
}

// GetWeeklyData
func (c *Controller) GetWeeklyData(ctx *gin.Context) {
	var req model.WeeklyDataReq

	if err := _http.ReadRequest(ctx, &req); err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(http.StatusBadRequest, ParseMessage(err.Error()))
		return
	}

	weeklyData, err := c.userUseCase.GetWeeklyData(req)
	if err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(httpErrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, weeklyData)
}

func (c *Controller) GetPageBufferPercentages(ctx *gin.Context) {

	var req model.WeeklyDataReq

	if err := _http.ReadRequest(ctx, &req); err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(http.StatusBadRequest, ParseMessage(err.Error()))
		return
	}

	buffer, err := c.userUseCase.GetPageBufferPercentages(req)
	if err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(httpErrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, buffer)
}

// Subscribe
func (c *Controller) Subscribe(ctx *gin.Context) {

	var req model.SubscribeReq

	if err := _http.ReadRequest(ctx, &req); err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(http.StatusBadRequest, ParseMessage(err.Error()))
		return
	}

	err := c.userUseCase.Subscribe(req)
	if err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(httpErrors.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetAllSubscribers
func (c *Controller) GetAllSubscribers(ctx *gin.Context) {
	var req model.GetSubscriberReq

	if err := _http.ReadRequest(ctx, &req); err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(http.StatusBadRequest, ParseMessage(err.Error()))
		return
	}

	subscribers, err := c.userUseCase.GetAllSubscribers(req)

	if err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(httpErrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, subscribers)
}

// CalculatePercentageChange

func (c *Controller) CalculatePercentageChange(ctx *gin.Context) {
	var req model.DashboardData

	if err := _http.ReadRequest(ctx, &req); err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(http.StatusBadRequest, ParseMessage(err.Error()))
		return
	}

	percentage, err := c.userUseCase.CalculatePercentageChange(req)
	if err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(httpErrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, percentage)
}

// GetCountsWithPercentage

func (c *Controller) GetCountsWithPercentage(ctx *gin.Context) {
	var req model.Date

	if err := _http.ReadRequest(ctx, &req); err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(http.StatusBadRequest, ParseMessage(err.Error()))
		return
	}

	percentage, err := c.userUseCase.GetCountsWithPercentage(req)
	if err != nil {
		_http.LogResponseError(ctx, err)
		ctx.JSON(httpErrors.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, percentage)
}
