package news

import "net/http"

// Client struct representa o client para trabalhar com a news API.
// O campo httpClient aponta para o client HTTP que deve ser usado para fazer requisições,
// o campo apiKey contém a chave de API enquanto o campo PageSize contém o número de resultados a
// serem retornados por página (máximo de 100).
type Client struct {
	http     *http.Client
	key      string
	PageSize int
}

// NewClient cria e retorna uma nova instância de Client para fazer requisições à news API.
func NewClient(httpClient *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}

// iremos criar um client para trabalhar com a news API.
