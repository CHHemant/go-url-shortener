# Go URL Shortener API

## What is it?

Go URL Shortener API is a lightweight, fast, and simple RESTful service written in Go to convert long URLs into short codes. Users can share compact links, which redirect to the original addresses. It’s ideal for learning backend development, REST APIs, and Go concurrency, and can be used as a foundation for more advanced URL management systems.

---

## Why use this project?

- **Simplicity:**  
  Minimal setup; easy to understand for beginners and easy to extend for advanced users.

- **Performance:**  
  Built in Go, known for its speed and efficiency, especially under concurrent loads.

- **Modularity:**  
  Clean separation of logic makes it simple to add features like persistent storage, analytics, or custom domains.

- **Learning value:**  
  Teaches RESTful API design, URL validation, hash-based code generation, concurrent-safe data access, and modular Go programming.

- **Open source:**  
  Freely available for personal, educational, or commercial use and modification.

---

## How to use it

### 1. Install Go

- Download and install Go from [golang.org/dl](https://golang.org/dl/).

### 2. Clone the repository

```bash
git clone https://github.com/CHHemant/go-url-shortener.git
cd go-url-shortener
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the server

```bash
go run main.go
```
The server will start at `http://localhost:8080`.

---

### 5. Shorten a URL

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" \
     -d '{"url":"https://example.com"}' \
     http://localhost:8080/shorten
```

**Response:**

```json
{"code": "abc1234"}
```

---

### 6. Use the short URL

Open `http://localhost:8080/abc1234` in your browser, or use curl:

```bash
curl -v http://localhost:8080/abc1234
```
You will be redirected to `https://example.com`.

---

## API Endpoints

### POST /shorten

- **Description:** Receives a long URL, returns a short code.
- **Request:**  
  `{"url": "https://..."}`

- **Response:**  
  `{"code": "shortcode"}`

### GET /{code}

- **Description:** Redirects the browser or client to the original long URL.
- **Response:** HTTP 302 redirect.

---

## Project Structure

```
go-url-shortener/
├── main.go         # Entry point, router setup
├── handlers.go     # API endpoint logic
├── storage.go      # In-memory data handling (concurrent safe)
├── models.go       # Code generation, URL validation
├── go.mod          # Dependencies
├── README.md       # Documentation
└── .gitignore      # Ignore build artifacts
```

---

## How does it work?

- **ShortenHandler:** Accepts a JSON payload with a long URL, validates it, generates a short code (using SHA1 hash + base64), and saves the mapping.
- **RedirectHandler:** Receives a short code, looks up the original URL, and responds with a redirect.
- **Storage:** Uses a Go map with sync locks for concurrency safety. Easily replaceable with file or database storage.
- **Validation:** Ensures only valid HTTP/HTTPS URLs are accepted.

---

## Who should use this?

- **Students & beginners:**  
  Learn about Go HTTP servers, REST APIs, and concurrency.

- **Educators:**  
  Use as a classroom project to teach backend fundamentals.

- **Developers:**  
  Extend into a production-grade URL shortener with analytics, authentication, or persistent storage.

- **Teams:**  
  Collaborate to add features, refactor, and learn Go together.

---

## Why use this over other solutions?

- **No dependencies on external services.**
- **No signup, no ads, no tracking.**
- **Transparent code—review, audit, and modify freely.**
- **Great starting point for learning Go and backend APIs.**
- **Well-structured, easy to maintain and expand.**

---

## Extending the Project

- Add persistent storage (BoltDB, Redis, PostgreSQL).
- Build a web front-end for user interaction.
- Implement custom codes and analytics.
- Add authentication to manage personal links.

---

## License

MIT License—free for personal and commercial use.

---

## Support & Contribution

- Open issues for bugs or feature requests.
- Fork and submit pull requests for improvements.
- Discuss ideas with collaborators.

---

Happy learning and building!
