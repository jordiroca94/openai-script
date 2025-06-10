# Recommendation Script

A minimal Go script that uses a prompt template and OpenAI's GPT-4o model to generate high-quality movie or show recommendations based on an input title.

---

## Features

- Loads a customizable prompt template from `prompts/prompt.tpl`
- Sends the rendered prompt to OpenAI's GPT API (using GPT-4o)
- Returns **3 curated** movies or shows recommendations in structured JSON
- Measures and displays execution time for the request
- Logs OpenAI token usage

---

## Quick Start

### 1. Clone and Setup

```bash
git clone https://github.com/your-repo/openai-go-script.git
cd openai-go-script
```

### 2. Set OpenAI API Key

export OPENAI_API_KEY=your-api-key-here

### 3. Run the Script

Build and run the Go script with your desired input title. It will render the prompt, call the GPT-4o model, and return recommendations.

## Output Format

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
