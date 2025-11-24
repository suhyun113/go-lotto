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
