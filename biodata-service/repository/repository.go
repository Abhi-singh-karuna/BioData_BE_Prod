package repository

import (
	"database/sql"
	"myapp/config"
	"myapp/model"

	"github.com/Abhi-singh-karuna/my_Liberary/baselogger"
	"github.com/Abhi-singh-karuna/my_Liberary/cachehandler"

	"github.com/pkg/errors"
)

// Define the Database struct
type Database struct {
	db          *sql.DB
	logger      *baselogger.BaseLogger
	cfg         *config.Config
	redisClient cachehandler.CacheHandler
}

// Define the Repository interface
type Repository interface {
	ValidateUserVerified(userId string) (bool, *model.User, error)
	AddPassword(userId *model.PasswordReq) (*model.PasswordResp, error)
	CountVisitWebsite(count *model.BioDataCount) (error)
}

// NewRepository creates a new instance of Database
func NewRepository(db *sql.DB, logger *baselogger.BaseLogger, cfg *config.Config, redisClient cachehandler.CacheHandler) Repository {
	return &Database{db: db, logger: logger, cfg: cfg, redisClient: redisClient}
}

func (r *Database) AddPassword(pReq *model.PasswordReq) (*model.PasswordResp, error) {
	row, err := r.db.Query(QueryAddPassword, pReq.User_Id, pReq.Website_Name, pReq.Password)
	if err != nil {
		return nil, errors.Wrap(err, "repository.password.AddPassword.QueryAddPassword.sp.query")
	}
	var res model.PasswordResp
	for row.Next() {
		err := row.Scan(&res.ID, &res.Name)
		if err != nil {
			return nil, errors.Wrap(err, "repository.password.AddPassword.rows.Scan")
		}
	}
	r.logger.Infof("repository.user.AddPassword %v", res)

	return &res, nil
}

func (r *Database) CountVisitWebsite(count *model.BioDataCount) (error) {
	_, err := r.db.Query(QueryCountVisitWebsite, count.ID)
	if err != nil {
		return errors.Wrap(err, "repository.password.CountVisitWebsite.QueryCountVisitWebsite.sp.query")
	}

	return nil
}
