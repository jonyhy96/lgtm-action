ARG GO_VERSION=1.13.7
ARG DND_VERSION=19.03

FROM golang:${GO_VERSION} as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

FROM docker:${DND_VERSION}

COPY --from=builder /app/lgtm-action /lgtm-action

ENTRYPOINT ["/lgtm-action"]
