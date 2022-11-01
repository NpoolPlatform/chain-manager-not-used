FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY chain-manager /app

ENTRYPOINT ["/app/chain-manager"]

EXPOSE 50330 50331 50332 50333

CMD ["run"]
