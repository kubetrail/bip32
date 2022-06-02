package keys

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"path"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/tyler-smith/go-bip32"
)

// base58CharMap is the lookup hashmap for base58 char set
var base58CharMap map[rune]struct{}

func init() {
	base58CharMap = make(map[rune]struct{})
	for _, r := range base58CharSet {
		base58CharMap[r] = struct{}{}
	}
}

var netParams = map[string]*chaincfg.Params{
	NetworkTypeMainnet: &chaincfg.MainNetParams,
	NetworkTypeTestnet: &chaincfg.TestNet3Params,
}

var (
	keyVersions       map[string][]byte
	mainnetVersions   map[string]struct{}
	testnetVersions   map[string]struct{}
	versionToVersions map[string][]string
	versionToAddrType map[string]string
)

func mustDecodeHex(input string) []byte {
	b, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	return b
}

func init() {
	// https://electrum.readthedocs.io/en/latest/xpub_version_bytes.html#specification
	keyVersions = map[string][]byte{
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2pkhOrP2sh, KeyTypePub): mustDecodeHex(xpub),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2pkhOrP2sh, KeyTypePrv): mustDecodeHex(xprv),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2pkhOrP2sh, KeyTypePub): mustDecodeHex(tpub),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2pkhOrP2sh, KeyTypePrv): mustDecodeHex(tprv),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2wpkhP2sh, KeyTypePub):  mustDecodeHex(ypub),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2wpkhP2sh, KeyTypePrv):  mustDecodeHex(yprv),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2wpkhP2sh, KeyTypePub):  mustDecodeHex(upub),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2wpkhP2sh, KeyTypePrv):  mustDecodeHex(uprv),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2wshP2sh, KeyTypePub):   mustDecodeHex(Ypub),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2wshP2sh, KeyTypePrv):   mustDecodeHex(Yprv),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2wshP2sh, KeyTypePub):   mustDecodeHex(Upub),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2wshP2sh, KeyTypePrv):   mustDecodeHex(Uprv),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2wpkh, KeyTypePub):      mustDecodeHex(zpub),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2wpkh, KeyTypePrv):      mustDecodeHex(zprv),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2wpkh, KeyTypePub):      mustDecodeHex(vpub),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2wpkh, KeyTypePrv):      mustDecodeHex(vprv),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2wsh, KeyTypePub):       mustDecodeHex(Zpub),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, AddrTypeP2wsh, KeyTypePrv):       mustDecodeHex(Zprv),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2wsh, KeyTypePub):       mustDecodeHex(Vpub),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, AddrTypeP2wsh, KeyTypePrv):       mustDecodeHex(Vprv),
	}

	mainnetVersions = map[string]struct{}{
		xpub: {},
		xprv: {},
		ypub: {},
		yprv: {},
		Ypub: {},
		Yprv: {},
		zpub: {},
		zprv: {},
		Zpub: {},
		Zprv: {},
	}

	testnetVersions = map[string]struct{}{
		tpub: {},
		tprv: {},
		upub: {},
		uprv: {},
		Upub: {},
		Uprv: {},
		vpub: {},
		vprv: {},
		Vpub: {},
		Vprv: {},
	}

	// versionToVersions is used to detect input extended key version
	// and thereby assign bip32 pkg key versions, hence each detected
	// key version, whether from private key or public key, is matched
	// against a tuple of corresponding public and private key versions
	versionToVersions = map[string][]string{
		xpub: {xpub, xprv},
		xprv: {xpub, xprv},
		ypub: {ypub, yprv},
		yprv: {ypub, yprv},
		Ypub: {Ypub, Yprv},
		Yprv: {Ypub, Yprv},
		zpub: {zpub, zprv},
		zprv: {zpub, zprv},
		Zpub: {Zpub, Zprv},
		Zprv: {Zpub, Zprv},
		tpub: {tpub, tprv},
		tprv: {tpub, tprv},
		upub: {upub, uprv},
		uprv: {upub, uprv},
		Upub: {Upub, Uprv},
		Uprv: {Upub, Uprv},
		vpub: {vpub, vprv},
		vprv: {vpub, vprv},
		Vpub: {Vpub, Vprv},
		Vprv: {Vpub, Vprv},
	}

	versionToAddrType = map[string]string{
		xpub: AddrTypeP2pkhOrP2sh,
		xprv: AddrTypeP2pkhOrP2sh,
		ypub: AddrTypeP2wpkhP2sh,
		yprv: AddrTypeP2wpkhP2sh,
		Ypub: AddrTypeP2wshP2sh,
		Yprv: AddrTypeP2wshP2sh,
		zpub: AddrTypeP2wpkh,
		zprv: AddrTypeP2wpkh,
		Zpub: AddrTypeP2wsh,
		Zprv: AddrTypeP2wsh,
		tpub: AddrTypeP2pkhOrP2sh,
		tprv: AddrTypeP2pkhOrP2sh,
		upub: AddrTypeP2wpkhP2sh,
		uprv: AddrTypeP2wpkhP2sh,
		Upub: AddrTypeP2wshP2sh,
		Uprv: AddrTypeP2wshP2sh,
		vpub: AddrTypeP2wpkh,
		vprv: AddrTypeP2wpkh,
		Vpub: AddrTypeP2wsh,
		Vprv: AddrTypeP2wsh,
	}
}

