package contracts

type GetPaymentResponse struct {
	Id                   int64                   `json:"id"`
	CompanyId            int64                   `json:"companyId"`
	Name                 string                  `json:"name"`
	Bank                 *string                 `json:"bank"`
	AccountNumber        *string                 `json:"accountNumber"`
	AccountType          *string                 `json:"accountType"`
	IdentificationNumber *string                 `json:"identificationNumber"`
	IdentificationType   *string                 `json:"identificationType"`
	Phone                *string                 `json:"phone"`
	Email                *string                 `json:"email"`
	PaymentTypeId        int64                   `json:"paymentTypeId"`
	PaymentType          *GetPaymentTypeResponse `json:"paymentType"`
	IsActive             bool                    `json:"isActive"`
}
