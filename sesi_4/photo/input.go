package photo

type CreatePhotoInput struct {
	Title    string `json:"title" example:"" binding:"required"`
	Caption  string `json:"caption" example:""`
	PhotoUrl string `json:"photo_url" example:"" binding:"required"`
}