// IsValidBase58String checks if all chars in input string
// belong to valid base58 char set
func IsValidBase58String(input string) bool {
	if len(input) == 0 {
		return false
	}

	for _, r := range input {
		if _, ok := base58CharMap[r]; !ok {
			return false
		}
	}

	return true
}

// Key represents BIP32 key components that are presented
// to the user
type Key struct {
	Seed           string `json:"seed,omitempty" yaml:"seed,omitempty"`
	XPrv           string `json:"xPrv,omitempty" yaml:"xPrv,omitempty"`
	XPub           string `json:"xPub,omitempty" yaml:"xPub,omitempty"`
	PubKeyHex      string `json:"pubKeyHex,omitempty" yaml:"pubKeyHex,omitempty"`
	PrvKeyWif      string `json:"prvKeyWif,omitempty" yaml:"prvKeyWif,omitempty"`
	Addr           string `json:"addr,omitempty" yaml:"addr,omitempty"`
	AddrType       string `json:"addrType,omitempty" yaml:"addrType,omitempty"`
	DerivationPath string `json:"derivationPath,omitempty" yaml:"derivationPath,omitempty"`
	CoinType       string `json:"coinType,omitempty" yaml:"coinType,omitempty"`
	Network        string `json:"network,omitempty" yaml:"network,omitempty"`
	segWitNested   string
	segWitBech32   string
}

type Config struct {
	Seed           []byte
	Network        string
	DerivationPath string
	AddrType       string
}

// New generates a new key pair with a seed. The derivation paths
// can be successive derivation indices such as m, 0, 0h etc.
// or can be provided as m/0/0h.
func New(config *Config) (*Key, error) {
	seed, network, derivationPath, addrType :=
		config.Seed,
		strings.ToLower(config.Network),
		strings.ToLower(config.DerivationPath),
		strings.ToLower(config.AddrType)

	switch addrType {
	case AddrTypeLegacy, AddrTypeBip44:
		addrType = AddrTypeP2pkhOrP2sh
	case AddrTypeP2sh, AddrTypeSegWitCompatible, AddrTypeBip49:
		addrType = AddrTypeP2wpkhP2sh
	case AddrTypeSegWitNative, AddrTypeBech32, AddrTypeBip84:
		addrType = AddrTypeP2wpkh
	}

	if derivationPath == "auto" {
		switch addrType {
		case AddrTypeP2pkhOrP2sh:
			derivationPath = "m/44h/0h/0h/0/0"
		case AddrTypeP2wpkhP2sh, AddrTypeP2wshP2sh:
			derivationPath = "m/49h/0h/0h/0/0"
		case AddrTypeP2wpkh, AddrTypeP2wsh:
			derivationPath = "m/84h/0h/0h/0/0"
		}
	}

	switch network {
	case NetworkTypeMainnet, NetworkTypeTestnet:
	default:
		return nil, fmt.Errorf("invalid or unsupported network: %s. allowed networks are %v", network,
			[]string{NetworkTypeMainnet, NetworkTypeTestnet},
		)
	}

	// setup key versions based on network
	var ok bool
	bip32.PublicWalletVersion, ok = keyVersions[path.Join(CoinTypeBtc, network, addrType, KeyTypePub)]
	if !ok {
		return nil, fmt.Errorf("failed to get key version for pubic key")
	}

	bip32.PrivateWalletVersion, ok = keyVersions[path.Join(CoinTypeBtc, network, addrType, KeyTypePrv)]
	if !ok {
		return nil, fmt.Errorf("failed to get key version for private key")
	}

	xKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, fmt.Errorf("failed to generate root key: %w", err)
	}

	xKey, err = extendedKeyToDerivedExtendedKey(xKey, derivationPath)
	if err != nil {
		return nil, fmt.Errorf("failed to derive extended key: %w", err)
	}

	key, err := extendedKeyToKey(xKey)
	if err != nil {
		return nil, fmt.Errorf("failed to convert extended key for output: %w", err)
	}

	key.Seed = hex.EncodeToString(seed)
	key.DerivationPath = derivationPath

	switch addrType {
	case AddrTypeP2pkhOrP2sh:
		key.segWitNested, key.segWitBech32 = "", ""
		key.AddrType = AddrTypeLegacy
	case AddrTypeP2wpkhP2sh, AddrTypeP2wshP2sh:
		key.Addr, key.segWitNested, key.segWitBech32 = key.segWitNested, "", ""
		key.AddrType = fmt.Sprintf("%s, %s", AddrTypeSegWitCompatible, AddrTypeP2sh)
	case AddrTypeP2wpkh, AddrTypeP2wsh:
		key.Addr, key.segWitNested, key.segWitBech32 = key.segWitBech32, "", ""
		key.AddrType = fmt.Sprintf("%s, %s", AddrTypeSegWitNative, AddrTypeBech32)
	default:
		return nil, fmt.Errorf("invalid addr type")
	}

	return key, nil
}

