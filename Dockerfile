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
WORKDIR /First_Project
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .
# Stage 2 — minimal final image (~20MB instead of ~1GB)
FROM alpine:3.19
WORKDIR /First_Project
COPY --from=builder /First_Project/server .
COPY --from=builder /First_Project/Setting ./Setting
EXPOSE 8081
CMD ["./server"]
