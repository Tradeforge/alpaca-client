package model

type GetCalendarParams struct {
	Since *string `query:"start"`
	Until *string `query:"end"`
}

type GetCalendarResponse = []CalendarDay

type CalendarDay struct {
	// Date is a date in YYYY-MM-DD format.
	Date string `json:"date"`
	// Open is the time the market opens in HH:MM format.
	Open string `json:"open"`
	// Close is the time the market closes in HH:MM format.
	Close string `json:"close"`
	// SettlementDate is the date of the settlement in YYYY-MM-DD format.
	SettlementDate string `json:"settlement_date"`
}
