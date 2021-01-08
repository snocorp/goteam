FROM golang:1.13
ARG VERSION
WORKDIR /app
ADD ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X \"gitlab.com/isotask/trek/version.VERSION=${VERSION}\"" -a -installsuffix cgo -o dbtrek .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/dbtrek .
CMD ["./dbtrek"]
