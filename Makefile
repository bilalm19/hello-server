GO_BUILD = go build -o
GO_TEST = go test
TEST_DIR = ./internal/server/...
MAIN = cmd/hello-server/main.go
EXECUTABLE = hello
DOCKER_REPOSITORY = bilalmahmood19/hello-server
IMAGE_TAG = 1.0
CGO = 0

.PHONY: build
build:
	CGO_ENABLED=$(CGO) $(GO_BUILD) $(EXECUTABLE) $(MAIN)

.PHONY: test
test:
	$(GO_TEST) $(TEST_DIR) -v

.PHONY: docker-build-push
docker-build-push: docker-build
	docker push $(DOCKER_REPOSITORY):$(IMAGE_TAG)

.PHONY: docker-build
docker-build:
	DOCKER_BUILDKIT=1 docker build --target prod -t $(DOCKER_REPOSITORY):$(IMAGE_TAG) .
