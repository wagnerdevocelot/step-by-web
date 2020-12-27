package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"

	// importação do pacote time
	"time"

	"github.com/joho/godotenv"
	// importação do pacote news.go
	"github.com/wagnerdevocelot/step-by-web/news"
)

var newsapi *news.Client

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

// Outra abordagem seria utilizar uma clojure para acessar o client newsapi. Esta é potencialmente uma solução melhor
// pois torna o teste muito mais fácil.
func searchHandler(newsapi *news.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

	}
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

	// Precisamos acessar a variável "newsapi" dentro de searchHandler para que possamos usá-la para fazer requisições
	// ao newsapi.org. Poderíamos tornar newsapi uma variável de package level scope e atribuir o valor de retorno de NewClient()
	// a ela para que possamos acessá-la de qualquer lugar no packge main
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}
	newsapi := news.NewClient(myClient, apiKey, 20)

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// A função searchHandler agora aceita um ponteiro para news.Client e retorna uma função anônima
	// que satisfaz o tipo http.HandlerFunc. Esta função fecha sobre o parâmetro newsapi, o que significa que terá acesso
	// a ele sempre que for chamado.
	mux.HandleFunc("/search", searchHandler(newsapi))
	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)
}
