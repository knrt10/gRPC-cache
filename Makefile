build:
	- chmod 777 scripts/protoc-gen.sh && chmod 777 scripts/build.sh
	- scripts/build.sh

docker:
	- chmod 777 scripts/build-docker.sh
	- scripts/build-docker.sh

protoc:
	- scripts/protoc-gen.sh

server:
	- go run api/main.go

dockerServer:
	- docker run -it -p 5001:5001 knrt10/grpc-cache

client:
	- go run examples/client.go

test:
	- go test api/server/* -v -cover -race

run-k8s-server:
	- chmod 777 scripts/run-k8-server.sh
	- scripts/run-k8-server.sh

run-helm-server:
	- chmod 777 scripts/run-helm.sh
	- scripts/run-helm.sh

run-terraform-server:
	- chmod 777 scripts/run-terraform.sh
	- scripts/run-terraform.sh
