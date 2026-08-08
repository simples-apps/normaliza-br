// Package financeiro oferece funções para normalizar e validar valores
// monetários no formato brasileiro.
package financeiro

import (
	"regexp"
	"strings"
)

// Resultado representa o resultado padronizado de uma validação ou normalização.
type Resultado struct {
	// Valido indica se a entrada passou na validação.
	Valido bool
	// Valor contém o dado normalizado quando Valido é true; caso contrário, fica vazio.
	Valor string
	// Erro descreve o motivo da rejeição quando Valido é false; caso contrário, fica vazio.
	Erro string
}

// ValidarMoeda verifica se um valor monetário possui estrutura básica aceitável.
//
// O parâmetro valor aceita números no padrão brasileiro, opcionalmente com "R$",
// pontos de milhar, vírgula decimal e sinal negativo (ex.: "R$ 1.234,56").
// Em caso de sucesso, Resultado.Valor preserva a entrada já com trim; em caso de
// falha, Resultado.Erro é preenchido.
func ValidarMoeda(valor string) Resultado {
	valorTrim := strings.TrimSpace(valor)
	if valorTrim == "" {
		return Resultado{Valido: false, Valor: "", Erro: "Valor monetário inválido"}
	}
	if regexp.MustCompile(`^[0-9.,R$\s-]+$`).MatchString(valorTrim) {
		return Resultado{Valido: true, Valor: valorTrim, Erro: ""}
	}
	return Resultado{Valido: false, Valor: "", Erro: "Valor monetário inválido"}
}

// NormalizarMoeda converte um valor monetário brasileiro para string numérica com ponto decimal.
//
// O parâmetro valor aceita entradas como "R$ 1.234,56". Retorna a forma compacta
// (ex.: "1234.56"), removendo símbolo de moeda e separadores de milhar. Não valida
// a entrada antes de converter.
func NormalizarMoeda(valor string) string {
	valorTrim := strings.TrimSpace(valor)
	valorTrim = strings.ReplaceAll(valorTrim, ".", "")
	valorTrim = strings.ReplaceAll(valorTrim, ",", ".")
	valorTrim = strings.ReplaceAll(valorTrim, "R$", "")
	valorTrim = strings.TrimSpace(valorTrim)
	re := regexp.MustCompile(`[^0-9.-]`)
	return re.ReplaceAllString(valorTrim, "")
}

// NormalizarMoedaSeValido normaliza e devolve o valor monetário somente quando ele é válido.
//
// O parâmetro valor é o texto monetário a processar. Em caso de sucesso, Resultado.Valor
// contém a forma numérica com ponto decimal; em caso de falha, devolve o Resultado
// de ValidarMoeda.
func NormalizarMoedaSeValido(valor string) Resultado {
	resultado := ValidarMoeda(valor)
	if !resultado.Valido {
		return resultado
	}
	return Resultado{Valido: true, Valor: NormalizarMoeda(valor), Erro: ""}
}
