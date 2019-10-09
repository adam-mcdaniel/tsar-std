package tsar_std_foreign

import (
	"fmt"

	. "github.com/adam-mcdaniel/xgopher"
)

func Xasm_println(m *Machine) {
	fmt.Println(m.Pop())
}

func Xasm_print(m *Machine) {
	fmt.Print(m.Pop())
}
