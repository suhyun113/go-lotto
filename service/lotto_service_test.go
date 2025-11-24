package service

import (
	"testing"

	"github.com/suhyun113/go-lotto/constants"
	"github.com/suhyun113/go-lotto/domain"
)

func TestIssueTickets_Success(t *testing.T) {
	svc := NewLottoService()
	tickets, err := svc.IssueTickets(5000)
	if err != nil {
		t.Fatalf("[ERROR] 티켓 발행 실패: %v", err)
	}
	if len(tickets) != 5 {
		t.Fatalf("[ERROR] 티켓 개수가 잘못됨: got=%d want=5", len(tickets))
	}
}

func TestIssueTickets_Fail_InvalidMoney(t *testing.T) {
	svc := NewLottoService()
	_, err := svc.IssueTickets(1500)
	if err == nil {
		t.Fatalf("[ERROR] 잘못된 금액(1000원 단위 아님)을 허용함")
	}
}

func TestEvaluate_CorrectRanksAndPrize(t *testing.T) {
	svc := NewLottoService()

	winning, _ := domain.NewLotto([]int{1, 2, 3, 4, 5, 6})
	bonus := 7

	t1, _ := domain.NewLotto([]int{1, 2, 3, 4, 5, 6}) // 1등
	t2, _ := domain.NewLotto([]int{1, 2, 3, 4, 5, 7}) // 2등
	t3, _ := domain.NewLotto([]int{10, 11, 12, 13, 14, 15}) // 꽝

	rankCounts, totalPrize := svc.Evaluate([]*domain.Lotto{t1, t2, t3}, winning, bonus)

	if rankCounts[constants.RankFirst] != 1 {
		t.Fatalf("[ERROR] 1등 개수 잘못됨: %d", rankCounts[constants.RankFirst])
	}
	if rankCounts[constants.RankSecond] != 1 {
		t.Fatalf("[ERROR] 2등 개수 잘못됨: %d", rankCounts[constants.RankSecond])
	}
	if rankCounts[constants.RankNone] != 1 {
		t.Fatalf("[ERROR] 꽝 개수 잘못됨: %d", rankCounts[constants.RankNone])
	}

	wantPrize := constants.Prize[constants.RankFirst] + constants.Prize[constants.RankSecond]
	if totalPrize != wantPrize {
		t.Fatalf("[ERROR] 총 상금 계산 오류: got=%d want=%d", totalPrize, wantPrize)
	}
}

func TestCalcProfitRate(t *testing.T) {
	svc := NewLottoService()
	rate := svc.CalcProfitRate(5000, 5000)
	if rate != 100.0 {
		t.Fatalf("[ERROR] 수익률 계산 오류: got=%f want=100.0", rate)
	}
}
