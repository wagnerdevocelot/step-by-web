package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

// Essa rota espera dois parâmetros de consulta: "q" representa a consulta do usuário e "page" é usado para
// percorrer os resultados. Este parâmetro page é opcional. Se não estiver incluído no URL, presumiremos
// apenas que a página é 1.
func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	searchQuery := params.Get("q")
	page := params.Get("page")
	if page == "" {
		page = "1"
	}

	fmt.Println("Search Query is: ", searchQuery)
	fmt.Println("Page is: ", page)

	// O código acima extrai os parâmetros "q" e "page" da URL de requisição e os imprime na saída padrão.
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

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// Registre a função searchHandler como handler para o padrão /search conforme mostrado abaixo
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)
}

//faça o build rode o código e faça uma busca, verifique no terminal o resultado com o nome da query + numero da página
