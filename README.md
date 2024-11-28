# Receipt Processor
Project Structure

receipt-processor/
│
├── main.go                 # Main application entry point      
├── handlers/               # HTTP route handlers               
│   ├── process_receipt.go  # Handler for processing receipts  
│   └── get_points.go       # Handler for retrieving points     
│
├── models/                 # Data models                       
│   ├── receipt.go          # Receipt struct definitions
│   └── item.go             # Item struct definitions
│
├── service/                # Point Calc logic                  
│   └── point_calculator.go # Logic for calculating receipt points
│
├── storage/                # In-memory storage                 
│   └── receipt_store.go    # In-memory receipt storage
│
├── Dockerfile              # Docker configuration
└── go.mod                  # Go module dependencies



# Dependency Management:

go mod init receipt-processor
go get github.com/google/uuid  # For generating unique IDs
go get github.com/gorilla/mux  # For routing 

# Running the Application
Running the Application:

# Build Docker image
docker build -t receipt-processor .

# Run the Docker container
docker run -p 8080:8080 receipt-processor


# Implementation 

1. Point Calculation                          - complete
2. Models                                     - complete
3. Storage Mechanism                          - complete
4. Main App Structure                         - complete?
5. Handlers                                   - complete
6. Routing                                    - complete?
    - POST endpoint at /receipts/process
    - GET endpoint at /receipts/{id}/points
7. Error Handling                             - complete?

additional implementation to-do:

Add input validation (matching API spec)
Write unit tests
Add logging
Implement graceful shutdown
Add middleware for logging and recovery
integration tests for routing endpoints

