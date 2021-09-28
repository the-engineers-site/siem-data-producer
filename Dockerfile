FROM golang:1.16 as builder
WORKDIR /home
ADD . .
RUN go build -o siem-data-producer main.go

FROM centos:8
RUN useradd -ms /bin/bash app

USER app
WORKDIR /home/app

USER root

COPY --from=builder /home/siem-data-producer siem-data-producer

ADD --from=builder /home/static static
ADD --from=builder /home/docs docs

RUN chown -R app:app /home/app

ENTRYPOINT ["/home/app/siem-data-producer"]