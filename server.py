from flask import Flask
import sys
app = Flask(__name__)

#Servidor simples que recebe requisições no path '/' e retorna o nome do servidor

#nome é definido quando o script é executado
serverName = sys.argv[1]

@app.route('/')
def hello():
    return serverName

if __name__ == '__main__':
    #port é definido quando o script é executado
    app.run(port=sys.argv[2])