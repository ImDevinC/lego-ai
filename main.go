package main

import (
	"encoding/base64"
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
	oai := imagegenerators.NewOpenAIGenerator(apiKey, "gpt-4o")
	req := models.NewImageToImageRequest(
		512,
		512,
		"https://platform.vox.com/wp-content/uploads/sites/2/chorus/author_profile_images/195199/Screen_Shot_2021-12-13_at_12.25.33_PM.0.png?quality=90&strip=all&crop=0%2C0%2C100%2C100&w=256",
	)
	resp, err := oai.GenerateImageFromImage(req)
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
