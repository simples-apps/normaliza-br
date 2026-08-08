# Localização

[← Voltar ao README principal](../README.md)

Normalização e validação de dados geográficos brasileiros: CEP e unidade federativa (UF).

O retorno de validação usa [`comum.Resultado`](../comum/README.md).

**Import:** `github.com/simples-apps/normaliza-br/localizacao`

## Escopo

| Tipo | Normalização | Validação |
|------|--------------|-----------|
| CEP | Remove `.` e `-` | Exatamente 8 caracteres |
| Estado (UF) | Maiúsculas; mapeia nomes conhecidos (ex.: São Paulo → `SP`) | Código com 2 letras `A-Z` |

## CEP

```go
localizacao.NormalizarCEP("13.080-300") // "13080300"

resultado := localizacao.ValidarCEP("13.080-300")
// resultado.Valido == true
// resultado.Valor  == "13080300"

localizacao.NormalizarCEPSeValido("123") // inválido
```

## Estado (UF)

```go
localizacao.NormalizarEstado("são paulo") // "SP"
localizacao.NormalizarEstado("sp")       // "SP"

resultado := localizacao.ValidarEstado("SP")
// resultado.Valido == true
// resultado.Valor  == "SP"
```

`NormalizarEstado` reconhece as 27 UFs pela sigla e também o nome “São Paulo” / “Sao Paulo”. Outros nomes por extenso ainda não são mapeados e retornam o texto original (após trim/maiúsculas).

`ValidarEstado` exige exatamente duas letras `A-Z` após a normalização básica (trim + maiúsculas).

## Funções

| Função | Retorno |
|--------|---------|
| `NormalizarCEP(valor string)` | `string` |
| `ValidarCEP(valor string)` | `comum.Resultado` |
| `NormalizarCEPSeValido(valor string)` | `comum.Resultado` |
| `NormalizarEstado(valor string)` | `string` |
| `ValidarEstado(valor string)` | `comum.Resultado` |

## Referências

- [Correios](https://www.correios.com.br/)
- [IBGE — Unidades da Federação](https://www.ibge.gov.br/explica/codigos-dos-municipios.php)
