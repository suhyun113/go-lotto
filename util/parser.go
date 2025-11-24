package util

import (
	"fmt"
	"regexp"
	"strconv"
	"string"

	"go-lotto/constants"
)

func ParseAmount(amountStr string) (int, error) {
	amount := strings.TrimSpace(amountStr)
	if !onlyDigits.MatchString(amount) {
		return 0, fmt.Errorf("구입 금액은 숫자여야 합니다.")
	}

	n, err := strconv.Atoi(amount)
	if err != nil {
		return 0, fmt.Errorf("구입 금액은 숫자여야 합니다.")
	}
	return n, nil
}

func ParseWinningNumbers(numbersStr string) ([]int, error) {
	raw := strings.TrimSpace(numbersStr)
	tokens := strings.Split(raw, ",")

	if len(tokens) != constants.LottoSize {
		return nil, fmt.Errorf("당첨 번호는 쉼표로 구분된 %d개의 숫자여야 합니다.", constants.LottoSize)
	}

	nums := make([]int, 0, constants.LottoSize)
	for _, t := range tokens {
		t = strings.TrimSpace(t)
		if !onlyDigits.MatchString(t) {
			return nil, fmt.Errorf("당첨 번호는 쉼표로 구분된 %d개의 숫자여야 합니다.", constants.LottoSize)
		}
		n, err := strconv.Atoi(t)
		if err != nil {
			return nil, fmt.Errorf("당첨 번호는 쉼표로 구분된 %d개의 숫자여야 합니다.", constants.LottoSize)
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func ParseBonusNumber(bonusStr string) (int, error) {
	bonus := strings.TrimSpace(bonusStr)
	if !onlyDigits.MatchString(bonus) {
		return 0, fmt.Errorf("보너스 번호는 숫자여야 합니다.")
	}

	n, err := strconv.Atoi(bonus)
	if err != nil {
		return 0, fmt.Errorf("보너스 번호는 숫자여야 합니다.")
	}
	return n, nil
}