# RSS reader package

RSS reader package, which parses concurrently multiple RSS/Atom feeds.

- To install package: go get github.com/nestorov88/rss_reader
- The package has only exportable RssItem struct and func Parse, which takes slice of url strings and returns a slice of RssItem structs.
- In order to run test navigate to main dir and run go test -v ./...
