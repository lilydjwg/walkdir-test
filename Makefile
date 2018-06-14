.PHONY: all clean

languages = Go Go_3rd Rust

all: $(languages)

.PHONY: $(languages)
$(languages): %:
	$(MAKE) -C $*

clean:
	for d in $(languages); do $(MAKE) clean -C "$$d"; done
