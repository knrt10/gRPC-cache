all:
	- make build && make protoc

build:
	- chmod 777 scripts/protoc-gen.sh && chmod 777 scripts/build.sh
	- scripts/build.sh

docker:
	- chmod 777 scripts/build-docker.sh
	- scripts/build-docker.sh

protoc:
	- scripts/protoc-gen.sh

server:
	- go run pkg/server/main.go

dockerServer:
	- docker run -it -p 5001:5001 knrt10/grpc-cache

client:
	- go run pkg/client/client.go

test:
	- go test pkg/server/v1/* -v -cover
