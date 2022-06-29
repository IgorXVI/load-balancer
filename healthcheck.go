package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

//parte que confere a saúde dos servidores

//inicia a checagem de saúde dos servidores,
//a checagem vai acontecer no intervalo de tempo em segundos passado como parametro,
//a cada x segundos a função faz um loop pela lista de servidores e confere se cada um deles
//está saudavel
func startHealthCheck(interval int) {
	scheduler := gocron.NewScheduler(time.Local)

	scheduler.Every(interval).Seconds().Do(func() {
		for _, server := range serverList {
			health := server.CheckHealth()

			if health {
				fmt.Printf("HEALTH CHECK: %v is healthy!\n", server.URL)
			} else {
				fmt.Printf("HEALTH CHECK: %v is NOT healthy!\n", server.URL)
			}
		}
	})

	scheduler.StartAsync()
}

//pega o próximo servidor no round robin, 
//se ele não estiver tenta pegar o próximo,
//continua até atingir o limite de iterações determinado pelo parametro
func getHealthyServer(maxIterations int) (*Server, error) {
	count := 0

	for {
		if count > maxIterations {
			return nil, errors.New("reached limit for healthy hosts search, no healthy hosts found")
		}

		server := getRoundRobinServer()

		if server.Health {
			return server, nil
		}

		count++
	}
}
