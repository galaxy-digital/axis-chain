package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://8f58673cb8f1a8bc8d109c1112034d45dc379c3c702769ccd1f561e3b68d52a3680c15f660ee5f308a53c60bc1d46511abd762cb6931c4f953fa7863d813e79b@185.64.104.17:15060",
		"enode://e8514c480b21d722ba99a74e290bb69a7b333ce8cac86f99f8897e678190b7f445fea0925bfac232cf31d5d5115d30f2cf8a06abfc728119fff6d794d129d223@185.25.50.199:15060",
		"enode://af9ad92f3004d220c2271b3cbf7e4095c08e07e6ad9d5edb725efbae9e4069a027ea0ed53343f5d280adcc43d8b92f30620d66f7b6d9b0df0a767df7ea34b3a2@185.25.50.202:15060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
