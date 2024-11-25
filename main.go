package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Address struct {
	OriginAPI  string
	Url        string
	JsonReturn string
}

func fetchAddress(channel chan<- Address, name string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyRes, err := io.ReadAll(resp.Body)

	var addr Address
	addr.OriginAPI = name
	addr.JsonReturn = string(bodyRes)
	addr.Url = url

	channel <- addr
}
func main() {

	channelViaCep := make(chan Address)
	channelBrasilApi := make(chan Address)

	go fetchAddress(channelBrasilApi, "BrasilAPI", "https://brasilapi.com.br/api/cep/v1/01001000")
	go fetchAddress(channelViaCep, "ViaCEP", "https://viacep.com.br/ws/01001000/json")

	select {
	case resultChannelViaCep := <-channelViaCep:
		fmt.Printf("The fastest is: %+v with URL: %s \n", resultChannelViaCep.OriginAPI, resultChannelViaCep.Url)
		fmt.Printf("Returned data: %+v\n", resultChannelViaCep.JsonReturn)
	case resultChannelBrasilApi := <-channelBrasilApi:
		fmt.Printf("The fastest is: %+v with URL: %s \n", resultChannelBrasilApi.OriginAPI, resultChannelBrasilApi.Url)
		fmt.Printf("Returned data: %+v\n", resultChannelBrasilApi.JsonReturn)
	case <-time.After(time.Second * 1):
		println("Timeout")
	}
}
