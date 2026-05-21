# Aprendendo Go 🚀

Repositório para aprendizado prático de Go, com exemplos progressivos de complexidade.

## 📋 Índice

- [Como Executar](#como-executar)
- [Projetos](#projetos)
- [Pré-requisitos](#pré-requisitos)
- [Comandos Go Essenciais](#comandos-go-essenciais)

## Como Executar

### Executar um Programa

```bash
# Navegar até o diretório do projeto
cd <projeto>

# Executar diretamente
go run main.go
# ou
go run .

# Compilar para um executável
go build

# Executar o executável gerado
./main.exe  # Windows
./main      # Linux/Mac
```

### Build e Run

```bash
# Build gerando executável
go build -o nome_executavel

# Executar com argumentos
go run main.go arg1 arg2
```

## Projetos

### 1. **hello-world** - Primeiro Programa
📁 `hello-world/`

**Descrição:** Projeto básico introdutório ao Go. Imprime uma mensagem simples.

**Conceitos:**
- Estrutura básica de um programa Go
- Função `main()`
- Import de pacotes
- Função `fmt.Println()`

**Como executar:**
```bash
cd hello-world
go run hello.go
```

---

### 2. **funcao** - Funções e Lógica
📁 `funcao/`

**Descrição:** Projeto focado em definição e manipulação de funções em Go.

**Conceitos:**
- Declaração de funções
- Parâmetros e retorno de valores
- Múltiplos retornos
- Tratamento de erros

**Como executar:**
```bash
cd funcao
go run funcao.go
```

---

### 3. **web-service-gin** - Web Service REST
📁 `web-service-gin/`

**Descrição:** Aplicação web completa usando o framework **Gin**. Implementa um CRUD de álbuns com arquitetura em camadas.

**Estrutura do Projeto:**

```
web-service-gin/
├── main.go                          # Entry point da aplicação
├── go.mod                           # Gerenciador de dependências
├── controller/
│   └── album_controller.go          # Handlers das rotas HTTP
├── service/
│   └── album_service.go             # Lógica de negócio
├── repository/
│   └── inmemory_album_Repository.go # Acesso aos dados (in-memory)
├── domain/
│   └── album.go                     # Modelo de domínio
└── dto/
    ├── album_create_request.go      # DTO para criação
    ├── album_response.go            # DTO para resposta
    └── album_update_request.go      # DTO para atualização
```

**Conceitos:**

- **Arquitetura em Camadas:**
  - `Controller` → Recebe requisições HTTP
  - `Service` → Implementa regras de negócio
  - `Repository` → Acessa dados
  - `Domain` → Modelos de negócio
  - `DTO` → Transferência de dados

- **Framework Gin:** Roteamento e handling de requisições HTTP
- **Tipos e Structs** em Go
- **Interfaces** e contrato de dados
- **Padrão Repository** para acesso a dados

**Endpoints Disponíveis:**

| Método | Rota | Descrição |
|--------|------|-----------|
| GET | `/albums` | Lista todos os álbuns |
| GET | `/albums/:id` | Obtém um álbum por ID |
| POST | `/albums` | Cria um novo álbum |
| PUT | `/albums/:id` | Atualiza um álbum |
| DELETE | `/albums/:id` | Deleta um álbum |

**Como executar:**

```bash
cd web-service-gin

# Instalar dependências (se necessário)
go mod download

# Executar a aplicação
go run main.go

# A aplicação ficará disponível em http://localhost:8080
```

**Testando os Endpoints:**

```bash
# GET - Listar todos os álbuns
curl http://localhost:8080/albums

# POST - Criar novo álbum
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Album Teste","artist":"Artista"}'

# GET - Obter álbum específico
curl http://localhost:8080/albums/1

# PUT - Atualizar álbum
curl -X PUT http://localhost:8080/albums/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Novo Título","artist":"Novo Artista"}'

# DELETE - Deletar álbum
curl -X DELETE http://localhost:8080/albums/1
```

---

## Pré-requisitos

### Instalação do Go

1. Baixe e instale em [golang.org](https://golang.org/dl/)
2. Configure variáveis de ambiente:
   - `GOPATH`: Diretório de trabalho do Go
   - `GOROOT`: Diretório de instalação do Go

**Verificar instalação:**
```bash
go version
go env
```

### Dependências

O projeto utiliza:
- **Gin Framework** (web-service-gin)
  ```bash
  go get -u github.com/gin-gonic/gin
  ```

## Comandos Go Essenciais

```bash
# Verificar versão
go version

# Visualizar variáveis de ambiente
go env

# Executar programa
go run <arquivo.go>

# Compilar para executável
go build

# Compilar e renomear
go build -o <nome>

# Gerenciar dependências
go mod init <nome-modulo>      # Inicializar módulo
go mod tidy                     # Limpar dependências não usadas
go mod download                 # Baixar dependências

# Testes
go test ./...                   # Rodar todos os testes
go test -v                      # Testes com verbose

# Formatação
go fmt ./...                    # Formatar código
go vet ./...                    # Verificar erros potenciais

# Instalar pacote global
go install <pacote>
```

## 📚 Recursos Adicionais

- [Documentação Oficial Go](https://golang.org/doc/)
- [Go Playground](https://play.golang.org/)
- [Gin Framework](https://gin-gonic.com/)
- [Effective Go](https://golang.org/doc/effective_go)

---

**Desenvolvido durante o aprendizado de Go** 🎓
