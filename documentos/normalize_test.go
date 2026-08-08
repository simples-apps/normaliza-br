package documentos

import "testing"

func TestNormalizarCPF(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "com mascara", valor: "123.456.789-09", want: "12345678909"},
		{nome: "sem mascara", valor: "52998224725", want: "52998224725"},
		{nome: "com espacos", valor: " 123 456 789 09 ", want: "12345678909"},
		{nome: "com letras misturadas", valor: "123.ABC.789-09", want: "12378909"},
		{nome: "vazio", valor: "", want: ""},
		{nome: "somente letras", valor: "abcdef", want: ""},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarCPF(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarCPF(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestNormalizarCNPJ(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		want  string
	}{
		{nome: "alfanumerico oficial", valor: "12.ABC.345/01DE-35", want: "12ABC34501DE35"},
		{nome: "numerico classico", valor: "12.345.678/0001-95", want: "12345678000195"},
		{nome: "minusculas", valor: "12.abc.345/01de-35", want: "12ABC34501DE35"},
		{nome: "remove caracteres invalidos", valor: "12@ABC#345$01DE%35", want: "12ABC34501DE35"},
		{nome: "vazio", valor: "   ", want: ""},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			got := NormalizarCNPJ(caso.valor)
			if got != caso.want {
				t.Fatalf("NormalizarCNPJ(%q) = %q, want %q", caso.valor, got, caso.want)
			}
		})
	}
}

func TestValidarCPF(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		ok    bool
		want  string
	}{
		{nome: "valido formatado", valor: "529.982.247-25", ok: true, want: "52998224725"},
		{nome: "valido sem mascara", valor: "39053344705", ok: true, want: "39053344705"},
		{nome: "valido com espacos", valor: " 111.444.777-35 ", ok: true, want: "11144477735"},

		{nome: "vazio", valor: "", ok: false},
		{nome: "somente espacos", valor: "   ", ok: false},
		{nome: "curto", valor: "123", ok: false},
		{nome: "longo", valor: "123456789091", ok: false},
		{nome: "10 digitos", valor: "1234567890", ok: false},
		{nome: "somente letras", valor: "abcdefghijk", ok: false},
		{nome: "mascara incompleta", valor: "123.456.789", ok: false},

		// Tentativas comuns de fraude / placeholders
		{nome: "sequencial ascendente com dv fraudado", valor: "123.456.789-00", ok: false},
		{nome: "sequencial ascendente com dv invertido", valor: "123.456.789-90", ok: false},
		{nome: "sequencial descendente com dv fraudado", valor: "987.654.321-11", ok: false},
		{nome: "sequencial com digito a mais no meio", valor: "123.456.78901-9", ok: false},
		{nome: "zeros a esquerda incompletos", valor: "000.000.001", ok: false},

		{nome: "dv primeiro incorreto", valor: "529.982.247-15", ok: false},
		{nome: "dv segundo incorreto", valor: "529.982.247-20", ok: false},
		{nome: "transposicao dos dvs", valor: "529.982.247-52", ok: false},

		{nome: "sql injection textual", valor: "12345678900' OR '1'='1", ok: false},
		{nome: "html/script com dv fraudado", valor: "<script>12345678900</script>", ok: false},
		{nome: "null byte e lixo", valor: "11111111111\x00extra", ok: false},
		{nome: "cpf invalido escondido em texto", valor: "documento:12345678900-filial", ok: false},
		{nome: "cpf valido embutido em texto e limpo", valor: "documento:52998224725", ok: true, want: "52998224725"},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := ValidarCPF(caso.valor)
			if resultado.Valido != caso.ok {
				t.Fatalf("ValidarCPF(%q) Valido = %v, want %v (erro=%q valor=%q)", caso.valor, resultado.Valido, caso.ok, resultado.Erro, resultado.Valor)
			}
			if caso.ok && resultado.Valor != caso.want {
				t.Fatalf("ValidarCPF(%q) Valor = %q, want %q", caso.valor, resultado.Valor, caso.want)
			}
			if !caso.ok {
				if resultado.Erro == "" {
					t.Fatalf("ValidarCPF(%q) deve retornar mensagem de erro", caso.valor)
				}
				if resultado.Valor != "" {
					t.Fatalf("ValidarCPF(%q) Valor deve ficar vazio quando inválido, got %q", caso.valor, resultado.Valor)
				}
			}
		})
	}
}

func TestValidarCPFDigitosRepetidos(t *testing.T) {
	// Placeholders clássicos usados para burlar formulários.
	repetidos := []string{
		"000.000.000-00",
		"111.111.111-11",
		"222.222.222-22",
		"333.333.333-33",
		"444.444.444-44",
		"555.555.555-55",
		"666.666.666-66",
		"777.777.777-77",
		"888.888.888-88",
		"999.999.999-99",
		"00000000000",
		"11111111111",
		"99999999999",
	}

	for _, valor := range repetidos {
		t.Run(valor, func(t *testing.T) {
			resultado := ValidarCPF(valor)
			if resultado.Valido {
				t.Fatalf("ValidarCPF(%q) deveria rejeitar dígitos repetidos", valor)
			}
			if resultado.Erro == "" {
				t.Fatalf("ValidarCPF(%q) deve retornar mensagem de erro", valor)
			}
		})
	}
}

