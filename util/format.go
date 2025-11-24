package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func FormatTicket(numbers []int) string {
	if len(numbers) == 0 {
		return "[]"
	}

	var b strings.Builder
	b.WriteString("[")

	for i, n := range numbers {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(strconv.Itoa(n))
	}

	b.WriteString("]")
	return b.String()
}

func FormatRate(value float64) string {
	rounded := math.Round(value*10) / 10
	// 불필요한 소수 .0은 제거
	s := strconv.FormatFloat(rounded, 'f', -1, 64)
	return fmt.Sprintf("%s%%", s)
}