package printer

import (
	"fmt"

	"github.com/abdfnx/gosh"
)

type Printer interface {
	Print(filePath string) (string, string, error)
}

type PrinterForWin struct {
}

func NewPrinterForWin() *PrinterForWin {
	return &PrinterForWin{}
}

func (p *PrinterForWin) Print(filePath string) (string, string, error) {
	cmd := fmt.Sprintf("%s %s %s", "Start-Process", filePath, "-Verb print")
	// run a command
	// gosh.Run(cmd)

	// run a command with output
	err, out, errout := gosh.RunOutput(cmd)
	if err != nil {
		return out, errout, err
	}

	return "", "", err
}
