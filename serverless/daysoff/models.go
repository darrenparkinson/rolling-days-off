package main

import "time"

// BankHolidays represents to top level bank holiday response from gov.uk
type BankHolidays struct {
	EnglandAndWales struct {
		Division string
		Events   []BankHoliday
	} `json:"england-and-wales"`
}

// BankHoliday represents a single Bank Holiday from gov.uk
type BankHoliday struct {
	Title   string
	Date    string
	Notes   string
	Bunting bool
}

// DayOff represents a day off along with BankHoliday information
type DayOff struct {
	Date             time.Time `json:"date"`
	BankHoliday      bool      `json:"bankHoliday"`
	BankHolidayTitle string    `json:"bankHolidayTitle"`
	BankHolidayNotes string    `json:"bankHolidayNotes"`
}
