package main

import (
	// importação do pacote de template, html/template é uma opção melhor pois evita code injection
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// tpl é uma variável de package level scope que aponta para uma definição de template a partir dos arquivos fornecidos.
// A chamada para template.ParseFiles analisa e valida o arquivo index.html na raiz do nosso diretório de projeto.

// A invocação de template.ParseFiles é envolvida com template.Must para que o código entre em panic se um erro for
// obtido durante a análise do arquivo de template. O motivo de panic aqui, em vez de tentar lidar com o erro
// é porque um webApp com um template corrompido não é exatamente um webApp. É um problema que
// deve ser corrigido antes de tentar reiniciar o servidor.
var tpl = template.Must(template.ParseFiles("index.html"))

// Na função indexHandler, o template tpl é executado fornecendo dois argumentos: onde queremos escrever a saída
// e os dados que queremos passar para o template.

// No caso acima, estamos escrevendo a saída para a interface ResponseWriter e, uma vez que não temos nenhum dado para
// passar para nosso template no momento, nil é passado como o segundo argumento.
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

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}

// faça o build novamente e suba o server na porta 3000 para verificar o estado da aplicação
