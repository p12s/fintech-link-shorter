package repository

import (
	"math/big"
)

// NOTATION - основание целевой системы счисления, использующей 0-9a-zA-Z
const NOTATION = 62

// Convert - конвертер 10-чной системы счисления в 62-ричную
type Convert struct {
	Notation map[int64]rune
}

// NewConvert - конструктор
func NewConvert() *Convert {
	return &Convert{}
}

// Convert - перевод 10-чной системы счисления в 62-ричную: 0-9a-zA-Z
// Чтобы не псиать самому метод перевода, воспользуемся реализацией пакета "big"
func (n *Convert) Convert(number int64) string {
	return big.NewInt(number).Text(NOTATION)
}
