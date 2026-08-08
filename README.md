# Normaliza BR

Biblioteca Go para normalização e validação de dados brasileiros.

Repositório: [github.com/simples-apps/normaliza-br](https://github.com/simples-apps/normaliza-br)

Há **um único módulo** Go na raiz. Cada pasta é um **pacote** independente: você importa só o que precisa (`documentos`, `temporal`, etc.) sem escrever o código dos demais no seu projeto.

## Por que existe

Cadastros, APIs e integrações no Brasil lidam com formatos inconsistentes — CPF com e sem máscara, CNPJ alfanumérico, CEP pontuado, valores em `R$`, datas `DD/MM/AAAA`. Este repositório concentra regras de limpeza e validação estrutural para reduzir retrabalho e inconsistências.

## Pacotes

| Pacote | Escopo |
|--------|--------|
| [comum](comum/README.md) | Tipos compartilhados (`Resultado`) |
| [documentos](documentos/README.md) | CPF e CNPJ (numérico e alfanumérico) |
| [localizacao](localizacao/README.md) | CEP e UF |
| [contatos](contatos/README.md) | Telefone e e-mail |
| [financeiro](financeiro/README.md) | Valores monetários |
| [temporal](temporal/README.md) | Data e hora |

## Instalação

```bash
go get github.com/simples-apps/normaliza-br@latest
```

Depois importe apenas os pacotes necessários:

```go
import (
	"github.com/simples-apps/normaliza-br/comum"
	"github.com/simples-apps/normaliza-br/documentos"
)
```

## Uso rápido

```go
import (
	"github.com/simples-apps/normaliza-br/contatos"
	"github.com/simples-apps/normaliza-br/documentos"
	"github.com/simples-apps/normaliza-br/financeiro"
	"github.com/simples-apps/normaliza-br/localizacao"
	"github.com/simples-apps/normaliza-br/temporal"
)

cpf := documentos.NormalizarCPF("529.982.247-25")
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
| `ValidarX` | Valida e devolve `comum.Resultado` (`Valido`, `Valor`, `Erro`) |
| `NormalizarXSeValido` | Normaliza só se a validação passar |

O tipo `Resultado` vive em [comum](comum/README.md) e é reutilizado por todos os demais:

```go
import "github.com/simples-apps/normaliza-br/comum"

type Resultado struct {
	Valido bool
	Valor  string
	Erro   string
}
```

## Testes

```bash
go test ./...
```

## Referências

- [Receita Federal](https://www.gov.br/receitafederal)
- [Cálculo do DV do CNPJ alfanumérico](https://www.gov.br/receitafederal/pt-br/centrais-de-conteudo/publicacoes/documentos-tecnicos/cnpj)
- [Correios](https://www.correios.com.br/)
- [Anatel](https://www.gov.br/anatel)
- [Banco Central do Brasil](https://www.bcb.gov.br/)
- [ABNT](https://www.abnt.org.br/)
