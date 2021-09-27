FROM centos:8
RUN useradd -ms /bin/bash app

USER app
WORKDIR /home/app

USER root

ARG JAR_FILE=siem-data-producer
COPY ${JAR_FILE} siem-data-producer

ADD static static
ADD docs docs

RUN chown -R app:app /home/app

USER app

ENTRYPOINT ["/home/app/siem-data-producer"]