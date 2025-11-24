package util

import (
	"fmt"

	"go-lotto/constants"
)

func inRange(n int) bool {
	return n >= constants.LottoMin && n <= constants.LottoMax
}

func ValidateAmount(amount int) (int, error) {
	if amount == 0 {
		return 0, fmt.Errorf("구입 금액은 숫자여야 합니다.")
	}

	if amount < constants.LottoPrice || amount%constants.LottoPrice != 0 {
		return 0, fmt.Errorf("구입 금액은 1000원 단위여야 합니다.")
	}

	return amount, nil
}

func ValidateWinningNumbers(numbers []int) ([]int, error) {
	if len(numbers) != constants.LottoSize {
		return nil, fmt.Errorf("로또 번호는 %d개의 숫자여야 합니다.", constants.LottoSize)
	}

	seen := make(map[int]bool, len(numbers))
	for _, n := range numbers {
		if !inRange(n) {
			return nil, fmt.Errorf(
				"로또 번호는 %d부터 %d 사이의 숫자여야 합니다.",
				constants.LottoMin, constants.LottoMax,
			)
		}
		if seen[n] {
			return nil, fmt.Errorf("로또 번호는 중복될 수 없습니다.")
		}
		seen[n] = true
	}

	return numbers, nil
}

func ValidateBonusNumber(bonus int, winningNumbers []int) (int, error) {
	if !inRange(bonus) {
		return 0, fmt.Errorf(
			"보너스 번호는 %d부터 %d 사이의 숫자여야 합니다.",
			constants.LottoMin, constants.LottoMax,
		)
	}

	for _, w := range winningNumbers {
		if w == bonus {
			return 0, fmt.Errorf("보너스 번호는 당첨 번호와 중복될 수 없습니다.")
		}
	}

	return bonus, nil
}