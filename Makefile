build: build-indexer build-crawler build-spider build-tocrawl build-mergedocs build-server 

build-indexer:
	go build -o $(PWD)/bin/indexer $(PWD)/cmd/indexer/*.go

build-crawler:
	go build -o $(PWD)/bin/crawler $(PWD)/cmd/crawler/*.go

build-tocrawl:
	go build -o $(PWD)/bin/crawler $(PWD)/cmd/tocrawl/*.go

build-mergedocs:
	go build -o $(PWD)/bin/server $(PWD)/cmd/mergedocs/*.go

build-spider:
	go build -o $(PWD)/bin/server $(PWD)/cmd/spider/*.go

build-server:
	go build -o $(PWD)/bin/server $(PWD)/cmd/server/*.go
