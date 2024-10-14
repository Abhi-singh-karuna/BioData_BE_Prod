package model

import "github.com/golang-jwt/jwt/v4"

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
	ID int `json:"Id" validate:"required"` // ex: 1- Visit Count , 2 :- ClickGenerateCount , 3 :- Fill Form Count , 4 :- download bio-data count
}
