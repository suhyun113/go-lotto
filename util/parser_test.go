package util

import (
	"testing"

	"github.com/suhyun113/go-lotto/constants"
)

func TestParseAmount_Success(t *testing.T) {
	got, err := ParseAmount("5000")
	if err != nil {
		t.Fatalf("[ERROR] 성공해야 하는 입력인데 에러 발생: %v", err)
	}
	if got != 5000 {
		t.Fatalf("[ERROR] 금액 파싱 결과가 잘못됨: got=%d want=5000", got)
	}
}

func TestParseAmount_Fail_NonDigit(t *testing.T) {
	_, err := ParseAmount("1,000")
	if err == nil {
		t.Fatalf("[ERROR] 숫자가 아닌 입력을 허용함 (콤마 포함 입력)")
	}
}

func TestParseWinningNumbers_Success(t *testing.T) {
	input := "1, 2, 3, 4, 5, 6"
	got, err := ParseWinningNumbers(input)
	if err != nil {
		t.Fatalf("[ERROR] 당첨 번호 파싱 실패: %v", err)
	}
	if len(got) != constants.LottoSize {
		t.Fatalf("[ERROR] 파싱된 번호 길이가 잘못됨: got=%d want=%d",
			len(got), constants.LottoSize)
	}
}

func TestParseWinningNumbers_Fail_WrongCount(t *testing.T) {
	_, err := ParseWinningNumbers("1,2,3")
	if err == nil {
		t.Fatalf("[ERROR] 번호 개수가 잘못됐는데 에러가 발생하지 않음")
	}
}

func TestParseWinningNumbers_Fail_NonDigit(t *testing.T) {
	_, err := ParseWinningNumbers("1,2,a,4,5,6")
	if err == nil {
		t.Fatalf("[ERROR] 숫자가 아닌 입력을 허용함")
	}
}

func TestParseBonusNumber_Success(t *testing.T) {
	got, err := ParseBonusNumber("7")
	if err != nil {
		t.Fatalf("[ERROR] 보너스 번호 파싱 실패: %v", err)
	}
	if got != 7 {
		t.Fatalf("[ERROR] 보너스 번호 파싱 결과가 잘못됨: got=%d want=7", got)
	}
}

func TestParseBonusNumber_Fail_NonDigit(t *testing.T) {
	_, err := ParseBonusNumber("7a")
	if err == nil {
		t.Fatalf("[ERROR] 숫자가 아닌 보너스 입력을 허용함")
	}
}
