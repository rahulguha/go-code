FROM golang:1.12.6-alpine3.9
RUN apk add git
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get -d github.com/gorilla/mux
RUN go build -o status .
CMD ["/app/status"]