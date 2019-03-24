FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -u github.com/gin-gonic/gin
RUN go build *.go
EXPOSE 8080
RUN export GIN_MODE=release
CMD ["/app/main"]