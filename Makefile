.PHONY: clean install test

install:
	go install github.com/Fire-Dragon-DoL/lab/cli/lab

test: install
	go test -v -json | lab

clean:
	go clean -i github.com/Fire-Dragon-DoL/lab...
