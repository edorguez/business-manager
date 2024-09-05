package contracts

type GetPaymentsTypesResponse struct {
	Id            int64                   `json:"id"`
	CompanyId     int64                   `json:"companyId"`
	PaymentTypeId int64                   `json:"paymentTypeId"`
	IsActive      bool                    `json:"isActive"`
	PaymentType   *GetPaymentTypeResponse `json:"paymentType"`
}
