package imagegenerators

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"

	"github.com/imdevinc/ghibli-ai/pkg/models"
)

type OpenAIGenerator struct {
	urlBase string
	apiKey  string
	model   string
}

type openAIRequest struct {
	Model          string `json:"model"`
	Prompt         string `json:"prompt"`
	Count          int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
}

type openAIResponse struct {
	Data []struct {
		B64JSON string `json:"b64_json"`
	} `json:"data"`
}

var _ Generator = (*OpenAIGenerator)(nil)

func NewOpenAIGenerator(apiKey string, model string) OpenAIGenerator {
	return OpenAIGenerator{
		urlBase: "https://api.openai.com/v1/images/",
		model:   model,
		apiKey:  apiKey,
	}
}

func (g *OpenAIGenerator) do(endpoint string, request []byte, extraHeaders map[string]string) (openAIResponse, error) {
	req, err := http.NewRequest(http.MethodPost, g.urlBase+endpoint, bytes.NewReader(request))
	if err != nil {
		return openAIResponse{}, fmt.Errorf("failed to generate request. %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+g.apiKey)
	for k, v := range extraHeaders {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return openAIResponse{}, fmt.Errorf("failed to perform request. %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return openAIResponse{}, fmt.Errorf("failed to read body. %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return openAIResponse{}, fmt.Errorf("unexpected response code: %d. %s. %s", resp.StatusCode, resp.Status, string(body))
	}
	var r openAIResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return openAIResponse{}, fmt.Errorf("failed to unmarshal response. %w", err)
	}
	return r, nil
}

func (g *OpenAIGenerator) GenerateImageFromText(request models.TextToImageRequest) (string, error) {
	req := openAIRequest{
		Model:          g.model,
		Prompt:         strings.Join(request.TextPrompts, "\n"),
		Count:          request.Samples,
		Size:           fmt.Sprintf("%dx%d", request.Width, request.Height),
		ResponseFormat: "b64_json",
	}
	bytes, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal input. %w", err)
	}
	resp, err := g.do("generations", bytes, nil)
	if err != nil {
		return "", err
	}
	return resp.Data[0].B64JSON, nil
}

func (g *OpenAIGenerator) GenerateImageFromImage(request models.ImageToImageRequest) (string, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	partHeader := make(textproto.MIMEHeader)
	partHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="image"; filename="%s"`, "image.png"))
	partHeader.Set("Content-Type", "image/png")
	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, bytes.NewReader(request.Image))
	if err != nil {
		return "", fmt.Errorf("failed to copy image to buffer. %w", err)
	}
	if err := writer.WriteField("model", g.model); err != nil {
		return "", err
	}
	if err := writer.WriteField("prompt", strings.Join(request.TextPrompts, "\n")); err != nil {
		return "", err
	}
	samples := strconv.Itoa(request.Samples)
	if err := writer.WriteField("n", samples); err != nil {
		return "", err
	}
	if err := writer.WriteField("size", fmt.Sprintf("%dx%d", request.Width, request.Height)); err != nil {
		return "", err
	}
	if err := writer.WriteField("response_format", "b64_json"); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}
	payload, err := io.ReadAll(bufio.NewReader(&body))
	if err != nil {
		return "", fmt.Errorf("failed to read buffer. %w", err)
	}
	resp, err := g.do("edits", payload, map[string]string{"Content-Type": writer.FormDataContentType()})
	if err != nil {
		return "", err
	}
	return resp.Data[0].B64JSON, nil
}
