RusticSearch, a HTTP search server
============

RusticSearch is an open source general purpose HTTP search server that is minimalist and built with commitment to ruthless practicality. It makes it easy to build in the search functionality for your website that is simple, reliable and efficient.

# Install
The project is currently under development and should not be used for anything else but idle fiddling.
The easiest way to install the tool would be using [go get](http://golang.org/cmd/go/) command (you may need to install `go` tools first):
```bash
    go install git@github.com:dborzov/rusticsearch
```
This will download the source files to your `$GOPATH/src/github.com/dborzov/rusticsearch`, compile them and put the binary `rusticsearch` into `$GOPATH/bin`.  

# Configure
Run `rusticsearch -help` to see the configurable flags:
```bash
    $ rusticsearch -help
    Hi, I am Rustic Search Server!
    Usage of rusticsearch:
      -DBaddr="username@tcp(host)/tablename": Database connection address
      -configFile="rusticsearch.config": configuration json filepath
      -port=8080: a serving TCP port
      -refreshTime=10: time period (in minutes) when the search index is refreshed
```

It should be pretty straightforward for the most part. `DBaddr` here defines the connection credentials for a MySQl database in the [following format](https://github.com/go-sql-driver/mysql#dsn-data-source-name):
```bash
    [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
```

# Understand
Here is the search algorithm to be implemented. Checkbox marks if it is actually implemented just yet:

- [x] Lower the register of the query and attempt looking up the exact match. If we get enough matches, stop.
- [ ] Search for error-corrected query with distane one from the correct one (including symbol permutations) using [this](http://norvig.com/spell-correct.html) approach. If we get enough matches, stop.
- [x] Split query into words and search for each word separately. Then count total frequency of matches for each item and return the responses in that order.
