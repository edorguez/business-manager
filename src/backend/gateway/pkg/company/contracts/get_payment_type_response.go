package contracts

type GetPaymentTypeResponse struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	ImagePath *string `json:"imagePath"`
}
