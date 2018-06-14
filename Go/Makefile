.PHONY: all clean

%: %.go
	go build -ldflags "-w -s" $<

all: walk

clean:
	-rm -f walk
