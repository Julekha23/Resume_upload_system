FROM golang:1.26.2

WORKDIR /First_Project
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8081
# Fixed: binary name matches 'go build -o server', and CMD runs it directly
RUN go build -o server .
CMD ["./server"]