func Prompt(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "Enter key: "); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}

func Read(r io.Reader) (string, error) {
	inputReader := bufio.NewReader(r)
	key, err := inputReader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read key from input: %w", err)
	}
	key = strings.Trim(key, "\n")

	return key, nil
}

func DecodePublicHex(keyString string) (*Key, error) {
	pubKeyBytes, err := hex.DecodeString(keyString)
	if err != nil {
		return nil, fmt.Errorf("failed to decode pub key: %w", err)
	}

	pub, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	if err != nil {
		return nil, fmt.Errorf("failed to parse pub key: %w", err)
	}

	addressPubKey, err := btcutil.NewAddressPubKey(pub.SerializeCompressed(), netParams[NetworkTypeMainnet])
	if err != nil {
		return nil, fmt.Errorf("failed to generate new address from pub key: %w", err)
	}

	addr := addressPubKey.EncodeAddress()

	key := &Key{
		XPrv:      "",
		XPub:      "",
		PrvKeyWif: "",
		PubKeyHex: keyString,
		Addr:      addr,
		Network:   NetworkTypeMainnet,
		CoinType:  CoinTypeBtc,
	}

	return key, nil
}

func DecodePrivateWifKey(keyString string) (*Key, error) {
	wif, err := btcutil.DecodeWIF(keyString)
	if err != nil {
		return nil, fmt.Errorf("failed to decode wif: %w", err)
	}

	var network string
	for k, v := range netParams {
		if wif.IsForNet(v) {
			network = k
			break
		}
	}

	if len(network) == 0 {
		return nil, fmt.Errorf("detected network is not supported, only btc mainnet and testnet keys are supported")
	}

	serializedPubKey := wif.SerializePubKey()
	addressPubKey, err := btcutil.NewAddressPubKey(serializedPubKey, netParams[network])
	if err != nil {
		return nil, fmt.Errorf("failed to generate new address from pub key: %w", err)
	}

	addr := addressPubKey.EncodeAddress()

	key := &Key{
		XPrv:      "",
		XPub:      "",
		PrvKeyWif: keyString,
		PubKeyHex: hex.EncodeToString(serializedPubKey),
		Addr:      addr,
		Network:   network,
		CoinType:  CoinTypeBtc,
	}

	return key, nil
}

func DecodeExtendedKey(keyString string) (*Key, error) {
	key, err := Derive(keyString, "m")
	if err != nil {
		return nil, fmt.Errorf("failed to self derive extended key: %w", err)
	}

	return key, nil
}

