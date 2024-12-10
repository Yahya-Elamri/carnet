# Carnet 🚗🛠️

**GitHub Repository Description**: Carnet: A car forum and marketplace for Moroccan enthusiasts. Connect, discuss, and trade all things automotive! 🚗🛠️

---

## Features

- **Car Marketplace**: Post cars, parts, or car rental offers with detailed descriptions and images.
- **Forum Discussions**: Join communities like car repairs or car meetings and participate in engaging discussions.
- **Voting System**: Upvote or downvote posts to highlight the best content.
- **Dynamic Filters**: Find listings or threads easily using filters by price, category, or other attributes.
- **Dark/Light Mode**: Switch between themes seamlessly with HTMX integration.
- **User Profiles**: Showcase your profile with a description, profile picture, and social links.
- **Mobile-Friendly UI**: Built with Tailwind CSS for a responsive and visually appealing design.

---

## Tech Stack

- **Backend**: [Go](https://golang.org/) with the [Echo framework](https://echo.labstack.com/)
- **Frontend**: [HTMX](https://htmx.org/) and [Tailwind CSS](https://tailwindcss.com/) and [templ](https://templ.guide/)
- **Database**: PostgreSQL with GORM for ORM.
- **Media Storage**: AWS S3 for images and videos.
- **Authentication**: JWT-based user authentication.
- **DevOps**: Dockerized setup for easy deployment.

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Yahya-Elamri/carnet
   cd carnet
   
2. **Install dependencies**:
  Go: Install the latest version from golang.org
  Mysql: Install the latest version from [MySQL](https://www.mysql.com/)
  
3. **Set up environment variables**:
  Create a .env file with the following keys:
    ```bash
      JWT_TOCKEN_KEY="your jwt secret tocken"
      DB_HOST=127.0.0.1
      DB_PORT=3306
      DB_DATABASE=your database name
      DB_USERNAME=your database username
      DB_PASSWORD=your database password
    
4. **Run the application**:
   ```bash
      cd file
      make up or go run migration/cmd/main.go -direction=up 
      make run or go run cmd/main.go

5. **Access the app:**:
   ```bash
      Open http://localhost:8080 in your browser.

## Authors

- [@Elamri Yahya](https://github.com/Yahya-Elamri)

### Commit Messages

- Follow [conventional commit messages](https://www.conventionalcommits.org/).

## Reporting Issues

Found a bug or have a feature request? [Open an issue](../../issues) and provide detailed information about the problem or suggestion.

## Contact

For further assistance or questions, you can contact me at [yahya.elamri.23@ump.ac.ma].

## License


**Contact me at the email in contact section**

