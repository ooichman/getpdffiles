FROM ubi8/go-toolset AS builder
RUN mkdir -p /opt/app-root/src/getFTPfile
WORKDIR /opt/app-root/src/getFTPfile
ENV GOPATH=/opt/app-root/
ENV PATH="${PATH}:/opt/app-root/src/go/bin/"
COPY  src/getpdffiles/ .
RUN set -x && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep init && \
    dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o getpdffiles

FROM scratch
USER 1001

COPY --from=builder  /opt/app-root/src/getFTPfile/getpdffiles /usr/bin/
COPY run.sh /usr/bin/
WORKDIR /data
ENTRYPOINT ["/usr/bin/run.sh"]