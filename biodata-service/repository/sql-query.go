package repository

const (
	QueryAddPassword = `call sp_CreateOrUpdatePassword(?,?,?)`

	QueryCountVisitWebsite = `call sp_CountVisitWebsite(?)`
)
