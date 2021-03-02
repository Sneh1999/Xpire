FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build .

EXPOSE 8000
CMD ["./Xpire"]
