# LearnGo

A personal repository to document my Go learning journey.

## Resources

Here are some excellent resources I'm using to learn Go:

- 📘 [Learn Go with Tests (quii)](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world) — Test-driven approach to learning Go.
- 🚀 [A Tour of Go](https://go.dev/tour/welcome/1) — Official interactive Go tour from the creators of Go.
- 💡 [Go by Example](https://gobyexample.com/values) — Practical, example-driven Go snippets and explanations.

---

# Go Learning Path: Backend Developer Edition 🐹

This guide outlines a structured path to learn **Go (Golang)** with the goal of eventually building web applications and microservices using the **Gin** framework. Whether you're transitioning from Python or expanding your backend dev toolkit, this plan will help you level up with practical, real-world skills.

---

## 📌 Learning Phases Overview

---

### **Phase 1: Go Fundamentals (1–2 Weeks)**

**Goals:**
- Learn the syntax and structure of Go
- Build small command-line programs

**Topics:**
- Variables, constants, data types
- Control flow: if, for, switch
- Functions (named returns, variadic)
- Arrays, slices, maps
- Structs and methods
- Pointers
- Packages and imports

---

### **Phase 2: Intermediate Go (1–2 Weeks)**

**Goals:**
- Understand Go’s memory and type system
- Master interfaces and concurrency

**Topics:**
- Error handling idioms
- Interfaces and type assertions
- Goroutines and channels
- `defer`, `panic`, and `recover`
- Go modules and dependency management

**Free Resources:**
- [GolangBot Learn Series](https://golangbot.com/learn-golang-series/)
- [Effective Go](https://golang.org/doc/effective_go.html)

---

### **Phase 3: Practical Go Projects (2–3 Weeks)**

**Goals:**
- Build mini-projects to reinforce skills
- Apply Go to files, APIs, and data formats

**Project Ideas:**
- CLI task manager or calculator
- JSON-based HTTP API with `net/http`
- Concurrent web scraper

**Free Resources:**
- [Go Web Programming Book](https://github.com/astaxie/build-web-application-with-golang)
- [Go Patterns](https://github.com/tmrts/go-patterns)

---

### **Phase 4: Web Development in Standard Library (1–2 Weeks)**

**Goals:**
- Understand how Go handles web servers
- Learn routing, request/response handling

**Topics:**
- `net/http` fundamentals
- Handlers and routers
- Writing middleware
- JSON parsing and rendering

**Free Resources:**
- [Go Web Examples](https://gowebexamples.com/)
- [JustForFunc (YouTube)](https://www.youtube.com/playlist?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ)

---

### **Phase 5: Gin Framework (Optional: 2+ Weeks)**

**Goals:**
- Build production-ready REST APIs using Gin
- Understand middleware, routing, validation, and auth

**Topics:**
- Routing and route parameters
- Middleware
- Request binding and validation
- JWT Auth
- Database integration (GORM, sqlx)

**Free Resources:**
- [Gin Official Docs](https://gin-gonic.com/docs/)
- [Gin Crash Course (YouTube)](https://www.youtube.com/watch?v=HvplcTy4VHs)

---

## 📆 Suggested Weekly Timeline

| Week | Focus Area                     | Deliverables                                 |
|------|--------------------------------|----------------------------------------------|
| 1    | Go Fundamentals                | CLI Calculator or To-Do App                  |
| 2    | Intermediate Go + Concurrency | Web scraper with goroutines and channels     |
| 3    | Practical Projects             | Basic HTTP API using `net/http`              |
| 4    | Web Dev in Standard Library    | Build a JSON CRUD Blog Server                |
| 5+   | Gin Framework (optional next) | REST API using Gin + PostgreSQL              |

---

## ✅ Notes & Tips

- Stick to writing Go code daily, even just 30 minutes.
- Don’t skip writing tests! `go test` is easy and helps build confidence.
- Use `go fmt`, `go vet`, and `golint` to get used to Go’s tooling early.
- Join the [Gophers Slack Community](https://invite.slack.golangbridge.org/) for help.

---

Happy Go-ing! 🚀🐹