func Derive(keyString string, derivationPath string) (*Key, error) {
	bip32Key, err := bip32.B58Deserialize(keyString)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize key: %w", err)
	}

	versions, ok := versionToVersions[hex.EncodeToString(bip32Key.Version)]
	if !ok {
		return nil, fmt.Errorf("failed to identity valid key version: %w", err)
	}

	bip32.PublicWalletVersion = mustDecodeHex(versions[0])
	bip32.PrivateWalletVersion = mustDecodeHex(versions[1])

	bip32Key, err = extendedKeyToDerivedExtendedKey(bip32Key, derivationPath)
	if err != nil {
		return nil, fmt.Errorf("failed to derive extended key: %w", err)
	}

	key, err := extendedKeyToKey(bip32Key)
	if err != nil {
		return nil, fmt.Errorf("failed to get key from extended key")
	}

	switch versionToAddrType[hex.EncodeToString(bip32Key.Version)] {
	case AddrTypeP2pkhOrP2sh:
		key.segWitNested, key.segWitBech32 = "", ""
	case AddrTypeP2wpkhP2sh, AddrTypeP2wshP2sh:
		key.Addr, key.segWitNested, key.segWitBech32 = key.segWitNested, "", ""
	case AddrTypeP2wpkh, AddrTypeP2wsh:
		key.Addr, key.segWitNested, key.segWitBech32 = key.segWitBech32, "", ""
	}

	switch versionToAddrType[hex.EncodeToString(bip32Key.Version)] {
	case AddrTypeP2pkhOrP2sh:
		key.AddrType = AddrTypeLegacy
	case AddrTypeP2wpkhP2sh, AddrTypeP2wshP2sh:
		key.AddrType = fmt.Sprintf("%s, %s", AddrTypeSegWitCompatible, AddrTypeP2sh)
	case AddrTypeP2wpkh, AddrTypeP2wsh:
		key.AddrType = fmt.Sprintf("%s, %s", AddrTypeSegWitNative, AddrTypeBech32)
	}

	return key, nil
}

func extendedKeyToDerivedExtendedKey(key *bip32.Key, derivationPath string) (*bip32.Key, error) {
	derivationPath = strings.Trim(strings.ToLower(derivationPath), "/")
	if len(derivationPath) == 0 {
		derivationPath = "m"
	}

	parts := strings.Split(derivationPath, "/")
	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid derivation path, must not be empty")
	}
	if parts[0] != "m" {
		return nil, fmt.Errorf("invalid derivation path, must start with m: %s", derivationPath)
	}

	for i, part := range parts {
		if i == 0 {
			continue
		}
		var idx uint32
		if part[len(part)-1] == '\'' || part[len(part)-1] == 'h' {
			idx = bip32.FirstHardenedChild
			part = part[:len(part)-1]
		}

		index, err := strconv.ParseInt(part, 10, 64)
		if err != nil || index < 0 {
			return nil, fmt.Errorf("invalid derivation path at index %d: %s, %w", i, derivationPath, err)
		}

		idx += uint32(index)
		key, err = key.NewChildKey(idx)
		if err != nil {
			return nil, fmt.Errorf("failed to generate %d child key: %w", i, err)
		}
	}

	return key, nil
}

