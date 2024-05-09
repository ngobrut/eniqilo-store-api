package request

type RegisterCustomer struct {
	Name  string `json:"name" validate:"required,min=5,max=50"`
	Phone string `json:"phoneNumber" validate:"required,min=10,max=16,phoneCode"`
}
