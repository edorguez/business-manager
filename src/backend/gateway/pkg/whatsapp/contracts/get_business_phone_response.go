package contracts

type GetBusinessPhoneResponse struct {
	CompanyId int64  `json:"companyId"`
	Phone     string `json:"phone"`
}
