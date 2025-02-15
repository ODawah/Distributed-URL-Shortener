# Distributed URL Shortener with Rate Limiting & Analytics

## Overview
This project is a distributed URL shortener service built using Golang, Gin, Redis, and MongoDB. It provides short URLs (8-character unique IDs) for long URLs, implements rate limiting, and logs analytics in MongoDB.

## Build and Run
To build and run the application, use Docker Compose:
```sh
# Build the containers
docker-compose build

# Start the API service
docker-compose up
```


## API Endpoints
- `POST /shorten` - Accepts a long URL and returns a shortened version.
- `GET /{shortID}` - Redirects to the original URL and logs analytics in MongoDB.

## Usage
#### Shorten a URL

###### POST /shorten

###### Request Body:
```json
{
"url": "www.example.com"
}
```

#### Success Response:
```json
{
"url": "www.h.com",
"id": "p1_dU71U"
}
```
#### Error Responses:

- 400 Bad Request: Invalid or missing URL.

- 500 Internal Server Error: Unexpected server issue.

- Retrieve a Shortened URL

###### GET /{shortID}

###### Success Response:
```json
{
"url": "www.h.com",
"id": "p1_dU71U"
}
```

- The request metadata (shortID, timestamp, user IP) is logged in MongoDB.

### Error Responses:

- 400 Bad Request: If the provided short ID is less than 8 characters.

- 404 Not Found: If the short ID does not exist.


## Features
- Shorten long URLs and retrieve the original URL using an 8-character unique short ID.
- Store URL mappings in Redis with an expiry of 30 days for unused URLs.
- Implement rate limiting (10 requests per minute per user).
- Log each request (shortID, timestamp, user IP) in MongoDB asynchronously using worker goroutines.
- Deployable using Kubernetes with Nginx.

## Use Cases
- User Send URL to be Shortened
- User hits a shortened URL


## Project Structure
```
-- models
  ├── requests.go      # Model for logging requests to MongoDB
  ├── url.go           # Model for shortened URLs

-- handlers
  ├── url              # Handlers for API endpoints

-- middlewares
  ├── limiter          # Middleware for rate limiting
  ├── logger           # Middleware for logging requests to MongoDB

-- persistence
  ├── mongo            # MongoDB initialization
  ├── redis            # Redis initialization

-- server
  ├── router           # Gin router setup
  ├── routes           # API route definitions

-- services
  ├── url              # Business logic for shortening and retrieving URLs
```

## API
We will use REST API Architecture using Gin framework as it's simple, flexible and scalable since it's stateless and easy to use
The API will contain One service:
- URL Service

### URL Service
The URL service has two functionalities.
1. URL Shortening:
    - When a user submits a long URL to the service, it undergoes a shortening process.
    - We First Validate if the user have sent a valid url or not to avoid having unwanted data.
    - URL is shortened using NanoID package that works with random generating.
    - Once a unique short URL is obtained, it is stored in redis and mapped to the original long URL.

2. Get the Original URL:

    - When a user accesses a shortened URL, the service handles the lookup.
    - First, the service checks if the shortened URL valid and equals to 8 characters as we store.
    - The then searches for it in Redis.
    - If the URL is found in Redis
    - Send The Request metadata to MongoDB for analytics

## Rate Limiting

- Rate limiting is implemented to control excessive API requests and ensure fair usage. The middleware:

- Allows up to 10 requests per minute per IP.

- Maintains a map of limiters per IP address.

- Periodically cleans up inactive limiters to free memory.

- Responds with HTTP 429 (Too Many Requests) when the limit is exceeded.

Implementation

- Each request's IP address is checked against a stored rate limiter. If no limiter exists, a new one is created. The limiter allows one request every 6 seconds, up to a burst of 10 requests.
- We implemented using bucket token solution so that if the user used all his tokens at once he doesn't need to wait whole minute to start again it's so on average user got request around every 6 seconds limiting him to 10 per minute
- A background goroutine cleans up limiters for inactive users every 5 minutes to optimize memory usage.




