proto-build: proto/tree.proto

docker-build: Dockerfile
	@docker build -t calc .

docker-run:
	@docker run calc -p 8080:8080
	
docker-clean:
	@docker image rm calc
	@docker image rm golang:1.19.6-bullseye