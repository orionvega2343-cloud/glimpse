package api

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"google.golang.org/genai"
)

func Send(base string, prompt string) (string, error) {
	//Создание клиента Gemini
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: os.Getenv("API_KEY")})
	if err != nil {
		log.Fatal(err)
	}

	//Декодирование base64 в байты
	b, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		return "", err
	}

	model := "gemini-2.5-flash"

	//Запрос к Gemini с изображением и промптом
	resp, err := client.Models.GenerateContent(
		ctx,
		model,
		[]*genai.Content{
			{
				Parts: []*genai.Part{
					{Text: prompt},
					{InlineData: &genai.Blob{
						MIMEType: "image/png",
						Data:     b,
					}},
				},
			},
		},
		nil,
	)
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
