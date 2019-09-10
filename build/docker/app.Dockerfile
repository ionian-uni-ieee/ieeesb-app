FROM golang:1.13.0-alpine3.10

# Default environment variables
ENV API_HOST=app
ENV API_PORT=5000
ENV API_DATABASE_HOST=mongodb://mongo
ENV API_DATABASE_NAME=local

# Create directory for application's sources
RUN mkdir -p /opt/go-app
WORKDIR /opt/go-app

# Get main file
COPY ../../main.go ./

# Copy packages
COPY ../../pkg ./pkg 

# Get go module files
COPY ../../go.mod ../../go.sum ./

# Copy vendors
COPY ../../vendor ./vendor

# Copy application
COPY ../../internal/app ./internal/app

EXPOSE 5000

CMD ["go", "run", "./main.go"]
