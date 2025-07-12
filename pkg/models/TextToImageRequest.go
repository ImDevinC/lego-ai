package models

// TextToImageRequest represents the request structure for generating images from text prompts
type TextToImageRequest struct {
	APIKey      string
	TextPrompts []string
	Height      int
	Width       int
}

// NewTextToImageRequest creates a new TextToImageRequest with the specified parameters
func NewTextToImageRequest(apikey string, prompt string, height int, width int, samples int) TextToImageRequest {
	return TextToImageRequest{
		APIKey:      apikey,
		TextPrompts: []string{prompt},
		Height:      height,
		Width:       width,
	}
}
