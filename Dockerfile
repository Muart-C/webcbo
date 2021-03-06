FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/Muart-C/project

ENV USER root254
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET WOMdx3gst6sSiqDy

# Replace this with actual PostgreSQL DSN.
ENV DSN $GO_BOOTSTRAP_MYSQL_DSN

WORKDIR /go/src/github.com/Muart-C/project

RUN godep go build

EXPOSE 8888
CMD ./project
