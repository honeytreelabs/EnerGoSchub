all: energoschub scan gui

energoschub: cmd/energoschub/energoschub.go
	go build $<

scan: cmd/scan/scan.go
	go build $<

gui: cmd/gui/gui.go cmd/gui/mac_storage.go
	go build $^

playground: cmd/playground/playground.go
	go build $<

.PHONY: run
run:
	go run cmd/gui/gui.go cmd/gui/mac_storage.go

.PHONY: clean
clean:
	-rm energoschub scan gui
