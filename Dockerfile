FROM debian

WORKDIR /var/www

RUN apt-get update

EXPOSE 1323

ARG SERVER_FILE

COPY server /var/www/

RUN chmod +x server

CMD /var/www/server
