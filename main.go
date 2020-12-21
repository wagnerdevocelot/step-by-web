package main

import (
	// importação do pacote log pra manipular erros
	"log"

	"net/http"
	"os"

	// importação do pacote para definir variavel de ambiente
	"github.com/joho/godotenv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {

	// O método Load lê o arquivo .env e carrega as variáveis de ambiente ​​definidas para que possam ser acessadas
	// por meio do método os.Getenv(). Isso é especialmente útil para armazenar credenciais secretas no ambiente,
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
