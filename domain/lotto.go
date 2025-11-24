package domain

import (
	"fmt"
	"sort"

	"go-lotto/constants"
)

type Lotto struct {
	numbers []int
}

func NewLotto(numbers []int) (*Lotto, error) {
	if err := validate(numbers); err != nil {
		return nil, err
	}

	cp := make([]int, len(numbers))
	copy(cp, numbers)
	sort.Ints(cp)

	return &Lotto{numbers: cp}, nil
}

func validate(numbers []int) error {
	if len(numbers) != constants.LottoSize {
		return fmt.Errorf("로또 번호는 %d개여야 합니다.", constants.LottoSize)
	}

	seen := make(map[int]bool, len(numbers))
	for _, n := range numbers {
		if n < constants.LottoMin || n > constants.LottoMax {
			return fmt.Errorf(
				"로또 번호는 %d부터 %d 사이의 숫자여야 합니다.",
				constants.LottoMin, constants.LottoMax,
			)
		}
		if seen[n] {
			return fmt.Errorf("로또 번호는 중복될 수 없습니다.")
		}
		seen[n] = true
	}

	return nil
}

func (l *Lotto) GetNumbers() []int {
	cp := make([]int, len(l.numbers))
	copy(cp, l.numbers)
	return cp
}

func (l *Lotto) Contains(n int) bool {
	for _, v := range l.numbers {
		if v == n {
			return true
		}
	}
	return false
}

func (l *Lotto) MatchCount(win *Lotto) int {
	cnt := 0
	for _, n := range l.numbers {
		if win.Contains(n) {
			cnt++
		}
	}
	return cnt
}