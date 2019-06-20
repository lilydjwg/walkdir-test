CRYSTAL_BIN ?= $(shell which crystal)

.PHONY: all clean

%: %.cr
	$(CRYSTAL_BIN) build --release --no-debug $<
	strip $@

all: walk

clean:
	-rm -f walk
