proto-gen: proto/*.proto
	protoc --go_out=. --go-grpc_out=. proto/*.proto

docker-build:
	docker build -t wtbackend .

docker-run:
	docker run -p 8080:8080 -p 8081:8081 wtbackend