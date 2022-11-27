BASENAME = $(shell basename ${PWD})
BUILD_DIR = bin
build:
	go build -o ${BUILD_DIR}/${BASENAME} ${PWD}

clean:
	rm -rf ${BUILD_DIR}/

run:
	go run bot.go $(ARGS)
