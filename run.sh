go mod init receipt-processor
go get github.com/google/uuid  # For generating unique IDs
go get github.com/gorilla/mux  # For routing
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
