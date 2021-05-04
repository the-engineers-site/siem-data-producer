FROM golang:1.16.0
ARG JAR_FILE=siem-data-producer
COPY ${JAR_FILE} siem-data-producer
ENTRYPOINT ["./siem-data-producer"]
