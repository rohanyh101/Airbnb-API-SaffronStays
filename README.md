# Airbnb API (SaffronStays)

This is a Go-based API that provides two main features for Airbnb room management:

1. **Occupancy Percentage**: Returns the occupancy percentage over the next 5 months, month-on-month, for a given room ID.
2. **Price Metrics**: Returns the average night rates, highest, and lowest rates over the next 30 days for a given room ID.

## Requirements

- Go 1.18 or higher
- Docker (optional, for containerization)
- Swagger for API documentation

## Deployed Application

The application has been deployed and is available for use at:

- **Deployed Application**: [https://airbnb-api-saffronstays.onrender.com](https://airbnb-api-saffronstays.onrender.com)

**Note**: Response times may take longer than usual because the servers of the platform "Render" are located in Singapore.

## Setup and Installation

### Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/airbnb-api.git
cd airbnb-api
```
### Running the Application Locally
1. Install Dependencies:

Run the following command to install the required dependencies:

```bash
go mod tidy
```

2. Run the Go Application:
To start the server, run:

```bash
go run main.go
```

This will start the API server on `http://localhost:8080`.

## Docker Hub (Via Remote Image)
You can also pull the pre-built Docker image from Docker Hub:
- Build the Docker Image: [https://hub.docker.com/r/rohanyh/airbnb_api_saffronstays](https://hub.docker.com/r/rohanyh/airbnb_api_saffronstays)
- Follow The Commands,
1. Pull the remote docker image:
```bash
docker pull rohanyh/airbnb_api_saffronstays
```
2. Run the Image locally,
```bash
docker run rohanyh/airbnb_api_saffronstays
```

## Running with Docker (Building Image locally)
If you prefer running the application in a containerized environment, you can use Docker. Here's how to do it:

Build the Docker image using the following command:

```bash
docker build -t airbnb-api .
```
2. Run the Application with Docker Compose:

Use Docker Compose to run the application in a container:

```bash
docker-compose up --build
```
This will start the application on `http://localhost:8080`.

## API Endpoints
### Health Check
GET `/api/v1/health`

- Returns the health status of the server.
```bash
> curl --location --request GET 'localhost:8080/api/v1/health' --header 'Content-Type: application/json' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    40  100    40    0     0    174      0 --:--:-- --:--:-- --:--:--   175
{
  "status": "server is up and running..."
}
```

### Get Metrics
GET `/api/v1/getMetrics/{room_id}`

- Fetches the occupancy percentage and price metrics for a given room ID.

Path Parameter:

- `room_id` (string): The unique identifier of the Airbnb room.

```bash
> curl --location --request GET 'localhost:8080/api/v1/getMetrics/619966061834034729' --header 'Content-Type: application/json' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   243  100   243    0     0     73      0  0:00:03  0:00:03 --:--:--    73
{
  "request_id": "619966061834034729",
  "occupancy_percentage": {
    "2024-12": 61.111111111111114,
    "2025-01": 32.25806451612903,
    "2025-02": 75,
    "2025-03": 100,
    "2025-04": 100,
    "2025-05": 100
  },
  "average_rate": 137.06666666666666,
  "highest_rate": 150,
  "lowest_rate": 123
}
```

## Swagger Docs

The API documentation is available via Swagger UI at:

- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- **Swagger JSON**: `http://localhost:8080/swagger/doc.json`

You can use the Swagger UI to explore and test all the available API endpoints.

## Challenges
During the development of this API, some of the challenges I faced included:

1. Handling Multiple API Calls: Data Structure Parsing: Parsing the data correctly from the external APIs and ensuring that it was structured appropriately for processing was one of the challenges. There were multiple fields and types to manage, and making sure that each piece of data was handled correctly required careful attention to detail.
2. Date Operations: A significant challenge was performing operations involving dates, such as calculating occupancy percentages for each month over the next 5 months, and determining average rates, highest, and lowest rates for the next 30 days. Ensuring the correct handling of dates across different formats, considering edge cases like leap years or time zones, and accurately performing calculations was complex.

## Contributions
Contributions are welcome! Feel free to submit issues and pull requests.
