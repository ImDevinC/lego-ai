package models

// UploadResponse represents the response structure for an image upload
type UploadResponse struct {
	Image  string `json:"image"`  // Base64 encoded image
	Prompt string `json:"prompt"` // Text prompt describing the image
}
