<h1 align="center">API ações sociais</h1>

## 📜 Sumário

- [Sobre](#Sobre)
- [Libs](#Libs)
- [Iniciar](#Iniciar)
- [Testes](#Testes)
- [Endpoints](#Endpoints)

---

<a id="Sobre"></a> 
## 📃 Sobre

API desenvolvida para realizar operações CRUD (cadastro, leitura, atualização, deleção) para o gerenciamento de ações sociais. O sistema comporta
a criação de voluntários e ações sociais que serão realizados em um determinado local. Antes de cadastrar as ações e seus voluntários para aquela ação,
é necessário o registro de voluntários. Após isso, no momento da criação de uma ação social, escolhe-se os voluntários dela. 

---

<a id="Libs"></a> 
## 🗄 Libs </br>

| Name        | Description | Documentation  |
| ----------- | ----------- | ------------- |    
| chi      | web framework      |  github.com/go-chi/chi |
| google/uuid | uuid generator     | github.com/google/uuid  | 
|  jackc/pgx  |   postgres database driver   | github.com/jackc/pgx/v4    |
| godotenv            |  .env variables reader                           | github.com/joho/godotenv  |
| go-cmp            |  lib tests assertions                          | github.com/google/go-cmp |

---

<a id="Iniciar"></a> 
## ⚙️ Iniciar

Clone o repositório do projeto:

```bash
    $ git clone https://github.com/MCarbono/go-social-action.git
``` 

Acesse a pasta do projeto

```bash
    $ cd go-social-action
```

Inicie o banco de dados. Digite um dos comandos abaixo:

```bash
    # docker comando
    $ docker compose up -d
```

```bash
    # Makefile comando
    $ make docker_up
```

depois, para rodar o projeto localmente, digite um dos comandos abaixo:

```bash
    # Go comando
    $ go run main.go
```

```bash
    # Makefile comando
    $ make run
```

Após digitar um dos comandos acima, você terá subido uma máquina docker com a api exposta na porta 3000.

---

<a id="Testes"></a> 
## 🧪 Testes 

### Unidade

Na raiz do projeto, digite um dos comandos abaixo:

```bash
    # Makefile comando
    $ make unit_test
```

```bash
    # Go comando
    $ go test ./domain/entity -v
```

### Integração

na raiz do projeto, após subir o docker do banco de dados, digite um dos comandos abaixo no terminal:

```bash
    # Makefile comando
    $ make test_integration
```

```bash
    # Go comando
    $ go test ./test/integration -v
```
---

<a id="Endpoints"></a> 
## 💻 Endpoints

### Criação de um voluntário: 

Request:

```bash
    $ curl --location 'http://localhost:3000/volunteers' \
    --header 'Content-Type: application/json' \
    --data '{
        "first_name": "name de teste",
        "last_name": "last name de teste",
        "neighborhood": "bairro de teste",
        "city": "cidade de teste"
    }'
```

Response:

Status code: 201<br>
body: 

```json
    {
        "id": "0a860831-04a1-4bfd-b2ac-6f1f167d502e",
        "first_name": "name de teste",
        "last_name": "last name de teste",
        "neighborhood": "bairro de teste",
        "city": "cidade de teste",
        "created_at": "2023-07-20T01:13:38.439164Z",
        "updated_at": "2023-07-20T01:13:38.439164Z"
    }
``` 
### Informações de um voluntário: 

Request:

```bash
    $ curl --location 'http://localhost:3000/volunteers/0a860831-04a1-4bfd-b2ac-6f1f167d502e'
```

Response:

Status code: 200<br>
body: 

```json
    {
        "id": "0a860831-04a1-4bfd-b2ac-6f1f167d502e",
        "first_name": "name de teste",
        "last_name": "last name de teste",
        "neighborhood": "bairro de teste",
        "city": "cidade de teste",
        "created_at": "2023-07-20T01:13:38.439164Z",
        "updated_at": "2023-07-20T01:13:38.439164Z"
    }
``` 

Status code: 404<br>
body: 

```json
    {
        "err": "volunteer not found"
    }
``` 


### Criação de uma ação social

Request:

```bash
    $ curl --location 'http://localhost:3000/social-actions' \
    --header 'Content-Type: application/json' \
    --data '
    {
        "name": "social action test",
        "organizer": "organizer test",
        "description": "description test",
        "street_line": "streetLine test",
        "street_number": "streetNumber test",
        "neighborhood": "neighborhoor test",
        "city": "city test",
        "social_action_volunteers": [
            "0a860831-04a1-4bfd-b2ac-6f1f167d502e"
        ]
    }'
```

Response:

Status code: 201<br>
body: 

```json
    {
    "id": "daa30269-fdd2-41f4-b374-665ead8c9307",
    "name": "social action test",
    "organizer": "organizer test",
    "description": "description test",
    "address": {
        "street_line": "streetLine test",
        "street_number": "streetNumber test",
        "neighborhood": "neighborhoor test",
        "city": "city test"
    },
    "social_action_volunteers": [
        {
            "id": "0a860831-04a1-4bfd-b2ac-6f1f167d502e",
            "social_action_id": "daa30269-fdd2-41f4-b374-665ead8c9307",
            "first_name": "name de teste",
            "last_name": "last name de teste",
            "neighborhoor": "bairro de teste",
            "city": "cidade de teste"
        }
    ],
    "created_at": "2023-07-20T14:48:26.218164Z",
    "updated_at": "2023-07-20T14:48:26.218164Z"
}
``` 

### Informações de uma ação social: 

Request:

```bash
    $ curl --location 'http://localhost:3000/social-actions/daa30269-fdd2-41f4-b374-665ead8c9307'
```

Response:

Status code: 200<br>
body: 

```json
   {
        "id": "daa30269-fdd2-41f4-b374-665ead8c9307",
        "name": "social action test",
        "organizer": "organizer test",
        "description": "description test",
        "address": {
            "street_line": "streetLine test",
            "street_number": "streetNumber test",
            "neighborhood": "neighborhoor test",
            "city": "city test"
        },
        "social_action_volunteers": [
            {
                "id": "0a860831-04a1-4bfd-b2ac-6f1f167d502e",
                "social_action_id": "daa30269-fdd2-41f4-b374-665ead8c9307",
                "first_name": "name de teste",
                "last_name": "last name de teste",
                "neighborhoor": "bairro de teste",
                "city": "cidade de teste"
            }
        ],
        "created_at": "2023-07-20T14:48:26.218164Z",
        "updated_at": "2023-07-20T14:48:26.218164Z"
    }
``` 

Status code: 404<br>
body: 

```json
    {
        "err": "social action not found"
    }
``` 
### Informações de todas as ações sociais: 

Request:

```bash
    $ curl --location 'http://localhost:3000/social-actions/'
```

Response:

Status code: 200<br>
body: 

```json
   {
        [
            {
                "id": "daa30269-fdd2-41f4-b374-665ead8c9307",
                "name": "social action test",
                "organizer": "organizer test",
                "description": "description test",
                "address": {
                    "street_line": "streetLine test",
                    "street_number": "streetNumber test",
                    "neighborhood": "neighborhoor test",
                    "city": "city test"
                },
                "social_action_volunteers": [
                    {
                        "id": "0a860831-04a1-4bfd-b2ac-6f1f167d502e",
                        "social_action_id": "daa30269-fdd2-41f4-b374-665ead8c9307",
                        "first_name": "name de teste",
                        "last_name": "last name de teste",
                        "neighborhoor": "bairro de teste",
                        "city": "cidade de teste"
                    }
                ],
                "created_at": "2023-07-20T14:48:26.218164Z",
                "updated_at": "2023-07-20T14:48:26.218164Z"
            }
        ]
    }
``` 

### Exclusão de uma ação social: 

Request:

```bash
    $ curl --location --request DELETE 'http://localhost:3000/social-actions/daa30269-fdd2-41f4-b374-665ead8c9307'
```

Response:

Status code: 204 - No content<br>

### Edição de uma ação social:

Requests:

Edição dos dados de uma ação social: 

```bash
    $ curl --location --request PATCH 'http://localhost:3000/social-actions/fcef499e-2db0-4620-ad08-947141694b94' \
    --header 'Content-Type: application/json' \
    --data '
    {
        "name": "social action test updated",
        "organizer": "organizer test updated",
        "description": "description test updated",
        "street_line": "streetLine test updated",
        "street_number": "streetNumber test updated",
        "neighborhood": "neighborhoor test updated",
        "city": "city test updated"
    }'
```

#### Edição ou inclusão de um novo voluntário a uma determinada ação social.

Para incluir um voluntário é necessário apenas enviar
no corpo da requisição o id. Caso seja uma atualização de um um voluntário que já está cadastrado, é preciso enviar o id com o 
atributo que é necessário atualizar.

```bash
    $ curl --location --request PATCH 'http://localhost:3000/social-actions/49f946eb-faf5-4eff-bea0-88dfff3c3ee4' \
        --header 'Content-Type: application/json' \
        --data '
        {
            "social_action_volunteers": [
                {
                    "id": "7c01c7c3-e1dd-4f91-9003-5e1642975cf6"
                },
                {
                    "id": "4fb6791c-fc79-4220-9b08-87a14588659e",
                    "first_name": "updated name de teste",
                    "last_name": "updated last name de teste",
                    "neighborhood": "updated bairro de teste",
                    "city": "updated cidade de teste"
                }
            ]
        }'
```

Responses:

Status code: 204 - No content<br>
Status code: 404<br>
body: 

```json
    {
        "err": "social action not found"
    }
``` 