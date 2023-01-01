#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

#copy all files
COPY . .

#build the binary
RUN mkdir static/logs
RUN go clean
RUN go build
RUN zip -r -X cheesy-arena.zip LICENSE README.md access_point_config.tar.gz cheesy-arena-lite.exe db font schedules static switch_config.txt templates