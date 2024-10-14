FROM alpine:latest

RUN mkdir /app

COPY biodataApp /app
COPY ./config/ /app/config/
# only for used the local
COPY ./config/ /config/   
COPY ./migration/ /app/migration/
# only for used the local
COPY ./migration/ /migration/
COPY ./templates/ /app/templates/
# only for used the local
COPY ./templates/ /templates/

COPY ./run.sh /app/run.sh

RUN chmod +x /app/run.sh

# Debugging: Run a shell to test
ENTRYPOINT ["/bin/sh"]
CMD ["/app/run.sh"]
