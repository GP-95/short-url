FROM golang:1.22-bookworm

WORKDIR /app

ENV PORT=3000

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN touch /app/database.db

RUN go build -v -o /usr/local/bin/app ./cmd/app/main.go

EXPOSE 3000

CMD ["app"]
