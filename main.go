package main

import (
	// net/http fornece implementações de cliente e servidor http
	"net/http"
	// os nos da acesso a funcionalidade do sistema operacional
	// vá para a linha 22
	"os"
)

// O parâmetro w é a estrutura que usamos para enviar respostas a uma requisição HTTP. Ele implementa um método Write()
// que aceita um slice de bytes e grava os dados na conexão como parte de uma resposta HTTP.

// o parâmetro r representa a requisição HTTP recebida do cliente. É como acessamos os dados enviados por um navegador
// para o servidor. Ainda não o estamos usando aqui, mas com certeza faremos uso dele mais tarde.
// vá para a linha 37
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
	// usa o pacote os para setar a porta em 3000 caso não seja passada uma opção
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Essencialmente, um multiplexer de requisição compara a URL de requisição recebida com uma lista de padrões
	// registrados e chama o handler associado para ao padrão sempre que uma correspondência é encontrada.
	mux := http.NewServeMux()

	// O registro de requisição HTTP é feito por meio do método HandleFunc, que recebe a string
	// como seu primeiro argumento e uma função
	// vá para a linha 11
	mux.HandleFunc("/", indexHandler)

	// Método http.ListenAndServe() que inicia o servidor na porta determinada pela variável port.
	http.ListenAndServe(":"+port, mux)
}

// go build
// ./step-by-web
// abra o borwser na porta 3000 do localhost
