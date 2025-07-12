package models

// ImageToTextRequest represents the request structure for generating text from an image
type ImageToTextRequest struct {
	APIKey string
	Image  string
}

// NewImageToTextRequest creates a new ImageToTextRequest with the specified API key and image
func NewImageToTextRequest(apikey string, image string) ImageToTextRequest {
	return ImageToTextRequest{
		APIKey: apikey,
		Image:  image,
	}
}
