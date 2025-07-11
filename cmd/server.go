package main

import (
	"log/slog"
	"os"

	"github.com/imdevinc/lego-ai/pkg/app"
	"github.com/imdevinc/lego-ai/pkg/imagegenerators"
	_ "github.com/joho/godotenv/autoload"
)

const (
	templatesDir = "templates"
	port         = "8080"
)

func main() {
	systemPrompt, err := os.ReadFile("prompts/system.txt")
	if err != nil {
		slog.Error("failed to get systemprompt", "error", err)
		os.Exit(1)
	}
	userPrompt, err := os.ReadFile("prompts/user.txt")
	if err != nil {
		slog.Error("failed to get userprompt", "error", err)
		os.Exit(1)
	}
	legoPrompt, err := os.ReadFile("prompts/lego.txt")
	if err != nil {
		slog.Error("failed to get legoprompt", "error", err)
		os.Exit(1)
	}
	imageModel := os.Getenv("LEGOAI_IMAGE_MODEL")
	if imageModel == "" {
		slog.Error("missing LEGOAI_IMAGE_MODEL")
		os.Exit(1)
	}
	chatModel := os.Getenv("LEGOAI_CHAT_MODEL")
	if chatModel == "" {
		slog.Error("missing LEGOAI_CHAT_MODEL")
		os.Exit(1)
	}
	generator := imagegenerators.NewOpenAIGenerator(imageModel, chatModel, string(systemPrompt), string(userPrompt), string(legoPrompt))
	server := app.Server{
		Port:        8080,
		TemplateDir: "templates",
		Generator:   &generator,
	}
	if err := server.Start(); err != nil {
		slog.Error(err.Error())
	}
}
