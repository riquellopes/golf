FROM golang:1.7

# Install dependencies projects.

RUN go get -u github.com/labstack/echo
