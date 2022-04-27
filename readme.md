# sample usage of https://api.hitbtc.com

- run: `go run main.go`
- endpoint1: `curl localhost:9999/currency/{symbol} | jq`
	- `curl localhost:9999/currency/ETHBTC | jq`
	- `curl localhost:9999/currency/ETHUST | jq`
- endpoint2: `curl localhost:9999/currency/all | jq | more`

# third party
- using `github.com/codemodify/systemkit*` framework (migrating to goframework.io)

# see screenshots
![](screen1.png?raw=true "")
![](screen2.png?raw=true "")
