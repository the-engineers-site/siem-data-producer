FROM golang:1.16 as builder
WORKDIR /home
ADD . .
RUN go build -o siem-data-producer main.go
RUN go build -o producerctl producectl/main.go

FROM centos:8
RUN useradd -ms /bin/bash app

USER app
WORKDIR /home/app

USER root

COPY --from=builder /home/siem-data-producer siem-data-producer

COPY --from=builder /home/static static
COPY --from=builder /home/docs docs
COPY --from=builder /home/producerctl /usr/local/bin/producerctl
COPY --from=builder /home/producerctl /home/producerctl
RUN chmod 777 /usr/local/bin/producerctl
RUN chmod 777 /home/producerctl
RUN chown -R app:app /home

ENTRYPOINT ["/home/app/siem-data-producer"]