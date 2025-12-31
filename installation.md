# Chi router
go get -u github.com/go-chi/chi/v5

# Chi's JWT middleware (jwtauth)
go get -u github.com/go-chi/jwtauth/v5

# Chi render (useful for JSON responses)
go get -u github.com/go-chi/render

# Chi CORS middleware
go get -u github.com/go-chi/cors

# JWT
go get github.com/golang-jwt/jwt/v5

# pgx (recommended - pure Go, better performance)
go get -u github.com/jackc/pgx/v5

# Install sqlc CLI
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# godotenv for environment variables
go get -u github.com/joho/godotenv

# bcrypt for password hashing
go get -u golang.org/x/crypto/bcrypt

# validator for struct validation
go get -u github.com/go-playground/validator/v10
