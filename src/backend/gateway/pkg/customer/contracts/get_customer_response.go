package contracts

type GetCustomerResponse struct {
	Id                   int64   `json:"id"`
	CompanyId            int64   `json:"companyId"`
	FirstName            string  `json:"firstName"`
	LastName             *string `json:"lastName"`
	Email                *string `json:"email"`
	Phone                *string `json:"phone"`
	IdentificationNumber string  `json:"identificationNumber"`
	IdentificationType   string  `json:"identificationType"`
}
