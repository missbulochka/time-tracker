FROM base-time-tracker:1.0 AS builder
WORKDIR /usr/src/timetracker
COPY . .
RUN go build -v -o /usr/src/timetracker/bin/res ./cmd/main.go

FROM ubuntu:22.04 AS runner
WORKDIR /usr/src/timetracker/bin
COPY --from=builder /usr/src/timetracker/bin/res .
ENTRYPOINT [ "/res" ]
