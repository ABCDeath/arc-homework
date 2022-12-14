PROJECT_ROOT=$(shell pwd)
SRC_PATH=/src
MOCK_CODEGEN=docker run -v $(shell pwd):$(SRC_PATH) -w $(SRC_PATH)  vektra/mockery

codegen:
	PROJECT_ROOT=$(PROJECT_ROOT) go generate ./...

test:
	go test -v ./space-game/...

build: codegen
	go build -o build/binary -v ./space-game/...

mock:
	# Movable
	$(MOCK_CODEGEN) \
		--name=Movable \
		--case=underscore \
		--dir=space-game/moving/movable \
		--output=space-game/moving/movable/mocks

	# Rotatable
	$(MOCK_CODEGEN) \
		--name=Rotatable \
		--case=underscore \
		--dir=space-game/moving/rotatable \
		--output=space-game/moving/rotatable/mocks

	# Object
	$(MOCK_CODEGEN) \
		--name=Object \
		--case=underscore \
		--dir=space-game/moving/object \
		--output=space-game/moving/object/mocks

	# Command
	$(MOCK_CODEGEN) \
		--name=Command \
		--case=underscore \
		--dir=space-game/command \
		--output=space-game/command/mocks

	# Error Handler
	$(MOCK_CODEGEN) \
		--name=Handler \
		--case=underscore \
		--dir=space-game/error_handler \
		--output=space-game/error_handler/mocks

	# Fuel
	$(MOCK_CODEGEN) \
		--name=Fuel \
		--case=underscore \
		--dir=space-game/moving/engine \
		--output=space-game/moving/engine/mocks

	# Direction
	$(MOCK_CODEGEN) \
		--name=Direction \
		--case=underscore \
		--dir=space-game/moving/direction \
		--output=space-game/moving/direction/mocks

	# Queue
	$(MOCK_CODEGEN) \
		--name=Queue \
		--case=underscore \
		--dir=space-game/queue \
		--output=space-game/queue/mocks
