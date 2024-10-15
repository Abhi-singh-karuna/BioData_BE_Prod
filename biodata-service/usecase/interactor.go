package usecase

import (
	"myapp/config"
	"myapp/model"
	"myapp/repository"

	"github.com/Abhi-singh-karuna/my_Liberary/baselogger"

	email "github.com/Abhi-singh-karuna/my_Liberary/email"
)

// Define the UserInteractor struct
type UserInteractor struct {
	userRepo     repository.Repository
	logger       *baselogger.BaseLogger
	cfg          *config.Config
	emailService *email.EmailService
}

// Define the UserUseCase interface
type UserUseCase interface {
	ValidateUserVerified(userId string) (bool, *model.User, error)
	AddPassword(userId *model.PasswordReq) (*model.PasswordResp, error)
	CountVisitWebsite(count *model.BioDataCount) error
	GetBioDataTrackerInfo() (*model.BioDataTrackerInfo, error)
	GetWeeklyData(weeklyData model.WeeklyDataReq) ([]*model.WeeklyData, error)
	GetPageBufferPercentages(weeklyData model.WeeklyDataReq) (*model.PageBufferPercentagesResponse, error)
	Subscribe(subscribeReq model.SubscribeReq) error
	GetAllSubscribers(req model.GetSubscriberReq) (*model.GetAllSubscribers, error)
	CalculatePercentageChange(req model.DashboardData) (*model.DashboardData, error)
	GetCountsWithPercentage(req model.Date) (*model.CountsWithPercentage, error)
}

// NewUserInteractor creates a new instance of UserInteractor
func NewUserInteractor(userRepo repository.Repository, logger *baselogger.BaseLogger, cfg *config.Config, emailService *email.EmailService) UserUseCase {
	return &UserInteractor{userRepo: userRepo, logger: logger, cfg: cfg, emailService: emailService}
}

func (i *UserInteractor) ValidateUserVerified(userId string) (bool, *model.User, error) {
	return i.userRepo.ValidateUserVerified(userId)
}

func (i *UserInteractor) AddPassword(pass *model.PasswordReq) (*model.PasswordResp, error) {
	res, err := i.userRepo.AddPassword(pass)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *UserInteractor) CountVisitWebsite(count *model.BioDataCount) error {
	return i.userRepo.CountVisitWebsite(count)
}

func (i *UserInteractor) GetBioDataTrackerInfo() (*model.BioDataTrackerInfo, error) {
	return i.userRepo.GetBioDataTrackerInfo()
}

func (i *UserInteractor) GetWeeklyData(weeklyData model.WeeklyDataReq) ([]*model.WeeklyData, error) {
	return i.userRepo.GetWeeklyData(weeklyData)
}

func (i *UserInteractor) GetPageBufferPercentages(weeklyData model.WeeklyDataReq) (*model.PageBufferPercentagesResponse, error) {
	return i.userRepo.GetPageBufferPercentages(weeklyData)
}

func (i *UserInteractor) Subscribe(subscribeReq model.SubscribeReq) error {
	return i.userRepo.Subscribe(subscribeReq)
}

func (i *UserInteractor) GetAllSubscribers(req model.GetSubscriberReq) (*model.GetAllSubscribers, error) {
	return i.userRepo.GetAllSubscribers(req)
}

func (i *UserInteractor) CalculatePercentageChange(req model.DashboardData) (*model.DashboardData, error) {
	return i.userRepo.CalculatePercentageChange(req)
}

func (i *UserInteractor) GetCountsWithPercentage(req model.Date) (*model.CountsWithPercentage, error) {
	return i.userRepo.GetCountsWithPercentage(req)
}
