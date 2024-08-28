package contracts

import "time"

type GetCustomerByMonthsResponse struct {
	MonthInterval time.Time `json:"monthInterval"`
	RecordCount   int64     `json:"recordCount"`
}
