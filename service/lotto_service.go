package service

import (
	"math/rand"
	"time"
	"sort"

	"go-lotto/constants"
	"go-lotto/domain"
)

type LottoService struct {
	rng *rand.Rand
}

func NewLottoService() *LottoService {
	return &LottoService{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *LottoService) GenerateTicket() (*domain.Lotto, error) {
	nums := make([]int, 0, constants.LottoSize)
	used := make(map[int]bool, constants.LottoSize)

	for len(nums) < constants.LottoSize {
		n := s.rng.Intn(constants.LottoMax-constants.LottoMin+1) + constants.LottoMin
		if used[n] {
			continue
		}
		used[n] = true
		nums = append(nums, n)
	}

	sort.Ints(nums)
	return domain.NewLotto(nums)
}