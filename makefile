PROJECT_NAME = app
BIN_DIR = ./bin/
MAIN_SRC_FILE = app.go

COMP = go

TEST_FLAGS = -v


all:
	@ ${COMP} run ./*.go


build:
	@ ${COMP} build -o ${BIN_DIR}${PROJECT_NAME}.exe ./*.go

tests:
	@ ${COMP} test ${TEST_FLAGS}
