build-windows:
	env GOOS=windows go build -o out/spacemath.exe cmd/spacemath.go

build-darwin:
	env GOOS=darwin go build -o out/spacemath cmd/spacemath.go