package utils

import "math/big"

// ToAxis number of AXIS to Wei
func ToAxis(axis uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(axis), big.NewInt(1e18))
}
