package models

// cfg_scale = 7
// height = 512
// width = 768
// samples = 1
// steps = 30

type ImageToImageRequest struct {
	TextPrompts []string
	CfgScale    int
	Height      int
	Width       int
	Samples     int
	Steps       int
	Preset      string
	Image       []byte
}

func NewImageToImageRequest(text string, style string, cfgScale int, height int, width int, samples int, steps int, image []byte) ImageToImageRequest {
	return ImageToImageRequest{
		TextPrompts: []string{text},
		Preset:      style,
		CfgScale:    cfgScale,
		Height:      height,
		Width:       width,
		Samples:     samples,
		Steps:       steps,
		Image:       image,
	}
}
