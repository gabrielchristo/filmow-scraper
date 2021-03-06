## Filmow Scraper

Um web scraper desenvolvido em Go para salvar todos os filmes e séries marcados como assistidos no [Filmow](https://filmow.com/).

A motivação é tanto a falta desta funcionalidade de maneira nativa quanto a falta de uma API oficial.

### Modus Operandi

O primeiro passo é obter todos os filmes e séries vistos pelo usuário, através da `<ul>` com ID `movies-list`.

Isso é feito pelo colly, tomando como base as URLs `https://filmow.com/usuario/%USERNAME%/filmes/ja-vi/` e `https://filmow.com/usuario/%USERNAME%/series/ja-vi/`, e visitando todas as páginas de filmes ou séries assistidas a partir de ambas.

Para cada item da lista acima iremos salvar o `data-movie-pk` correspondente, que será utilizado no endpoint `https://filmow.com/async/tooltip/movie/?movie_pk=` para obter informações.

O procedimento é o mesmo tanto para filmes quanto para séries.

### Utilização

- Instalar Go

- Instalar framework goColly
	- `go get github.com/gocolly/colly`

- Executar o script `run.sh`

### Saída

A saída do programa é uma tabela no formato CSV com as seguintes colunas:

- Nome Traduzido
- Nome Original
- Ano
- Nota Filmow

### Trabalhos Relacionados

A aplicação foi inspirada e baseada nos seguintes trabalhos:

- https://github.com/pauloralves/filmow_to_letterboxd_csv

- https://github.com/larissamoreira/filmow_to_letterboxd

- https://github.com/alvarofpp/imd0105-web-scraping-cinemas-natal

### ToDo

- salvar Nota IMDB ?

- salvar em csv, sqlite ou html ?

- paralelizar (goroutines) ou async ?

- sort movies by name ?

- pegar todos os diretores

- problema com ano