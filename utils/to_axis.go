package utils

import "math/big"

// ToAxis number of RLV to Wei
func ToAxis(rlv uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(rlv), big.NewInt(1e18))
}
