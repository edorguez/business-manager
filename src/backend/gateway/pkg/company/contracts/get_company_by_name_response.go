package contracts

import "time"

type GetCompanyByNameResponse struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	ImageUrl        *string   `json:"imageUrl"`
	PlanId          int64     `json:"planId"`
	LastPaymentDate time.Time `json:"lastPaymentDate"`
}
