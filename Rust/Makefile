.PHONY: all clean

target = $(shell \
  if type jshon >/dev/null 2>&1; then \
    cargo metadata --format-version=1 | jshon -e target_directory -u; \
  elif type jq >/dev/null 2>&1; then \
    cargo metadata --format-version=1 | jq -r .target_directory; \
  else \
    echo target; \
  fi; \
)

walk: src/main.rs Cargo.toml
	cargo build --release
	strip $(target)/release/walk
	ln -srf $(target)/release/walk

all: walk

clean:
	-rm -rf $(target)/release/swapview
