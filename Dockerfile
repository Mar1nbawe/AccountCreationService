FROM golang:1.23.2
WORKDIR /app
COPY go.mod go.sum .env ./

ENV TZ=Europe/Romania
RUN go mod download
COPY *.go ./
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /AccountCreationService

EXPOSE 8080
CMD ["/AccountCreationService"]

#ENTRYPOINT ["top", "-b"]