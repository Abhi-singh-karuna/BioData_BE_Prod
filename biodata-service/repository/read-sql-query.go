package repository

const (
	QueryValidateUserVerified = `SELECT IFNULL(USER_ID,""), NAME, EMAIL, PHONE_NO, IS_VERIFIED FROM UserManagement.users WHERE USER_ID = ?`
)
