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
	systemPrompt, err := os.ReadFile("prompts/system.txt")
	if err != nil {
		log.Fatal(err)
	}
	userPrompt, err := os.ReadFile("prompts/user.txt")
	if err != nil {
		log.Fatal(err)
	}
	oai := imagegenerators.NewOpenAIGenerator(apiKey, "dall-e-3", "gpt-4o", string(systemPrompt), string(userPrompt))
	req := models.NewImageToImageRequest(
		1024,
		1024,
		"https://lh3.googleusercontent.com/pw/AP1GczPpcHP2mtRHd3CdU42DfP9SBzAC54YSpfAhRjhabW_KvAriMwGfI2McBv3ImkaDdGStuPo62SGPJtsCHlbN9ESi8ruASLPEqXEL_b2aIKaD_bkc8TIV9vew7w0Lp__1npNJzn3Vv5QKN6LLSb0FKdXqwA=w512-h911-s-no-gm?authuser=0",
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
