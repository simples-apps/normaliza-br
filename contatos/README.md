# Contatos

[← Voltar ao README principal](../README.md)

Normalização e validação básica de meios de comunicação: telefone e e-mail.

**Import:** `github.com/simples-apps/normaliza-br/contatos`

## Escopo

| Tipo | Normalização | Validação |
|------|--------------|-----------|
| Telefone | Remove tudo que não for dígito | — |
| E-mail | Trim + minúsculas | Contém `@` e `.` |

A validação de e-mail é estrutural e deliberadamente simples (não consulta DNS nem RFC completa).

## Telefone

```go
contatos.NormalizarTelefone("(11) 99999-8888") // "11999998888"
contatos.NormalizarTelefone("+55 11 99999-8888") // "5511999998888"
```

## E-mail

```go
contatos.NormalizarEmail(" Usuario@Exemplo.COM ") // "usuario@exemplo.com"

resultado := contatos.ValidarEmail(" Usuario@Exemplo.COM ")
// resultado.Valido == true
// resultado.Valor  == "usuario@exemplo.com"

contatos.NormalizarEmailSeValido("emailinvalido") // inválido
```

## Funções

| Função | Retorno |
|--------|---------|
| `NormalizarTelefone(valor string)` | `string` |
| `NormalizarEmail(valor string)` | `string` |
| `ValidarEmail(valor string)` | `Resultado` |
| `NormalizarEmailSeValido(valor string)` | `Resultado` |

## Referências

- [Anatel](https://www.gov.br/anatel)
