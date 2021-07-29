FROM golang:alpine

ENV GO111MODULE=on
ENV PORT=8080
ARG API_KEY=123
ENV API_KEY ${API_KEY}

WORKDIR /app
ADD . /app
RUN cd /app && go mod download && go build -o goapp .

EXPOSE 8080

ENTRYPOINT ./goapp
