
## 🌐 Socials:
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230077B5.svg?logo=linkedin&logoColor=white)](https://linkedin.com/in/www.linkedin.com/in/ayanahmad15) [![X](https://img.shields.io/badge/X-black.svg?logo=X&logoColor=white)](https://x.com/ayanAhm4d) 

# 💻 Tech Stack:
- Backend: Go (Gin framework)

- Database: Redis

- Containerization: Docker, Docker Compose
# 📊 GitHub Stats:
![](https://github-readme-stats.vercel.app/api?username=ayanAhm4d&theme=dark&hide_border=false&include_all_commits=false&count_private=false)<br/>
![](https://github-readme-streak-stats.herokuapp.com/?user=ayanAhm4d&theme=dark&hide_border=false)<br/>
![](https://github-readme-stats.vercel.app/api/top-langs/?username=ayanAhm4d&theme=dark&hide_border=false&include_all_commits=false&count_private=false&layout=compact)

---
[![](https://visitcount.itsvg.in/api?id=ayanAhm4d&icon=0&color=0)](https://visitcount.itsvg.in)

<!-- Proudly created with GPRM ( https://gprm.itsvg.in ) -->

# URL Shortener

The URL Shortener project is a high-performance web application built with Go, Docker and Redis that allows users to shorten URLs and redirect to the original links. This project uses Redis for fast data storage and retrieval, ensuring optimal performance.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Installation](#installation)
- [Usage](#usage)


## Features

- Shorten URLs: Generate shortened links for long URLs.

- Custom Short URLs: Users can create their custom short links.

- Automatic Expiry: URLs expire after a specified duration (default: 24 hours).

- Rate Limiting: Prevent abuse by limiting requests to 10 per user every 30 minutes.

- HTTPS Enforcement: Ensures all URLs are served with HTTP or HTTPS.

- Redis Integration: Uses Redis for high-speed storage and retrieval.

- Dockerized Deployment: Simplified deployment with Docker and Docker Compose.

## Project Structure


```
URL-shortener/
├── api/
    ├── database/
	  ├──database.go
    ├── helpers/
	  ├─ helpers.go
    ├── routes/
	  ├──resolve.go
	  ├──shorten.go
├── main.go
├── Dockerfile
├── docker-compose.yml
├── .env
└── README.md
```
## Environment Variables

Create a .env file in the project root directory with the following variables:
```
DB_ADDR=localhost:6379
DB_PASSWORD=
DOMAIN=localhost:3000
API_QUOTA=10
APP_PORT=:3000
```


## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ayanAhm4d/URL-shortener.git
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```
Using Docker

1. Build and start the services:
   ```
   docker-compose build
   ```

   ```
   docker-compose up -d
   ```
2. Access the application at http://localhost:3000.

## Usage

API Endpoints

POST /api/v1: Shorten a URL.

Request Body:
```
{
  "url": "https://example.com",
  "short": "customShort",
  "expiry": 24
}
```
Response:
```
{
  "url": "https://example.com",
  "short": "http://localhost:3000/customShort",
  "expiry": 24,
  "rate_limit": 9,
  "rate_limit_reset": 29
}
```
GET /:url: Redirect to the original URL.

Example: http://localhost:3000/4fg redirects to https://example.com.


You can use tools like curl, Postman, or your browser to interact with the API.

## Dockerfile Explanation

- Builder Stage: Builds the Go binary for the application.

- Runtime Stage: Runs the application with a minimal Alpine image to ensure a lightweight container.

## Docker-Compose Explanation

- API Service: Builds and runs the Go application.

- DB Service: Runs a Redis instance.


Contributing
If you'd like to contribute, feel free to open an issue or submit a pull request.

Contact

For any queries, reach out at www.ayan007ahmad@gmail.com.
