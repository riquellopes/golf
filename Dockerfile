FROM golang:alpine

# Install dependencies projects.

RUN apk add --update -t build-deps curl go git libc-dev gcc libgcc
