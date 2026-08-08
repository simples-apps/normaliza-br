# Normaliza BR

Biblioteca Go modular para normalização e validação de dados brasileiros.

Repositório: [github.com/simples-apps/normaliza-br](https://github.com/simples-apps/normaliza-br)

Cada pacote é um módulo Go independente (`github.com/simples-apps/normaliza-br/<pacote>`): importe só o que precisar, sem puxar o restante do projeto.

## Por que existe

Cadastros, APIs e integrações no Brasil lidam com formatos inconsistentes — CPF com e sem máscara, CNPJ alfanumérico, CEP pontuado, valores em `R$`, datas `DD/MM/AAAA`. Este repositório concentra regras de limpeza e validação estrutural para reduzir retrabalho e inconsistências.

## Módulos

| Módulo | Pacote | Escopo |
|--------|--------|--------|
| [Documentos](documentos/README.md) | `documentos` | CPF e CNPJ (numérico e alfanumérico) |
| [Localização](localizacao/README.md) | `localizacao` | CEP e UF |
| [Contatos](contatos/README.md) | `contatos` | Telefone e e-mail |
| [Financeiro](financeiro/README.md) | `financeiro` | Valores monetários |
| [Temporal](temporal/README.md) | `temporal` | Data e hora |

## Instalação

```bash
go get github.com/simples-apps/normaliza-br/documentos
go get github.com/simples-apps/normaliza-br/localizacao
go get github.com/simples-apps/normaliza-br/contatos
go get github.com/simples-apps/normaliza-br/financeiro
go get github.com/simples-apps/normaliza-br/temporal
```

No desenvolvimento local, o repositório usa um `go.work` com todos os módulos.

## Uso rápido

```go
import (
	"github.com/simples-apps/normaliza-br/contatos"
	"github.com/simples-apps/normaliza-br/documentos"
	"github.com/simples-apps/normaliza-br/financeiro"
	"github.com/simples-apps/normaliza-br/localizacao"
	"github.com/simples-apps/normaliza-br/temporal"
)

cpf := documentos.NormalizarCPF("123.456.789-09")
cnpj := documentos.ValidarCNPJ("12.ABC.345/01DE-35")
cep := localizacao.NormalizarCEP("13.080-300")
email := contatos.NormalizarEmail(" Usuario@Exemplo.COM ")
valor := financeiro.NormalizarMoeda("R$ 1.234,56")
data := temporal.NormalizarData("07/08/2026")
```

## Convenções da biblioteca

Os pacotes seguem o mesmo padrão:

| Função | Comportamento |
|--------|----------------|
| `NormalizarX` | Limpa/padroniza e devolve `string` |
| `ValidarX` | Valida e devolve `Resultado` (`Valido`, `Valor`, `Erro`) |
| `NormalizarXSeValido` | Normaliza só se a validação passar |

```go
type Resultado struct {
	Valido bool
	Valor  string
	Erro   string
}
```

## Testes

```bash
go test ./documentos ./localizacao ./contatos ./financeiro ./temporal
```

Ou, dentro de cada módulo:

```bash
cd documentos && go test ./...
```

## Referências

- [Receita Federal](https://www.gov.br/receitafederal)
- [Cálculo do DV do CNPJ alfanumérico](https://www.gov.br/receitafederal/pt-br/centrais-de-conteudo/publicacoes/documentos-tecnicos/cnpj)
- [Correios](https://www.correios.com.br/)
- [Anatel](https://www.gov.br/anatel)
- [Banco Central do Brasil](https://www.bcb.gov.br/)
- [ABNT](https://www.abnt.org.br/)
