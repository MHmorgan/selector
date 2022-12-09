EXE := selector
PLATFORMS := windows linux darwin
ARCHITECTURES := amd64 arm64

# -s : Dead code elimination
# -w : Disable DWARF generation
# netgo : Pure Go implementation of network-related packages
# static : Create a fully statically-linked executable
FLAGS = -tags netgo,static -ldflags "-s -w" -trimpath

all: clean $(PLATFORMS)

$(PLATFORMS):
	GOOS=$@ GOARCH=arm64 go build -o $(EXE)_$@_arm64 $(FLAGS)
	GOOS=$@ GOARCH=amd64 go build -o $(EXE)_$@_amd64 $(FLAGS)

clean:
	rm $(EXE)*
