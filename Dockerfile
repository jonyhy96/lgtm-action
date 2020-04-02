FROM golang:1.13 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

# final stage
FROM scratch
WORKDIR /lgtm-action
COPY --from=builder /app /lgtm-action/
ENTRYPOINT ["./lgtm-action"]
