# 🔗 Encurtador de URL (Go)

Serviço simples de encurtamento de URLs escrito em Go puro (sem frameworks web), usando apenas a `net/http` da standard library. As URLs originais são criptografadas com **AES-CTR** antes de serem armazenadas em memória.

## Como funciona

```
POST/GET /shorten?url=https://exemplo.com/pagina-longa
          │
          ├─ valida protocolo (http/https)
          ├─ criptografa a URL original (AES-CTR + IV aleatório)
          ├─ gera um ID curto aleatório (6 caracteres)
          └─ guarda {shortID: urlCriptografada} em um map protegido por mutex

GET /{shortID}
          │
          ├─ busca a URL criptografada pelo shortID
          ├─ descriptografa
          └─ redireciona (302) para a URL original
```

## Stack

- **Go** 1.26
- [`github.com/joho/godotenv`](https://github.com/joho/godotenv) — carrega variáveis de ambiente do `.env`
- `crypto/aes` + `crypto/cipher` — criptografia simétrica das URLs armazenadas
- Armazenamento **em memória** (`map[string]string`), sem banco de dados

## Estrutura do projeto

```
.
├── main.go                       # bootstrap do servidor HTTP e rotas
├── src/
│   ├── config/
│   │   └── config.go              # carrega .env, porta, secret e o "store" em memória
│   ├── handlers/
│   │   ├── shortURL.go            # POST/GET /shorten — cria a URL curta
│   │   └── redirectURL.go         # GET /{id} — redireciona para a URL original
│   └── utils/
│       ├── encrypt-url.go         # criptografa a URL original (AES-CTR)
│       ├── decrypt-url.go         # descriptografa a URL original
│       └── generate-short-id.go   # gera o ID curto aleatório
├── .env.example
└── go.mod
```

## Rodando localmente

1. Clone o repositório e entre na pasta do projeto.
2. Copie o arquivo de exemplo de variáveis de ambiente:

   ```bash
   cp .env.example .env
   ```

3. Preencha o `.env`:

   ```env
   PORT=5000
   SECRET=uma-chave-secreta-com-16-24-ou-32-bytes
   ```

   > ⚠️ O `SECRET` é usado como chave AES e **precisa ter exatamente 16, 24 ou 32 bytes** (AES-128/192/256), senão o servidor falha ao iniciar.

4. Baixe as dependências e rode:

   ```bash
   go mod tidy
   go run main.go
   ```

5. O servidor sobe em `http://localhost:5000` (ou na porta definida em `PORT`).

## Endpoints

### Encurtar uma URL

```
GET /shorten?url=https://www.google.com
```

**Resposta:**
```
A URL encurtada desta url original é: http://localhost:5000/aB3xY9
```

### Acessar a URL encurtada

```
GET /{shortID}
```

Redireciona (HTTP 302) para a URL original.

## Licença

Projeto pessoal / de estudo — sem licença definida.
