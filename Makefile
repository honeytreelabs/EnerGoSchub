all: energotritt scan

energotritt: cmd/energotritt/energotritt.go
	go build $<

scan: cmd/scan/scan.go
	go build $<

playground: cmd/playground/playground.go
	go build $<

.PHONY: clean
clean:
	-rm energotritt scan
