package utils

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
	"github.com/sabareeswaran-sakthivel/e-commerce/constants"
	"github.com/sabareeswaran-sakthivel/e-commerce/services"
)

func parseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}

func IsBetween(startDate, endDate, date string) bool {
	startDateTime, err := parseDate(startDate)
	if err != nil {
		fmt.Printf(constants.DATE_PARSE_ERR_MESSAGE, err)
		return false
	}

	endDateTime, err := parseDate(endDate)
	fmt.Printf(constants.DATE_PARSE_ERR_MESSAGE, err)
	if err != nil {
		return false
	}

	dateTime, err := parseDate(date)
	fmt.Printf(constants.DATE_PARSE_ERR_MESSAGE, err)
	if err != nil {
		return false
	}

	isBetween := dateTime.After(startDateTime) && dateTime.Before(endDateTime) || dateTime.Equal(startDateTime) || dateTime.Equal(endDateTime)
	return isBetween
}

func RunCronJob() {
	cronService := services.UploadCSVDataService{}

	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		cronService.UploadCSVData("data.csv")
	})

	c.Start()

	select {}
}
