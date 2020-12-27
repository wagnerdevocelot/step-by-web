package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Article é uma representação em go do que seria o retorno em json da api
type Article struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`

	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

// Results é exatamente a mesma que a mostrada anteriormente, exceto que Article agora é
// parte de sua própria strutct em vez de ser definida inline como antes.
type Results struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// Client ...
type Client struct {
	http     *http.Client
	key      string
	PageSize int
}

// NewClient ...
func NewClient(httpClient *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}

// FetchEverything endpoint aceita dois argumentos: a query de pesquisa e page.
// Eles são anexados ao URL da requisição, além da API key e do tamanho da página. Observe que a query de
// é codificada em URL por meio do método QueryEscape().
func (c *Client) FetchEverything(query, page string) (*Results, error) {
	endpoint := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&pageSize=%d&page=%s&apiKey=%s&sortBy=publishedAt&language=en", url.QueryEscape(query), c.PageSize, page, c.key)
	// A requisição HTTP é feita por meio do cliente HTTP personalizado que criamos anteriormente.
	// Este cliente personalizado foi definido para tempo limite após 10 segundos. O cliente padrão não tem nenhum tempo
	// limite, portanto, não é recomendado para uso em produção.
	resp, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	// Se a resposta da API não for 200 OK, um erro genérico será retornado.
	defer resp.Body.Close()

	// Caso contrário, o body da requisição é convertido em um byte slice usando o método ioutil.ReadAll()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	// e subsequentemente decodificado na struct Result por meio do método json.Unmarshal().
	res := &Results{}
	return res, json.Unmarshal(body, res)
}
