# Financeiro

[← Voltar ao README principal](../README.md)

Normalização e validação básica de valores monetários no formato brasileiro.

O retorno de validação usa [`comum.Resultado`](../comum/README.md).

**Import:** `github.com/simples-apps/normaliza-br/financeiro`

## Escopo

| Tipo | Normalização | Validação |
|------|--------------|-----------|
| Moeda | Remove `R$`, troca `.` de milhar e `,` decimal por `.` | Aceita dígitos, `.`, `,`, `R$`, espaços e `-` |

A saída normalizada é uma string numérica com ponto decimal (ex.: `"1234.56"`), adequada para parsing em `float`/`decimal` em camadas superiores.

## Moeda

```go
financeiro.NormalizarMoeda("R$ 1.234,56") // "1234.56"
financeiro.NormalizarMoeda("-1.000,00")   // "-1000.00"

resultado := financeiro.ValidarMoeda("R$ 1.234,56")
// resultado.Valido == true

financeiro.NormalizarMoedaSeValido("abc") // inválido
```

Ordem da normalização:

1. Trim
2. Remove pontos de milhar
3. Troca vírgula decimal por ponto
4. Remove prefixo `R$`
5. Descarta demais caracteres que não sejam dígito, ponto ou sinal `-`

## Funções

| Função | Retorno |
|--------|---------|
| `NormalizarMoeda(valor string)` | `string` |
| `ValidarMoeda(valor string)` | `comum.Resultado` |
| `NormalizarMoedaSeValido(valor string)` | `comum.Resultado` |

## Referências

- [Banco Central do Brasil](https://www.bcb.gov.br/)
