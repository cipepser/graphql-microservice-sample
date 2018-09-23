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

→別のディレクトリで行ったら問題なく、実行できた。

## debug memo

`bash`がなかったので以下でログインして、`app`を実行したら`postgres`のdriverがないエラーだった。エラーログをどうやって吐かせるかは別途課題。

```
❯ docker-compose exec account sh
```

ちなみに`account`が`service`を書く欄で、`docker-compose.yaml`で定義したやつ。

## 実行

### アカウントを作る

`http://localhost:8000/graphql`で以下を実行

```graphql
mutation {
  createAccount(account: {name: "John"}) {
    id
    name
  }
}
```

レスポンス

```json
{
  "data": {
    "createAccount": {
      "id": "1AKrImOn6QwwL2zBtaA2PeFxIBr",
      "name": "John"
    }
  }
}
```

### 商品を作る

```graphql
mutation {
  a: createProduct(product: {name: "Kindle Oasis", description: "Kindle Oasis is the first waterproof Kindle with our largest 7-inch 300 ppi display, now with Audible when paired with Bluetooth.", price: 300}) { id },
  b: createProduct(product: {name: "Samsung Galaxy S9", description: "Discover Galaxy S9 and S9+ and the revolutionary camera that adapts like the human eye.", price: 720}) { id },
  c: createProduct(product: {name: "Sony PlayStation 4", description: "The PlayStation 4 is an eighth-generation home video game console developed by Sony Interactive Entertainment", price: 300}) { id },
  d: createProduct(product: {name: "ASUS ZenBook Pro UX550VE", description: "Designed to entice. Crafted to perform.", price: 300}) { id },
  e: createProduct(product: {name: "Mpow PC Headset 3.5mm", description: "Computer Headset with Microphone Noise Cancelling, Lightweight PC Headset Wired Headphones, Business Headset for Skype, Webinar, Phone, Call Center", price: 43}) { id }
}
```


```json
{
  "data": {
    "a": {
      "id": "1AKBtPreNArQjpOk6pBqtXgbdin"
    },
    "b": {
      "id": "1AKBtbTA3wfac4yFEx62np5G1sC"
    },
    "c": {
      "id": "1AKBtcdccsUcOoXttIvUCYAoM25"
    },
    "d": {
      "id": "1AKBtYSsODfyrFBIjNf5xBAMeLd"
    },
    "e": {
      "id": "1AKBtbn0gFPYIkxeeaHgxOhlI8J"
    }
  }
}
```

### 注文してみる

`accountId`や`id`は上記で得たレスポンスに合わせる。ここでは`John`が`a`、`b`、`e`を注文してみる。

```graphql
mutation {
  createOrder(order: { accountId: "1AKrImOn6QwwL2zBtaA2PeFxIBr", products: [
    { id: "1AKBtPreNArQjpOk6pBqtXgbdin", quantity: 2 },
    { id: "1AKBtbTA3wfac4yFEx62np5G1sC", quantity: 1 },
    { id: "1AKBtbn0gFPYIkxeeaHgxOhlI8J", quantity: 5 }
  ]}) {
    id
    createdAt
    totalPrice
  }
}
```

TODO: accountsのクエリが空のせいかうまくいかない。これが原因じゃないかも。

```json
{
  "data": {
    "createOrder": null
  },
  "errors": [
    {
      "message": "rpc error: code = Unknown desc = pq: relation \"orders\" does not exist",
      "path": [
        "createOrder"
      ]
    }
  ]
}
```

```graphql
query {
	accounts {
		id
    name
  }
}
```

```json
{
  "data": {
    "accounts": []
  }
}
```

## References
* [Using GraphQL with Microservices in Go](https://outcrawl.com/go-graphql-gateway-microservices/)
* [proposal: cmd/go: module semantics in $GOPATH/src #26377](https://github.com/golang/go/issues/26377)