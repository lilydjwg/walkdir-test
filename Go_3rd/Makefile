.PHONY: all clean depends

default: depends walk

depends:
	go get github.com/karrick/godirwalk

%: %.go
	go build -ldflags "-w -s" $<

all: walk

clean:
	-rm -f walk
