# go-trie
go-trie is a simple prefix trie implementation that supports any UTF-8 phrase insertions and lookups.

# Installation
go-trie can be installed using Go releases in module mode:
```
go get github.com/skykistler/go-trie@latest
```

# Usage
```go
import "github.com/skykistler/go-trie"

myTrie := trie.NewTrie()
```
Import the module and create a new Trie instance

```go
myTrie.Insert("My phrase.")
```
Insert words or phrases as a string.

```go
myTrie.Contains("my phrase") // false
myTrie.Contains("My phrase.") // true
```
`Contains` will return true only if there's an exact match.

# Contributing
This repository is mainly for demonstration but contributions are welcome!

# License
This library is distributed under the [Apache 2.0 license](./LICENSE). 