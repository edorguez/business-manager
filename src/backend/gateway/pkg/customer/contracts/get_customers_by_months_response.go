package contracts

import "time"

type GetCustomerByMonthsResponse struct {
	Dates []time.Time `json:"dates"`
}
