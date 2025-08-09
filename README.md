# 🎭 Quotify MCP: The Revolutionary AI Quote Engine

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-00ADD8.svg)](https://golang.org/doc/devel/release.html)
[![MCP Compatible](https://img.shields.io/badge/MCP-Compatible-brightgreen.svg)](https://modelcontextprotocol.io/)
[![Claude Desktop](https://img.shields.io/badge/Claude%20Desktop-Ready-orange.svg)](https://claude.ai/desktop)
[![Famingo Labs](https://img.shields.io/badge/Famingo%20Labs-Approved-ff69b4.svg)](https://github.com/almiche/quotify-mcp)
[![Quotes](https://img.shields.io/badge/Quotes-60%2B-blue.svg)](#)
[![Chaos Level](https://img.shields.io/badge/Chaos%20Level-Maximum-red.svg)](#)

> *"The ode lives upon the ideal, the epic upon the grandiose, the drama upon the real." - Satan*

## 🚀 Welcome to the Future of Wisdom Distribution

In a world where artificial intelligence has revolutionized everything from your morning coffee recommendations to predicting the next viral TikTok dance, one critical piece was missing from the AI ecosystem: **truly chaotic, beautifully random, and utterly unpredictable quote generation**.

**Famingo Labs Inc.** is proud to present the most groundbreaking advancement in quote technology since Socrates first said "I know that I know nothing" (which he probably never actually said, but that's beside the point).

### 🎪 The Genesis of Greatness

This project resurrects the legendary [**Quotify Ruby Gem**](https://github.com/jusleg/quotify-ruby) - a masterpiece of comedic quote generation that has brought joy, confusion, and occasional existential crises to developers worldwide. We've taken this Ruby gem and catapulted it into the **AI Age** with full Model Context Protocol (MCP) support.

Now, instead of manually running a Ruby script like some kind of caveman, you can seamlessly integrate the wisdom of Master Yoda, the business acumen of Logan Paul, and the philosophical depths of Satan directly into your Claude Desktop conversations. This is not just progress - this is **evolution**.

## ✨ What Makes This Revolutionary?

- 🎲 **60+ Carefully Curated Quotes** - From profound wisdom to absolute chaos
- 🎭 **29 Distinguished Authors** - Including Dog The Bounty Hunter, The Red Power Ranger, and Albus Dumbledore
- 🧠 **AI-Native Architecture** - Built with the official Anthropic MCP Go SDK
- ⚡ **Zero Configuration** - Works out of the box with Claude Desktop
- 🎪 **Maximum Chaos** - Because predictability is the enemy of innovation

## 🛠️ Installation & Setup

### Prerequisites
- Go 1.21 or higher
- Claude Desktop
- A sense of humor (mandatory)
- An appreciation for chaos (highly recommended)

### Build the Server
```bash
git clone https://github.com/almiche/quotify-mcp.git
cd quotify-mcp
go build -o bin/quotify-server cmd/quotify_server.go
```

### Configure Claude Desktop

1. Create or edit your Claude Desktop configuration file:
   - **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
   - **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

2. Add the Quotify MCP server:
```json
{
  "mcpServers": {
    "quotify": {
      "command": "/path/to/your/quotify-mcp/bin/quotify-server",
      "args": [],
      "env": {}
    }
  }
}
```

3. **Completely restart Claude Desktop** (this is crucial!)

## 🎯 Usage

Once configured, you can use the `quotify` tool directly in Claude Desktop:

```
Please use the quotify tool to generate a random quote for me!
```

### Output Formats

- **Default (Text)**: `"May the force be with you" - Master Yoda`
- **JSON Format**: 
```json
{
  "text": "Don't ever play yourself",
  "author": "Albus Dumbledore"
}
```

### Advanced Usage
```
Use the quotify tool with JSON format to get a structured quote
```

## 🎪 The Quotify Experience

Prepare yourself for profound wisdom such as:

- *"C++ supports OOP"* - attributed to Sarah Palin
- *"You can't see me"* - John Cena (ironically)
- *"Wingardium leviosa"* - The Undertaker
- *"Those were alternative facts"* - Soulja Boy
- *"i think dreams are a socialist construct"* - Judge Tyco

Each quote is randomly paired with a random author, creating a beautiful symphony of chaos that challenges conventional wisdom and breaks down the barriers between reality and absurdity.

## 🏢 About Famingo Labs Inc.

At Famingo Labs, we believe that the future belongs to those brave enough to embrace the beautiful randomness of existence. Our mission is to bring order to chaos, wisdom to confusion, and laughter to the serious business of artificial intelligence.

This project represents our commitment to pushing the boundaries of what's possible when you combine cutting-edge AI technology with the timeless art of random quote generation.

## 🤝 Contributing

We welcome contributions from fellow chaos enthusiasts! Whether you want to add more quotes, more authors, or more ways to randomly combine them, we're here for it.

## 📜 License

This project is licensed under the MIT License - because freedom should be free, and quotes should flow like water (or chaos, depending on your perspective).

---

*"you're born and then you die that's all there is to it" - Dj Khaled*

**Made with 💖 and maximum chaos by Famingo Labs Inc.**