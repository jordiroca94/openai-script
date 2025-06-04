package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/joho/godotenv"
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

	movieName := "Enemy"

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
		Model: "gpt-4",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: renderedPrompt.String(),
			},
		},
		MaxTokens: 1500,
	}

	// Timer goroutine
	start := time.Now()
	done := make(chan struct{})

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				elapsed := time.Since(start).Seconds()
				fmt.Printf("\rWaiting for response… %.1fs", elapsed)
			case <-done:
				return
			}
		}
	}()

	// Make the OpenAI API call
	resp, err := client.CreateChatCompletion(ctx, chatReq)
	close(done)   // Stop timer
	fmt.Println() // Move to new line
	if err != nil {
		log.Fatalf("OpenAI API error: %v\n", err)
	}

	// Print result
	if len(resp.Choices) > 0 {
		fmt.Println("=== 🎬 OpenAI Movie Recommendations ===")
		fmt.Println(resp.Choices[0].Message.Content)
	} else {
		fmt.Println("No choices returned.")
	}
}
