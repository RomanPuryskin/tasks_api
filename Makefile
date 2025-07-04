.PHONY:
.SILENT:

TARGET = .bin/cmd.exe

SOURCE = cmd/main.go

build:
	go build -o ${TARGET} ${SOURCE}

run: build
	${TARGET}

swag:
	swag init -g ./cmd/main.go --output ./docs --parseDependency --parseInternal