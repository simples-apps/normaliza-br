# Comum

[← Voltar ao README principal](../README.md)

Tipos compartilhados pelos demais pacotes da biblioteca.

**Import:** `github.com/simples-apps/normaliza-br/comum`

## Escopo

| Tipo | Uso |
|------|-----|
| `Resultado` | Retorno padronizado de `ValidarX` e `NormalizarXSeValido` |

Os pacotes `documentos`, `localizacao`, `contatos`, `financeiro` e `temporal` devolvem `comum.Resultado`.

## Resultado

```go
import "github.com/simples-apps/normaliza-br/comum"

type Resultado struct {
	Valido bool
	Valor  string
	Erro   string
}
```

| Campo | Significado |
|-------|-------------|
| `Valido` | `true` quando a entrada passou na validação |
| `Valor` | Dado normalizado quando válido; vazio quando inválido |
| `Erro` | Motivo da rejeição quando inválido; vazio quando válido |

## Exemplo

```go
import (
	"github.com/simples-apps/normaliza-br/comum"
	"github.com/simples-apps/normaliza-br/documentos"
)

var resultado comum.Resultado = documentos.ValidarCPF("529.982.247-25")
if !resultado.Valido {
	// resultado.Erro
}
```
