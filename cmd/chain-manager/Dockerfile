FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY chain-manager /app

ENTRYPOINT ["/app/chain-manager"]

EXPOSE 50560 50561 50562 50563

CMD ["run"]
