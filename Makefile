run-sets: main.go
	@go run . sim sets -f ./pulpcalc.yml

proto-build: proto/tree.proto

docker-build: Dockerfile
	@echo "Building 'pulpcalc' docker image"
	@docker build -t pulpcalc .

docker-run:
	@echo "Running 'pulpcalc' docker image"
	@docker run pulpcalc -p 8080:8080
	
docker-clean:
	@docker image rm pulpcalc
	@docker image rm golang:1.19.6-bullseye