<<<<<<< HEAD
FROM golang:1.26.2

=======
# FROM golang:1.25.0

# WORKDIR /First_Project

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# EXPOSE 8081

# RUN go build -o server .

# CMD ["./server"]
# Stage 1 — build the binary
FROM golang:1.25.0 AS builder
>>>>>>> a1f77ceab5fe9317d46f796ad157c7ae4e4b40e1
WORKDIR /First_Project
COPY go.mod go.sum ./
RUN go mod download
COPY . .
<<<<<<< HEAD
EXPOSE 8081
# Fixed: binary name matches 'go build -o server', and CMD runs it directly
RUN go build -o server .
=======
RUN CGO_ENABLED=0 GOOS=linux go build -o server .
# Stage 2 — minimal final image (~20MB instead of ~1GB)
FROM alpine:3.19
WORKDIR /First_Project
COPY --from=builder /First_Project/server .
COPY --from=builder /First_Project/Setting ./Setting
EXPOSE 8081
>>>>>>> a1f77ceab5fe9317d46f796ad157c7ae4e4b40e1
CMD ["./server"]
