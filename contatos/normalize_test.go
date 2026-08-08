package contatos

import "testing"

func TestNormalizarTelefone(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "celular com mascara", valor: "(11) 99999-8888", want: "11999998888"},
		{nome: "com DDI", valor: "+55 11 99999-8888", want: "5511999998888"},
		{nome: "somente digitos", valor: "1133334444", want: "1133334444"},
		{nome: "com texto", valor: "tel: (11) 3333-4444", want: "1133334444"},
		{nome: "vazio", valor: "", want: ""},
		{nome: "somente simbolos", valor: "()-", want: ""},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarTelefone(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarTelefone(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestNormalizarEmail(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "maiusculas e espacos", valor: " Usuario@Exemplo.COM ", want: "usuario@exemplo.com"},
		{nome: "ja normalizado", valor: "a@b.co", want: "a@b.co"},
		{nome: "vazio", valor: "   ", want: ""},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarEmail(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarEmail(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestValidarEmail(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		ok    bool
		want  string
	}{
		{nome: "valido", valor: " Usuario@Exemplo.COM ", ok: true, want: "usuario@exemplo.com"},
		{nome: "valido simples", valor: "a@b.co", ok: true, want: "a@b.co"},
		{nome: "vazio", valor: "", ok: false},
		{nome: "sem arroba", valor: "emailinvalido", ok: false},
		{nome: "sem ponto", valor: "usuario@exemplo", ok: false},
		{nome: "somente arroba", valor: "@", ok: false},
		{nome: "somente ponto", valor: ".", ok: false},
		{nome: "espacos", valor: "   ", ok: false},
		{nome: "injection sem estrutura", valor: "admin'--", ok: false},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := ValidarEmail(caso.valor)
			if resultado.Valido != caso.ok {
				t.Fatalf("ValidarEmail(%q) Valido = %v, want %v", caso.valor, resultado.Valido, caso.ok)
			}
			if caso.ok && resultado.Valor != caso.want {
				t.Fatalf("ValidarEmail(%q) Valor = %q, want %q", caso.valor, resultado.Valor, caso.want)
			}
			if !caso.ok && resultado.Erro == "" {
				t.Fatalf("ValidarEmail(%q) deve retornar erro", caso.valor)
			}
		})
	}
}

func TestNormalizarEmailSeValido(t *testing.T) {
	t.Run("aceita valido", func(t *testing.T) {
		resultado := NormalizarEmailSeValido(" Usuario@Exemplo.COM ")
		if !resultado.Valido || resultado.Valor != "usuario@exemplo.com" {
			t.Fatalf("NormalizarEmailSeValido() = %+v", resultado)
		}
	})

	t.Run("rejeita invalido", func(t *testing.T) {
		resultado := NormalizarEmailSeValido("sem-arroba")
		if resultado.Valido {
			t.Fatalf("NormalizarEmailSeValido() deve rejeitar e-mail inválido")
		}
	})
}
