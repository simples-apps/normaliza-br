// Package documentos oferece funções para normalizar e validar documentos
// oficiais brasileiros, como CPF e CNPJ.
package documentos

import (
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

// NormalizarCPF remove caracteres não numéricos de um CPF.
//
// O parâmetro valor aceita CPF com ou sem máscara (ex.: "529.982.247-25" ou "52998224725").
// Retorna apenas os dígitos encontrados, sem validar dígitos verificadores.
func NormalizarCPF(valor string) string {
	return normalizarDigitos(valor)
}

// NormalizarCNPJ remove separadores e caracteres inválidos de um CNPJ,
// preservando o formato alfanumérico oficial (A-Z e 0-9) em maiúsculas.
//
// O parâmetro valor aceita CNPJ numérico ou alfanumérico, com ou sem máscara
// (ex.: "12.ABC.345/01DE-35"). Retorna a forma compacta sem validar os DVs.
func NormalizarCNPJ(valor string) string {
	valor = strings.ToUpper(strings.TrimSpace(valor))
	var b strings.Builder
	b.Grow(len(valor))
	for _, r := range valor {
		if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// normalizarDigitos remove todos os caracteres não numéricos de valor.
func normalizarDigitos(valor string) string {
	var b strings.Builder
	b.Grow(len(valor))
	for _, r := range valor {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// NormalizarDocumento remove espaços em branco no início e no fim de um documento genérico.
//
// O parâmetro valor é qualquer texto de documento. Retorna valor após TrimSpace.
func NormalizarDocumento(valor string) string {
	return strings.TrimSpace(valor)
}

// ValidarCPF valida um CPF conforme o padrão oficial da Receita Federal:
// 11 dígitos, rejeição de sequências repetidas e dígitos verificadores (módulo 11).
//
// O parâmetro valor aceita CPF com ou sem máscara. Em caso de sucesso, Resultado.Valor
// contém os 11 dígitos normalizados; em caso de falha, Resultado.Erro é preenchido.
func ValidarCPF(valor string) Resultado {
	valorLimpo := normalizarDigitos(valor)
	if len(valorLimpo) != 11 || digitosIguais(valorLimpo) || !cpfDigitosValidos(valorLimpo) {
		return Resultado{Valido: false, Valor: "", Erro: "CPF inválido"}
	}
	return Resultado{Valido: true, Valor: valorLimpo, Erro: ""}
}

// NormalizarCPFSeValido normaliza e devolve o CPF somente quando ele é válido.
//
// O parâmetro valor aceita CPF com ou sem máscara. Equivale a ValidarCPF(valor).
func NormalizarCPFSeValido(valor string) Resultado {
	return ValidarCPF(valor)
}

// ValidarCNPJ valida um CNPJ conforme o padrão oficial da Receita Federal
// (numérico ou alfanumérico, vigente a partir de julho/2026):
// 14 posições no formato [A-Z0-9]{12}[0-9]{2} e dígitos verificadores (módulo 11
// com conversão ASCII-48).
//
// O parâmetro valor aceita CNPJ com ou sem máscara. Em caso de sucesso, Resultado.Valor
// contém as 14 posições normalizadas em maiúsculas; em caso de falha, Resultado.Erro
// é preenchido.
func ValidarCNPJ(valor string) Resultado {
	valorLimpo := NormalizarCNPJ(valor)
	if len(valorLimpo) != 14 || !cnpjFormatoValido(valorLimpo) || !cnpjDigitosValidos(valorLimpo) {
		return Resultado{Valido: false, Valor: "", Erro: "CNPJ inválido"}
	}
	return Resultado{Valido: true, Valor: valorLimpo, Erro: ""}
}

// NormalizarCNPJSeValido normaliza e devolve o CNPJ somente quando ele é válido.
//
// O parâmetro valor aceita CNPJ numérico ou alfanumérico, com ou sem máscara.
// Equivale a ValidarCNPJ(valor).
func NormalizarCNPJSeValido(valor string) Resultado {
	return ValidarCNPJ(valor)
}

// digitosIguais reporta se todos os caracteres de valor são idênticos.
func digitosIguais(valor string) bool {
	for i := 1; i < len(valor); i++ {
		if valor[i] != valor[0] {
			return false
		}
	}
	return true
}

// cpfDigitosValidos confere os dois dígitos verificadores de um CPF já normalizado
// com exatamente 11 dígitos (parâmetro cpf).
func cpfDigitosValidos(cpf string) bool {
	nums := make([]int, 11)
	for i := 0; i < 11; i++ {
		nums[i] = int(cpf[i] - '0')
	}

	soma := 0
	for i := 0; i < 9; i++ {
		soma += nums[i] * (10 - i)
	}
	resto := soma % 11
	dv1 := 0
	if resto >= 2 {
		dv1 = 11 - resto
	}
	if nums[9] != dv1 {
		return false
	}

	soma = 0
	for i := 0; i < 10; i++ {
		soma += nums[i] * (11 - i)
	}
	resto = soma % 11
	dv2 := 0
	if resto >= 2 {
		dv2 = 11 - resto
	}
	return nums[10] == dv2
}

// cnpjFormatoValido verifica se cnpj normalizado tem 14 posições no formato
// [A-Z0-9]{12}[0-9]{2}.
func cnpjFormatoValido(cnpj string) bool {
	for i := 0; i < 12; i++ {
		c := cnpj[i]
		if !((c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z')) {
			return false
		}
	}
	return cnpj[12] >= '0' && cnpj[12] <= '9' && cnpj[13] >= '0' && cnpj[13] <= '9'
}

// cnpjDigitosValidos confere os dois dígitos verificadores de um CNPJ já normalizado
// com 14 posições (parâmetro cnpj).
func cnpjDigitosValidos(cnpj string) bool {
	dv1 := calcularDVCNPJ(cnpj[:12])
	if cnpj[12] != byte('0'+dv1) {
		return false
	}
	dv2 := calcularDVCNPJ(cnpj[:12] + string(rune('0'+dv1)))
	return cnpj[13] == byte('0'+dv2)
}

// calcularDVCNPJ calcula um dígito verificador pelo módulo 11 da Receita Federal,
// convertendo cada caractere com valor ASCII - 48.
//
// O parâmetro base deve conter 12 caracteres (1º DV) ou 13 caracteres (2º DV),
// já normalizados. Retorna um inteiro de 0 a 9.
func calcularDVCNPJ(base string) int {
	soma := 0
	peso := 2
	for i := len(base) - 1; i >= 0; i-- {
		soma += (int(base[i]) - 48) * peso
		peso++
		if peso > 9 {
			peso = 2
		}
	}
	resto := soma % 11
	if resto <= 1 {
		return 0
	}
	return 11 - resto
}
