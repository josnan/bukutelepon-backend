FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get -u github.com/lib/pq

RUN go build src/*.go

EXPOSE 8080
RUN export GIN_MODE=release

CMD ["/app/main"]
