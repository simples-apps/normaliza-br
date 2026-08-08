# Temporal

[← Voltar ao README principal](../README.md)

Normalização e validação de datas e horas no padrão brasileiro, com saída em formatos ISO comuns.

O retorno de validação usa [`comum.Resultado`](../comum/README.md).

**Import:** `github.com/simples-apps/normaliza-br/temporal`

## Escopo

| Tipo | Entrada típica | Saída normalizada |
|------|----------------|-------------------|
| Data | `DD/MM/AAAA` | `AAAA-MM-DD` |
| Hora | `HH:MM` | `HH:MM:00` |
| Data e hora | `DD/MM/AAAA HH:MM` | `AAAA-MM-DDTHH:MM:00` |

## Data

```go
temporal.NormalizarData("07/08/2026") // "2026-08-07"

resultado := temporal.ValidarData("07/08/2026")
// resultado.Valido == true
// resultado.Valor  == "07/08/2026"  // ValidarData preserva a entrada válida

temporal.NormalizarDataSeValido("07/08/2026")
// Valor == "2026-08-07"

temporal.ValidarData("32/13/2026") // inválido (calendário)
```

`ValidarData` usa `time.Parse` com layout `02/01/2006`, então rejeita datas impossíveis no calendário.

## Hora

```go
temporal.NormalizarHora("14:30") // "14:30:00"

resultado := temporal.ValidarHora("14:30")
// resultado.Valido == true
```

`ValidarHora` exige exatamente duas partes numéricas separadas por `:` (`HH:MM`).

## Data e hora

```go
temporal.NormalizarDataHora("07/08/2026 14:30") // "2026-08-07T14:30:00"
```

## Funções

| Função | Retorno |
|--------|---------|
| `NormalizarData(valor string)` | `string` |
| `ValidarData(valor string)` | `comum.Resultado` |
| `NormalizarDataSeValido(valor string)` | `comum.Resultado` |
| `NormalizarHora(valor string)` | `string` |
| `ValidarHora(valor string)` | `comum.Resultado` |
| `NormalizarDataHora(valor string)` | `string` |

## Referências

- [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html)
- [ABNT](https://www.abnt.org.br/)
