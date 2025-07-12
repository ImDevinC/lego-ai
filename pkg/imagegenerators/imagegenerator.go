package imagegenerators

import "github.com/imdevinc/lego-ai/pkg/models"

type Generator interface {
	// Takes a text prompt and returns a base64 representation of the image
	GenerateImageFromText(request models.TextToImageRequest) (string, error)
	// Takes an image and returns a text prompt that describes the image
	GenerateDescriptionFromImage(request models.ImageToTextRequest) (string, error)
}
