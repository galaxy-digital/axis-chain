package utils

import "math/big"

// ToRlv number of RLV to Wei
func ToRlv(rlv uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(rlv), big.NewInt(1e18))
}
