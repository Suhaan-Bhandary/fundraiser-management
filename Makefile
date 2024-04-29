run: ## Run project on host machine
	go run cmd/main.go

create-admin: ## create admin in db 
	go run cmd/main.go create-admin admin password

clean: ## Clean database file for a fresh start
	rm test.db

test: ## Run all unit tests in the project
	go test -v ./...

test-cover: ## Run all unit tests in the project with test coverage
	go test -v ./... -covermode=count -coverprofile=coverage.out

html-cover: test-cover
	go tool cover -html="coverage.out"

mock-app:
	find ./internal/app -mindepth 1 -maxdepth 1 -type d -execdir sh -c 'cd "{}" && mockery --name=Service' \;

mock-repository:
	cd internal/repository && mockery --all

mock: mock-app mock-repository
	@echo "Mocks Generated"
