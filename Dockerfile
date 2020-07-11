FROM golang:1.14 as build
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

FROM build as server
WORKDIR /app/server
RUN go build -o .

FROM build as worker
WORKDIR /app/worker
RUN go build -o .

FROM scratch
WORKDIR /app
COPY --from=server /app/server /app/server
COPY --from=worker /app/worker /app/worker

CMD  ["/app/server"]


