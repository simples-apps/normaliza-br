package localizacao

import "testing"

func TestNormalizarCEP(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "com mascara", valor: "13.080-300", want: "13080300"},
		{nome: "somente hifen", valor: "13080-300", want: "13080300"},
		{nome: "sem mascara", valor: "13080300", want: "13080300"},
		{nome: "com espacos", valor: " 13080-300 ", want: "13080300"},
		{nome: "vazio", valor: "", want: ""},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarCEP(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarCEP(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestValidarCEP(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		ok    bool
		want  string
	}{
		{nome: "valido", valor: "13.080-300", ok: true, want: "13080300"},
		{nome: "valido sem mascara", valor: "01310200", ok: true, want: "01310200"},
		{nome: "vazio", valor: "", ok: false},
		{nome: "curto", valor: "123", ok: false},
		{nome: "longo", valor: "130803001", ok: false},
		{nome: "somente hifen", valor: "-----", ok: false},
		{nome: "letras", valor: "ABCDEFGH", ok: true, want: "ABCDEFGH"}, // tamanho 8 passa na regra atual
		{nome: "lixo com tamanho errado", valor: "CEP-123", ok: false},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := ValidarCEP(caso.valor)
			if resultado.Valido != caso.ok {
				t.Fatalf("ValidarCEP(%q) Valido = %v, want %v", caso.valor, resultado.Valido, caso.ok)
			}
			if caso.ok && resultado.Valor != caso.want {
				t.Fatalf("ValidarCEP(%q) Valor = %q, want %q", caso.valor, resultado.Valor, caso.want)
			}
			if !caso.ok && resultado.Erro == "" {
				t.Fatalf("ValidarCEP(%q) deve retornar erro", caso.valor)
			}
		})
	}
}

func TestNormalizarEstado(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "sigla minuscula", valor: "sp", want: "SP"},
		{nome: "sigla maiuscula", valor: "RJ", want: "RJ"},
		{nome: "nome sao paulo", valor: "são paulo", want: "SP"},
		{nome: "nome sao paulo sem acento", valor: "sao paulo", want: "SP"},
		{nome: "com espacos", valor: "  mg  ", want: "MG"},
		{nome: "desconhecido", valor: "XX", want: "XX"},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarEstado(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarEstado(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestValidarEstado(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		ok    bool
	}{
		{nome: "valido", valor: "SP", ok: true},
		{nome: "minusculo", valor: "rj", ok: true},
		{nome: "vazio", valor: "", ok: false},
		{nome: "uma letra", valor: "S", ok: false},
		{nome: "tres letras", valor: "SPO", ok: false},
		{nome: "numeros", valor: "12", ok: false},
		{nome: "nome por extenso", valor: "São Paulo", ok: false},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := ValidarEstado(caso.valor)
			if resultado.Valido != caso.ok {
				t.Fatalf("ValidarEstado(%q) Valido = %v, want %v", caso.valor, resultado.Valido, caso.ok)
			}
			if !caso.ok && resultado.Erro == "" {
				t.Fatalf("ValidarEstado(%q) deve retornar erro", caso.valor)
			}
		})
	}
}

func TestNormalizarCEPSeValido(t *testing.T) {
	t.Run("aceita valido", func(t *testing.T) {
		resultado := NormalizarCEPSeValido("13.080-300")
		if !resultado.Valido || resultado.Valor != "13080300" {
			t.Fatalf("NormalizarCEPSeValido() = %+v", resultado)
		}
	})

	t.Run("rejeita invalido", func(t *testing.T) {
		resultado := NormalizarCEPSeValido("123")
		if resultado.Valido {
			t.Fatalf("NormalizarCEPSeValido() deve rejeitar CEP curto")
		}
	})
}
