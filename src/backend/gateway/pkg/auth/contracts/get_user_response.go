package contracts

type GetUserResponse struct {
	Id        int64          `json:"id"`
	CompanyId int64          `json:"companyId"`
	RoleId    int64          `json:"roleId"`
	Email     string         `json:"email"`
	Rol       GetRolResponse `json:"rol"`
}
