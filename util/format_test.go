package util

import "testing"

func TestCommaInt64(t *testing.T) {
	if got := CommaInt64(999); got != "999" {
		t.Fatalf("[ERROR] 3자리 이하 숫자 포맷 실패: got=%s want=999", got)
	}
	if got := CommaInt64(1234567); got != "1,234,567" {
		t.Fatalf("[ERROR] 천 단위 콤마 포맷 실패: got=%s want=1,234,567", got)
	}
}

func TestFormatTicket(t *testing.T) {
	if got := FormatTicket([]int{}); got != "[]" {
		t.Fatalf("[ERROR] 빈 배열 포맷 실패: got=%s want=[]", got)
	}
	if got := FormatTicket([]int{1, 2, 3}); got != "[1, 2, 3]" {
		t.Fatalf("[ERROR] 숫자 배열 포맷 실패: got=%s want=[1, 2, 3]", got)
	}
}

func TestFormatRate(t *testing.T) {
	if got := FormatRate(62.34); got != "62.3%" {
		t.Fatalf("[ERROR] 수익률 포맷 실패: got=%s want=62.3%%", got)
	}
	if got := FormatRate(50.0); got != "50%" {
		t.Fatalf("[ERROR] 정수 수익률 포맷 실패: got=%s want=50%%", got)
	}
}
