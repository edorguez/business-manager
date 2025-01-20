package contracts

import "time"

type GetCompanyByNameUrlResponse struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	NameFormatUrl   string    `json:"nameFormatUrl"`
	ImageUrl        *string   `json:"imageUrl"`
	PlanId          int64     `json:"planId"`
	LastPaymentDate time.Time `json:"lastPaymentDate"`
}
