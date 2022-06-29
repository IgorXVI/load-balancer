package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

//classe que contém os dados de cada servidor,
//com uma flag para indicar se ele está saudavel ou não

type Server struct {
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}

func NewServer(serverUrl string) *Server {
	serverUrlObj, _ := url.Parse(serverUrl)

	reverseProxy := httputil.NewSingleHostReverseProxy(serverUrlObj)

	return &Server{
		URL:          serverUrl,
		ReverseProxy: reverseProxy,
		Health:       true,
	}
}

//confere a integridade do servidor fazendo uma
//requsição para ele, se não retornar erro
//o servidor está saudável
func (server *Server) CheckHealth() bool {
	resp, err := http.Head(server.URL)
	if err != nil {
		fmt.Println(err)
		server.Health = false
		return false
	}

	if resp.StatusCode != http.StatusOK {
		server.Health = false
		return false
	}

	server.Health = true
	return true
}
