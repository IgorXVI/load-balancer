package main

//parte que implementa round robin

//index do último servidor que foi chamado
var lastServedIndex = -1

//retorna o próximo servidor incrementado o último index
//e tirando o modulo do tamanho do array para voltar a 0 quando chegar no último index
func getRoundRobinServer() *Server {
	lastServedIndex++

	nextIndex := lastServedIndex % len(serverList)

	server := serverList[nextIndex]

	return server
}