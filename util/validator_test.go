package util

import (
	"testing"

	"github.com/suhyun113/go-lotto/constants"
)

func TestValidateAmount_Success(t *testing.T) {
	got, err := ValidateAmount(5000)
	if err != nil {
		t.Fatalf("[ERROR] 금액 검증 실패: %v", err)
	}
	if got != 5000 {
		t.Fatalf("[ERROR] 금액 검증 결과가 잘못됨: got=%d want=5000", got)
	}
}

func TestValidateAmount_Fail_Zero(t *testing.T) {
	_, err := ValidateAmount(0)
	if err == nil {
		t.Fatalf("[ERROR] 0원을 허용함")
	}
}

func TestValidateAmount_Fail_NotMultiple(t *testing.T) {
	_, err := ValidateAmount(1500)
	if err == nil {
		t.Fatalf("[ERROR] 1000원 단위가 아닌 금액을 허용함")
	}
}

func TestValidateWinningNumbers_Success(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	got, err := ValidateWinningNumbers(nums)
	if err != nil {
		t.Fatalf("[ERROR] 당첨 번호 검증 실패: %v", err)
	}
	if len(got) != constants.LottoSize {
		t.Fatalf("[ERROR] 번호 개수가 잘못됨: got=%d want=%d", len(got), constants.LottoSize)
	}
}

func TestValidateWinningNumbers_Fail_Duplicate(t *testing.T) {
	_, err := ValidateWinningNumbers([]int{1, 1, 2, 3, 4, 5})
	if err == nil {
		t.Fatalf("[ERROR] 중복 번호를 허용함")
	}
}

func TestValidateWinningNumbers_Fail_OutOfRange(t *testing.T) {
	_, err := ValidateWinningNumbers([]int{0, 2, 3, 4, 5, 6})
	if err == nil {
		t.Fatalf("[ERROR] 범위 밖 번호를 허용함")
	}
}

func TestValidateBonusNumber_Success(t *testing.T) {
	bonus := 7
	winning := []int{1, 2, 3, 4, 5, 6}
	got, err := ValidateBonusNumber(bonus, winning)
	if err != nil {
		t.Fatalf("[ERROR] 보너스 번호 검증 실패: %v", err)
	}
	if got != bonus {
		t.Fatalf("[ERROR] 보너스 번호 반환값이 잘못됨: got=%d want=%d", got, bonus)
	}
}

func TestValidateBonusNumber_Fail_Duplicate(t *testing.T) {
	_, err := ValidateBonusNumber(6, []int{1, 2, 3, 4, 5, 6})
	if err == nil {
		t.Fatalf("[ERROR] 보너스 번호가 당첨 번호와 중복되는데 허용됨")
	}
}

func TestValidateBonusNumber_Fail_OutOfRange(t *testing.T) {
	_, err := ValidateBonusNumber(46, []int{1, 2, 3, 4, 5, 6})
	if err == nil {
		t.Fatalf("[ERROR] 범위 밖 보너스 번호를 허용함")
	}
}
