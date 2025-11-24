package integration_test

import (
	"testing"

	"github.com/suhyun113/go-lotto/constants"
	"github.com/suhyun113/go-lotto/domain"
	"github.com/suhyun113/go-lotto/service"
	"github.com/suhyun113/go-lotto/util"
)

func TestApplicationFlow_FunctionalLikeJS(t *testing.T) {
	svc := service.NewLottoService()

	fixedTickets := [][]int{
		{8, 21, 23, 41, 42, 43},
		{3, 5, 11, 16, 32, 38},
		{7, 11, 16, 35, 36, 44},
		{1, 8, 11, 31, 41, 42},
		{13, 14, 16, 38, 42, 45},
		{7, 11, 30, 40, 42, 43},
		{2, 13, 22, 32, 38, 45},
		{1, 3, 5, 14, 22, 45},
	}

	tickets := make([]*domain.Lotto, 0, len(fixedTickets))
	for _, nums := range fixedTickets {
		tk, err := domain.NewLotto(nums)
		if err != nil {
			t.Fatalf("[ERROR] 고정 티켓 생성 실패: %v", err)
		}
		tickets = append(tickets, tk)
	}

	money := 8000
	winning, _ := domain.NewLotto([]int{1, 2, 3, 4, 5, 6})
	bonus := 7

	rankCounts, totalPrize := svc.Evaluate(tickets, winning, bonus)

	if rankCounts[constants.RankFifth] != 1 {
		t.Fatalf("[ERROR] 5등 개수 결과가 다름: got=%d want=1", rankCounts[constants.RankFifth])
	}
	if rankCounts[constants.RankNone] != 7 {
		t.Fatalf("[ERROR] 꽝 개수 결과가 다름: got=%d want=7", rankCounts[constants.RankNone])
	}

	wantPrize := constants.Prize[constants.RankFifth]
	if totalPrize != wantPrize {
		t.Fatalf("[ERROR] 총 상금 결과가 다름: got=%d want=%d", totalPrize, wantPrize)
	}

	rate := svc.CalcProfitRate(totalPrize, money)
	formattedRate := util.FormatRate(rate)
	if formattedRate != "62.5%" {
		t.Fatalf("[ERROR] 수익률 결과가 다름: got=%s want=62.5%%", formattedRate)
	}
}

func TestApplicationFlow_ExceptionLikeJS(t *testing.T) {
	if _, err := util.ParseAmount("1000j"); err == nil {
		t.Fatalf("[ERROR] 숫자가 아닌 구입 금액을 허용하면 안 된다")
	}
}
