package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	appServer "github.com/codemodify/systemkit-appserver"
	httpServer "github.com/codemodify/systemkit-appserver-http"
	mixedServer "github.com/codemodify/systemkit-appserver-mixed"
)

func main() {
	const port = 9999

	mixedServer.NewMixedServer([]appServer.IServer{
		httpServer.NewHTTPServer([]httpServer.HTTPHandler{
			httpServer.HTTPHandler{
				Route:   "/currency/all",
				Verb:    "GET",
				Handler: currencyAllRequestHandler,
			},
			httpServer.HTTPHandler{
				Route:   "/currency/{symbol}",
				Verb:    "GET",
				Handler: currencyRequestHandler,
			},
		}),
	}).Run(fmt.Sprintf(":%d", port), true)
}

func currencyRequestHandler(rw http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	url := fmt.Sprintf("https://api.hitbtc.com/api/3/public/symbol/%s", symbol)

	client := http.Client{}

	queryResponse, err := client.Get(url)
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}

	if queryResponse != nil || queryResponse.StatusCode != http.StatusOK {
		body, err := io.ReadAll(queryResponse.Body)
		if err != nil {
			rw.Write([]byte(err.Error()))
			return
		}

		rw.Write(body)
	}

	r.Body.Close()
}

func currencyAllRequestHandler(rw http.ResponseWriter, r *http.Request) {
	client := http.Client{}

	queryResponse, err := client.Get("https://api.hitbtc.com/api/3/public/symbol")
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}

	if queryResponse != nil || queryResponse.StatusCode != http.StatusOK {
		body, err := io.ReadAll(queryResponse.Body)
		if err != nil {
			rw.Write([]byte(err.Error()))
			return
		}

		rw.Write(body)
	}

	r.Body.Close()
}
