# Documentos

[← Voltar ao README principal](../README.md)

Normalização e validação estrutural de documentos oficiais brasileiros, alinhadas à Receita Federal.

O retorno de validação usa [`comum.Resultado`](../comum/README.md).

**Import:** `github.com/simples-apps/normaliza-br/documentos`

## Escopo

| Tipo | Normalização | Validação |
|------|--------------|-----------|
| CPF | Remove não dígitos | 11 dígitos, rejeita sequências iguais, DV módulo 11 |
| CNPJ | Maiúsculas, mantém `A-Z`/`0-9` | 14 posições `[A-Z0-9]{12}[0-9]{2}`, DV módulo 11 (ASCII − 48) |

Também há `NormalizarDocumento`, que apenas remove espaços nas extremidades.

## CPF

```go
documentos.NormalizarCPF("123.456.789-09") // "12345678909"

resultado := documentos.ValidarCPF("123.456.789-09")
// resultado.Valido == true
// resultado.Valor  == "12345678909"

documentos.NormalizarCPFSeValido("111.111.111-11") // inválido (dígitos repetidos)
```

Regras aplicadas:

1. Exatamente 11 dígitos após a limpeza
2. Rejeição de valores com todos os dígitos iguais
3. Conferência dos dois dígitos verificadores (módulo 11)

## CNPJ

Aceita o formato numérico clássico e o **alfanumérico oficial** (em produção desde julho/2026). Os CNPJs numéricos existentes continuam válidos; o novo algoritmo de DV é compatível com ambos.

```go
documentos.NormalizarCNPJ("12.ABC.345/01DE-35") // "12ABC34501DE35"

resultado := documentos.ValidarCNPJ("12.345.678/0001-95")
// resultado.Valido == true
// resultado.Valor  == "12345678000195"

documentos.ValidarCNPJ("12.ABC.345/01DE-35") // válido (exemplo oficial da RF)
```

Regras aplicadas:

1. Normalização para maiúsculas e remoção de separadores/caracteres inválidos
2. Exatamente 14 posições: 12 alfanuméricas (`A-Z`, `0-9`) + 2 dígitos verificadores numéricos
3. Dígitos verificadores pelo módulo 11, convertendo cada caractere com `ASCII − 48` (conforme manual da Receita Federal)

Máscara de referência: `AA.AAA.AAA/AAAA-DV`

## Funções

| Função | Retorno |
|--------|---------|
| `NormalizarCPF(valor string)` | `string` |
| `ValidarCPF(valor string)` | `comum.Resultado` |
| `NormalizarCPFSeValido(valor string)` | `comum.Resultado` |
| `NormalizarCNPJ(valor string)` | `string` |
| `ValidarCNPJ(valor string)` | `comum.Resultado` |
| `NormalizarCNPJSeValido(valor string)` | `comum.Resultado` |
| `NormalizarDocumento(valor string)` | `string` |

## Referências

- [Receita Federal](https://www.gov.br/receitafederal)
- [Manual de cálculo do DV do CNPJ alfanumérico](https://www.gov.br/receitafederal/pt-br/centrais-de-conteudo/publicacoes/documentos-tecnicos/cnpj/manual-dv-cnpj.pdf)
- [Perguntas e respostas — CNPJ alfanumérico](https://www.gov.br/receitafederal/pt-br/centrais-de-conteudo/publicacoes/perguntas-e-respostas/cnpj/cnpj-alfanumerico.pdf)
