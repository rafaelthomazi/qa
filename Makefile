
dep:
	@go get -v -u -d ./...

vet: build_testdocker
	@docker run --rm -i qatest:latest go vet ./...

lint: build_testdocker
	@docker run --rm -i qatest:latest golint --min_confidence=0.9 -set_exit_status ./...

test: build_testdocker
	@docker run --rm -i qatest:latest go test -short ./...

build_docker:
	@docker build --pull --no-cache --rm --target runtime -t qaservice:latest .

build_testdocker:
	@docker build --target test -t qatest:latest .