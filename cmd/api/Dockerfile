FROM golang:latest

WORKDIR /app

# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

COPY . .

ENV PORT=8080
ENV ALLOWED_ORIGINS=http://localhost:3000

RUN go build -C cmd/api -mod=readonly -v -o bin

EXPOSE 8080

CMD ["/app/cmd/api/bin"]
