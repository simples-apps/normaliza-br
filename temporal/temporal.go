// Package temporal oferece funções para normalizar e validar datas e horas
// no padrão brasileiro, com saída em formatos ISO comuns.
package temporal

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/simples-apps/normaliza-br/comum"
)

// ValidarData verifica se uma data está no formato DD/MM/AAAA e é válida no calendário.
//
// O parâmetro valor deve estar em "02/01/2006" (ex.: "07/08/2026"). Em caso de sucesso,
// Resultado.Valor preserva a entrada com trim; em caso de falha, Resultado.Erro é preenchido.
func ValidarData(valor string) comum.Resultado {
	valorTrim := strings.TrimSpace(valor)
	if _, err := time.Parse("02/01/2006", valorTrim); err != nil {
		return comum.Resultado{Valido: false, Valor: "", Erro: "Data inválida"}
	}
	return comum.Resultado{Valido: true, Valor: valorTrim, Erro: ""}
}

// NormalizarData converte uma data no formato brasileiro (DD/MM/AAAA) para ISO 8601 (AAAA-MM-DD).
//
// O parâmetro valor deve conter três partes separadas por "/" (ex.: "07/08/2026").
// Se o formato não for reconhecido, retorna valor apenas com trim.
func NormalizarData(valor string) string {
	valorTrim := strings.TrimSpace(valor)
	partes := strings.Split(valorTrim, "/")
	if len(partes) == 3 {
		return partes[2] + "-" + partes[1] + "-" + partes[0]
	}
	return valorTrim
}

// NormalizarHora padroniza uma hora para o formato HH:MM:SS.
//
// O parâmetro valor aceita hora com pelo menos horas e minutos (ex.: "14:30").
// Quando há ao menos duas partes separadas por ":", os segundos são fixados em "00".
// Caso contrário, retorna valor apenas com trim.
func NormalizarHora(valor string) string {
	valorTrim := strings.TrimSpace(valor)
	partes := strings.Split(valorTrim, ":")
	if len(partes) >= 2 {
		return partes[0] + ":" + partes[1] + ":00"
	}
	return valorTrim
}

// NormalizarDataSeValido normaliza a data para ISO 8601 somente quando ela é válida.
//
// O parâmetro valor deve estar em DD/MM/AAAA. Em caso de sucesso, Resultado.Valor
// fica em AAAA-MM-DD; em caso de falha, devolve o Resultado de ValidarData.
func NormalizarDataSeValido(valor string) comum.Resultado {
	resultado := ValidarData(valor)
	if !resultado.Valido {
		return resultado
	}
	return comum.Resultado{Valido: true, Valor: NormalizarData(valor), Erro: ""}
}

// NormalizarDataHora junta data e hora em um único texto no formato AAAA-MM-DDTHH:MM:00.
//
// O parâmetro valor deve conter data e hora separados por espaço
// (ex.: "07/08/2026 14:30"). Se não houver duas partes, retorna valor apenas com trim.
func NormalizarDataHora(valor string) string {
	valorTrim := strings.TrimSpace(valor)
	re := regexp.MustCompile(`\s+`)
	partes := re.Split(valorTrim, 2)
	if len(partes) == 2 {
		return NormalizarData(partes[0]) + "T" + NormalizarHora(partes[1])
	}
	return valorTrim
}

// ValidarHora verifica se uma hora está no formato básico HH:MM com partes numéricas.
//
// O parâmetro valor deve ter exatamente duas partes separadas por ":" (ex.: "14:30").
// Em caso de sucesso, Resultado.Valor preserva a entrada com trim; em caso de falha,
// Resultado.Erro é preenchido. Não valida faixas de 0–23 / 0–59.
func ValidarHora(valor string) comum.Resultado {
	valorTrim := strings.TrimSpace(valor)
	partes := strings.Split(valorTrim, ":")
	if len(partes) != 2 {
		return comum.Resultado{Valido: false, Valor: "", Erro: "Hora inválida"}
	}
	if _, err := strconv.Atoi(partes[0]); err != nil {
		return comum.Resultado{Valido: false, Valor: "", Erro: "Hora inválida"}
	}
	if _, err := strconv.Atoi(partes[1]); err != nil {
		return comum.Resultado{Valido: false, Valor: "", Erro: "Hora inválida"}
	}
	return comum.Resultado{Valido: true, Valor: valorTrim, Erro: ""}
}
