package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-lotto/util"
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