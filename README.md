# Comparador de APIs de CEP

Este projeto é uma aplicação simples em Go que realiza requisições para diferentes APIs de consulta de CEP (BrasilAPI e ViaCEP) e identifica qual delas responde mais rápido.

## Descrição

O programa faz requisições simultâneas para duas APIs de CEP:
- [BrasilAPI](https://brasilapi.com.br/api/cep/v1/01001000)
- [ViaCEP](https://viacep.com.br/ws/01001000/json)

Utilizando canais (`channels`), o programa recebe as respostas e determina qual API foi a mais rápida a retornar os dados. Em caso de atraso de ambas, um **timeout** é acionado após 1 segundo.

## Funcionamento

1. **Requisições assíncronas:** 
   Utiliza `goroutines` para realizar requisições HTTP de forma simultânea.

2. **Canais para comunicação:**
   As respostas são enviadas pelos canais para o programa principal.

3. **Timeout:**
   Caso nenhuma das APIs responda em 1 segundo, o programa finaliza com a mensagem "Timeout".

4. **Exibição de resultados:**
   O programa exibe:
   - Nome da API que respondeu mais rápido.
   - URL usada na requisição.
   - Dados retornados em formato JSON.

## Exemplo de Saída

```plaintext
The fastest is: ViaCEP with URL: https://viacep.com.br/ws/01001000/json 
Returned data: {"cep":"01001-000","logradouro":"Praça da Sé","complemento":"lado ímpar","bairro":"Sé","localidade":"São Paulo","uf":"SP","ibge":"3550308","gia":"1004","ddd":"11","siafi":"7107"}
```

## Pré-requisitos

Para executar a aplicação, é necessário ter o Go instalado na máquina. A instalação pode ser feita conforme as instruções no [site oficial do Go](https://golang.org/dl/).

## Instruções de Execução

1. Clone o repositório e navegue até o diretório do projeto.
   
   ```bash
   git clone https://github.com/brunoofgod/goexpert-lesson-2.git
   cd goexpert-lesson-2
   ```

2. Execute o código:
   
   ```bash
   go run main.go
   ```

