package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	startDate := request.QueryStringParameters["startDate"]
	endDate := request.QueryStringParameters["endDate"]
	dayOff := request.QueryStringParameters["dayOff"]

	layoutISO := "2006-01-02"

	start, err := time.Parse(layoutISO, startDate)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: ""}, errors.New("error parsing startDate")
	}

	end, err := time.Parse(layoutISO, endDate)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: ""}, errors.New("error parsing endDate")
	}

	off, err := strconv.Atoi(dayOff)

	result := CalculateRollingDaysOff(start, end, time.Weekday(off))
	b, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

// FetchBankHolidays retrieves a list of bank holidays from gov.uk
func FetchBankHolidays() ([]BankHoliday, error) {
	govURL := "https://www.gov.uk/bank-holidays.json"
	resp, err := http.Get(govURL)
	if err != nil {
		return []BankHoliday{}, err
	}
	var govResp BankHolidays
	json.NewDecoder(resp.Body).Decode(&govResp)
	return govResp.EnglandAndWales.Events, nil
}

// CalculateRollingDaysOff takes a start date and a day off
// and then calculates rolling days off for the year.
func CalculateRollingDaysOff(startDate, endDate time.Time, dayOff time.Weekday) []DayOff {
	var returnDates []DayOff
	//
	const layoutISO = "2006-01-02"
	publicHolidays, _ := FetchBankHolidays()
	// turn the bank holidays into a map for easier checking
	bankHolidays := make(map[string]BankHoliday)
	for _, bh := range publicHolidays {
		bankHolidays[bh.Date] = bh
	}

	d := startDate
	for {
		do := DayOff{}
		newDate := d.AddDate(0, 0, 8)
		// Go back a week so we don't miss a day off
		if newDate.Weekday() == dayOff {
			newDate = newDate.AddDate(0, 0, -6)
		}
		do.Date = newDate
		// Add BankHoliday information
		if val, ok := bankHolidays[newDate.Format(layoutISO)]; ok {
			do.BankHoliday = true
			do.BankHolidayTitle = val.Title
			do.BankHolidayNotes = val.Notes
		} else {
			// fmt.Println(d.Weekday(), d.Format(layoutISO))
			do.BankHoliday = false
		}

		returnDates = append(returnDates, do)
		d = newDate

		if newDate.After(endDate) {
			break
		}
	}

	return returnDates
}
