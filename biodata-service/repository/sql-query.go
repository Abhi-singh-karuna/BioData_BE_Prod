package repository

const (
	QueryAddPassword = `call sp_CreateOrUpdatePassword(?,?,?)`

	QueryCountVisitWebsite = `call sp_CountVisitWebsite(?)`

	QueryGetBioDataTrackerInfo = "SELECT ID, WEBSITE_NAME, VISIT_COUNT, CLICK_GENERATE_COUNT, FILL_FORM_COUNT, DOWNLOAD_BIO_DATA_COUNT, SUBSCRIBER_COUNT, CUSTOMIZED_TEMPLATE_COUNT, DATE, UPDATED_AT FROM `visit-bio-data`"

	QueryGetWeeklyData = "call sp_GetWeeklyData(?,?)"

	QueryGetPageBufferPercentages = "call sp_GetPageBufferPercentages(?)"

	QuerySubscribe = "INSERT INTO `subscribers` (`APPLICATION_ID`, `EMAIL`) VALUES (?, ?)"

	QueryGetAllSubscribers = `
	SELECT 
		Application.ID AS ApplicationID,
		Application.NAME AS ApplicationName,
		COUNT(subscribers.ID) OVER (PARTITION BY Application.ID) AS TotalSubscribers,
		subscribers.ID,
		subscribers.EMAIL,
		subscribers.SUBSCRIBED_AT,
		subscribers.IS_EMAIL_SENT 
	FROM subscribers
	LEFT JOIN Application ON subscribers.APPLICATION_ID = Application.ID
	WHERE Application.ID = ?
	AND (? = '' OR DATE(subscribers.SUBSCRIBED_AT) = STR_TO_DATE(?, '%m/%d/%Y'))
	ORDER BY subscribers.SUBSCRIBED_AT DESC
	`

	QueryCalculatePercentageChange = "CALL sp_GetTotalCountsWithPercentage(?, ?, ?, ?, ?, ?)"

	QueryGetCountsWithPercentage = "CALL sp_GetCountsWithPercentage(?)"
)
