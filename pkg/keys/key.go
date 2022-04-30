package keys

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"path"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/tyler-smith/go-bip32"
	"gopkg.in/yaml.v3"
)

// base58CharMap is the lookup hashmap for base58 char set
var base58CharMap map[rune]struct{}

func init() {
	base58CharMap = make(map[rune]struct{})
	for _, r := range base58CharSet {
		base58CharMap[r] = struct{}{}
	}
}

var (
	keyVersions map[string][]byte
)

func init() {
	mustDecodeHex := func(input string) []byte {
		b, err := hex.DecodeString(input)
		if err != nil {
			panic(err)
		}
		return b
	}

	keyVersions = map[string][]byte{
		path.Join(CoinTypeBtc, NetworkTypeMainnet, KeyTypePub): mustDecodeHex("0488B21E"),
		path.Join(CoinTypeBtc, NetworkTypeMainnet, KeyTypePrv): mustDecodeHex("0488ADE4"),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, KeyTypePub): mustDecodeHex("043587CF"),
		path.Join(CoinTypeBtc, NetworkTypeTestnet, KeyTypePrv): mustDecodeHex("04358394"),
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
	XPrv      string `json:"xPrv,omitempty" yaml:"xPrv,omitempty"`
	XPub      string `json:"xPub,omitempty" yaml:"xPub,omitempty"`
	Addr      string `json:"addr,omitempty" yaml:"addr,omitempty"`
	PrvKeyWif string `json:"prvKeyWif,omitempty" yaml:"prvKeyWif,omitempty"`
	PubKeyHex string `json:"pubKeyHex,omitempty" yaml:"pubKeyHex,omitempty"`
}

func (g *Key) String() string {
	jb, err := json.Marshal(g)
	if err != nil {
		return err.Error()
	}

	return string(jb)
}

func (g *Key) Print() string {
	b, err := yaml.Marshal(g)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

// New generates a new key pair with a seed. The derivation paths
// can be successive derivation indices such as m, 0, 0h etc.
// or can be provided as m/0/0h.
func New(seed []byte, network, derivationPath string) (*Key, error) {
	network = strings.ToLower(network)
	switch network {
	case NetworkTypeMainnet, NetworkTypeTestnet:
	default:
		return nil, fmt.Errorf("invalid or unsupported network: %s. allowed networks are %v", network,
			[]string{NetworkTypeMainnet, NetworkTypeTestnet},
		)
	}

	// setup key versions based on network
	bip32.PrivateWalletVersion = keyVersions[path.Join(CoinTypeBtc, network, KeyTypePrv)]
	bip32.PublicWalletVersion = keyVersions[path.Join(CoinTypeBtc, network, KeyTypePub)]

	key, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, fmt.Errorf("failed to generate root key: %w", err)
	}

	key, err = extendedKeyToDerivedExtendedKey(key, derivationPath)
	if err != nil {
		return nil, fmt.Errorf("failed to derive extended key: %w", err)
	}

	return extendedKeyToKey(key)
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

func DecodeToJson(keyString string) ([]byte, error) {
	key, err := bip32.B58Deserialize(keyString)
	if err != nil {
		return nil, fmt.Errorf("failed to decode key: %w", err)
	}

	jb, err := json.MarshalIndent(key, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to serialize key: %w", err)
	}

	return jb, nil
}

func Derive(keyString string, derivationPath string) (*Key, error) {
	key, err := bip32.B58Deserialize(keyString)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize key: %w", err)
	}

	key, err = extendedKeyToDerivedExtendedKey(key, derivationPath)
	if err != nil {
		return nil, fmt.Errorf("failed to derive extended key: %w", err)
	}

	return extendedKeyToKey(key)
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
	params := &chaincfg.MainNetParams
	if bytes.Equal(key.Version, keyVersions[path.Join(CoinTypeBtc, NetworkTypeTestnet, KeyTypePub)]) ||
		bytes.Equal(key.Version, keyVersions[path.Join(CoinTypeBtc, NetworkTypeTestnet, KeyTypePrv)]) {
		params = &chaincfg.TestNet3Params
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

	if prvKey != nil {
		prvKeyString = fmt.Sprintf("%s", prvKey)

		prv, _ := btcec.PrivKeyFromBytes(btcec.S256(), prvKey.Key)
		wif, err := btcutil.NewWIF(prv, params, true)
		if err != nil {
			return nil, fmt.Errorf("failed to generate wif formatted prv key: %w", err)
		}
		prvKeyWif = wif.String()

		serializedPubKey := wif.SerializePubKey()
		addressPubKey, err := btcutil.NewAddressPubKey(serializedPubKey, params)
		if err != nil {
			return nil, fmt.Errorf("failed to generate new address from pub key: %w", err)
		}

		addr = addressPubKey.EncodeAddress()
	} else {
		p, err := btcec.ParsePubKey(pubKey.Key, btcec.S256())
		if err != nil {
			return nil, fmt.Errorf("failed to parse pubkey: %w", err)
		}
		addressPubKey, err := btcutil.NewAddressPubKey(p.SerializeCompressed(), params)
		if err != nil {
			return nil, fmt.Errorf("failed to generate new address from pub key: %w", err)
		}

		addr = addressPubKey.EncodeAddress()
	}

	return &Key{
		XPrv:      prvKeyString,
		XPub:      pubKeyString,
		Addr:      addr,
		PrvKeyWif: prvKeyWif,
		PubKeyHex: hex.EncodeToString(pubKey.Key),
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
