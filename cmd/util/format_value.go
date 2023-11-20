package util

import (
	"fmt"
	"strings"
)

func FormatarValor(valor float64) string {

	valorStr := fmt.Sprintf("%.2f", valor)
	valorStr = strings.Replace(valorStr, ".", ",", -1)

	n := len(valorStr)
	for i := n - 6; i > 0; i -= 3 {
		valorStr = valorStr[:i] + "." + valorStr[i:]
	}

	return "R$ " + valorStr
}
