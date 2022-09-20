FROM golang:1.17-alpine as builder
LABEL maintainer="Dion Aditya<dionjatingaleh@gmail.com>"

WORKDIR /app
COPY . .


RUN go build -o day7 . && \
    chmod 777 day7

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/day7 .
COPY --from=builder /app/.env .
EXPOSE 1323
CMD [ "./day7" ]