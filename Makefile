# Go parameters
GOCMD :=		go
GOBUILD :=		$(GOCMD) build
GOCLEAN :=		$(GOCMD) clean
GOGET :=		$(GOCMD) get
GOTEST :=		$(GOCMD) test

# Binary parameters
BINDIR :=		./bin/
CLINAME :=		nPuzzleCli
SERVERNAME :=	nPuzzleServer

# Sources
CLIDIR :=		./cli/
SERVERDIR :=	./server/
SRCDIR :=		./src/

# Rules
all: test build

$(BINDIR):
	@mkdir -p $@

build: cli server

clean:
	@$(GOCLEAN)
	@/bin/rm -rf $(BINDIR)

cli: | $(BINDIR)
	@printf "Compiling cli..."
	@$(GOBUILD) -o $(BINDIR)$(CLINAME) $(CLIDIR)*
	@printf " done\n"

server: | $(BINDIR)
	@printf "Compiling server..."
	@$(GOBUILD) -o $(BINDIR)$(SERVERNAME) $(SERVERDIR)*
	@printf " done\n"

test:
	@$(GOTEST) -v $(SRCDIR)

.PHONY: all build clean cli server test
