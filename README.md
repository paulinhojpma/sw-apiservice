# SW-API


O SW-API  é um aplicação que recebe requisições HTTP, em que é possível cadastra, remover, consultar, listar os planetas da saga Star Wars. Ele utiliza como banco de dados o MongoDB.

# REST API

O recurso principal é o `planet` ele  é formatado em JSON. Exemplo:

 ```json
   { "id": "608d8cce6c52289e68175c7b",
    "nome": "Hoth",
    "terreno": "tundra",
    "clima": "congelado",
    "aparicoes": "integer"  }
```

# Métodos

## Listar planetas

### Request

`GET /planetas/`



### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    []

**Exemplo: **

```json
{
   "planets":[
      {
         "id":"608f0c657a7ef3a683c79a0b",
         "nome":"Tatooine",
         "clima":"arido",
         "terreno":"deserto",
         "aparicoes":5
      },
      {
         "id":"608f0c667a7ef3a683c79a0c",
         "nome":"Alderaan",
         "clima":"temperado",
         "terreno":"montanhoso",
         "aparicoes":2
      },
      {
         "id":"608f0c667a7ef3a683c79a0d",
         "nome":"Hoth",
         "clima":"congelado",
         "terreno":"tundra",
         "aparicoes":1
      }
   ],
   "codResposta":200,
   "mensagem":"Planetas retornado com sucesso"
}
```
## Cadastrar novo planeta

### Request

`POST /planetas/`

`Request Body:`

```json
		{"nome": "Tatooine",
		"clima": "arido",
		"terreno": "deserto"}
```

### Response

    HTTP/1.1 201 Created
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 201 Created
    Connection: close
    Content-Type: application/json
    Location: /thing/1
    Content-Length: 36

  
**Exemplo:**
```json
{
   "planet":{
      "id":"608f0c657a7ef3a683c79a0b",
      "nome":"Tatooine",
      "clima":"arido",
      "terreno":"deserto",
      "aparicoes":5
   },
   "codResposta":201,
   "mensagem":"Planeta cadastrado com sucesso"
}
```
## Pegar um planeta pelo ID

### Request

`GET /planetas/id`



### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 36


**Exemplo:**
```json
{
   "planet":{
      "id":"608f0c657a7ef3a683c79a0b",
      "nome":"Tatooine",
      "clima":"arido",
      "terreno":"deserto",
      "aparicoes":5
   },
   "codResposta":200,
   "mensagem":"Planeta retornado com sucesso"
}
```

## Consutar planetas pelo nome

### Request

`GET /planetas?nome=####`



### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    

**Exemplo: **
```json
{
   "planets":[
      {
         "id":"608f0c657a7ef3a683c79a0b",
         "nome":"Tatooine",
         "clima":"arido",
         "terreno":"deserto",
         "aparicoes":5
      }
   ],
   "codResposta":200,
   "mensagem":"Planetas retornado com sucesso"
}
```

## Remover um planeta pelo ID

### Request

`DELETE /planetas/id`



### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 36


**Exemplo:**
```json
{
   "planet":null,
   "codResposta":200,
   "mensagem":"Planeta deletado com sucesso"
}
```
## Configuração

Antes de rodar você deve atualizar o path do mongoDB no arquivo  `conf.json` no atributo `db-host` para o url do seu banco MongoDB:

## Testes

Para rodar os testes e popular o banco com os primeiros planetas execute os comandos

```bash
$ go build
$ go test
```
 ## Iniciar a aplicação
 
 Após ter feito o build e os testes, para a rodar a aplicação rode o comando:
 
 ```bash
$ ./api-service
```
Para acessar localmente acesse:

```bash
http://localhost:8890/
```

