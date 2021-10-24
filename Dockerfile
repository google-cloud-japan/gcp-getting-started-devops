## 1st builder for backend
# Use golang image as a builder for backend
FROM golang:1.16.7-alpine3.14 AS backend_builder

# Create and set workdir
WORKDIR /backend

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY ./backend/go.mod ./backend/go.sum ./

# Install git to be used "go mod download"
RUN apk add --no-cache git

# Download dependencies
RUN go mod download

# Copy all files and build an executable
COPY ./backend .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o devops_handson

## 2nd builder for frontend
# Use node image as a builder for frontend
FROM node:14.17.5-alpine3.14 AS frontend_builder

# Create and set workdir
WORKDIR /frontend

# Copy `package*.json` to invalidate the next layer
# in case of a change in the dependencies
COPY ./frontend/package*.json ./

# Install dependencies
RUN npm install

# Copy all files for frontend and compiles an Angular app
COPY ./frontend .
RUN npm run build

## Image to run the backend and the frontend
# Use a Docker multi-stage build to create a lean production image
FROM alpine:3.14.1
RUN apk add --no-cache ca-certificates
COPY --from=backend_builder /backend/devops_handson ./
COPY --from=frontend_builder /frontend/dist/api-counter ./static
EXPOSE 8080
ENTRYPOINT ["/devops_handson"]
