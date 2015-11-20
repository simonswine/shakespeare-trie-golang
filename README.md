# Shakespeare's Tries

Generates an autocomplete Trie from Shakespeare's full corpus.

Written after November 2015's [West London Hack Night](http://www.meetup.com/West-London-Hack-Night/).

## Building & Running

You'll need [golang](https://golang.org/doc/install). Then call

``` sh
# get/install source code
go get github.com/simonswine/shakespeare-trie-golang

# run tests
go test github.com/simonswine/shakespeare-trie-golang

# run program
shakespeare-trie-golang
```

You can enter queries and it will show the lines that begin with that string.
