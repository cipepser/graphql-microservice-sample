# graphql-microservice-sample

[Using GraphQL with Microservices in Go](https://outcrawl.com/go-graphql-gateway-microservices/)の写経

基本的には[github](https://github.com/tinrab/spidey)に上がっているコードを写経していく。

`go1.11`を使っているので`vgo`がモジュール化されていて動かない。

```sh
❯ vgo vendor
go vendor is now go mod -vendor
```

`go mod -vendor`も現時点(2018/09/17)では、`go mod vendor`らしいので、これで実行する。
ちなみに`$GOPATH/src`で実行しようとすると怒られるので注意。なぜ怒られるかは[proposal: cmd/go: module semantics in $GOPATH/src #26377](https://github.com/golang/go/issues/26377)に既存に影響を与えないようにというのが議論されている。


## References
* [Using GraphQL with Microservices in Go](https://outcrawl.com/go-graphql-gateway-microservices/)
* [proposal: cmd/go: module semantics in $GOPATH/src #26377](https://github.com/golang/go/issues/26377)