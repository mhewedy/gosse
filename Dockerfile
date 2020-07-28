FROM alpine:latest
WORKDIR /root/
COPY index.html .
COPY gosse .
CMD ["./gosse"]
