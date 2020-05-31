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

create-k8resources:
	- chmod 777 scripts/create-k8resources.sh
	- scripts/create-k8resources.sh $(hostPort) $(containerPort)

run-kubernetes:
	- chmod 777 scripts/run-kubernetes.sh
	- scripts/run-kubernetes.sh $(hostPort) $(containerPort)
