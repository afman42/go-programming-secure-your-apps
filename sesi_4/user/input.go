package user

type RegisterUserInput struct {
	Username string `json:"username" binding:"required" example:""`
	Age      uint   `json:"age" binding:"required,gt=8" example:"0"`
	Email    string `json:"email" binding:"required,email" example:""`
	Password string `json:"password" binding:"required,min=6" example:""`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email" example:""`
	Password string `json:"password" form:"password" binding:"required" example:""`
}
