NAME := nPuzzle

all: $(NAME)

$(NAME):
	@mkdir -p /bin/
	@go build -o ./bin/nPuzzle ./src/*.go

clean:
	@/bin/rm -rf ./bin/

test:
	@go test ./...

.PHONY: all clean test
