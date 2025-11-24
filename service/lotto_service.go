package service

import (
	"fmt"
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

func (s *LottoService) IssueTickets(money int) ([]*domain.Lotto, error) {
	if money < constants.LottoPrice || money%constants.LottoPrice != 0 {
		return nil, fmt.Errorf("구입 금액은 %d원 단위여야 합니다.", constants.LottoPrice)
	}

	count := money / constants.LottoPrice
	tickets := make([]*domain.Lotto, 0, count)

	for i := 0; i < count; i++ {
		t, err := s.GenerateTicket()
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}

func (s *LottoService) MatchCounter(ticket *domain.Lotto, winning *domain.Lotto) int {
	return ticket.MatchCount(winning)
}

func (s *LottoService) RankResolver(ticket *domain.Lotto, winning *domain.Lotto, bonus int) constants.Rank {
	matchCount := ticket.MatchCount(winning)
	bonusMatch := ticket.Contains(bonus)
	return domain.DetermineRank(matchCount, bonusMatch)
}

func (s *LottoService) Evaluate(
	tickets []*domain.Lotto,
	winning *domain.Lotto,
	bonus int,
) (map[constants.Rank]int, int64) {

	rankCounts := map[constants.Rank]int{
		constants.RankFirst:  0,
		constants.RankSecond: 0,
		constants.RankThird:  0,
		constants.RankFourth: 0,
		constants.RankFifth:  0,
		constants.RankNone:   0,
	}

	var totalPrize int64 = 0

	for _, t := range tickets {
		rank := s.RankResolver(t, winning, bonus)
		rankCounts[rank]++
		totalPrize += constants.Prize[rank]
	}

	return rankCounts, totalPrize
}

func (s *LottoService) CalcProfitRate(totalPrize int64, money int) float64 {
	if money == 0 {
		return 0
	}
	return float64(totalPrize) / float64(money) * 100.0
}