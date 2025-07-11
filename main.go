package main

import (
	"encoding/base64"
	"io"
	"log"
	"os"

	"github.com/imdevinc/ghibli-ai/pkg/imagegenerators"
	"github.com/imdevinc/ghibli-ai/pkg/models"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("missing OPENAI_API_KEY")
	}
	input, err := os.Open("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	bytes, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	oai := imagegenerators.NewOpenAIGenerator(apiKey, "dall-e-2")
	resp, err := oai.GenerateImageFromImage(models.NewImageToImageRequest(
		"Studio Ghibli Style",
		"Comic Book",
		7, 512, 512, 1, 30,
		bytes,
	))
	//resp, err := oai.GenerateImageFromText(models.TextToImageRequest{
	//	TextPrompts: []string{"A cute baby sea otter"},
	//	CfgScale:    7,
	//	Height:      512,
	//	Width:       512,
	//	Samples:     1,
	//	Steps:       30,
	//})
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	dec, err := base64.StdEncoding.DecodeString(resp)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(dec)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}
