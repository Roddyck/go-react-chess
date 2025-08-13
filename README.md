# GRChess

Minimal chess website, built mostly for my own fun and learning.

For now still in progress.

## Features
* **User authentication and authorization**: using JWT access tokens and refresh tokens
* **Game session management**: Users can create and join game sessions

* **Move validation**: server checks each move to ensure they are legal according
to chess rules

* **Real-Time updates**: Session updates using WebSocket

* **Saving games to database** (later will be added that users can watch through their games)

## Quick start

Clone repo:
```bash
git clone https://github.com/Roddyck/go-react-chess.git
```

Create `.env` file in the root of the project with following contents
```env
DB_URL=postgresql://postgres:postgres@postgres:5432/grchess?sslmode=disable
DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=grchess
TOKEN_SECRET=your-jwt-secret
```

Start docker containers with docker compose
```bash
# build only first time, btw
docker compose build
docker compose up
```

## Development
### Backend (Go)

Create `.env` file in `backend` directory with following contents
```env
PORT="8080"
TOKEN_SECRET=your-jwt-secret
DB_URL=postgres://postgres:postgres@localhost:5432/grchess?sslmode=disable
DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=grchess
```

Then inside same directory run:
```bash
go mod download
go run .
```

### Frontend (React)
```bash
cd frontend
npm install
npm run dev
```

## Tech Stack
* **Backend**: Go, net/http, Gorilla WebSocket, PostgreSQL
* **Frontend**: React, TypeScript, Vite, Tailwind CSS
* **Infrastructure**: Docker, Docker Compose

## API
### Not protected paths
* `POST /api/users`: To register users with provided name, email and password
* `POST /api/login`: To login user with email and password
* `POST /api/refresh`: To refresh access token with provided refresh token inside request headers
* `POST /api/revoke`: To revoke given refresh token

### Protected paths (requires auth)
* `GET /api/users`: To get current user info
* `POST /api/games`: To get game info with provided game id

### WebSocket and Sessions
* `POST /ws/sessions`: To create session, game created automaticly (requires authenticated users)
* `GET /ws/sessions`: To get all active sessions
* `/ws/sessions/{roomID}`: To join session

