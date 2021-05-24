package authreqbodies

type SignUpReq struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Email    string `json:"email"  validate:"required,email,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
}
