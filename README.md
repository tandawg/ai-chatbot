# AI Chatbot (Mock Version)

A simple mock AI chatbot written in Go with Gin framework.  
Includes a web interface and an HTTP API endpoint for easy local testing and demonstration.

---

## Features

- Written in Go (Gin framework)
- Single endpoint: `/chat`
- Mock logic with preset answers
- Simple web UI (HTML + JS)
- English-language Q&A
- Supports sending messages by pressing Enter
- Logs chat history to `chatlog.txt`


---

## How to run

1. Clone the repository:

```bash
git clone https://github.com/yourusername/ai-chatbot
cd ai-chatbot
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run main.go
```

4. Open your browser:

```
http://localhost:8080/
```

Type your question and hit **Enter** or click **Send**!

---

## Example Questions

How old is the Sun?  
What is the shape of the Earth?  
Hello  

---

## License

MIT