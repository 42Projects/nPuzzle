# Go parameters
GOCMD :=		go
GOBUILD :=		$(GOCMD) build
GOCLEAN :=		$(GOCMD) clean
GOFMT :=		$(GOCMD)fmt
GOGET :=		$(GOCMD) get
GOLINT :=		golint
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
all: valid build

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

fmt:
	@$(GOFMT) -d $(CLIDIR) $(SERVERDIR) $(SRCDIR)
	@$(GOFMT) -w $(CLIDIR) $(SERVERDIR) $(SRCDIR)

lint:
	@$(GOLINT) $(CLIDIR) $(SERVERDIR) $(SRCDIR)

server: test | $(BINDIR)
	@printf "Compiling server..."
	@$(GOBUILD) -o $(BINDIR)$(SERVERNAME) $(SERVERDIR)*
	@printf " done\n"

test:
	@$(GOTEST) -v $(SRCDIR)

valid: fmt lint

.PHONY: all build clean cli fmt lint server test valid
