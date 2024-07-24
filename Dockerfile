FROM golang:1.22.4 AS base
RUN mkdir /app
WORKDIR /app
COPY . /app/
RUN go mod download
RUN make build
FROM alpine AS release
WORKDIR /app
COPY --from=base /app/bin/main /app/
RUN mkdir templates
COPY --from=base /app/templates /app/templates
ENTRYPOINT [ "./main" ]