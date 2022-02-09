package models

type UserReport struct {
	Reason       string
	NumOfReports float64
	SnoozeStatus bool
	CanSnooze    bool
}
