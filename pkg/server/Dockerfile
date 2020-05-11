FROM golang

ADD . /go/src/github.com/knrt10/gRPC-cache/

WORKDIR /go/src/github.com/knrt10/gRPC-cache/

RUN chmod 777 scripts/protoc-gen.sh && chmod 777 scripts/build.sh && scripts/build.sh

ENTRYPOINT [ "./server-cache" ]

EXPOSE 5001
