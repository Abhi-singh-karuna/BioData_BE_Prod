package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type PasswordReq struct {
	User_Id      string `json:"UserId" `
	Website_Name string `json:"Name" validate:"required"`
	Password     string `json:"Password" validate:"required"`
}

type PasswordResp struct {
	ID   string `json:"Id"`
	Name string `json:"Name" validate:"required"`
}

type ErrorMessage struct {
	Message string `json:"Message"`
}

type JwtCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type HeaderId struct {
	USER_ID       string
	USER_EMAIL    string
	USER_PHONE_NO string
	USER_NAME     string
}

type User struct {
	ID         string `json:"Id"`
	Name       string `json:"Name" validate:"required"`
	Email      string `json:"Email" validate:"required,email"`
	Password   string `json:"Password,omitempty" validate:"required"`
	PhoneNo    int64  `json:"PhoneNo" validate:"required"`
	OTP        string `json:"Otp,omitempty"`
	IsVerified bool   `json:"IsVerified"`
}

// Bio-Data

type BioDataCount struct {
	ID int `json:"Id" validate:"required,oneof=1 2 3 4 5 6"` // ex: 1- Visit Count , 2 :- ClickGenerateCount , 3 :- Fill Form Count , 4 :- download bio-data count , 5 :- Subscriber Count , 6 :- Customized Template Count
}

type BioDataTrackerInfo struct {
	ID          int                          `json:"Id" validate:"required"`
	WebsiteName string                       `json:"WebsiteName"`
	Data        []BioDataTrackerInfoResponse `json:"Data"`
}

type BioDataTrackerInfoResponse struct {
	VisitCount              int       `json:"VisitCount"`
	ClickGenerateCount      int       `json:"ClickGenerateCount"`
	FillFormCount           int       `json:"FillFormCount"`
	DownloadBioDataCount    int       `json:"DownloadBioDataCount"`
	SubscriberCount         int       `json:"SubscriberCount"`
	CustomizedTemplateCount int       `json:"CustomizedTemplateCount"`
	Date                    string    `json:"Date"`
	UpdatedAt               time.Time `json:"UpdatedAt"`
}

type WeeklyDataReq struct {
	Date      string `json:"Date"`
	Frequency string `json:"Frequency" validate:"required,oneof=Weekly Monthly Yearly"`
}

type WeeklyData struct {
	Date        string  `json:"Date"`
	Value1      int     `json:"Value1"`
	PercentHike float64 `json:"PercentHike"`
}

// sp_GetPageBufferPercentages

type PageBufferPercentagesResponse struct {
	Date                         string  `json:"Date"`
	VisitCount                   int     `json:"VisitCount"`
	GeneratePagePercentage       float64 `json:"GeneratePagePercentage"`
	FillFormPercentage           float64 `json:"FillFormPercentage"`
	DownloadBioDataPercentage    float64 `json:"DownloadBioDataPercentage"`
	SubscriberPercentage         float64 `json:"SubscriberPercentage"`
	CustomizedTemplatePercentage float64 `json:"CustomizedTemplatePercentage"`
}

// Subscribe

type SubscribeReq struct {
	Email string `json:"Email" validate:"required,email"`
}

// Application ID

type GetSubscriberReq struct {
	ID   int    `json:"Id" validate:"required"`
	Date string `json:"Date"` // mm-dd-yyyy
}

// GetAllSubscribers

type GetAllSubscribers struct {
	ApplicationID    int                `json:"ApplicationId"`
	Application      string             `json:"Application"`
	TotalSubscribers int                `json:"TotalSubscribers"`
	Subscribers      []SubscribeDetails `json:"Subscribers"`
}

type SubscribeDetails struct {
	ID           int       `json:"Id"`
	Email        string    `json:"Email"`
	SubscribedAt time.Time `json:"SubscribedAt"`
	IsEmailSent  bool      `json:"IsEmailSent"`
}

type DashboardMetric struct {
	Value      int     `json:"Value" `
	Percentage float64 `json:"percentage" `
	Date       string  `json:"Date" validate:"required"`
}

type DashboardData struct {
	TotalVisitCount              DashboardMetric `json:"totalVisitCount" validate:"required"`
	TotalClickGenerateCount      DashboardMetric `json:"totalClickGenerateCount" validate:"required"`
	TotalFillFormCount           DashboardMetric `json:"totalFillFormCount" validate:"required"`
	TotalDownloadBioDataCount    DashboardMetric `json:"totalDownloadBioDataCount" validate:"required"`
	TotalSubscriberCount         DashboardMetric `json:"totalSubscriberCount" validate:"required"`
	TotalCustomizedTemplateCount DashboardMetric `json:"totalCustomizedTemplateCount" validate:"required"`
}

type Date struct {
	Date string `json:"Date" validate:"required"`
}

type CountsWithPercentage struct {
	Date                    string  `json:"Date"`
	VisitCount              int     `json:"VisitCount"`
	ClickGenerateCount      int     `json:"ClickGenerateCount"`
	FillFormCount           int     `json:"FillFormCount"`
	DownloadBioDataCount    int     `json:"DownloadBioDataCount"`
	SubscriberCount         int     `json:"SubscriberCount"`
	CustomizedTemplateCount int     `json:"CustomizedTemplateCount"`
	TotalCount              int     `json:"TotalCount"`
	Percentage              float64 `json:"Percentage"`
}
