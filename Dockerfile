FROM golang:1.22.1-alpine AS builder

WORKDIR /

COPY . .

RUN GOOS=linux GOARCH=amd64 PATH="/tmp/go/bin:$PATH" go build ./cmd/gosanta/.

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder --chown=nonroot /gosanta /usr/local/bin/

USER nonroot

ENTRYPOINT [ "gosanta" ]
