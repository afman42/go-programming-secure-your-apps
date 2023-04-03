package socialmedia

type CreateSocialMediaInput struct {
	Name           string `json:"name" binding:"required" example:""`
	SocialMediaUrl string `json:"social_media_url" binding:"required" example:""`
}
