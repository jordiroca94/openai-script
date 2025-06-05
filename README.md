# 🎬 OpenAI Movie Recommendation Script

A minimal Go script that uses a prompt template and OpenAI's GPT-4o to generate high-quality movie recommendations based on an input film. The recommendations mimic the style and insight of a seasoned cinema director.

---

## 📦 Features

- Loads a customizable prompt template from `prompts/prompt.tpl`
- Sends the rendered prompt to OpenAI's GPT API (using `gpt-4o`)
- Returns 3 curated movie recommendations in structured JSON
- Measures and displays request execution time
- Logs OpenAI token usage

---

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone https://github.com/your-repo/openai-go-script.git
cd openai-go-script
```
