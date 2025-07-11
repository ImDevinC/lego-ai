package models

type TextToImageRequest struct {
	APIKey      string
	Model       string
	TextPrompts []string
	Height      int
	Width       int
}

func NewTextToImageRequest(apikey string, model string, text string, height int, width int, samples int) TextToImageRequest {
	return TextToImageRequest{
		APIKey:      apikey,
		Model:       model,
		TextPrompts: []string{text},
		Height:      height,
		Width:       width,
	}
}
