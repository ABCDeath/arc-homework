DOCKER_PATH=/src
MOCK_CODEGEN=docker run -v $(shell pwd):$(DOCKER_PATH) -w $(DOCKER_PATH)  vektra/mockery

test:
	go test -v ./...

build:
	go build -o build/binary -v ./...

mock:
	# Movable
	$(MOCK_CODEGEN) \
		--name=Movable \
		--case=underscore \
		--dir=space-game/moving/movable \
		--output=space-game/moving/movable/mocks

	# movable Object
	$(MOCK_CODEGEN) \
		--name=Object \
		--case=underscore \
		--dir=space-game/moving/object \
		--output=space-game/moving/object/mocks
