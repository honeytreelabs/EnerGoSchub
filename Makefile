all: energoschub scan gui

energoschub: cmd/energoschub/energoschub.go
	go build $<

scan: cmd/scan/scan.go
	go build $<

gui: cmd/gui/gui.go
	go build $<

playground: cmd/playground/playground.go
	go build $<

.PHONY: clean
clean:
	-rm energoschub scan gui
