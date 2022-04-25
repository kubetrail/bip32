package run

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
	"github.com/wemeetagain/go-hdwallet"
)

func Validate(cmd *cobra.Command, args []string) error {
	prompt, err := getPromptStatus()
	if err != nil {
		return fmt.Errorf("failed to get prompt status: %w", err)
	}

	if prompt {
		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter key: "); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	}

	inputReader := bufio.NewReader(cmd.InOrStdin())
	keyString, err := inputReader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read key from input: %w", err)
	}
	keyString = strings.Trim(keyString, "\n")

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

	if prompt {
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), "key is valid"); err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}
	}

	return nil
}
