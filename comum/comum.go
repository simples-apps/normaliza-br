// Package comum reúne tipos compartilhados pelos demais pacotes da biblioteca.
package comum

// Resultado representa o resultado padronizado de uma validação ou normalização.
type Resultado struct {
	// Valido indica se a entrada passou na validação.
	Valido bool
	// Valor contém o dado normalizado quando Valido é true; caso contrário, fica vazio.
	Valor string
	// Erro descreve o motivo da rejeição quando Valido é false; caso contrário, fica vazio.
	Erro string
}
