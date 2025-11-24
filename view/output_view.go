package view

import (
	"fmt"

	"go-lotto/domain"
	"go-lotto/util"
	"go-lotto/constants"
)

type OutputView struct{}

func NewOutputView() *OutputView { return &OutputView{} }

func (ov *OutputView) PrintTickets(tickets []*domain.Lotto) {
	fmt.Printf("\n%d개를 구매했습니다.\n", len(tickets))
	for _, t := range tickets {
		fmt.Println(util.FormatTicket(t.GetNumbers()))
	}
}

func (ov *OutputView) PrintEvaluationResults(
	rankCounts map[constants.Rank]int,
	totalPrize int64,
	purchaseAmount int,
) {
	fmt.Println("\n당첨 통계")
	fmt.Println("---")

	order := []constants.Rank{
		constants.RankFifth,
		constants.RankFourth,
		constants.RankThird,
		constants.RankSecond,
		constants.RankFirst,
	}

	for _, r := range order {
		c := 0
		if rankCounts != nil {
			c = rankCounts[r]
		}
		fmt.Printf("%s - %d개\n", constants.RankLabel[r], c)
	}

	rate := 0.0
	if purchaseAmount > 0 {
		rate = float64(totalPrize) / float64(purchaseAmount) * 100.0
	}
	fmt.Printf("총 수익률은 %s입니다.\n", util.FormatRate(rate))
	fmt.Printf("(총 당첨금: %s원)\n", util.CommaInt64(totalPrize))
}