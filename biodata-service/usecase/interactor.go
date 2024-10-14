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
