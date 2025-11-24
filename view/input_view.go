package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-lotto/util"
	"go-lotto/domain"
)

type InputView struct {
	reader *bufio.Reader
}

func NewInputView() *InputView {
	return &InputView{reader: bufio.NewReader(os.Stdin)}
}

func (iv *InputView) readLine() string {
	line, _ := iv.reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func (iv *InputView) ReadPurchaseMoney() int {
	for {
		fmt.Println("구입금액을 입력해 주세요.")
		moneyStr := iv.readLine()

		amount, err := util.ParseAmount(moneyStr)
		if err != nil {
			PrintError(err)
			continue
		}

		amount, err = util.ValidateAmount(amount)
		if err != nil {
			PrintError(err)
			continue
		}

		return amount
	}
}

func (iv *InputView) ReadWinningNumbers() *domain.Lotto {
	for {
		fmt.Println("\n당첨 번호를 입력해 주세요.")
		line := iv.readLine()

		nums, err := util.ParseWinningNumbers(line)
		if err != nil {
			PrintError(err)
			continue
		}

		nums, err = util.ValidateWinningNumbers(nums)
		if err != nil {
			PrintError(err)
			continue
		}

		lotto, err := domain.NewLotto(nums)
		if err != nil {
			PrintError(err)
			continue
		}

		return lotto
	}
}

func (iv *InputView) ReadBonusNumber(winning *domain.Lotto) int {
	for {
		fmt.Println("\n보너스 번호를 입력해 주세요.")
		line := iv.readLine()

		bonus, err := util.ParseBonusNumber(line)
		if err != nil {
			PrintError(err)
			continue
		}

		bonus, err = util.ValidateBonusNumber(bonus, winning.GetNumbers())
		if err != nil {
			PrintError(err)
			continue
		}

		return bonus
	}
}