# Go parameters
GOCMD :=	go
GOBUILD :=	$(GOCMD) build
GOCLEAN :=	$(GOCMD) clean
GOGET :=	$(GOCMD) get
GOTEST :=	$(GOCMD) test

# Binary parameters
BINDIR :=	./bin/
BINNAME :=	nPuzzleSolver

# Sources
SRCDIR	:=	./src/

# Rules
all: test build

build: | $(BINDIR)
	@$(GOBUILD) -o $(BINDIR)$(BINNAME) -v $(SRCDIR)*

$(BINDIR):
	@mkdir -p $@

benchmark:
	@$(GOTEST) -bench=. -v ./...

clean:
	@$(GOCLEAN)
	@/bin/rm -rf $(BINDIR)

test:
	@$(GOTEST) -v ./...

.PHONY: all clean test
