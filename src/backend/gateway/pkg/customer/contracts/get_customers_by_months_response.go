package contracts

type GetCustomerByMonthsResponse struct {
	MonthInterval int64 `json:"monthInterval"`
	RecordCount   int64 `json:"recordCount"`
}
