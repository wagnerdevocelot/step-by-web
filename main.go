package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// instancia um objeto file server passando o diretório onde todos os nossos arquivos estáticos estão
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()

	// precisamos dizer ao nosso router para usar este objeto de file server para todos
	// os caminhos que começam com o prefixo /assets/:

	// O método http.StripPrefix() modifica o URL da requisição removendo o prefixo especificado
	// antes de encaminhar a tratamento da requisição ao http.Handler no segundo parâmetro.

	// Se uma requisição for feita para o arquivo /assets/style.css, StripPrefix() cortará o /assets/
	// e encaminhará a requisição modificada para o handler retornado por http.FileServer()
	// para que ele veja o resource solicitado como style.css. Em seguida, ele procurará e servirá o resource
	// relativo à pasta especificada como o diretório raiz para o arquivo estático.

	// O uso de Handle em vez de HandleFunc aqui, ocorre porque o método http.FileServer()
	// retorna um tipo http.Handler em vez de um HandlerFunc
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)
}
