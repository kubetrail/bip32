package keys

// https://electrum.readthedocs.io/en/latest/xpub_version_bytes.html#specification
const (
	ScriptTypeP2pkhOrP2sh = "p2pkh-or-p2sh" // mainnet: [xpub, xprv], testnet: [tpub, tprv]
	ScriptTypeP2wpkhP2sh  = "p2wpkh-p2sh"   // mainnet: [ypub, yprv], testnet: [upub, uprv]
	ScriptTypeP2wshP2sh   = "p2wsh-p2sh"    // mainnet: [Ypub, Yprv], testnet: [Upub, Uprv]
	ScriptTypeP2wpkh      = "p2wpkh"        // mainnet: [zpub, zprv], testnet: [vpub, vprv]
	ScriptTypeP2wsh       = "p2wsh"         // mainnet: [Zpub, Zprv], testnet: [Vpub, Vprv]

	ScriptTypeLegacy           = "legacy"            // same as ScriptTypeP2pkhOrP2sh, xpub, xprv etc.
	ScriptTypeP2sh             = "p2sh"              // same as ScriptTypeP2wpkhP2sh, ypub, yprv etc.
	ScriptTypeSegWitCompatible = "segwit-compatible" // same as ScriptTypeP2wpkhP2sh, ypub, yprv etc.
	ScriptTypeSegWitNative     = "segwit-native"     // same as ScriptTypeP2wpkh, zpub, zprv etc.
	ScriptTypeBech32           = "bech32"            // same as ScriptTypeP2wpkh, zpub, zprv etc.
)

// key versions
const (
	xpub = "0488b21e"
	xprv = "0488ade4"
	tpub = "043587cf"
	tprv = "04358394"
	ypub = "049d7cb2"
	yprv = "049d7878"
	upub = "044a5262"
	uprv = "044a4e28"
	Ypub = "0295b43f"
	Yprv = "0295b005"
	Upub = "024289ef"
	Uprv = "024285b5"
	zpub = "04b24746"
	zprv = "04b2430c"
	vpub = "045f1cf6"
	vprv = "045f18bc"
	Zpub = "02aa7ed3"
	Zprv = "02aa7a99"
	Vpub = "02575483"
	Vprv = "02575048"
)
