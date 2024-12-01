# Receipt Processor
Project Structure
```markdown
receipt-processor/
│
├── main.go                 # Main application entry point
├── handlers/               # HTTP route handlers               
│   ├── process_receipt.go  # Handler for processing receipts
│   └── get_points.go       # Handler for retrieving points
│
├── models/                 # struct definitions 
│   ├── receipt.go          
│
├── service/                
│   └── point_calculator.go # Logic for calculating receipt points
│
├── storage/                # In-memory storage
│   └── receipt_store.go    # In-memory receipt storage
|
├── Dockerfile              # Docker configuration
└── go.mod                  # Go module dependencies
```

Run the application with: 
```markdown 
$ ./run.sh
or
$ go mod init receipt-processor
$ go get github.com/google/uuid  # For generating unique IDs
$ go get github.com/gorilla/mux  # For routing 

$ docker build -t receipt-processor .
$ docker run -p 8080:8080 receipt-processor
``` 
Example: 
```markdown 
$ curl -X POST http://localhost:8080/process-receipt -d '{
  "Retailer": "My Retailer",
  "PurchaseDate": "2024-11-28T10:30:00Z", 
  "PurchaseTime": "2024-11-28T10:30:00Z",
  "Items": [
    {
      "shortDescription": "Item 1",
      "price": "10.00"
    },
    {
      "shortDescription": "Item 2",
      "price": "20.50"
    }
  ],
  "Total": "30.50",
  "RetailerId": "123"
}' -H "Content-Type: application/json"

{"id":"22205988-20ad-42ce-9c43-804766bd019e"}

$ curl http://localhost:8080/get-points/22205988-20ad-42ce-9c43-804766bd019e

{"points":47}

