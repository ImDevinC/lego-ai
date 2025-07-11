package imagegenerators

import "github.com/imdevinc/ghibli-ai/pkg/models"

type Generator interface {
	GenerateImageFromText(request models.TextToImageRequest) (string, error)
	GenerateImageFromImage(request models.ImageToImageRequest) (string, error)
}
