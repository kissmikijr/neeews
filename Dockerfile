FROM golang:1.14 as build
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download


FROM build as server-dev
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build server/main.go" --command=./main

FROM build as prod
COPY . .

FROM prod as server
WORKDIR /app/server
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o .

FROM prod as worker
WORKDIR /app/worker
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o .

FROM alpine
WORKDIR /app
COPY --chown=0:0 --from=server /app/server/server /app/server
COPY --chown=0:0 --from=worker /app/worker/worker /app/worker
RUN chmod +x /app/server
RUN chmod +x /app/worker

CMD  ["/app/server"]


