FROM base-time-tracker:1.0 AS builder
WORKDIR /go
COPY . .
RUN go build -o /go/bin/res ./cmd/main.go

FROM ubuntu:22.04 AS runner
WORKDIR /go/bin
COPY --from=builder /go/bin/res .
ENTRYPOINT [ "./res" ]
