FROM debian

WORKDIR /var/www

RUN apt-get update

EXPOSE 1323

ARG SERVER_FILE

COPY ${SERVER_FILE} /var/www/

RUN chmod +x ${SERVER_FILE}

CMD /var/www/${SERVER_FILE}
