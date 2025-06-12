# 🎬 OpenAI Movie Recommendation Script

A modular Go application that leverages OpenAI's GPT-4o model to generate high-quality movie recommendations based on an input title. The application uses customizable prompt templates and provides structured JSON output.

## ✨ Features

- **Modular Architecture**: Separation of concerns with abstracted OpenAI client logic
- **Customizable Templates**: Loads prompt templates from `prompts/prompt.tpl`
- **Advanced Prompting**: Uses role-playing and structured output format for optimal results
- **Performance Monitoring**: Real-time progress indicator and execution time tracking
- **Token Usage Analytics**: Logs detailed OpenAI API token consumption
- **Structured Output**: Returns recommendations in clean, parseable JSON format

## 🚀 Quick Start

### Prerequisites

- Go 1.18 or higher
- OpenAI API key

### 1. Clone the Repository

```bash
git clone https://github.com/jordiroca94/openai-go-script.git
cd openai-go-script
```

### 2. Set Up Your API Key

Create a `.env` file in the root directory:

```bash
echo "OPENAI_API_KEY=your-api-key-here" > .env
```

Alternatively, export it directly in your shell:

```bash
export OPENAI_API_KEY=your-api-key-here
```

### 3. Run the Script

```bash
go run ./cmd/main.go
```

## 🎯 How It Works

1. The application loads your OpenAI API key from environment variables
2. It reads the prompt template from `prompts/prompt.tpl`
3. The template is rendered with your specified movie title
4. The prompt is sent to OpenAI's GPT-4o model through the client
5. The response is processed and displayed in a structured format

## 📝 Customizing the Prompt Template

Edit `prompts/prompt.tpl` to customize how recommendations are generated. The template uses Go's template syntax and currently supports the following variables:

- `{{.MovieName}}` - The name of the movie for which recommendations are sought

## 📊 Sample Output

```json
[
  {
    "title": "The Green Mile",
    "year": 1999,
    "rating": 8.6,
    "type": "movie",
    "genre": ["Drama", "Crime", "Fantasy"],
    "plot": "A death row corrections officer forms a unique bond with a mysterious inmate who possesses a supernatural gift.",
    "tmdb_id": 497
  },
  {
    "title": "Schindler's List",
    "year": 1993,
    "rating": 8.9,
    "type": "movie",
    "genre": ["Drama", "History", "War"],
    "plot": "The true story of Oskar Schindler, a German industrialist who saves the lives of over a thousand Polish Jews during the Holocaust by employing them in his factories.",
    "tmdb_id": 424
  },
  {
    "title": "American History X",
    "year": 1998,
    "rating": 8.5,
    "type": "movie",
    "genre": ["Drama"],
    "plot": "A former neo-nazi skinhead tries to prevent his younger brother from going down the same wrong path that he did.",
    "tmdb_id": 73
  }
]
```

## 🧩 Extending the Application

### Adding New Templates

Create additional template files in the `prompts/` directory and modify the `main.go` file to use your new template:

```go
content, usage, err := client.ProcessPromptTemplate(ctx, "prompts/your-new-template.tpl", data)
```

### Customizing the OpenAI Client

You can modify the OpenAI client options in `main.go`:

```go
client := openai.NewClient(openai.ClientOptions{
    APIKey:    apiKey,
    Model:     "gpt-4-turbo",  // Change model
    MaxTokens: 2000,           // Increase token limit
})
```

## 📄 License

MIT License

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
