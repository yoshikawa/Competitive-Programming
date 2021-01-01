OJ=oj
ACC=acc

all: test

.PHONY: test
test: ## go run test with oj
	$(OJ) t -d ./a/tests -c "go run a/main.go"
	$(OJ) t -d ./b/tests -c "go run b/main.go"