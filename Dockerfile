FROM golang as server
WORKDIR /server
COPY server .
RUN go get ./...
RUN go build .

FROM golang as worker
WORKDIR /worker

RUN go build ./worker

FROM alpine
WORKDIR /app
COPY --from=server /server/server /server/server
COPY --from=worker /worker/worker /worker/worker

CMD  ["/app/server/server"]


