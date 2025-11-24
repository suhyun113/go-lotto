package main

import (
	"go-lotto/service"
	"go-lotto/view"
)

func main() {
	iv := view.NewInputView()
	ov := view.NewOutputView()
	svc := service.NewLottoService()

	// 구입 금액 입력
	money := iv.ReadPurchaseMoney()

	// 티켓 발행
	tickets, err := svc.IssueTickets(money)
	if err != nil {
		view.PrintError(err)
		return
	}
	ov.PrintTickets(tickets)

	// 당첨 번호/보너스 입력
	winning := iv.ReadWinningNumbers()
	bonus := iv.ReadBonusNumber(winning)

	// 평가
	rankCounts, totalPrize := svc.Evaluate(tickets, winning, bonus)

	// 결과 출력
	ov.PrintEvaluationResults(rankCounts, totalPrize, money)
}
