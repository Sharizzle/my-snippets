# My Snippets

### Go-SQL-React Admin-Nuxt-Vue-Project

A Full Stack Web Application that is used to display code snippets for quick ease of use.

## Technologies Used

### Front-end

- Vue 3
- Nuxt (Server Side Rendering)
- Babel
- Webpack
- Prettier
- Highlight.js

### Back-end

- Mux (Routing)
- Crypto (Hashing passwords)
- Gorm (Go ORM)
- SQL Lite
- Cors Library

### Admin Panel

- React Admin

## Routes

Basic CRUD Routes for User Data

- Show Codes `GET /codes`
- Create Code `POST /codes`
- Show One Code `GET /codes/{codeId}`
- Delete Code `DELETE /codes/{codeId}`
- Update Code `PUT /codes/{codeId}`

## Configuration

Create a `.env` file with the following parameters.

PORT = ...  
ENVIRONMENT = ...

## Getting Started

### Server

Run the command to install all dependancies and start server.

```go
go mod download
go run main.go
```

### Nuxt Client

```bash
yarn dev
```

### React Admin Dashboard

```bash
yarn start
```

## Authors

- **Sharif Kanaan** - [GitHub](https://github.com/Sharizzle) | | [Website](https://sharif.thekanaan.com/) | | [Email](mailto:sharif@thekanaan.com) || [Linkedin](https://www.linkedin.com/in/SharifKanaan/)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
