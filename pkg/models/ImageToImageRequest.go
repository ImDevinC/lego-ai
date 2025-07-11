package models

type ImageToImageRequest struct {
	APIKey string
	Height int
	Width  int
	Image  string
}

func NewImageToImageRequest(apikey string, height int, width int, image string) ImageToImageRequest {
	return ImageToImageRequest{
		APIKey: apikey,
		Height: height,
		Width:  width,
		Image:  image,
	}
}
