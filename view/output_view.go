package view

import (
	"fmt"

	"go-lotto/domain"
	"go-lotto/util"
)

type OutputView struct{}

func NewOutputView() *OutputView { return &OutputView{} }

func (ov *OutputView) PrintTickets(tickets []*domain.Lotto) {
	fmt.Printf("\n%d개를 구매했습니다.\n", len(tickets))
	for _, t := range tickets {
		fmt.Println(util.FormatTicket(t.GetNumbers()))
	}
}