func extendedKeyToKey(key *bip32.Key) (*Key, error) {
	var network string
	var params *chaincfg.Params

	if _, ok := mainnetVersions[hex.EncodeToString(key.Version)]; ok {
		params = &chaincfg.MainNetParams
		network = NetworkTypeMainnet
	} else {
		if _, ok := testnetVersions[hex.EncodeToString(key.Version)]; ok {
			params = &chaincfg.TestNet3Params
			network = NetworkTypeTestnet
		}
	}

	if len(network) == 0 {
		return nil, fmt.Errorf("unsupported network and/or coin type, accepted values are BTC:%v",
			[]string{NetworkTypeMainnet, NetworkTypeTestnet})
	}

	var pubKey *bip32.Key
	var prvKey *bip32.Key

	var prvKeyString string
	var pubKeyString string

	var addr string
	var prvKeyWif string

	if key.IsPrivate {
		prvKey = key
		pubKey = key.PublicKey()
	} else {
		pubKey = key
	}

	pubKeyString = fmt.Sprintf("%s", pubKey)
	var serializedPubKey []byte

	if prvKey != nil {
		prvKeyString = fmt.Sprintf("%s", prvKey)

		prv, _ := btcec.PrivKeyFromBytes(btcec.S256(), prvKey.Key)

		wif, err := btcutil.NewWIF(prv, params, true)
		if err != nil {
			return nil, fmt.Errorf("failed to generate wif formatted prv key: %w", err)
		}
		prvKeyWif = wif.String()

		serializedPubKey = wif.SerializePubKey()
	} else {
		p, err := btcec.ParsePubKey(pubKey.Key, btcec.S256())
		if err != nil {
			return nil, fmt.Errorf("failed to parse pubkey: %w", err)
		}

		serializedPubKey = p.SerializeCompressed()
	}

	addressPubKey, err := btcutil.NewAddressPubKey(serializedPubKey, params)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new address from pub key: %w", err)
	}

	addr = addressPubKey.EncodeAddress()

	// generate a normal p2wkh address from the pubkey hash
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, params)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new address witness pub key hash: %w", err)
	}

	segwitBech32 := addressWitnessPubKeyHash.EncodeAddress()

	// generate an address which is
	// backwards compatible to Bitcoin nodes running 0.6.0 onwards, but
	// allows us to take advantage of segwit's scripting improvments,
	// and malleability fixes.
	serializedScript, err := txscript.PayToAddrScript(addressWitnessPubKeyHash)
	if err != nil {
		return nil, fmt.Errorf("failed to generate pay to addr script: %w", err)
	}

	addressScriptHash, err := btcutil.NewAddressScriptHash(serializedScript, params)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new address script hash: %w", err)
	}

	segwitNested := addressScriptHash.EncodeAddress()

	return &Key{
		XPrv:         prvKeyString,
		XPub:         pubKeyString,
		PrvKeyWif:    prvKeyWif,
		PubKeyHex:    hex.EncodeToString(pubKey.Key),
		Addr:         addr,
		segWitNested: segwitNested,
		segWitBech32: segwitBech32,
		Network:      network,
		CoinType:     CoinTypeBtc,
	}, nil
}

func Validate(keyString string) error {
	key, err := bip32.B58Deserialize(keyString)
	if err != nil {
		return fmt.Errorf("failed to decode key: %w", err)
	}

	versionFound := false
	for k, version := range keyVersions {
		if bytes.Equal(key.Version, version) {
			switch path.Base(k) {
			case KeyTypePub:
				if key.IsPrivate {
					return fmt.Errorf("key is marked private, however, key version is public")
				}
			case KeyTypePrv:
				if !key.IsPrivate {
					return fmt.Errorf("key is marked public, however, key version is private")
				}
			}
			versionFound = true
			break
		}
	}
	if !versionFound {
		return fmt.Errorf("unknown key version found")
	}

	if !key.IsPrivate && key.Key[0] == 4 {
		return fmt.Errorf("invalid public key prefix 04")
	}

	if key.IsPrivate && key.Key[0] == 4 {
		return fmt.Errorf("invalid private key prefix 04")
	}

	if !key.IsPrivate && key.Key[0] == 1 {
		return fmt.Errorf("invalid public key prefix 01")
	}

	if key.IsPrivate && key.Key[0] == 1 {
		return fmt.Errorf("invalid private key prefix 01")
	}

	if key.Depth == 0 {
		for _, fp := range key.FingerPrint {
			if fp > 0 {
				return fmt.Errorf("key depth is zero, however, parent non-zero fingerprint exists")
			}
		}
	}

	if key.Depth == 0 {
		for _, fp := range key.ChildNumber {
			if fp > 0 {
				return fmt.Errorf("key depth is zero, however, non-zero child index exists")
			}
		}
	}

	if key.IsPrivate {
		n := new(big.Int)
		var z *big.Int
		var acc big.Accuracy

		if f, _, err := big.ParseFloat(BigZ, 10, 0, big.ToNearestEven); err != nil {
			return fmt.Errorf("failed to big parse float 0")
		} else {
			z, acc = f.Int(z)
			if acc != big.Exact {
				return fmt.Errorf("exact accuracy not found in computing z")
			}
		}

		bigN, err := base64.StdEncoding.DecodeString(BigN)
		if err != nil {
			return fmt.Errorf("failed to base64 decode big N")
		}
		n.SetBytes(bigN)

		x := new(big.Int)
		x.SetBytes(key.Key)

		if x.Cmp(n) != -1 {
			return fmt.Errorf("key is not in 1:n-1, key is too large")
		}

		if x.Cmp(z) != 1 {
			return fmt.Errorf("key is not in 1:n-1, key is too small")
		}
	}

	return nil
}
