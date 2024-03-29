# Filmow Scraper

Um web scraper desenvolvido em Go para salvar todos os filmes e séries marcados como `Já vi` ou `Quero ver` no [Filmow](https://filmow.com/).

A motivação é tanto a falta desta funcionalidade de maneira nativa quanto a falta de uma API oficial.

### Modus Operandi

O primeiro passo é obter todos os filmes e séries vistos pelo usuário, através da `<ul>` com ID `movies-list`.

Isso é feito pelo colly, tomando como base as URLs `https://filmow.com/usuario/%USERNAME%/filmes/ja-vi/` e `https://filmow.com/usuario/%USERNAME%/series/ja-vi/`, e visitando todas as páginas de filmes ou séries assistidas a partir de ambas.

Para cada item da lista acima iremos salvar o `data-movie-pk` correspondente, que será utilizado no endpoint `https://filmow.com/async/tooltip/movie/?movie_pk=` para obter informações.

O procedimento é o mesmo tanto para filmes quanto para séries.

### Utilização

- Instalar Go

- Clonar repositório
	- `git clone https://github.com/gabrielchristo/filmow-scraper.git`

- Instalar framework [goColly](https://github.com/gocolly/colly)
	- `go get github.com/gocolly/colly`

- Executar o script `run.sh` com o nome de usuário como argumento
	- `./run.sh nome_de_usuario`

### Saída

A saída do programa é uma tabela no formato CSV com as seguintes colunas:

- Filmow ID
- Nome Traduzido em pt-br
- Nome Original
- Número de comentários
- Ano
- Nota Filmow
- Diretor(es)

### Trabalhos Relacionados

A aplicação foi inspirada e baseada nos seguintes trabalhos:

- [Filmow to Letterboxd (Ruby)](https://github.com/pauloralves/filmow_to_letterboxd_csv)
- [Filmow to Letterboxd (Python)](https://github.com/larissamoreira/filmow_to_letterboxd)
- [Web Scraping Cinema](https://github.com/alvarofpp/imd0105-web-scraping-cinemas-natal)

### ToDo

- salvar user_rate
- salvar imdb ID/Nota IMDB ?
- salvar html ?
- sort movies by name ?

