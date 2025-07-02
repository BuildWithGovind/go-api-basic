# 🧰 BuildwithGovind

# 🧰 Basic Go API Setup (Beginner Friendly)

# http://localhost:8080/users

This is a minimal Go API starter template for beginners, using:

- ⚡ Gin (Fast HTTP web framework)
- 🛢️ GORM (ORM for MySQL)
- 🔐 godotenv (Load `.env` variables)

---

# 🚀 Basic Go API Setup (Beginner Friendly & Interview-Ready)

This is a clean and minimal **GoLang REST API** boilerplate designed for **absolute beginners** who want to:

- Learn Go with real-world API structure
- Build their first backend service in Go
- Understand Go for job interviews
- Use best practices for long-term maintainability

---

## 👤 About Me

Hi, I'm **Govind** — a backend developer with 4+ years of experience in Node.js, Go, PHP, and full-stack tools.  
I’m passionate about clean architecture, performance, and designing scalable systems using modern tech like Go, Elasticsearch, and microservices.

📌 Follow my journey or reach out: [github.com/BuildwithGovind](https://github.com/BuildwithGovind)

---

## 🔧 Tech Stack

- ⚡ **Gin** – Minimal & fast Go web framework
- 🛢️ **GORM** – Elegant ORM for MySQL
- 🔐 **godotenv** – Load `.env` files easily

---

## 🔧 Setup

```bash
git clone https://github.com/BuildwithGovind/go-api-basic.git
cd go-api-basic
go mod tidy
go run main.go




## 📁 Folder Structure

go-basic-api-setup/
├── .env
├── go.mod
├── go.sum
├── main.go
├── models/
│   └── user.go
├── routes/
│   └── user.go
├── controllers/
│   └── userController.go
├── database/
│   └── connection.go
├── README.md



go mod init go-basic-api-setup
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/joho/godotenv


DB_USER=root
DB_PASS=password
DB_NAME=testdb
DB_HOST=127.0.0.1
DB_PORT=3306


🧠 What This Project Does
    Sets up a basic REST API in Go using Gin (web framework)
    Connects to MySQL using GORM (ORM)
    Uses .env file to keep secrets/configs separate
    Structure is easy to scale for real projects
```
