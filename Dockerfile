FROM gliderlabs/alpine

CMD ["redis-ranking-demo"]

COPY ./bin/redis-ranking-demo  /usr/local/bin/

EXPOSE 8080
