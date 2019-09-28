FROM golang:latest AS builder

RUN mkdir /nPuzzle
ADD . /nPuzzle
WORKDIR /nPuzzle
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/nPuzzleCli server/main.go

FROM scratch

COPY --from=builder /bin/nPuzzleCli /bin/nPuzzleCli
ENTRYPOINT ["/bin/nPuzzleCli"]
