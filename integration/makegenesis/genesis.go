package makegenesis

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/galaxy-digital/lachesis-base/hash"
	"github.com/galaxy-digital/lachesis-base/inter/idx"

	galaxy "github.com/galaxy-digital/axis-chain/galaxy"
	"github.com/galaxy-digital/axis-chain/galaxy/genesis"
	"github.com/galaxy-digital/axis-chain/galaxy/genesis/driver"
	"github.com/galaxy-digital/axis-chain/galaxy/genesis/driverauth"
	"github.com/galaxy-digital/axis-chain/galaxy/genesis/evmwriter"
	"github.com/galaxy-digital/axis-chain/galaxy/genesis/gpos"
	"github.com/galaxy-digital/axis-chain/galaxy/genesis/netinit"
	"github.com/galaxy-digital/axis-chain/galaxy/genesis/sfc"
	"github.com/galaxy-digital/axis-chain/galaxy/genesisstore"
	"github.com/galaxy-digital/axis-chain/inter"
	"github.com/galaxy-digital/axis-chain/inter/validatorpk"
	futils "github.com/galaxy-digital/axis-chain/utils"
)

var (
	FakeGenesisTime = inter.Timestamp(1608600000 * time.Second)
)

// FakeKey gets n-th fake private key.
func FakeKey(n int) *ecdsa.PrivateKey {
	reader := rand.New(rand.NewSource(int64(n)))

	key, err := ecdsa.GenerateKey(crypto.S256(), reader)

	fmt.Printf("\nYour new privatekey was generated %x\n", key.D)

	if err != nil {
		panic(err)
	}

	return key
}

type ValidatorAccount struct {
	address   string
	validator string
}

