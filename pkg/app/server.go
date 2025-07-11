package app

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/imdevinc/ghibli-ai/pkg/imagegenerators"
	"github.com/imdevinc/ghibli-ai/pkg/models"
)

const MAX_SIZE = 50 * 1024 * 1024 // 50MB

type Server struct {
	TemplateDir string
	Port        int
	Generator   imagegenerators.Generator
	template    *template.Template
}

func (s *Server) Start() error {
	if _, err := os.Stat(s.TemplateDir); os.IsNotExist(err) {
		return fmt.Errorf("failed to find template directory. %w", err)
	}
	tmpl, err := template.ParseGlob(s.TemplateDir + "/*.html")
	if err != nil {
		return fmt.Errorf("failed to parse templates. %w", err)
	}
	s.template = tmpl

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.indexHandler)
	mux.HandleFunc("POST /upload", s.uploadHandler)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	slog.Info("server started", "port", s.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), mux); err != nil {
		return fmt.Errorf("server failed. %w", err)
	}
	return nil
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	err := s.template.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		slog.Error("failed to execute template", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) uploadHandler(w http.ResponseWriter, r *http.Request) {
	b64image, apiKey, err := getUploadData(r)
	if err != nil {
		slog.Error("failed to get user data", "error", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	request := models.NewImageToImageRequest(apiKey, 1024, 1024, b64image)
	imageResult, err := s.Generator.GenerateImageFromImage(request)
	if err != nil {
		slog.Error("failed to generate image", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := base64.StdEncoding.DecodeString(imageResult)
	if err != nil {
		slog.Error("failed to decode image", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	reader := bytes.NewReader(data)
	if _, err := io.Copy(w, reader); err != nil {
		slog.Error("failed to write image", "error", err)
		http.Error(w, "invalid image", http.StatusInternalServerError)
		return
	}
}

func getUploadData(r *http.Request) (string, string, error) {
	err := r.ParseMultipartForm(50 << 20) // 50 MB
	if err != nil {
		return "", "", fmt.Errorf("failed to parse form. %w", err)
	}
	slog.Info("form values", "form", r.Form)
	apiKey := r.Form.Get("apiKey")
	if apiKey == "" {
		return "", "", errors.New("missing apiKey")
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		return "", "", fmt.Errorf("failed to get image from form. %w", err)
	}
	if header.Size > MAX_SIZE {
		return "", "", errors.New("image too big")

	}
	data, err := io.ReadAll(file)
	if err != nil {
		return "", "", fmt.Errorf("failed to read file. %w", err)
	}
	b64Data := base64.StdEncoding.EncodeToString(data)
	return b64Data, apiKey, nil
}
