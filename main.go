package main

import (
	"fmt"
	"log"
	"net/http"
)

//parte que lida com requisições

//lista dos servidores python

var serverList = []*Server{
	NewServer("http://127.0.0.1:4041/"),
	NewServer("http://127.0.0.1:4042/"),
	NewServer("http://127.0.0.1:4043/"),
	NewServer("http://127.0.0.1:4044/"),
	NewServer("http://127.0.0.1:4045/"),
}

func main() {
	fmt.Println("running...")

	//inicia a checagem de integridade em uma thread separada
	//recebe o intervalo em segundos que cada chegagem será feita

	go startHealthCheck(5)

	//recebe requisições no path '/' e manda elas para um servidor python
	//depois pega a resposta e devolve ela

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		//procura um servidor que esteja "saudavel"
		//função recebe a quantidade máxima de iterações
		//para procurar um servidor saudavel na serverList
		server, err := getHealthyServer(100)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadGateway)
		}

		fmt.Printf("Making request to %v\n", server.URL)

		server.ReverseProxy.ServeHTTP(res, req)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
