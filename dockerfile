FROM public.ecr.aws/docker/library/golang:latest as builder

WORKDIR /go/src/app
COPY . .
RUN make build

FROM scratch
WORKDIR /

COPY --from=builder /go/src/app/telebot_* .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./telebot"]