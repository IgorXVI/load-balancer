executar os comandos em ordem:

python server.py "server 1" 4041

python server.py "server 2" 4042

python server.py "server 3" 4043

python server.py "server 4" 4044

python server.py "server 5" 4045

go build . && ./load-balancer.exe

for i in {1..20}; do curl http://localhost:8080/; echo; done