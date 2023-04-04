package comments

type CreateCommentInput struct {
	Message string `json:"message" example:"" binding:"required"`
}
