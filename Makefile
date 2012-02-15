all: echo-pkg echo-cmd

echo-pkg:
	go build ./pkg/echo

echo-cmd:
	go build ./cmd/echo
