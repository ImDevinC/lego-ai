package imagegenerators

// https://github.com/infosecak/DeGhiblify/blob/main/src/openai_client.py

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/imdevinc/ghibli-ai/pkg/models"
)

type OpenAIGenerator struct {
	urlBase      string
	imageModel   string
	chatModel    string
	systemPrompt string
	userPrompt   string
	legoPrompt   string
}

type openAIImageRequest struct {
	Model          string `json:"model"`
	Prompt         string `json:"prompt"`
	Count          int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format,omitempty"`
}

type openAIImageResponse struct {
	Data []struct {
		B64JSON string `json:"b64_json"`
	} `json:"data"`
}

type openAIUploadResponse struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
	CreatedAt int64  `json:"created_at"`
	Bytes     int64  `json:"bytes"`
}

type openAIChatRequest struct {
	Model    string              `json:"model"`
	Messages []openAIChatMessage `json:"messages"`
}

type openAIChatMessage struct {
	Role    string              `json:"role"`
	Content []openAIChatContent `json:"content"`
}

type openAIChatContent struct {
	ImageURL *openAIChatImageURL `json:"image_url,omitempty"`
	Text     string              `json:"text,omitempty"`
	Type     string              `json:"type"`
}

type openAIChatImageURL struct {
	URL string `json:"url"`
}

type openAIChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

var _ Generator = (*OpenAIGenerator)(nil)

func NewOpenAIGenerator(imageModel string, chatModel string, systemPrompt string, userPrompt string, legoPrompt string) OpenAIGenerator {
	return OpenAIGenerator{
		urlBase:      "https://api.openai.com/v1/",
		imageModel:   imageModel,
		chatModel:    chatModel,
		systemPrompt: systemPrompt,
		userPrompt:   userPrompt,
		legoPrompt:   legoPrompt,
	}
}

func (g *OpenAIGenerator) do(apikey string, endpoint string, request []byte, extraHeaders map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, g.urlBase+endpoint, bytes.NewReader(request))
	if err != nil {
		return nil, fmt.Errorf("failed to generate request. %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+apikey)
	for k, v := range extraHeaders {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request. %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body. %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d. %s. %s", resp.StatusCode, resp.Status, string(body))
	}
	return body, nil
}

func (g *OpenAIGenerator) GenerateImageFromText(request models.TextToImageRequest) (string, error) {
	req := openAIImageRequest{
		Model:          request.Model,
		Prompt:         strings.Join(request.TextPrompts, "\n"),
		Size:           fmt.Sprintf("%dx%d", request.Width, request.Height),
		ResponseFormat: "b64_json",
		Count:          1,
	}
	bytes, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal input. %w", err)
	}
	resp, err := g.do(request.APIKey, "images/generations", bytes, nil)
	if err != nil {
		return "", err
	}
	var respData openAIImageResponse
	if err := json.Unmarshal(resp, &respData); err != nil {
		return "", fmt.Errorf("failed to unmarshal response. %w", err)
	}
	return respData.Data[0].B64JSON, nil
}

func (g *OpenAIGenerator) GenerateImageFromImage(request models.ImageToImageRequest) (string, error) {
	img, err := g.convertImage(request)
	if err != nil {
		return "", fmt.Errorf("failed to convert image. %w", err)
	}
	return img, nil
}

func (g *OpenAIGenerator) convertImage(request models.ImageToImageRequest) (string, error) {
	payload := openAIChatRequest{
		Model: g.chatModel,
		Messages: []openAIChatMessage{
			{
				Role: "system",
				Content: []openAIChatContent{
					{
						Type: "text",
						Text: g.systemPrompt,
					},
				},
			},
			{
				Role: "user",
				Content: []openAIChatContent{
					{
						Type: "text",
						Text: g.userPrompt,
					},
					{
						Type: "image_url",
						ImageURL: &openAIChatImageURL{
							URL: "data:image/png;base64," + request.Image,
						},
					},
				},
			},
		},
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal chat request. %w", err)
	}
	resp, err := g.do(request.APIKey, "chat/completions", bytes, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return "", fmt.Errorf("failed to perform chat request. %w", err)
	}
	var chatResp openAIChatResponse
	if err := json.Unmarshal(resp, &chatResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal chat response. %w", err)
	}
	//log.Println(chatResp.Choices[0].Message.Content)
	imageRequest := models.TextToImageRequest{
		APIKey:      request.APIKey,
		Model:       g.imageModel,
		TextPrompts: []string{g.legoPrompt, chatResp.Choices[0].Message.Content},
		Height:      request.Height,
		Width:       request.Width,
	}
	img, err := g.GenerateImageFromText(imageRequest)
	if err != nil {
		return "", fmt.Errorf("failed to generate image from text. %w", err)
	}
	return img, nil
}
