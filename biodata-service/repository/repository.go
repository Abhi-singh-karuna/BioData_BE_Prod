package repository

import (
	"database/sql"
	"myapp/config"
	"myapp/model"

	"github.com/Abhi-singh-karuna/my_Liberary/baselogger"
	"github.com/Abhi-singh-karuna/my_Liberary/cachehandler"

	"github.com/pkg/errors"
)

const (
	BioDataApplicationID = 1
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
	CountVisitWebsite(count *model.BioDataCount) error
	GetBioDataTrackerInfo() (*model.BioDataTrackerInfo, error)
	GetWeeklyData(weeklyData model.WeeklyDataReq) ([]*model.WeeklyData, error)
	GetPageBufferPercentages(weeklyData model.WeeklyDataReq) (*model.PageBufferPercentagesResponse, error)
	Subscribe(subscribeReq model.SubscribeReq) error
	GetAllSubscribers(req model.GetSubscriberReq) (*model.GetAllSubscribers, error)
	CalculatePercentageChange(req model.DashboardData) (*model.DashboardData, error)
	GetCountsWithPercentage(req model.Date) (*model.CountsWithPercentage, error)
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

func (r *Database) CountVisitWebsite(count *model.BioDataCount) error {
	_, err := r.db.Query(QueryCountVisitWebsite, count.ID)
	if err != nil {
		return errors.Wrap(err, "repository.password.CountVisitWebsite.QueryCountVisitWebsite.sp.query")
	}

	return nil
}

func (r *Database) GetBioDataTrackerInfo() (*model.BioDataTrackerInfo, error) {
	row, err := r.db.Query(QueryGetBioDataTrackerInfo)
	if err != nil {
		return nil, errors.Wrap(err, "repository.password.GetBioDataTrackerInfo.QueryGetBioDataTrackerInfo.query")
	}

	var res model.BioDataTrackerInfo
	res.Data = make([]model.BioDataTrackerInfoResponse, 0)

	for row.Next() {
		var data model.BioDataTrackerInfoResponse
		err := row.Scan(&res.ID, &res.WebsiteName, &data.VisitCount, &data.ClickGenerateCount, &data.FillFormCount, &data.DownloadBioDataCount, &data.SubscriberCount, &data.CustomizedTemplateCount, &data.Date, &data.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "repository.password.GetBioDataTrackerInfo.rows.Scan")
		}
		res.Data = append(res.Data, data)
	}

	return &res, nil
}

func (r *Database) GetWeeklyData(weeklyData model.WeeklyDataReq) ([]*model.WeeklyData, error) {
	row, err := r.db.Query(QueryGetWeeklyData, weeklyData.Date, weeklyData.Frequency)
	if err != nil {
		return nil, errors.Wrap(err, "repository.password.GetWeeklyData.QueryGetWeeklyData.query")
	}

	var res []*model.WeeklyData

	for row.Next() {
		var data model.WeeklyData
		err := row.Scan(&data.Date, &data.Value1, &data.PercentHike)
		if err != nil {
			return nil, errors.Wrap(err, "repository.password.GetWeeklyData.rows.Scan")
		}
		res = append(res, &data)
	}

	return res, nil
}

func (r *Database) GetPageBufferPercentages(weeklyData model.WeeklyDataReq) (*model.PageBufferPercentagesResponse, error) {
	row, err := r.db.Query(QueryGetPageBufferPercentages, weeklyData.Date)
	if err != nil {
		return nil, errors.Wrap(err, "repository.password.GetPageBufferPercentages.QueryGetPageBufferPercentages.query")
	}

	var res model.PageBufferPercentagesResponse

	for row.Next() {
		err := row.Scan(&res.Date, &res.VisitCount, &res.GeneratePagePercentage, &res.FillFormPercentage, &res.DownloadBioDataPercentage, &res.SubscriberPercentage, &res.CustomizedTemplatePercentage)
		if err != nil {
			return nil, errors.Wrap(err, "repository.password.GetPageBufferPercentages.rows.Scan")
		}
	}

	return &res, nil
}

func (r *Database) Subscribe(subscribeReq model.SubscribeReq) error {

	_, err := r.db.Query(QuerySubscribe, BioDataApplicationID, subscribeReq.Email)
	if err != nil {
		return errors.Wrap(err, "repository.password.Subscribe.QuerySubscribe.query")
	}

	return nil
}

