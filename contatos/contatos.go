// Package contatos oferece funções para normalizar e validar dados de
// comunicação, como telefone e e-mail.
package contatos

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

// NormalizarTelefone remove caracteres de formatação de um telefone.
//
// O parâmetro valor aceita telefone com máscara, DDI ou texto auxiliar
// (ex.: "(11) 99999-8888" ou "+55 11 99999-8888"). Retorna somente os dígitos.
func NormalizarTelefone(valor string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(valor, "")
}

// NormalizarEmail remove espaços nas extremidades e converte o e-mail para minúsculas.
//
// O parâmetro valor é o endereço de e-mail informado pelo usuário
// (ex.: " Usuario@Exemplo.COM "). Retorna o e-mail padronizado, sem validar a estrutura.
func NormalizarEmail(valor string) string {
	return strings.ToLower(strings.TrimSpace(valor))
}

// ValidarEmail verifica se o e-mail possui estrutura básica (contém "@" e ".").
//
// O parâmetro valor é o endereço a validar. Em caso de sucesso, Resultado.Valor
// contém o e-mail em minúsculas; em caso de falha, Resultado.Erro é preenchido.
// A validação é deliberadamente simples e não cobre a RFC completa.
func ValidarEmail(valor string) Resultado {
	valorTrim := strings.TrimSpace(valor)
	if !strings.Contains(valorTrim, "@") || !strings.Contains(valorTrim, ".") {
		return Resultado{Valido: false, Valor: "", Erro: "E-mail inválido"}
	}
	return Resultado{Valido: true, Valor: strings.ToLower(valorTrim), Erro: ""}
}

// NormalizarEmailSeValido normaliza e devolve o e-mail somente quando ele é válido.
//
// O parâmetro valor é o endereço a normalizar. Em caso de falha, devolve o mesmo
// Resultado de ValidarEmail.
func NormalizarEmailSeValido(valor string) Resultado {
	resultado := ValidarEmail(valor)
	if !resultado.Valido {
		return resultado
	}
	return Resultado{Valido: true, Valor: resultado.Valor, Erro: ""}
}
