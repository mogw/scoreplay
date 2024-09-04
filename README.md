# ScorePlay Media Indexing API

## Overview

This project is an HTTP API designed to help sports organizations manage their media content by tagging and indexing photos. It allows creating tags, uploading media, and retrieving specific media items.

## Features

- Create and list tags
- Upload and retrieve media with associated tags
- Elasticsearch integration for future search capabilities

## Technology Stack

- Golang
- PostgreSQL
- Elasticsearch
- Docker

## Setup Instructions

### Prerequisites

- Docker and Docker Compose installed
- Golang installed

### Running the Application

1. Clone the repository:
```bash
git clone 
cd scoreplay
```

2. Build and run the application using Docker Compose:
```bash
docker compose up --build
```
This will start the application, PostgreSQL database, and Elasticsearch.

3. Access the API:
The API will be running on http://localhost:8080.

### API Endpoints
1. Create a Tag
- Endpoint: POST /tags
- Request: { "name": "string" }
- Response: 201 Created

2. List All Tags
- Endpoint: GET /tags
- Response: 200 OK

3. Create Media
- Endpoint: POST /media
- Request: Form-data with fields name (string), tags (array of tag IDs), and file (binary).
- Response: 201 Created

4. Retrieve a Media
- Endpoint: GET /media/{id}
- Response: 200 OK

### Running Tests
To run the tests, use the following command:
```bash
go test ./tests
```

### Future Improvements
- Implement pagination and filtering for listing tags and media.
- Implement user authentication and authorization.
- Add caching for frequently accessed data.
