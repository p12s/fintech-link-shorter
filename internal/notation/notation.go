package notation

import (
	"math/big"
)

const NOTATION = 62

type Convert struct {
	Notation map[int64]rune
}

func NewConvert() *Convert {
	return &Convert{}
}

// Convert - перевод 10-чной системы счисления в 62-ричную: 0-9a-zA-Z
// Чтобы не псиать самому метод перевода, воспользуемся реализацией пакета "big"
func (n *Convert) Convert(number int64) string {
	return big.NewInt(number).Text(NOTATION)
}
