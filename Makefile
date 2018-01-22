.PHONY: all

GIT_BRANCH := $(subst heads/,,$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null))

BIND_DIR := "dist"
SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

TARGETS := "indexer crawler spider tocrawl mergedocs server"

print-%: ; @echo $*=$($*)

build: build-indexer build-crawler build-spider build-tocrawl build-mergedocs build-server 

build-dist-%:
	go build -o $(PWD)/dist/$* $(PWD)/cmd/$*/*.go

build-%:
	go build -o $(PWD)/$* $(PWD)/cmd/$*/*.go

run-%:
	# @echo $*=$($*)
	# @echo "$@"
	cd $(PWD)/cmd/$* && go run -race *.go

info:
	@echo "Important: Create a basic \`conf.json\` file, limiting the crawler to a one minute run: \`{ \"crawler\": { \"due_per_run\": \"1m\" } }\`"
	@echo ""
	@echo "Available targets for build/run"
	@echo " 1. tocrawl: Run the package finder: \`go run tocrawl/*.go\`"
	@echo " 2. crawler: Run the crawler: \`go run crawler/*.go\`"
	@echo " 3. mergedocs: Merge the crawled docs: \`go run mergedocs/*.go\`"
	@echo " 4. indexer: Run the indexer: \`go run indexer/*.go\`"
	@echo " 5. server: Run the server: \`go run server/*.go\`"
	@echo ""
	@echo "Visit [http://localhost:8080](http://localhost:8080) in your browser"
	@echo ""

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
