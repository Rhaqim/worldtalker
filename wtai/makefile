PROTOC = python -m grpc_tools.protoc
PROTO_DIR = protos
PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)
PYTHON_OUT = ./wtai/proto

all: $(PROTO_FILES)
	@mkdir -p $(PYTHON_OUT)
	$(PROTOC) -I$(PROTO_DIR) --python_out=$(PYTHON_OUT) --grpc_python_out=$(PYTHON_OUT) $^

run:
	# python -m wtai.server
	poetry run python -m wtai.server

clean:
	rm -rf $(PYTHON_OUT)/*.py

docker-build:
	docker build -t wtai .

docker-run:
	docker run -p 50051:50051 wtai

.PHONY: all run clean