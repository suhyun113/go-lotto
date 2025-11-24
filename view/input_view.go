package view

import (
	"bufio"
	"os"
	"strings"
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
