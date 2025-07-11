package models

// cfg_scale = 7
// height = 512
// width = 768
// samples = 1
// steps = 30

type TextToImageRequest struct {
	TextPrompts []string
	CfgScale    int
	Height      int
	Width       int
	Samples     int
	Steps       int
	Preset      string
}

func NewTextToImageRequest(text string, style string, cfgScale int, height int, width int, samples int, steps int) TextToImageRequest {
	return TextToImageRequest{
		TextPrompts: []string{text},
		Preset:      style,
		CfgScale:    cfgScale,
		Height:      height,
		Width:       width,
		Samples:     samples,
		Steps:       steps,
	}
}
