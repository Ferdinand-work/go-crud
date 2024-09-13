FROM golang:1.22 AS builder
WORKDIR /app
COPY . /app
RUN go build /app
EXPOSE 9090
ENTRYPOINT [ "./go-crud" ]