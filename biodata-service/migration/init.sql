-- DROP TABLE IF EXISTS `visit-bio-data`;


CREATE TABLE IF NOT EXISTS `Application` (
    `ID` INT NOT NULL AUTO_INCREMENT,
    `NAME` VARCHAR(255) NOT NULL,
    `DESCRIPTION` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `visit-bio-data` (
    `ID` INT NOT NULL AUTO_INCREMENT,
    `WEBSITE_NAME` VARCHAR(255) NOT NULL,
    `VISIT_COUNT` INT NOT NULL DEFAULT 0,
    `CLICK_GENERATE_COUNT` INT NOT NULL DEFAULT 0,
    `FILL_FORM_COUNT` INT NOT NULL DEFAULT 0,
    `DOWNLOAD_BIO_DATA_COUNT` INT NOT NULL DEFAULT 0,
    `SUBSCRIBER_COUNT` INT NOT NULL DEFAULT 0,
    `CUSTOMIZED_TEMPLATE_COUNT` INT NOT NULL DEFAULT 0,
    `DATE` DATE DEFAULT (CURRENT_DATE()),
    `UPDATED_AT` DATETIME DEFAULT (UTC_TIMESTAMP),
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `subscribers` (
    `ID` INT NOT NULL AUTO_INCREMENT,
    `APPLICATION_ID` INT NOT NULL,
    `EMAIL` VARCHAR(255) NOT NULL,
    `IS_EMAIL_SENT` BOOLEAN NOT NULL DEFAULT FALSE,
    `SUBSCRIBED_AT` DATETIME DEFAULT (UTC_TIMESTAMP),
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`APPLICATION_ID`) REFERENCES `Application`(`ID`)
);

-- put data in Application table
INSERT INTO Application (ID, NAME, DESCRIPTION)
SELECT 1, 'bio-data', 'bio-data'
WHERE NOT EXISTS (
    SELECT 1 FROM Application WHERE NAME = 'bio-data'
);

DROP PROCEDURE IF EXISTS sp_CountVisitWebsite;

CREATE PROCEDURE sp_CountVisitWebsite(IN ID INT)
BEGIN
    DECLARE today DATE;
    DECLARE row_exists INT;

    SET today = CURRENT_DATE();

    -- Check if a row exists for today
    SELECT COUNT(*) INTO row_exists
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = today;

    -- If no row exists for today, insert a new row
    IF row_exists = 0 THEN
        INSERT INTO `visit-bio-data` (WEBSITE_NAME, DATE)
        VALUES ('bio-data', today);
    END IF;

    -- Update the corresponding count based on the ID
    CASE ID
        WHEN 1 THEN
            UPDATE `visit-bio-data` 
            SET VISIT_COUNT = VISIT_COUNT + 1, 
                UPDATED_AT = UTC_TIMESTAMP() 
            WHERE WEBSITE_NAME = 'bio-data' AND DATE = today;
        WHEN 2 THEN
            UPDATE `visit-bio-data` 
            SET CLICK_GENERATE_COUNT = CLICK_GENERATE_COUNT + 1, 
                UPDATED_AT = UTC_TIMESTAMP() 
            WHERE WEBSITE_NAME = 'bio-data' AND DATE = today;
        WHEN 3 THEN
            UPDATE `visit-bio-data` 
            SET FILL_FORM_COUNT = FILL_FORM_COUNT + 1, 
                UPDATED_AT = UTC_TIMESTAMP() 
            WHERE WEBSITE_NAME = 'bio-data' AND DATE = today;
        WHEN 4 THEN
            UPDATE `visit-bio-data` 
            SET DOWNLOAD_BIO_DATA_COUNT = DOWNLOAD_BIO_DATA_COUNT + 1, 
                UPDATED_AT = UTC_TIMESTAMP() 
            WHERE WEBSITE_NAME = 'bio-data' AND DATE = today;
        WHEN 5 THEN
            UPDATE `visit-bio-data` 
            SET SUBSCRIBER_COUNT = SUBSCRIBER_COUNT + 1, 
                UPDATED_AT = UTC_TIMESTAMP() 
            WHERE WEBSITE_NAME = 'bio-data' AND DATE = today;
        WHEN 6 THEN
            UPDATE `visit-bio-data` 
            SET CUSTOMIZED_TEMPLATE_COUNT = CUSTOMIZED_TEMPLATE_COUNT + 1, 
                UPDATED_AT = UTC_TIMESTAMP() 
            WHERE WEBSITE_NAME = 'bio-data' AND DATE = today;
    END CASE;
END;

-- get page buffer percentages

DROP PROCEDURE IF EXISTS sp_GetPageBufferPercentages;

CREATE PROCEDURE sp_GetPageBufferPercentages(IN input_date VARCHAR(10))
BEGIN
    DECLARE target_date DATE;
    DECLARE total_actions INT;

    -- Convert input date from mm/dd/yyyy to yyyy-mm-dd
    SET target_date = STR_TO_DATE(input_date, '%m/%d/%Y');

    -- Calculate total actions
    SELECT 
        CLICK_GENERATE_COUNT + FILL_FORM_COUNT + DOWNLOAD_BIO_DATA_COUNT + 
        SUBSCRIBER_COUNT + CUSTOMIZED_TEMPLATE_COUNT INTO total_actions
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = target_date;

    -- Calculate and return percentages
    SELECT
        DATE_FORMAT(DATE, '%m/%d/%Y') AS formatted_date,
        VISIT_COUNT,
        CASE 
            WHEN total_actions = 0 THEN 0
            ELSE ROUND((CLICK_GENERATE_COUNT / total_actions) * 100, 2)
        END AS generate_page_percentage,
        CASE 
            WHEN total_actions = 0 THEN 0
            ELSE ROUND((FILL_FORM_COUNT / total_actions) * 100, 2)
        END AS fill_form_percentage,
        CASE 
            WHEN total_actions = 0 THEN 0
            ELSE ROUND((DOWNLOAD_BIO_DATA_COUNT / total_actions) * 100, 2)
        END AS download_page_percentage,
        CASE 
            WHEN total_actions = 0 THEN 0
            ELSE ROUND((SUBSCRIBER_COUNT / total_actions) * 100, 2)
        END AS subscriber_page_percentage,
        CASE 
            WHEN total_actions = 0 THEN 0
            ELSE ROUND((CUSTOMIZED_TEMPLATE_COUNT / total_actions) * 100, 2)
        END AS customized_template_percentage
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = target_date;
END;

-- get count percentage increase

DROP PROCEDURE IF EXISTS sp_GetCountPercentageIncrease;

CREATE PROCEDURE sp_GetCountPercentageIncrease(IN input_date VARCHAR(10))
BEGIN
    DECLARE target_date DATE;
    DECLARE previous_date DATE;

    -- Convert input date from mm/dd/yyyy to yyyy-mm-dd
    SET target_date = STR_TO_DATE(input_date, '%m/%d/%Y');
    SET previous_date = DATE_SUB(target_date, INTERVAL 1 DAY);

    -- Calculate percentage increase for each count and include current counts
    SELECT 
        DATE_FORMAT(target_date, '%m/%d/%Y') AS target_date,
        DATE_FORMAT(previous_date, '%m/%d/%Y') AS previous_date,
        t.VISIT_COUNT AS current_visit_count,
        COALESCE(ROUND(((t.VISIT_COUNT - IFNULL(p.VISIT_COUNT, 0)) / GREATEST(IFNULL(p.VISIT_COUNT, 0), 1)) * 100, 2), 0) AS visit_count_increase,
        t.CLICK_GENERATE_COUNT AS current_click_generate_count,
        COALESCE(ROUND(((t.CLICK_GENERATE_COUNT - IFNULL(p.CLICK_GENERATE_COUNT, 0)) / GREATEST(IFNULL(p.CLICK_GENERATE_COUNT, 0), 1)) * 100, 2), 0) AS click_generate_increase,
        t.FILL_FORM_COUNT AS current_fill_form_count,
        COALESCE(ROUND(((t.FILL_FORM_COUNT - IFNULL(p.FILL_FORM_COUNT, 0)) / GREATEST(IFNULL(p.FILL_FORM_COUNT, 0), 1)) * 100, 2), 0) AS fill_form_increase,
        t.DOWNLOAD_BIO_DATA_COUNT AS current_download_count,
        COALESCE(ROUND(((t.DOWNLOAD_BIO_DATA_COUNT - IFNULL(p.DOWNLOAD_BIO_DATA_COUNT, 0)) / GREATEST(IFNULL(p.DOWNLOAD_BIO_DATA_COUNT, 0), 1)) * 100, 2), 0) AS download_increase,
        t.SUBSCRIBER_COUNT AS current_subscriber_count,
        COALESCE(ROUND(((t.SUBSCRIBER_COUNT - IFNULL(p.SUBSCRIBER_COUNT, 0)) / GREATEST(IFNULL(p.SUBSCRIBER_COUNT, 0), 1)) * 100, 2), 0) AS subscriber_increase,
        t.CUSTOMIZED_TEMPLATE_COUNT AS current_customized_template_count,
        COALESCE(ROUND(((t.CUSTOMIZED_TEMPLATE_COUNT - IFNULL(p.CUSTOMIZED_TEMPLATE_COUNT, 0)) / GREATEST(IFNULL(p.CUSTOMIZED_TEMPLATE_COUNT, 0), 1)) * 100, 2), 0) AS customized_template_increase
    FROM 
        `visit-bio-data` t
    LEFT JOIN 
        `visit-bio-data` p ON p.DATE = previous_date AND p.WEBSITE_NAME = 'bio-data'
    WHERE 
        t.DATE = target_date AND t.WEBSITE_NAME = 'bio-data';
END;

-- get total counts by period and type

DROP PROCEDURE IF EXISTS sp_GetTotalCountsByPeriodAndType;

CREATE PROCEDURE sp_GetTotalCountsByPeriodAndType(
    IN input_date VARCHAR(10),
    IN period_type VARCHAR(10)
)
BEGIN
    DECLARE target_date DATE;
    DECLARE start_date DATE;
    DECLARE end_date DATE;

    -- Convert input date from mm/dd/yyyy to yyyy-mm-dd
    SET target_date = STR_TO_DATE(input_date, '%m/%d/%Y');

    -- Set the start and end dates based on the period_type
    CASE period_type
        WHEN 'day' THEN
            SET start_date = target_date;
            SET end_date = target_date;
        WHEN 'month' THEN
            SET start_date = DATE_FORMAT(target_date, '%Y-%m-01');
            SET end_date = LAST_DAY(target_date);
        WHEN 'year' THEN
            SET start_date = DATE_FORMAT(target_date, '%Y-01-01');
            SET end_date = DATE_FORMAT(target_date, '%Y-12-31');
        ELSE
            SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'Invalid period_type. Use "day", "month", or "year".';
    END CASE;

    -- Calculate and return total counts, using COALESCE to replace NULL with 0
    SELECT 
        DATE_FORMAT(start_date, '%m/%d/%Y') AS start_date,
        DATE_FORMAT(end_date, '%m/%d/%Y') AS end_date,
        period_type AS period,
        COALESCE(SUM(VISIT_COUNT), 0) AS total_visit_count,
        COALESCE(SUM(CLICK_GENERATE_COUNT), 0) AS total_click_generate_count,
        COALESCE(SUM(FILL_FORM_COUNT), 0) AS total_fill_form_count,
        COALESCE(SUM(DOWNLOAD_BIO_DATA_COUNT), 0) AS total_download_count,
        COALESCE(SUM(SUBSCRIBER_COUNT), 0) AS total_subscriber_count,
        COALESCE(SUM(CUSTOMIZED_TEMPLATE_COUNT), 0) AS total_customized_template_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data'
    AND DATE BETWEEN start_date AND end_date;
END;

--  GET WEEKLY DATA 

DROP PROCEDURE IF EXISTS sp_GetWeeklyData;

CREATE PROCEDURE sp_GetWeeklyData(IN input_date VARCHAR(10), IN frequency VARCHAR(10))
BEGIN
    DECLARE end_date DATE;
    DECLARE start_date DATE;
    DECLARE prev_start_date DATE;
    DECLARE prev_end_date DATE;
    DECLARE current_total INT;
    DECLARE previous_total INT;
    DECLARE percent_hike DECIMAL(10, 2);

    -- Convert input date from mm/dd/yyyy to yyyy-mm-dd
    SET end_date = STR_TO_DATE(input_date, '%m/%d/%Y');

    -- Set date ranges based on frequency
    CASE frequency
        WHEN 'Weekly' THEN
            SET start_date = DATE_SUB(end_date, INTERVAL 6 DAY);
            SET prev_end_date = DATE_SUB(start_date, INTERVAL 1 DAY);
            SET prev_start_date = DATE_SUB(prev_end_date, INTERVAL 6 DAY);
        WHEN 'Monthly' THEN
            SET start_date = DATE_FORMAT(end_date, '%Y-%m-01');
            SET prev_end_date = DATE_SUB(start_date, INTERVAL 1 DAY);
            SET prev_start_date = DATE_SUB(prev_end_date, INTERVAL 5 DAY);
        WHEN 'Yearly' THEN
            SET start_date = DATE_FORMAT(end_date, '%Y-01-01');
            SET prev_end_date = DATE_SUB(start_date, INTERVAL 1 DAY);
            SET prev_start_date = DATE_FORMAT(DATE_SUB(end_date, INTERVAL 1 YEAR), '%Y-12-26');
    END CASE;

    -- Calculate total for current period
    SELECT SUM(CLICK_GENERATE_COUNT + FILL_FORM_COUNT + DOWNLOAD_BIO_DATA_COUNT + 
               SUBSCRIBER_COUNT + CUSTOMIZED_TEMPLATE_COUNT)
    INTO current_total
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data'
    AND DATE BETWEEN start_date AND end_date;

    -- Calculate total for previous period
    SELECT SUM(CLICK_GENERATE_COUNT + FILL_FORM_COUNT + DOWNLOAD_BIO_DATA_COUNT + 
               SUBSCRIBER_COUNT + CUSTOMIZED_TEMPLATE_COUNT)
    INTO previous_total
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data'
    AND DATE BETWEEN prev_start_date AND prev_end_date;

    -- Calculate percent hike
    IF previous_total > 0 THEN
        SET percent_hike = ((current_total - previous_total) / previous_total) * 100;
    ELSE
        SET percent_hike = 100; -- If previous total is 0, consider it as 100% increase
    END IF;

    -- Select data for the specified period
    SELECT 
        DATE_FORMAT(DATE, '%b %d') AS date,
        (CLICK_GENERATE_COUNT + FILL_FORM_COUNT + DOWNLOAD_BIO_DATA_COUNT + 
         SUBSCRIBER_COUNT + CUSTOMIZED_TEMPLATE_COUNT) AS value1,
        percent_hike AS percent_hike
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data'
    AND DATE BETWEEN 
        CASE 
            WHEN frequency = 'Weekly' THEN start_date
            WHEN frequency = 'Monthly' THEN DATE_SUB(start_date, INTERVAL 6 DAY)
            WHEN frequency = 'Yearly' THEN DATE_SUB(end_date, INTERVAL 6 DAY)
        END 
        AND end_date
    ORDER BY DATE ASC;
END;

-- CalculatePercentageChange 


DROP FUNCTION IF EXISTS CalculatePercentageChange;
CREATE FUNCTION CalculatePercentageChange(current_value INT, previous_value INT)
RETURNS DECIMAL(10, 2) DETERMINISTIC
BEGIN
    IF previous_value = 0 THEN
        RETURN 100;
    ELSE
        RETURN ROUND(((current_value - previous_value) / previous_value) * 100, 2);
    END IF;
END;

DROP PROCEDURE IF EXISTS sp_GetTotalCountsWithPercentage;

CREATE PROCEDURE sp_GetTotalCountsWithPercentage(
    IN visit_date VARCHAR(10),
    IN click_generate_date VARCHAR(10),
    IN fill_form_date VARCHAR(10),
    IN download_date VARCHAR(10),
    IN subscriber_date VARCHAR(10),
    IN customized_template_date VARCHAR(10)
)
BEGIN
    -- Declare variables for current and previous counts
    DECLARE current_visit_count INT;
    DECLARE current_click_generate_count INT;
    DECLARE current_fill_form_count INT;
    DECLARE current_download_count INT;
    DECLARE current_subscriber_count INT;
    DECLARE current_customized_template_count INT;

    DECLARE prev_visit_count INT;
    DECLARE prev_click_generate_count INT;
    DECLARE prev_fill_form_count INT;
    DECLARE prev_download_count INT;
    DECLARE prev_subscriber_count INT;
    DECLARE prev_customized_template_count INT;

    -- Get current counts for each specific date
    SELECT COALESCE(VISIT_COUNT, 0) INTO current_visit_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = STR_TO_DATE(visit_date, '%m/%d/%Y');

    SELECT COALESCE(CLICK_GENERATE_COUNT, 0) INTO current_click_generate_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = STR_TO_DATE(click_generate_date, '%m/%d/%Y');

    SELECT COALESCE(FILL_FORM_COUNT, 0) INTO current_fill_form_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = STR_TO_DATE(fill_form_date, '%m/%d/%Y');

    SELECT COALESCE(DOWNLOAD_BIO_DATA_COUNT, 0) INTO current_download_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = STR_TO_DATE(download_date, '%m/%d/%Y');

    SELECT COALESCE(SUBSCRIBER_COUNT, 0) INTO current_subscriber_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = STR_TO_DATE(subscriber_date, '%m/%d/%Y');

    SELECT COALESCE(CUSTOMIZED_TEMPLATE_COUNT, 0) INTO current_customized_template_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = STR_TO_DATE(customized_template_date, '%m/%d/%Y');

    -- Get previous day counts for each metric
    SELECT COALESCE(VISIT_COUNT, 0) INTO prev_visit_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = DATE_SUB(STR_TO_DATE(visit_date, '%m/%d/%Y'), INTERVAL 1 DAY);

    SELECT COALESCE(CLICK_GENERATE_COUNT, 0) INTO prev_click_generate_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = DATE_SUB(STR_TO_DATE(click_generate_date, '%m/%d/%Y'), INTERVAL 1 DAY);

    SELECT COALESCE(FILL_FORM_COUNT, 0) INTO prev_fill_form_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = DATE_SUB(STR_TO_DATE(fill_form_date, '%m/%d/%Y'), INTERVAL 1 DAY);

    SELECT COALESCE(DOWNLOAD_BIO_DATA_COUNT, 0) INTO prev_download_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = DATE_SUB(STR_TO_DATE(download_date, '%m/%d/%Y'), INTERVAL 1 DAY);

    SELECT COALESCE(SUBSCRIBER_COUNT, 0) INTO prev_subscriber_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = DATE_SUB(STR_TO_DATE(subscriber_date, '%m/%d/%Y'), INTERVAL 1 DAY);

    SELECT COALESCE(CUSTOMIZED_TEMPLATE_COUNT, 0) INTO prev_customized_template_count
    FROM `visit-bio-data`
    WHERE WEBSITE_NAME = 'bio-data' AND DATE = DATE_SUB(STR_TO_DATE(customized_template_date, '%m/%d/%Y'), INTERVAL 1 DAY);

    -- Return the result as a single row
    SELECT
        visit_date AS totalVisitCountDate,
        current_visit_count AS totalVisitCountValue,
        CalculatePercentageChange(current_visit_count, prev_visit_count) AS totalVisitCountpercentage,
        fill_form_date AS totalFillFormCountDate,
        current_fill_form_count AS totalFillFormCountValue,
        CalculatePercentageChange(current_fill_form_count, prev_fill_form_count) AS totalFillFormCountpercentage,
        subscriber_date AS totalSubscriberCountDate,
        current_subscriber_count AS totalSubscriberCountValue,
        CalculatePercentageChange(current_subscriber_count, prev_subscriber_count) AS totalSubscriberCountpercentage,
        click_generate_date AS totalClickGenerateCountDate,
        current_click_generate_count AS totalClickGenerateCountValue,
        CalculatePercentageChange(current_click_generate_count, prev_click_generate_count) AS totalClickGenerateCountpercentage,
        download_date AS totalDownloadBioDataCountDate,
        current_download_count AS totalDownloadBioDataCountValue,
        CalculatePercentageChange(current_download_count, prev_download_count) AS totalDownloadBioDataCountpercentage,
        customized_template_date AS totalCustomizedTemplateCountDate,
        current_customized_template_count AS totalCustomizedTemplateCountValue,
        CalculatePercentageChange(current_customized_template_count, prev_customized_template_count) AS totalCustomizedTemplateCountpercentage;
END;

-- Get Count With Percentage

DROP PROCEDURE IF EXISTS sp_GetCountsWithPercentage;

CREATE PROCEDURE sp_GetCountsWithPercentage(IN input_date VARCHAR(10))
BEGIN
    DECLARE target_date DATE;
    DECLARE previous_date DATE;
    DECLARE current_total INT;
    DECLARE previous_total INT;
    DECLARE percentage_change DECIMAL(10, 2);

    -- Convert input date from mm/dd/yyyy to yyyy-mm-dd
    SET target_date = STR_TO_DATE(input_date, '%m/%d/%Y');
    SET previous_date = DATE_SUB(target_date, INTERVAL 1 DAY);

    -- Calculate total for current date
    SELECT 
        COALESCE(SUM(VISIT_COUNT + CLICK_GENERATE_COUNT + FILL_FORM_COUNT + 
                     DOWNLOAD_BIO_DATA_COUNT + SUBSCRIBER_COUNT + CUSTOMIZED_TEMPLATE_COUNT), 0)
    INTO current_total
    FROM `visit-bio-data`
    WHERE DATE = target_date AND WEBSITE_NAME = 'bio-data';

    -- Calculate total for previous date
    SELECT 
        COALESCE(SUM(VISIT_COUNT + CLICK_GENERATE_COUNT + FILL_FORM_COUNT + 
                     DOWNLOAD_BIO_DATA_COUNT + SUBSCRIBER_COUNT + CUSTOMIZED_TEMPLATE_COUNT), 0)
    INTO previous_total
    FROM `visit-bio-data`
    WHERE DATE = previous_date AND WEBSITE_NAME = 'bio-data';

    -- Calculate percentage change
    IF previous_total = 0 THEN
        SET percentage_change = 100; -- Consider it as 100% increase if previous total was 0
    ELSE
        SET percentage_change = ((current_total - previous_total) / previous_total) * 100;
    END IF;

    -- Get counts and return result
    SELECT 
        DATE_FORMAT(target_date, '%m/%d/%Y') AS date,
        VISIT_COUNT,
        CLICK_GENERATE_COUNT,
        FILL_FORM_COUNT,
        DOWNLOAD_BIO_DATA_COUNT,
        SUBSCRIBER_COUNT,
        CUSTOMIZED_TEMPLATE_COUNT,
        current_total AS total_count,
        ROUND(percentage_change, 2) AS percentage_change
    FROM `visit-bio-data`
    WHERE DATE = target_date AND WEBSITE_NAME = 'bio-data';
END;


-- demo data 

DROP PROCEDURE IF EXISTS sp_InsertSampleData;


CREATE PROCEDURE sp_InsertSampleData()
BEGIN
    -- Use user-defined variables instead of DECLARE
    SET @i = 0;
    SET @current_date = DATE_SUB(CURDATE(), INTERVAL 99 DAY);
    SET @app_id = (SELECT ID FROM Application WHERE NAME = 'bio-data' LIMIT 1);

    -- Loop for 100 days
    WHILE @i < 100 DO
        -- Insert data into visit-bio-data table
        INSERT INTO `visit-bio-data` (
            WEBSITE_NAME,
            VISIT_COUNT,
            CLICK_GENERATE_COUNT,
            FILL_FORM_COUNT,
            DOWNLOAD_BIO_DATA_COUNT,
            SUBSCRIBER_COUNT,
            CUSTOMIZED_TEMPLATE_COUNT,
            DATE
        ) VALUES (
            'bio-data',
            FLOOR(RAND() * 1000),
            FLOOR(RAND() * 500),
            FLOOR(RAND() * 300),
            FLOOR(RAND() * 200),
            FLOOR(RAND() * 100),
            FLOOR(RAND() * 50),
            @current_date
        );

        -- Insert data into subscribers table
        INSERT INTO `subscribers` (
            APPLICATION_ID,
            EMAIL,
            IS_EMAIL_SENT,
            SUBSCRIBED_AT
        )
        SELECT
            @app_id,
            CONCAT('user', n, '@example.com'),
            CASE WHEN RAND() < 0.5 THEN TRUE ELSE FALSE END,
            DATE_ADD(@current_date, INTERVAL FLOOR(RAND() * 24 * 60 * 60) SECOND)
        FROM
            (SELECT @row := @row + 1 AS n FROM
                (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t1,
                (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t2,
                (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t3,
                (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t4,
                (SELECT @row:=0) r
            ) numbers
        WHERE n <= 100;

        SET @current_date = DATE_ADD(@current_date, INTERVAL 1 DAY);
        SET @i = @i + 1;
    END WHILE;
END ;

-- Call the procedure to insert sample data
-- CALL sp_InsertSampleData();