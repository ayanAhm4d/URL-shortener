
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

# URL Shortener

A simple URL shortener service built with Go (Golang) using the Repository Pattern, Singleton Pattern, and Dependency Injection. This project demonstrates how to create a basic URL shortener service with clean and maintainable code.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)


## Features

- Shorten long URLs
- Redirect to the original URL using the shortened link
- In-memory storage (can be replaced with database storage)
- Follows Repository Pattern and Singleton Pattern for clean architecture
- Easy-to-extend and maintain

## Project Structure


- **main.go**: Initializes the application, sets up routes, and starts the server.
- **api/handlers.go**: Contains the HTTP request handlers for shortening URLs and redirecting to original URLs.
- **shortener/service.go**: Core logic for generating and managing short URLs.
- **storage/storage.go**: In-memory storage for URLs (can be replaced with any persistent storage like a database).

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ayanAhm4d/URL-shortener.git
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Run the application:
   ```
   go run main.go
   ```

## Usage
You can use tools like curl, Postman, or your browser to interact with the API.

Shorten a URL
Make a POST request to shorten a URL:
```
curl -X POST http://localhost:1000/shorten -d '{"url": "https://github.com/"}' -H "Content-Type: application/json"
```

Redirect to Original URL
```
curl -X GET "http://localhost:1000/redirect?url={shortCode}"
```


Contributing
If you'd like to contribute, feel free to open an issue or submit a pull request.

Fork the repository
Create your feature branch (git checkout -b feature/my-new-feature)
Commit your changes (git commit -am 'Add some feature')
Push to the branch (git push origin feature/my-new-feature)
Create a new Pull Request
