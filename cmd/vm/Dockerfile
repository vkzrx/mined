FROM golang:latest

WORKDIR /app

# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

COPY . .

RUN go build -C cmd/vm -mod=readonly -v -o bin

EXPOSE 5001

CMD ["/app/cmd/vm/bin"]