func TestValidarCNPJ(t *testing.T) {
	casos := []struct {
		nome  string
		valor string
		ok    bool
		want  string
	}{
		{
			nome:  "alfanumerico oficial receita",
			valor: "12.ABC.345/01DE-35",
			ok:    true,
			want:  "12ABC34501DE35",
		},
		{
			nome:  "numerico classico",
			valor: "12.345.678/0001-95",
			ok:    true,
			want:  "12345678000195",
		},
		{
			nome:  "numerico conhecido",
			valor: "11.222.333/0001-81",
			ok:    true,
			want:  "11222333000181",
		},
		{
			nome:  "minusculas",
			valor: "12.abc.345/01de-35",
			ok:    true,
			want:  "12ABC34501DE35",
		},
		{
			nome:  "com espacos extras",
			valor: " 12.345.678/0001-95 ",
			ok:    true,
			want:  "12345678000195",
		},

		{nome: "vazio", valor: "", ok: false},
		{nome: "somente espacos", valor: "   ", ok: false},
		{nome: "curto", valor: "123", ok: false},
		{nome: "13 caracteres", valor: "1234567800019", ok: false},
		{nome: "15 caracteres", valor: "123456780001950", ok: false},
		{nome: "tamanho diferente com letras", valor: "12A345B678C0001D99", ok: false},

		{nome: "dv incorreto alfanumerico", valor: "12.ABC.345/01DE-00", ok: false},
		{nome: "dv incorreto numerico", valor: "12.345.678/0001-00", ok: false},
		{nome: "dv invertidos", valor: "12.345.678/0001-59", ok: false},
		{nome: "dv nao numerico", valor: "12ABC34501DEAB", ok: false},
		{nome: "letra no segundo dv", valor: "12ABC34501DE3A", ok: false},

		{nome: "matriz com ordem inventada", valor: "12.345.678/9999-95", ok: false},
		{nome: "sequencial numerico fraudado", valor: "12.345.678/0001-11", ok: false},
		{nome: "somente zeros com dv errado", valor: "00.000.000/0000-01", ok: false},
		{nome: "digitos repetidos", valor: "11.111.111/1111-11", ok: false},

		{nome: "caracteres especiais removiveis", valor: "12.345.678/0001-95!", ok: true, want: "12345678000195"},
		{nome: "sql injection textual", valor: "12345678000195; DROP TABLE", ok: false},
		{nome: "html/script com dv fraudado", valor: "<b>12345678000100</b>", ok: false},
		{nome: "cnpj valido escondido em texto longo", valor: "empresa12ABC34501DE35matriz", ok: false},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := ValidarCNPJ(caso.valor)
			if resultado.Valido != caso.ok {
				t.Fatalf("ValidarCNPJ(%q) Valido = %v, want %v (erro=%q valor=%q)", caso.valor, resultado.Valido, caso.ok, resultado.Erro, resultado.Valor)
			}
			if caso.ok && resultado.Valor != caso.want {
				t.Fatalf("ValidarCNPJ(%q) Valor = %q, want %q", caso.valor, resultado.Valor, caso.want)
			}
			if !caso.ok {
				if resultado.Erro == "" {
					t.Fatalf("ValidarCNPJ(%q) deve retornar mensagem de erro", caso.valor)
				}
				if resultado.Valor != "" {
					t.Fatalf("ValidarCNPJ(%q) Valor deve ficar vazio quando inválido, got %q", caso.valor, resultado.Valor)
				}
			}
		})
	}
}

func TestNormalizarCPFSeValido(t *testing.T) {
	t.Run("aceita valido", func(t *testing.T) {
		resultado := NormalizarCPFSeValido("529.982.247-25")
		if !resultado.Valido {
			t.Fatalf("NormalizarCPFSeValido() deve aceitar um CPF válido")
		}
		if resultado.Valor != "52998224725" {
			t.Fatalf("NormalizarCPFSeValido() = %q, want %q", resultado.Valor, "52998224725")
		}
	})

	t.Run("rejeita repetido", func(t *testing.T) {
		resultado := NormalizarCPFSeValido("111.111.111-11")
		if resultado.Valido {
			t.Fatalf("NormalizarCPFSeValido() deve rejeitar CPF repetido")
		}
	})

	t.Run("rejeita sequencial fraudado", func(t *testing.T) {
		resultado := NormalizarCPFSeValido("123.456.789-00")
		if resultado.Valido {
			t.Fatalf("NormalizarCPFSeValido() deve rejeitar sequencial com DV inválido")
		}
	})
}

func TestNormalizarCNPJSeValido(t *testing.T) {
	t.Run("aceita alfanumerico oficial", func(t *testing.T) {
		resultado := NormalizarCNPJSeValido("12.ABC.345/01DE-35")
		if !resultado.Valido {
			t.Fatalf("NormalizarCNPJSeValido() deve aceitar o CNPJ alfanumérico oficial")
		}
		if resultado.Valor != "12ABC34501DE35" {
			t.Fatalf("NormalizarCNPJSeValido() = %q, want %q", resultado.Valor, "12ABC34501DE35")
		}
	})

	t.Run("rejeita dv fraudado", func(t *testing.T) {
		resultado := NormalizarCNPJSeValido("12.345.678/0001-00")
		if resultado.Valido {
			t.Fatalf("NormalizarCNPJSeValido() deve rejeitar DV inválido")
		}
	})
}

func TestNormalizarDocumento(t *testing.T) {
	got := NormalizarDocumento("  ABC 123  ")
	want := "ABC 123"
	if got != want {
		t.Fatalf("NormalizarDocumento() = %q, want %q", got, want)
	}
}
