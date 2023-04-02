package user

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Age      uint   `json:"age" binding:"required,gt=8"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
