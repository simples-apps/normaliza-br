package temporal

import "testing"

func TestNormalizarData(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "brasileira", valor: "07/08/2026", want: "2026-08-07"},
		{nome: "com espacos", valor: " 01/01/2000 ", want: "2000-01-01"},
		{nome: "sem barras", valor: "07082026", want: "07082026"},
		{nome: "vazio", valor: "", want: ""},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarData(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarData(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestValidarData(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		ok    bool
	}{
		{nome: "valida", valor: "07/08/2026", ok: true},
		{nome: "ano bissexto", valor: "29/02/2024", ok: true},
		{nome: "vazio", valor: "", ok: false},
		{nome: "dia impossivel", valor: "32/01/2026", ok: false},
		{nome: "mes impossivel", valor: "15/13/2026", ok: false},
		{nome: "dia e mes impossiveis", valor: "32/13/2026", ok: false},
		{nome: "29 fev ano nao bissexto", valor: "29/02/2025", ok: false},
		{nome: "formato iso", valor: "2026-08-07", ok: false},
		{nome: "texto", valor: "ontem", ok: false},
		{nome: "zeros", valor: "00/00/0000", ok: false},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := ValidarData(caso.valor)
			if resultado.Valido != caso.ok {
				t.Fatalf("ValidarData(%q) Valido = %v, want %v", caso.valor, resultado.Valido, caso.ok)
			}
			if !caso.ok && resultado.Erro == "" {
				t.Fatalf("ValidarData(%q) deve retornar erro", caso.valor)
			}
		})
	}
}

func TestNormalizarHora(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "hhmm", valor: "14:30", want: "14:30:00"},
		{nome: "com espacos", valor: " 09:05 ", want: "09:05:00"},
		{nome: "incompleta", valor: "14", want: "14"},
		{nome: "vazio", valor: "", want: ""},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarHora(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarHora(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestValidarHora(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		ok    bool
	}{
		{nome: "valida", valor: "14:30", ok: true},
		{nome: "meia noite", valor: "00:00", ok: true},
		{nome: "vazio", valor: "", ok: false},
		{nome: "sem minutos", valor: "14", ok: false},
		{nome: "com segundos", valor: "14:30:00", ok: false},
		{nome: "letras", valor: "ab:cd", ok: false},
		{nome: "texto", valor: "meio-dia", ok: false},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := ValidarHora(caso.valor)
			if resultado.Valido != caso.ok {
				t.Fatalf("ValidarHora(%q) Valido = %v, want %v", caso.valor, resultado.Valido, caso.ok)
			}
			if !caso.ok && resultado.Erro == "" {
				t.Fatalf("ValidarHora(%q) deve retornar erro", caso.valor)
			}
		})
	}
}

func TestNormalizarDataHora(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "data e hora", valor: "07/08/2026 14:30", want: "2026-08-07T14:30:00"},
		{nome: "espacos multiplos", valor: "07/08/2026   14:30", want: "2026-08-07T14:30:00"},
		{nome: "somente data", valor: "07/08/2026", want: "07/08/2026"},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarDataHora(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarDataHora(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestNormalizarDataSeValido(t *testing.T) {
	t.Run("aceita valida", func(t *testing.T) {
		resultado := NormalizarDataSeValido("07/08/2026")
		if !resultado.Valido || resultado.Valor != "2026-08-07" {
			t.Fatalf("NormalizarDataSeValido() = %+v", resultado)
		}
	})

	t.Run("rejeita invalida", func(t *testing.T) {
		resultado := NormalizarDataSeValido("32/13/2026")
		if resultado.Valido {
			t.Fatalf("NormalizarDataSeValido() deve rejeitar data inválida")
		}
	})
}
