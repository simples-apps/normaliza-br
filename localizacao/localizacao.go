// Package localizacao oferece funções para normalizar e validar dados de
// localização brasileiros, como CEP e unidade federativa (UF).
package localizacao

import (
	"regexp"
	"strings"

	"github.com/simples-apps/normaliza-br/comum"
)

// NormalizarCEP remove caracteres de formatação de um CEP.
//
// O parâmetro valor aceita CEP com ou sem máscara (ex.: "13.080-300" ou "13080300").
// Retorna o valor sem espaços, pontos e hífens, sem validar o tamanho.
func NormalizarCEP(valor string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(valor), "-", ""), ".", "")
}

// ValidarCEP verifica se um CEP possui exatamente 8 caracteres após a normalização.
//
// O parâmetro valor aceita CEP com ou sem máscara. Em caso de sucesso, Resultado.Valor
// contém os 8 caracteres normalizados; em caso de falha, Resultado.Erro é preenchido.
func ValidarCEP(valor string) comum.Resultado {
	valorLimpo := NormalizarCEP(valor)
	if len(valorLimpo) != 8 {
		return comum.Resultado{Valido: false, Valor: "", Erro: "CEP inválido"}
	}
	return comum.Resultado{Valido: true, Valor: valorLimpo, Erro: ""}
}

// NormalizarCEPSeValido normaliza e devolve o CEP somente quando ele é válido.
//
// O parâmetro valor aceita CEP com ou sem máscara. Em caso de falha, devolve o
// mesmo Resultado de ValidarCEP.
func NormalizarCEPSeValido(valor string) comum.Resultado {
	resultado := ValidarCEP(valor)
	if !resultado.Valido {
		return resultado
	}
	return comum.Resultado{Valido: true, Valor: resultado.Valor, Erro: ""}
}

// NormalizarEstado padroniza o nome ou a sigla de um estado para a UF brasileira.
//
// O parâmetro valor aceita sigla (ex.: "sp") ou, para São Paulo, o nome por extenso
// ("São Paulo" / "Sao Paulo"). Retorna a sigla em maiúsculas quando reconhecida;
// caso contrário, devolve o texto já com trim e maiúsculas.
func NormalizarEstado(valor string) string {
	estado := strings.ToUpper(strings.TrimSpace(valor))
	estados := map[string]string{
		"AC": "AC",
		"AL": "AL",
		"AM": "AM",
		"AP": "AP",
		"BA": "BA",
		"CE": "CE",
		"DF": "DF",
		"ES": "ES",
		"GO": "GO",
		"MA": "MA",
		"MG": "MG",
		"MS": "MS",
		"MT": "MT",
		"PA": "PA",
		"PB": "PB",
		"PE": "PE",
		"PI": "PI",
		"PR": "PR",
		"RJ": "RJ",
		"RN": "RN",
		"RO": "RO",
		"RR": "RR",
		"RS": "RS",
		"SC": "SC",
		"SE": "SE",
		"SP": "SP",
		"TO": "TO",
	}
	if codigo, ok := estados[estado]; ok {
		return codigo
	}
	if len(estado) == 2 {
		return estado
	}
	if strings.Contains(estado, "SAO PAULO") || strings.Contains(estado, "SÃO PAULO") {
		return "SP"
	}
	return estado
}

// ValidarEstado verifica se o valor informado é uma UF com exatamente 2 letras A-Z.
//
// O parâmetro valor deve ser a sigla do estado (ex.: "SP" ou "sp"). Nomes por extenso
// são rejeitados. Em caso de sucesso, Resultado.Valor contém a UF em maiúsculas.
func ValidarEstado(valor string) comum.Resultado {
	estado := strings.ToUpper(strings.TrimSpace(valor))
	if len(estado) != 2 {
		return comum.Resultado{Valido: false, Valor: "", Erro: "Estado inválido"}
	}
	if regexp.MustCompile(`^[A-Z]{2}$`).MatchString(estado) {
		return comum.Resultado{Valido: true, Valor: estado, Erro: ""}
	}
	return comum.Resultado{Valido: false, Valor: "", Erro: "Estado inválido"}
}
