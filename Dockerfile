FROM debian

WORKDIR /var/www

RUN apt-get update

EXPOSE 1323

ARG SERVER_FILE

COPY ${SERVER_FILE} /var/www/

CMD ${SERVER_FILE}
