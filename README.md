# 📄 Resume Upload & Review Management System

A RESTful Resume Upload & Review Management System developed using **Golang (Gin Framework)** and **MongoDB**. The application enables candidates to upload resumes while allowing administrators to assign reviewers, track review status, download resumes, and manage reviewer information.

---

## ✨ Features

* Upload resumes to the server
* Store resume metadata in MongoDB
* Retrieve all uploaded resumes
* View resume details by ID
* Download uploaded resume files
* Delete resumes and associated files
* Assign reviewers to resumes
* Update resume review status
* Manage reviewer information
* RESTful API architecture
* Dockerized application for easy deployment

---

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Framework:** Gin
* **Database:** MongoDB
* **Containerization:** Docker & Docker Compose
* **Configuration:** JSON Configuration
* **API Testing:** Postman / Apidog

---

## 📂 Project Structure

```text
Resume_upload_system/
│
├── controller/
│   ├── controller.go
│   └── reviewer_controller.go
│
├── model/
│   ├── resume.go
│   ├── reviewer.go
│   └── response.go
│
├── upload/
│   └── up.go
│
├── set_up/
│   └── setup.go
│
├── Setting/
│   ├── setting.go
│   └── settings.jsonc
│
├── Dockerfile
├── docker-compose.yml
├── main.go
├── go.mod
└── README.md
```

---

## 🚀 Getting Started

### Clone the Repository

```bash
git clone https://github.com/Julekha23/Resume_upload_system.git
cd Resume_upload_system
```

---

## Run with Docker

```bash
docker compose up --build
```

---

## Run Without Docker

Install dependencies:

```bash
go mod tidy
```

Run the application:

```bash
go run main.go
```

---

## 📡 REST API Endpoints

| Method | Endpoint             | Description          |
| ------ | -------------------- | -------------------- |
| GET    | `/`                  | Welcome endpoint     |
| PUT    | `/resume`            | Upload a resume      |
| GET    | `/resume`            | Retrieve all resumes |
| GET    | `/resume/:id`        | Get resume by ID     |
| DELETE | `/resume/:id`        | Delete resume        |
| GET    | `/resume/:id/file`   | Download resume      |
| PATCH  | `/resume/:id/assign` | Assign reviewer      |
| PATCH  | `/resume/:id/status` | Update review status |
| POST   | `/reviewer`          | Add reviewer         |
| GET    | `/reviewer`          | Get all reviewers    |

---

## 🗄️ Database

MongoDB stores:

* Resume metadata
* Uploaded file path
* Original file name
* Review status
* Assigned reviewer
* Reviewer information

---

## 📦 Docker

Start the application using Docker Compose:

```bash
docker compose up --build
```

Stop the containers:

```bash
docker compose down
```

---

## Future Improvements

* JWT Authentication
* Role-Based Access Control
* Resume search and filtering
* Email notification system
* Resume versioning
* Resume parsing and AI-based screening
* Swagger/OpenAPI documentation
* Unit and Integration Tests

---

## Author

**Mosammat Julekha Molla**

Backend Developer | Golang Developer

GitHub: https://github.com/Julekha23

---

## License

This project is developed for learning and portfolio purposes.
