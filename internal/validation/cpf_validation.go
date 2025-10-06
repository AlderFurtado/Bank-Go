package validation

import (
	"strconv"
	"strings"
	"unicode"
)

// ValidaCPF recebe o CPF como string e retorna true se for válido
func ValidCpf(cpf string) bool {
	// Remove pontos e traços
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	if len(cpf) != 11 {
		return false
	}

	// Verifica se todos os dígitos são iguais
	allSame := true
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	// Converte para array de int
	nums := make([]int, 11)
	for i, r := range cpf {
		if !unicode.IsDigit(r) {
			return false
		}
		nums[i], _ = strconv.Atoi(string(r))
	}

	// Calcula o primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		sum += nums[i] * (10 - i)
	}
	firstDV := 11 - (sum % 11)
	if firstDV >= 10 {
		firstDV = 0
	}
	if nums[9] != firstDV {
		return false
	}

	// Calcula o segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		sum += nums[i] * (11 - i)
	}
	secondDV := 11 - (sum % 11)
	if secondDV >= 10 {
		secondDV = 0
	}
	if nums[10] != secondDV {
		return false
	}

	return true
}
