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

build: test | $(BINDIR)
	@printf "Compiling cli..."
	@$(GOBUILD) -o $(BINDIR)$(CLINAME) $(CLIDIR)*
	@printf " done\n"
	@printf "Compiling server..."
	@$(GOBUILD) -o $(BINDIR)$(SERVERNAME) $(SERVERDIR)*
	@printf " done\n"

clean:
	@$(GOCLEAN)
	@/bin/rm -rf $(BINDIR)

cli: test | $(BINDIR)
	@printf "Compiling cli..."
	@$(GOBUILD) -o $(BINDIR)$(CLINAME) $(CLIDIR)*
	@printf " done\n"

server: test | $(BINDIR)
	@printf "Compiling server..."
	@$(GOBUILD) -o $(BINDIR)$(SERVERNAME) $(SERVERDIR)*
	@printf " done\n"

test:
	@$(GOTEST) -v $(SRCDIR)

.PHONY: all build clean cli server test
