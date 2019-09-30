FROM golang as dev

ENV GO111MODULE=on

RUN mkdir -p $GOPATH/src/github.com/Alluxio/alluxio-csi
COPY . $GOPATH/src/github.com/Alluxio/alluxio-csi

FROM dev as build

RUN cd $GOPATH/src/github.com/Alluxio/alluxio-csi && \
    CGO_ENABLED=0 go build -o /usr/local/bin/alluxio-csi

FROM alluxio/alluxio-fuse:2.0.1 as final
COPY --from=build /usr/local/bin/alluxio-csi /usr/local/bin/