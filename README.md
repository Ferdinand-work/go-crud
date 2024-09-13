# Go CRUD Application

This project is a simple CRUD (Create, Read, Update, Delete) application built using Go and the Gin framework. It uses MongoDB for data storage and demonstrates how to integrate these technologies in a Go application.

## Features

- **User Management**: 
  - Create a new user
  - Retrieve user details by name
  - Retrieve all users
  - Update user details
  - Delete a user by name
  - Retrieve users by age
  - Add friends to a user's friend list

## Requirements

- Go 1.18+
- MongoDB
- Gin (Go web framework)
- `joho/godotenv` (for loading environment variables)

## Setup

### Environment Variables

Create a `.env` file in the root directory of the project with the following content:

```plaintext
  MONGO_URI=<your-mongodb-uri>
  PORT=<port-to-run-the-server>
```

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-repo/go-crud.git
   cd go-crud
   ```
