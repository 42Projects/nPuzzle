# Go parameters
GOCMD :=		go
GOBUILD :=		$(GOCMD) build
GOGET :=		$(GOCMD) get
GORUN :=		$(GOCMD) run
GOTEST :=		$(GOCMD) test
GOFMT :=		gofmt
GOLINT :=		golint

# Binary parameters
BINDIR :=		./bin/
CLINAME :=		nPuzzleCli

# Sources
CLIDIR :=		./cli/
SERVERDIR :=	./server/
SRCDIR :=		./src/

# Rules
all: build

$(BINDIR):
	@mkdir -p $@

build: valid test | $(BINDIR)
	@printf "Compiling cli..."
	@$(GOBUILD) -o $(BINDIR)$(CLINAME) $(CLIDIR)*
	@printf " done\n"

clean:
	@/bin/rm -rf $(BINDIR)

deploy: valid test | $(BINDIR)
	docker-compose -f deploy/docker-compose.yml up -d

down:
	docker-compose -f deploy/docker-compose.yml down

fmt:
	@$(GOFMT) -d $(CLIDIR) $(SERVERDIR) $(SRCDIR)
	@$(GOFMT) -w $(CLIDIR) $(SERVERDIR) $(SRCDIR)

lint:
	@$(GOLINT) $(CLIDIR) $(SERVERDIR) $(SRCDIR)

serve:
	@$(GORUN) $(SERVERDIR)/main.go

teardown:
	docker-compose -f deploy/docker-compose.yml down --rmi all

test:
	@$(GOTEST) -v $(SRCDIR)

valid: fmt lint

.PHONY: deploy
