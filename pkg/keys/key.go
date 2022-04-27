package keys

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"path"
	"strconv"
	"strings"

	"github.com/tyler-smith/go-bip32"
	"github.com/wemeetagain/go-hdwallet"
)

// Key represents BIP32 key components that are presented
// to the user
type Key struct {
	Addr string `json:"addr,omitempty"`
	Prv  string `json:"prv,omitempty"`
	Pub  string `json:"pub,omitempty"`
}

func (g *Key) String() string {
	jb, err := json.Marshal(g)
	if err != nil {
		return err.Error()
	}

	return string(jb)
}

// New generates a new key pair with a seed. The derivation paths
// can be successive derivation indices such as m, 0, 0h etc.
// or can be provided as m/0/0h.
func New(seed []byte, derivationPaths ...string) (*Key, error) {
	derivationPath := path.Join(derivationPaths...)
	if len(derivationPath) == 0 {
		derivationPath = "m"
	}

	key, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, fmt.Errorf("failed to generate root key: %w", err)
	}

	derivationPath = strings.Trim(derivationPath, "/")
	parts := strings.Split(derivationPath, "/")
	if len(parts) == 0 || parts[0] != "m" {
		return nil, fmt.Errorf("invalid derivation path: %s", derivationPath)
	}

	for i, part := range parts {
		if i == 0 {
			continue
		}

		if len(part) == 0 {
			return nil, fmt.Errorf("invalid derivation path at index %d: %s", i, derivationPath)
		}

		var idx uint32
		if part[len(part)-1] == '\'' || part[len(part)-1] == 'h' || part[len(part)-1] == 'H' {
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

	prv := fmt.Sprintf("%s", key)
	pub := fmt.Sprintf("%s", key.PublicKey())

	wallet, err := hdwallet.StringWallet(prv)
	if err != nil {
		return nil, fmt.Errorf("failed to create hdwallet: %w", err)
	}

	if wallet.String() != prv {
		return nil, fmt.Errorf("private key mismatch, \nwallet: %s\n bip32: %s", wallet.String(), prv)
	}

	if wallet.Pub().String() != pub {
		return nil, fmt.Errorf("private key mismatch, \nwallet: %s\n bip32: %s", wallet.Pub().String(), pub)
	}

	return &Key{
		Addr: wallet.Address(),
		Prv:  prv,
		Pub:  pub,
	}, nil
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

func Validate(keyString string) error {
	key, err := bip32.B58Deserialize(keyString)
	if err != nil {
		return fmt.Errorf("failed to decode key: %w", err)
	}

	if !bytes.Equal(key.Version, hdwallet.Private) &&
		!bytes.Equal(key.Version, hdwallet.Public) &&
		!bytes.Equal(key.Version, hdwallet.TestPrivate) &&
		!bytes.Equal(key.Version, hdwallet.TestPublic) {
		return fmt.Errorf("unknown key version")
	}

	if (bytes.Equal(key.Version, hdwallet.Private) ||
		bytes.Equal(key.Version, hdwallet.TestPrivate)) && !key.IsPrivate {
		return fmt.Errorf("public key with private key version mismatch")
	}

	if (bytes.Equal(key.Version, hdwallet.Public) ||
		bytes.Equal(key.Version, hdwallet.TestPublic)) && key.IsPrivate {
		return fmt.Errorf("private key with public key version mismatch")
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
