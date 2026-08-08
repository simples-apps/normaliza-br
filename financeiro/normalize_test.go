package financeiro

import "testing"

func TestNormalizarMoeda(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "real formatado", valor: "R$ 1.234,56", want: "1234.56"},
		{nome: "negativo", valor: "-1.000,00", want: "-1000.00"},
		{nome: "sem simbolo", valor: "99,9", want: "99.9"},
		{nome: "somente inteiro", valor: "1500", want: "1500"},
		{nome: "com espacos", valor: "  R$ 10,50  ", want: "10.50"},
		{nome: "lixo textual", valor: "R$ abc", want: ""},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarMoeda(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarMoeda(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestValidarMoeda(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		ok    bool
	}{
		{nome: "valido com real", valor: "R$ 1.234,56", ok: true},
		{nome: "valido negativo", valor: "-10,00", ok: true},
		{nome: "valido simples", valor: "0,01", ok: true},
		{nome: "vazio", valor: "", ok: false},
		{nome: "somente espacos", valor: "   ", ok: false},
		{nome: "letras", valor: "abc", ok: false},
		{nome: "sql injection", valor: "1; DROP TABLE", ok: false},
		{nome: "moeda estrangeira textual", valor: "USD 10", ok: false},
		{nome: "emoji", valor: "💰10", ok: false},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := ValidarMoeda(caso.valor)
			if resultado.Valido != caso.ok {
				t.Fatalf("ValidarMoeda(%q) Valido = %v, want %v", caso.valor, resultado.Valido, caso.ok)
			}
			if !caso.ok && resultado.Erro == "" {
				t.Fatalf("ValidarMoeda(%q) deve retornar erro", caso.valor)
			}
		})
	}
}

func TestNormalizarMoedaSeValido(t *testing.T) {
	t.Run("aceita valido", func(t *testing.T) {
		resultado := NormalizarMoedaSeValido("R$ 1.234,56")
		if !resultado.Valido || resultado.Valor != "1234.56" {
			t.Fatalf("NormalizarMoedaSeValido() = %+v", resultado)
		}
	})

	t.Run("rejeita invalido", func(t *testing.T) {
		resultado := NormalizarMoedaSeValido("abc")
		if resultado.Valido {
			t.Fatalf("NormalizarMoedaSeValido() deve rejeitar valor inválido")
		}
	})
}
