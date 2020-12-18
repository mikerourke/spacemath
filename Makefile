ifeq ($(OS),Windows_NT)
    SETENV_WIN := $(shell set GOOS=windows)
	SETENV_DARWIN := $(shell set GOOS=darwin)
else
    SETENV_WIN := $(shell env GOOS=windows)
	SETENV_DARWIN := $(shell set GOOS=darwin)
endif

out/spacemath.exe: cmd/spacemath.go
	$(SETENV_WIN) go build -v -o $@ $^

out/spacemath: cmd/spacemath.go
	$(SETENV_DARWIN) go build -v -o $@ $^
