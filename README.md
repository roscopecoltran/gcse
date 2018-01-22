Go Search [![GoSearch](http://go-search.org/badge?id=github.com%2Fdaviddengcn%2Fgcse)](http://go-search.org/view?id=github.com%2Fdaviddengcn%2Fgcse)
=========

A keyword search engine helping people to find popular and relevant Go packages.

Online service: [Go Search](http://go-search.org/)

This is the root package with shared functions.

Sub packages are commands for running:

* [HTTP Server](http://github.com/roscopecoltran/gcse/server): Searching and web service
* [ToCrawl](http://github.com/roscopecoltran/gcse/tocrawl): Find packages to crawl.
* [Crawler](http://github.com/roscopecoltran/gcse/crawler): Crawling package files.
* [MergeDocs](http://github.com/roscopecoltran/gcse/mergedocs): Merge crawled package files with doc DB.
* [Indexer](http://github.com/roscopecoltran/gcse/indexer): Analyzing package information and generating indexed data for searching.

Development
-----------

You'll need to perform the following steps to get a basic server running:

  1. Create a basic `conf.json` file, limiting the crawler to a one minute run: `{ "crawler": { "due_per_run": "1m" } }`
  1. Run the package finder: `go run cmd/tocrawl/*.go`
  1. Run the crawler: `go run cmd/crawler/*.go`
  1. Merge the crawled docs: `go run cmd/mergedocs/*.go`
  1. Run the indexer: `go run cmd/indexer/*.go`
  1. Run the server: ` go run cmd/server/*.go`
  1. Visit [http://localhost:8080](http://localhost:8080) in your browser


LICENSE
-------
BSD license.
