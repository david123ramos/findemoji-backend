# BuscaEmoji - Back-end

Scripts da API que recebe uma string (Português ou Inglês) como parâmetro e faz o webscrapping no [EmojiPedia](https://emojipedia.org/) 

Fiz uso do pacote [MUX](https://github.com/gorilla/mux) para definir a rota da API e do [GOQuery](https://github.com/PuerkitoBio/goquery) para parsear o HTML.

## Endpoint

Faça a chamada a essa url onde "{}" deve ser substituído pela string a ser buscada
###### Exemplo GET> https://buscaemoji.herokuapp.com/emoji?s=cachorro

```
  https://buscaemoji.herokuapp.com/emoji?s={}  
```
