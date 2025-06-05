package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	timer "github.com/jordiroca94/openai-go-script/internal"
	openai "github.com/sashabaranov/go-openai"
)

//  RUN script --> go run ./cmd/main.go

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or failed to load:", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalln("OPENAI_API_KEY environment variable not set")
	}

	tplBytes, err := os.ReadFile("prompts/prompt.tpl")
	if err != nil {
		log.Fatalf("Failed to read prompt: %v\n", err)
	}

	tpl, err := template.New("prompt").Parse(string(tplBytes))
	if err != nil {
		log.Fatalf("Failed to parse prompt: %v\n", err)
	}

	movieName := "Shawshank Redemption"

	var renderedPrompt bytes.Buffer
	err = tpl.Execute(&renderedPrompt, struct {
		MovieName string
	}{
		MovieName: movieName,
	})
	if err != nil {
		log.Fatalf("Failed to execute prompt template: %v\n", err)
	}

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	chatReq := openai.ChatCompletionRequest{
		Model: "gpt-4o",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: renderedPrompt.String(),
			},
		},
		MaxTokens: 1000,
	}

	stop := timer.StartTimer()

	resp, err := client.CreateChatCompletion(ctx, chatReq)
	stop()

	fmt.Println()
	if err != nil {
		log.Fatalf("OpenAI API error: %v\n", err)
	}

	// Log token usage
	if resp.Usage.TotalTokens > 0 {
		fmt.Printf("\n=== 🔢 Token Usage ===\n")
		fmt.Printf("Prompt Tokens: %d\n", resp.Usage.PromptTokens)
		fmt.Printf("Completion Tokens: %d\n", resp.Usage.CompletionTokens)
		fmt.Printf("Total Tokens: %d\n", resp.Usage.TotalTokens)
	}

	if len(resp.Choices) > 0 {
		fmt.Println("=== 🎬 OpenAI Movie Recommendations ===")
		fmt.Println(resp.Choices[0].Message.Content)
	} else {
		fmt.Println("No choices returned.")
	}
}
