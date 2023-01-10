BASENAME = $(shell basename ${PWD})
BUILD_DIR = bin

prepare:
	if [ ! -f parser/antlr4-4.11.1-complete.jar ]; then wget https://repo1.maven.org/maven2/org/antlr/antlr4/4.11.1/antlr4-4.11.1-complete.jar -O parser/antlr4-4.11.1-complete.jar -nc; fi
	go generate ./...

build : prepare
	go build -o ${BUILD_DIR}/${BASENAME} ${PWD}

clean:
	rm -rf ${BUILD_DIR}/

run:
	go run bot.go $(ARGS)
