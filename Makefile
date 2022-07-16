run:
	@go run cmd/main.go

build:
	@go build -o main cmd/main.go

docker:
	@docker image build --tag forum .
	@docker container run -d -p 8080:8080 --name forum forum

prune:
	@docker stop forum || true
	@docker system prune