.PHONY: run
run:
	go run ./cmd/demo/demo.go

.PHONY: mock
mock:
	mockgen -source=./internal/users/user_services/user_service.go \
		-destination=mocks/users/user_services/user_service_mock.go \
		-package=userservices_mock
	mockgen -source=./internal/users/user_repositories/user_repository.go \
		-destination=mocks/users/user_repositories/user_repository_mock.go \
		-package=userrepositories_mock

.PHONY: test-create
test-create:
	curl -X POST http://localhost:3210/v1/user -d '{"name":"Foobar"}'

.PHONY: test-all
test-all:
	curl -X GET http://localhost:3210/v1/users -H "Authorization: Bearer $(t)"

.PHONY: test-user
test-user:
	curl -X GET http://localhost:3210/v1/user -H "Authorization: Bearer $(t)" 

.PHONY: test-update
test-update:
	curl -X PATCH http://localhost:3210/v1/user -H "Authorization: Bearer $(t)"  -d '{"name":"$(n)"}'

.PHONY: test-login
test-login:
	curl -X POST http://localhost:3210/v1/login -d '{"id":1}'