func MakeGenesisStore() *genesisstore.Store {
	genStore := genesisstore.NewMemStore()
	genStore.SetRules(galaxy.MainNetRules())

	var validatorAccounts = []ValidatorAccount{
		{
			address:   "0x14fCa2361553F821b65812fD40ab7bDEF37de3c2",
			validator: "043d25984decfe1da8712fa974f84b776eb5c7389dfa91421c523bd6731348f6fa54bc2d43ff71ea562fe963206c220554763c6e8b20acdf671f180feb12ae6419",
		},
		{
			address:   "0xE285AC259F5F66E3580Db4dd2909A2E57fD1bF78",
			validator: "04c677dbb8836d031f6df7fce7e30ed9de9111fe8bfd17797208e380df76d017fadbe18181a0302acd8ea23dec49f98619b11764cd23c8b05f7a7e4fc7dd586d12",
		},
		{
			address:   "0xf31D5071964C7097e6a3aC6063e9CE437755c1Be",
			validator: "04a179f1099613066ffe442e2fc65c4cb09d8f257cd06fe76f9c95b01d88d5259b245dcbaa1e23f8b7ae7c882b04712dc0f31da605f9d7f46820f54ba21fed186a",
		},
		{
			address:   "0x57C79538af87DD5583c188A9424b9Fed0ae6b620",
			validator: "04a99c17d848575b7792010e91fc11d5a29047af6085036f21587ddf51a71548a052cb96a3fa9c7f5bd6d7514fa34d8ab3c49a029bbf66b9c69da53863045d4d01",
		},
		{
			address:   "0x0cc298DBA96D854d5626B22894D5c32634745f67",
			validator: "04bd2d38c0b432fd1dd7b4e93a4d33b39f790b903e50777c092f49b0af131ed014ed7f4302d93b54e3195b50b135c64e9dc0d280792c65994288167cba463f15b6",
		},
		{
			address:   "0xe91Ee1838772964dEFdB46E5346168119882aa46",
			validator: "049237fabf6617cb94283c646254a61dd774ea09761bfe15cd5b46adab720ed1e7dc8eca006904d9d7baecc12091143aa10605b3d7add17f5855b7d7079797f689",
		},
		{
			address:   "0x8B1E640DD01Cd5Fda611214a48545bCbb3C5A306",
			validator: "047b331e9078c50870e43a5e028f97d3fa1a7a2a3e36144499f9bee7d8383e1d823273137cdaf24fe13d45ce0c60168e45c6c9d1a2be323e26f23792b660266615",
		},
		{
			address:   "0x323B6412c4329B3ba46165c4ca60d6d4f00965D1",
			validator: "04e094d25fe3a4af2ddcdcf52183eacc2ae22be5ba3f879add94271fdab9a9d02e2462d7a41cc3416a02c39e563648221a4408d157c2f45bb70ca85eda3c95c82e",
		},
	}

	var initialAccounts = []string{
		"0xD22f71Fb6820366Fe21A481e970b88058A917f2F",
		"0x960c994FC3FB8D544d673C1a87335C0885D4DAe5",
	}
	num := len(validatorAccounts)

	_total := 150
	_validator := 0
	_staker := 3

	_initial := (_total - (_validator+_staker)*num) / len(initialAccounts)
	totalSupply := futils.ToAxis(uint64(_total) * 1e6)
	balance := futils.ToAxis(uint64(_validator) * 1e6)
	stake := futils.ToAxis(uint64(_staker) * 1e6)
	initialBalance := futils.ToAxis(uint64(_initial) * 1e6)

	validators := make(gpos.Validators, 0, num)

	now := time.Now() // current local time
	// sec := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsec := now.UnixNano()
	time := inter.Timestamp(nsec)
	for i := 1; i <= num; i++ {
		addr := common.HexToAddress(validatorAccounts[i-1].address)
		pubkeyraw := common.Hex2Bytes(validatorAccounts[i-1].validator)
		// fmt.Printf("\n# addr %x pubkeyraw %s len %d\n", addr, hex.EncodeToString(pubkeyraw), len(pubkeyraw))
		validatorID := idx.ValidatorID(i)
		pubKey := validatorpk.PubKey{
			Raw:  pubkeyraw,
			Type: validatorpk.Types.Secp256k1,
		}

		validators = append(validators, gpos.Validator{
			ID:               validatorID,
			Address:          addr,
			PubKey:           pubKey,
			CreationTime:     time,
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}
	for _, val := range initialAccounts {
		genStore.SetEvmAccount(common.HexToAddress(val), genesis.Account{
			Code:    []byte{},
			Balance: initialBalance,
			Nonce:   0,
		})
	}
	for _, val := range validators {
		genStore.SetEvmAccount(val.Address, genesis.Account{
			Code:    []byte{},
			Balance: balance,
			Nonce:   0,
		})
		genStore.SetDelegation(val.Address, val.ID, genesis.Delegation{
			Stake:              stake,
			Rewards:            new(big.Int),
			LockedStake:        new(big.Int),
			LockupFromEpoch:    0,
			LockupEndTime:      0,
			LockupDuration:     0,
			EarlyUnlockPenalty: new(big.Int),
		})
	}

	var owner common.Address
	if num != 0 {
		owner = validators[0].Address
	}

	genStore.SetMetadata(genesisstore.Metadata{
		Validators:    validators,
		FirstEpoch:    2,
		Time:          time,
		PrevEpochTime: time - inter.Timestamp(time.Time().Hour()),
		ExtraData:     []byte("galaxy"),
		DriverOwner:   owner,
		TotalSupply:   totalSupply,
	})
	genStore.SetBlock(0, genesis.Block{
		Time:        time - inter.Timestamp(time.Time().Minute()),
		Atropos:     hash.Event{},
		Txs:         types.Transactions{},
		InternalTxs: types.Transactions{},
		Root:        hash.Hash{},
		Receipts:    []*types.ReceiptForStorage{},
	})
	// pre deploy NetworkInitializer
	genStore.SetEvmAccount(netinit.ContractAddress, genesis.Account{
		Code:    netinit.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriver
	genStore.SetEvmAccount(driver.ContractAddress, genesis.Account{
		Code:    driver.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriverAuth
	genStore.SetEvmAccount(driverauth.ContractAddress, genesis.Account{
		Code:    driverauth.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy SFC
	genStore.SetEvmAccount(sfc.ContractAddress, genesis.Account{
		Code:    sfc.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// set non-zero code for pre-compiled contracts
	genStore.SetEvmAccount(evmwriter.ContractAddress, genesis.Account{
		Code:    []byte{0},
		Balance: new(big.Int),
		Nonce:   0,
	})

	return genStore
}
func MakeTestnetGenesisStore() *genesisstore.Store {
	genStore := genesisstore.NewMemStore()
	genStore.SetRules(galaxy.TestNetRules())
	var validatorAccounts = []ValidatorAccount{
		{
			address:   "0xC4A4393C956d76aCBB3F60c79dB83BD376708378",
			validator: "04cdbafa4f4d7d48f305a723b71b774fc45d54996acfe562f97867f6ae0cbaf21193cb553176a8833bef038284d73a9eecef996ddcfe5096ee08933e4ddffe6301",
		},
		{
			address:   "0xe8C8C6A989e4C8655fA3F11FB82FfA841B683CA4",
			validator: "046e4edacf9574ef8c14679776f4262a537c868e38777490076183b589a40f5f6ed2e0329ad9d8d46c4dc7bdc5d8de108c1a8c821c5af20688e89c0e16bc5dbbcf",
		},
		{
			address:   "0xE6D04Ec807b7683d2DfF94DeF41C19d7C6eb9FaA",
			validator: "04985b7f89a108bbe8912a0b0d76f5eae6d282a77694cf4f4900ace5c2aee03d0a23f1f836c447696b1d6d48b5ab36e1b5f13ebed521e5bbed20ebdc5def435d59",
		},
		{
			address:   "0x528b66B85ccE01Fa10A765567910A3824897Fd71",
			validator: "04309ca7d17172cbd477cde6122f58032452d6033b0752dc3288b4cc5afb0ca867e181ec0186b6daa3014c94266e58f466d5a7959072d5b5ebdcb6bf00865f6cd3",
		},
		{
			address:   "C111bC3108b04732EA5061bA0C126E3561DD3d",
			validator: "04f8be20597088549fa4954faa8e11c66c54f142cb9276e0a9de42d4ef1500c113d00fee4cf36973eab1035b0248da1cf035e9575b1de41c6720f3e4147046aae1",
		},
		{
			address:   "0x23279ad9F6B48d17D75CAD2dE3fEEeCd77966467",
			validator: "0429b6b9d9b6f9cdcc6caa0439f386614defa05bfcd55b34c2fe33bc8f090ffa68ec58809d5c2b092a856d35f05d95ea9440fa37ea78bf3c82b88d4464ba71526d",
		},
		{
			address:   "0x9c269BAAF93B5b6eb57198874c8Be42A3d445318",
			validator: "04c1668330a76794d54333700b833d772677827fc08dd6cd630508d1ebe914d574c9e2b0c9367ad090196e0217c23a4d3996f378d56723ad62b140743f0c00d445",
		},
		{
			address:   "0x0BD337Cb966151F425FE7801F540dc6735bF927F",
			validator: "04eb204027dd5f5b167d2528058d8fc80e43455e6f6774e0bd4474c141dbbdb6f16fd6e7db6b4f108f8c0bf4f337091fbd52ac2584f7f0a0cec3950b89ac3c686c",
		},
	}
	var initialAccounts = []string{
		"0x5c14C9Fe0754eb00aaD4d5522445BA890CB9E016",
		"0x00b9A99Cf4824FEb0bA5d80F7D10Fe1a16a4B6E2",
	}

	num := len(validatorAccounts)

	_total := 150
	_validator := 0
	_staker := 3

	_initial := (_total - (_validator+_staker)*num) / 10
	totalSupply := futils.ToAxis(uint64(_total) * 1e6)
	balance := futils.ToAxis(uint64(_validator) * 1e6)
	stake := futils.ToAxis(uint64(_staker) * 1e6)
	initialBalance := futils.ToAxis(uint64(_initial) * 1e6)

	validators := make(gpos.Validators, 0, num)

	now := time.Now() // current local time
	// sec := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsec := now.UnixNano()
	time := inter.Timestamp(nsec)
	for i := 1; i <= num; i++ {
		addr := common.HexToAddress(validatorAccounts[i-1].address)
		pubkeyraw := common.Hex2Bytes(validatorAccounts[i-1].validator)
		fmt.Printf("\n# addr %x pubkeyraw %s len %d\n", addr, hex.EncodeToString(pubkeyraw), len(pubkeyraw))
		validatorID := idx.ValidatorID(i)
		pubKey := validatorpk.PubKey{
			Raw:  pubkeyraw,
			Type: validatorpk.Types.Secp256k1,
		}

		validators = append(validators, gpos.Validator{
			ID:               validatorID,
			Address:          addr,
			PubKey:           pubKey,
			CreationTime:     time,
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}

	for _, val := range initialAccounts {
		genStore.SetEvmAccount(common.HexToAddress(val), genesis.Account{
			Code:    []byte{},
			Balance: initialBalance,
			Nonce:   0,
		})
	}

	for _, val := range validators {
		genStore.SetEvmAccount(val.Address, genesis.Account{
			Code:    []byte{},
			Balance: balance,
			Nonce:   0,
		})
		genStore.SetDelegation(val.Address, val.ID, genesis.Delegation{
			Stake:              stake,
			Rewards:            new(big.Int),
			LockedStake:        new(big.Int),
			LockupFromEpoch:    0,
			LockupEndTime:      0,
			LockupDuration:     0,
			EarlyUnlockPenalty: new(big.Int),
		})
	}

	var owner common.Address
	if num != 0 {
		owner = validators[0].Address
	}

	genStore.SetMetadata(genesisstore.Metadata{
		Validators:    validators,
		FirstEpoch:    2,
		Time:          time,
		PrevEpochTime: time - inter.Timestamp(time.Time().Hour()),
		ExtraData:     []byte("fake"),
		DriverOwner:   owner,
		TotalSupply:   totalSupply,
	})
	genStore.SetBlock(0, genesis.Block{
		Time:        time - inter.Timestamp(time.Time().Minute()),
		Atropos:     hash.Event{},
		Txs:         types.Transactions{},
		InternalTxs: types.Transactions{},
		Root:        hash.Hash{},
		Receipts:    []*types.ReceiptForStorage{},
	})
	// pre deploy NetworkInitializer
	genStore.SetEvmAccount(netinit.ContractAddress, genesis.Account{
		Code:    netinit.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriver
	genStore.SetEvmAccount(driver.ContractAddress, genesis.Account{
		Code:    driver.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriverAuth
	genStore.SetEvmAccount(driverauth.ContractAddress, genesis.Account{
		Code:    driverauth.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy SFC
	genStore.SetEvmAccount(sfc.ContractAddress, genesis.Account{
		Code:    sfc.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// set non-zero code for pre-compiled contracts
	genStore.SetEvmAccount(evmwriter.ContractAddress, genesis.Account{
		Code:    []byte{0},
		Balance: new(big.Int),
		Nonce:   0,
	})

	return genStore
}
func FakeGenesisStore(num int, balance, stake *big.Int) *genesisstore.Store {
	genStore := genesisstore.NewMemStore()
	genStore.SetRules(galaxy.FakeNetRules())

	validators := GetFakeValidators(num)

	totalSupply := new(big.Int)
	for _, val := range validators {
		genStore.SetEvmAccount(val.Address, genesis.Account{
			Code:    []byte{},
			Balance: balance,
			Nonce:   0,
		})
		genStore.SetDelegation(val.Address, val.ID, genesis.Delegation{
			Stake:              stake,
			Rewards:            new(big.Int),
			LockedStake:        new(big.Int),
			LockupFromEpoch:    0,
			LockupEndTime:      0,
			LockupDuration:     0,
			EarlyUnlockPenalty: new(big.Int),
		})
		totalSupply.Add(totalSupply, balance)
	}

	var owner common.Address
	if num != 0 {
		owner = validators[0].Address
	}

	genStore.SetMetadata(genesisstore.Metadata{
		Validators:    validators,
		FirstEpoch:    2,
		Time:          FakeGenesisTime,
		PrevEpochTime: FakeGenesisTime - inter.Timestamp(time.Hour),
		ExtraData:     []byte("fake"),
		DriverOwner:   owner,
		TotalSupply:   totalSupply,
	})
	genStore.SetBlock(0, genesis.Block{
		Time:        FakeGenesisTime - inter.Timestamp(time.Minute),
		Atropos:     hash.Event{},
		Txs:         types.Transactions{},
		InternalTxs: types.Transactions{},
		Root:        hash.Hash{},
		Receipts:    []*types.ReceiptForStorage{},
	})
	// pre deploy NetworkInitializer
	genStore.SetEvmAccount(netinit.ContractAddress, genesis.Account{
		Code:    netinit.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriver
	genStore.SetEvmAccount(driver.ContractAddress, genesis.Account{
		Code:    driver.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriverAuth
	genStore.SetEvmAccount(driverauth.ContractAddress, genesis.Account{
		Code:    driverauth.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy SFC
	genStore.SetEvmAccount(sfc.ContractAddress, genesis.Account{
		Code:    sfc.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// set non-zero code for pre-compiled contracts
	genStore.SetEvmAccount(evmwriter.ContractAddress, genesis.Account{
		Code:    []byte{0},
		Balance: new(big.Int),
		Nonce:   0,
	})

	return genStore
}

func GetFakeValidators(num int) gpos.Validators {
	validators := make(gpos.Validators, 0, num)

	for i := 1; i <= num; i++ {
		key := FakeKey(i)
		addr := crypto.PubkeyToAddress(key.PublicKey)
		pubkeyraw := crypto.FromECDSAPub(&key.PublicKey)

		validatorID := idx.ValidatorID(i)
		validators = append(validators, gpos.Validator{
			ID:      validatorID,
			Address: addr,
			PubKey: validatorpk.PubKey{
				Raw:  pubkeyraw,
				Type: validatorpk.Types.Secp256k1,
			},
			CreationTime:     FakeGenesisTime,
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}

	return validators
}

type Genesis struct {
	Nonce      uint64         `json:"nonce"`
	Timestamp  uint64         `json:"timestamp"`
	ExtraData  []byte         `json:"extraData"`
	GasLimit   uint64         `json:"gasLimit"   gencodec:"required"`
	Difficulty *big.Int       `json:"difficulty" gencodec:"required"`
	Mixhash    common.Hash    `json:"mixHash"`
	Coinbase   common.Address `json:"coinbase"`
	Alloc      GenesisAlloc   `json:"alloc"      gencodec:"required"`

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number     uint64      `json:"number"`
	GasUsed    uint64      `json:"gasUsed"`
	ParentHash common.Hash `json:"parentHash"`
	BaseFee    *big.Int    `json:"baseFeePerGas"`
}

type GenesisAlloc map[common.Address]GenesisAccount

type GenesisAccount struct {
	Code       []byte                      `json:"code,omitempty"`
	Storage    map[common.Hash]common.Hash `json:"storage,omitempty"`
	Balance    *big.Int                    `json:"balance" gencodec:"required"`
	Nonce      uint64                      `json:"nonce,omitempty"`
	PrivateKey []byte                      `json:"secretKey,omitempty"` // for tests
}
