GO_CMD = go
VARSION = -u

build:
	$(GO_CMD) build -o ./target/main ./main.go
clean:
	rm -rf ./target