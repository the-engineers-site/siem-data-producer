FROM centos:8
USER app
WORKDIR /home/app

USER root
ARG JAR_FILE=siem-data-producer
COPY ${JAR_FILE} siem-data-producer

RUN chown app:app siem-data-producer

ADD static static

RUN chown -R app:app static

USER app

ENTRYPOINT ["/home/app/siem-data-producer"]