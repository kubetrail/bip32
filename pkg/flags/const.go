package flags

const (
	DerivationPath         = "derivation-path"
	UsePassphrase          = "use-passphrase"
	SkipMnemonicValidation = "skip-mnemonic-validation"
	InputHexSeed           = "input-hex-seed"
	Network                = "network"
	MnemonicLanguage       = "mnemonic-language"
	AddrType               = "addr-type"
)

const (
	OutputFormat = "output-format"
)

const (
	OutputFormatNative = "native"
	OutputFormatJson   = "json"
	OutputFormatYaml   = "yaml"
)

const (
	NetworkMainnet = "mainnet"
	NetworkTestnet = "testnet"
)

// BIP-44 format m/purpose'/coinType'/account'/change/addressIndex
// 0h is coin type BTC
// 0h is account number, can be 1, 2, etc.
// 0 is the change
// 0, 1, etc. are address indices
const (
	DerivationPathAuto = "auto"
	DerivationPath0    = "m/44h/0h/0h/0/0"
	DerivationPath1    = "m/44h/0h/0h/0/1"
	DerivationPath2    = "m/49h/0h/0h/0/0"
	DerivationPath3    = "m/49h/0h/0h/0/1"
	DerivationPath4    = "m/84h/0h/0h/0/0"
	DerivationPath5    = "m/84h/0h/0h/0/1"
	DerivationPath6    = "m/44h/0h/0h"
	DerivationPath7    = "m/49h/0h/0h"
	DerivationPath8    = "m/84h/0h/0h"
)
