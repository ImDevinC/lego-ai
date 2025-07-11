package models

type TextToImageRequest struct {
	Model       string
	TextPrompts []string
	Height      int
	Width       int
}

func NewTextToImageRequest(model string, text string, height int, width int, samples int) TextToImageRequest {
	return TextToImageRequest{
		Model:       model,
		TextPrompts: []string{text},
		Height:      height,
		Width:       width,
	}
}
