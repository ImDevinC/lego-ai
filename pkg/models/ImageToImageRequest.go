package models

type ImageToImageRequest struct {
	Height int
	Width  int
	Image  string
}

func NewImageToImageRequest(height int, width int, image string) ImageToImageRequest {
	return ImageToImageRequest{
		Height: height,
		Width:  width,
		Image:  image,
	}
}
