.PHONY: run
run:
	go run ./cmd/demo/demo.go

.PHONY: mock
mock:
	mockgen -source=./internal/users/user_services/user_service.go -destination=mocks/users/user_services/user_service_mock.go -package=userservices_mock
	mockgen -source=./internal/users/user_repositories/user_repository.go -destination=mocks/users/user_repositories/user_repository_mock.go -package=userrepositories_mock

.PHONY: test-create
test-create:
	curl -X POST http://localhost:3210/v1/user -d '{"name":"Foobar"}'
