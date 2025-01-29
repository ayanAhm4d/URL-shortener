
## 🌐 Socials:
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230077B5.svg?logo=linkedin&logoColor=white)](https://linkedin.com/in/www.linkedin.com/in/ayanahmad15) [![X](https://img.shields.io/badge/X-black.svg?logo=X&logoColor=white)](https://x.com/ayanAhm4d) 

# 💻 Tech Stack:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)

# 📊 GitHub Stats:
![](https://github-readme-stats.vercel.app/api?username=ayanAhm4d&theme=dark&hide_border=false&include_all_commits=false&count_private=false)<br/>
![](https://github-readme-streak-stats.herokuapp.com/?user=ayanAhm4d&theme=dark&hide_border=false)<br/>
![](https://github-readme-stats.vercel.app/api/top-langs/?username=ayanAhm4d&theme=dark&hide_border=false&include_all_commits=false&count_private=false&layout=compact)

---
[![](https://visitcount.itsvg.in/api?id=ayanAhm4d&icon=0&color=0)](https://visitcount.itsvg.in)

<!-- Proudly created with GPRM ( https://gprm.itsvg.in ) -->
[https://roadmap.sh/projects/url-shortening-service]

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



## Project Structure


```
url-shortener/
├── .env
├── go.mod
├── go.sum
├── main.go
├── api/
     ├── database/
     │   ├── database.go
     │   └── schema.sql
     ├── helpers/
     │   └── helpers.go
     └── routes/
         ├── resolve.go
         └── shorten.go
```
## Environment Variables

Create a .env file in the project root directory with the following variables:
```
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=url_shortener
DOMAIN=localhost:3000
```


## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ayanAhm4d/URL-shortener.git
   ```
2. Database Setup.
```
   
   CREATE DATABASE IF NOT EXISTS url_shortener;
   USE url_shortener;

   CREATE TABLE urls (
    id INT AUTO_INCREMENT PRIMARY KEY,
    short_url VARCHAR(255) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at DATETIME NOT NULL
   );

   CREATE TABLE rate_limits (
    id INT AUTO_INCREMENT PRIMARY KEY,
    client_ip VARCHAR(45) NOT NULL,
    remaining INT NOT NULL,
    reset_at DATETIME NOT NULL,
    UNIQUE KEY unique_ip (client_ip)
   );

```
   
3. Install dependencies:
   ```
   go mod tidy
   ```


1. Start the server:
   ```
   go run main.go
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

Contributing
If you'd like to contribute, feel free to open an issue or submit a pull request.

Contact

For any queries, reach out at www.ayan007ahmad@gmail.com.
