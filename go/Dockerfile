FROM golang:1.21.1

WORKDIR /app

COPY src/go.mod ./
RUN go mod download

COPY ./src ./

RUN go build -o /go-docker-demo

EXPOSE 8080

CMD [ "/go-docker-demo" ]