func (r *Database) GetAllSubscribers(req model.GetSubscriberReq) (*model.GetAllSubscribers, error) {
	row, err := r.db.Query(QueryGetAllSubscribers, req.ID, req.Date, req.Date)
	if err != nil {
		return nil, errors.Wrap(err, "repository.password.GetAllSubscribers.QueryGetAllSubscribers.query")
	}

	var res model.GetAllSubscribers
	res.Subscribers = make([]model.SubscribeDetails, 0)

	for row.Next() {
		var data model.SubscribeDetails
		err := row.Scan(&res.ApplicationID, &res.Application, &res.TotalSubscribers, &data.ID, &data.Email, &data.SubscribedAt, &data.IsEmailSent)
		if err != nil {
			return nil, errors.Wrap(err, "repository.password.GetAllSubscribers.rows.Scan")
		}
		res.Subscribers = append(res.Subscribers, data)
	}

	return &res, nil

}

// CalculatePercentageChange

func (r *Database) CalculatePercentageChange(req model.DashboardData) (*model.DashboardData, error) {
	// Execute the stored procedure
	row, err := r.db.Query(QueryCalculatePercentageChange,
		req.TotalVisitCount.Date,
		req.TotalClickGenerateCount.Date,
		req.TotalFillFormCount.Date,
		req.TotalDownloadBioDataCount.Date,
		req.TotalSubscriberCount.Date,
		req.TotalCustomizedTemplateCount.Date,
	)
	if err != nil {
		return nil, errors.Wrap(err, "repository.password.CalculatePercentageChange.QueryCalculatePercentageChange.query")
	}

	// Prepare a struct to hold the result
	result := model.DashboardData{
		TotalVisitCount:             model.DashboardMetric{Date: req.TotalVisitCount.Date},
		TotalFillFormCount:          model.DashboardMetric{Date: req.TotalFillFormCount.Date},
		TotalSubscriberCount:        model.DashboardMetric{Date: req.TotalSubscriberCount.Date},
		TotalClickGenerateCount:     model.DashboardMetric{Date: req.TotalClickGenerateCount.Date},
		TotalDownloadBioDataCount:   model.DashboardMetric{Date: req.TotalDownloadBioDataCount.Date},
		TotalCustomizedTemplateCount: model.DashboardMetric{Date: req.TotalCustomizedTemplateCount.Date},
	}

	hasData := false
	for row.Next() {
		hasData = true
		err := row.Scan(
			&result.TotalVisitCount.Date,
			&result.TotalVisitCount.Value,
			&result.TotalVisitCount.Percentage,
			&result.TotalFillFormCount.Date,
			&result.TotalFillFormCount.Value,
			&result.TotalFillFormCount.Percentage,
			&result.TotalSubscriberCount.Date,
			&result.TotalSubscriberCount.Value,
			&result.TotalSubscriberCount.Percentage,
			&result.TotalClickGenerateCount.Date,
			&result.TotalClickGenerateCount.Value,
			&result.TotalClickGenerateCount.Percentage,
			&result.TotalDownloadBioDataCount.Date,
			&result.TotalDownloadBioDataCount.Value,
			&result.TotalDownloadBioDataCount.Percentage,
			&result.TotalCustomizedTemplateCount.Date,
			&result.TotalCustomizedTemplateCount.Value,
			&result.TotalCustomizedTemplateCount.Percentage,
		)

		if err != nil {
			return nil, errors.Wrap(err, "repository.CalculatePercentageChange.row.Scan")
		}
	}

	if !hasData {
		r.logger.Infof("No data found for the given dates, returning default values")
	}

	return &result, nil
}

func (r *Database) GetCountsWithPercentage(req model.Date) (*model.CountsWithPercentage, error) {
	row, err := r.db.Query(QueryGetCountsWithPercentage, req.Date)
	if err != nil {
		return nil, errors.Wrap(err, "repository.password.GetCountsWithPercentage.QueryGetCountsWithPercentage.query")
	}

	var res model.CountsWithPercentage

	for row.Next() {
		err := row.Scan(&res.Date, &res.VisitCount, &res.ClickGenerateCount, &res.FillFormCount, &res.DownloadBioDataCount, &res.SubscriberCount, &res.CustomizedTemplateCount, &res.TotalCount, &res.Percentage)
		if err != nil {
			return nil, errors.Wrap(err, "repository.password.GetCountsWithPercentage.rows.Scan")
		}
	}

	return &res, nil
}
