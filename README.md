# bip32
Key generation and validation based on BIP-32 spec

## disclaimer
> The use of this tool does not guarantee security or usability for any
> particular purpose. Please review the code and use at your own risk.

## installation
This step assumes you have [Go compiler toolchain](https://go.dev/dl/)
installed on your system.

Download this repo to a folder and cd to it. Make sure `go` toolchain
is installed
```bash
go install
```
Add autocompletion for `bash` to your `.bashrc`
```bash
source <(bip32 completion bash)
```

## generate keys
Keys can be generated using a mnemonic (with or without an additional passphrase),
or using a seed.

```bash
bip32 gen
Enter mnemonic: client sustain stumble prosper pepper maze prison view omit gold organ youth vintage tattoo practice mutual budget excite bubble economy quick conduct spot end
pub: xpub661MyMwAqRbcGVTuch4jj6mRfKpjGHJxMCJJgmurP9TQ6hqppgF644VGUevkyCztRpY4PjssirGR5LPpBSyr8BE8GGWev9qGihrfzGB7TpM
prv: xprv9s21ZrQH143K41PSWfXjMxph7HzErpb6yyNhtPWEpovRDuWgH8vqWGAndNz1oodj88J8JnaNyQMoL2yNKbYWCubfVTF9ux7aiNJCrF8thw7
```

Alternatively, pass mnemonic via STDIN pipe:
```bash
echo client sustain stumble prosper pepper maze prison view omit gold organ youth vintage tattoo practice mutual budget excite bubble economy quick conduct spot end | bip32 gen | jq '.'
{
  "prv": "xprv9s21ZrQH143K41PSWfXjMxph7HzErpb6yyNhtPWEpovRDuWgH8vqWGAndNz1oodj88J8JnaNyQMoL2yNKbYWCubfVTF9ux7aiNJCrF8thw7",
  "pub": "xpub661MyMwAqRbcGVTuch4jj6mRfKpjGHJxMCJJgmurP9TQ6hqppgF644VGUevkyCztRpY4PjssirGR5LPpBSyr8BE8GGWev9qGihrfzGB7TpM"
}
```

Or skip mnemonic validation when using an invalid mnemonic
```bash
bip32 gen
Enter mnemonic: this is an invalid mnemonic
Error: mnemonic is invalid or please use --skip-mnemonic-validation flag
Usage:
  bip32 gen [flags]

Flags:
      --chain string               Chain Derivation path (default "m")
  -h, --help                       help for gen
      --input-hex-seed             Treat input as hex seed instead of mnemonic
      --skip-mnemonic-validation   Skip mnemonic validation
      --use-passphrase             Prompt for secret passphrase

Global Flags:
      --config string   config file (default is $HOME/.bip32.yaml)

Error: mnemonic is invalid or please use --skip-mnemonic-validation flag
```

```bash
bip32 gen --skip-mnemonic-validation 
Enter mnemonic: this is an invalid mnemonic
pub: xpub661MyMwAqRbcGR8VcxX4xzHck2vkpqYsZpR2yuLLbSp9bLisq1EGB3XXeuz4xhRG5P92Witd9Qefo6qLaPmAXv8JPfcYwYdQMWU9g1DCAk1
prv: xprv9s21ZrQH143K3w42Wvz4brLtC16GRNq2CbVSBWvj37HAiYPjHTv1dFD3oecPnmevt1oRsvNJc8pKspRVvq2yoehVwbzkek1nHVwDraPCjvc
```

Use hex seed instead of a mnemonic to generate keys:
```bash
bip32 gen --input-hex-seed 
Enter seed in hex: 000102030405060708090a0b0c0d0e0f
pub: xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8
prv: xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi
```

## chain derivation path
A chain derivation path can be provided such as `m`, `m/0`, `m/0H`, `m/0/234` etc.

Generate root keys
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 | bip32 gen --input-hex-seed --chain=m | jq '.'
{
  "prv": "xprv9s21ZrQH143K48vGoLGRPxgo2JNkJ3J3fqkirQC2zVdk5Dgd5w14S7fRDyHH4dWNHUgkvsvNDCkvAwcSHNAQwhwgNMgZhLtQC63zxwhQmRv",
  "pub": "xpub661MyMwAqRbcGczjuMoRm6dXaLDEhW1u34gKenbeYqAix21mdUKJyuyu5F1rzYGVxyL6tmgBUAEPrEz92mBXjByMRiJdba9wpnN37RLLAXa"
}
```

Generate first hardened child of root key
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 | bip32 gen --input-hex-seed --chain=m/0h | jq '.'
{
  "prv": "xprv9vB7xEWwNp9kh1wQRfCCQMnZUEG21LpbR9NPCNN1dwhiZkjjeGRnaALmPXCX7SgjFTiCTT6bXes17boXtjq3xLpcDjzEuGLQBM5ohqkao9G",
  "pub": "xpub69AUMk3qDBi3uW1sXgjCmVjJ2G6WQoYSnNHyzkmdCHEhSZ4tBok37xfFEqHd2AddP56Tqp4o56AePAgCjYdvpW2PU2jbUPFKsav5ut6Ch1m"
}
```

Generate third child of second hardened child of root key
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 | bip32 gen --input-hex-seed --chain=m/2h/3 | jq '.'
{
  "prv": "xprv9xenovaMSsLaNKX8Yz2K1TEZ1b8VymYyji1SL6URcvAMT4EXKQTQxySayFFk2CA6BrhVaBkXWuzTSfNHMEuu1a6gCxZhdc5t9afpx7YRdq4",
  "pub": "xpub6Be9DS7FHEtsaobbf1ZKNbBHZcxzPEGq6vw38Ut3BFhLKrZfrwmfWmm4pWbqVMyPauABhiVdazRtW9ZBT7fpKR9Pbw5puUAsZaTSRhshGU4"
}
```

## decode keys
Keys have internal structure such as value, child index, parent signature, fingerprints etc.
```bash
echo xprv9xenovaMSsLaNKX8Yz2K1TEZ1b8VymYyji1SL6URcvAMT4EXKQTQxySayFFk2CA6BrhVaBkXWuzTSfNHMEuu1a6gCxZhdc5t9afpx7YRdq4 \
  | bip32 decode | jq '.'
{
  "Key": "ZhotNe3SInt/H1Hub0eW31jIj5OAJ8rblvh8BTB7Nlk=",
  "Version": "BIit5A==",
  "ChildNumber": "AAAAAw==",
  "FingerPrint": "/ofVQA==",
  "ChainCode": "gZ3lgh0ICcsIoOfU+Osi1mHwqZdvik8iyhPlBxLHhvc=",
  "Depth": 2,
  "IsPrivate": true
}
```

Similarly, a public key can be decoded as follows:
```bash
echo xpub661MyMwAqRbcGczjuMoRm6dXaLDEhW1u34gKenbeYqAix21mdUKJyuyu5F1rzYGVxyL6tmgBUAEPrEz92mBXjByMRiJdba9wpnN37RLLAXa \
  | bip32 decode | jq '.'
{
  "Key": "Am9v7ckkD2HaqccUS2gqQwo6E2ZXb4QL8tBwEB/LyaAt",
  "Version": "BIiyHg==",
  "ChildNumber": "AAAAAA==",
  "FingerPrint": "AAAAAA==",
  "ChainCode": "0Mih9u3yUAeYw+C1TxtW5F9tA+YHar025eL1QQHkTOY=",
  "Depth": 0,
  "IsPrivate": false
}
```

## key validation
Validity of the keys can be checked (for the most part)

For instance, key below is valid
```bash
bip32 validate 
Enter key: xpub661MyMwAqRbcGczjuMoRm6dXaLDEhW1u34gKenbeYqAix21mdUKJyuyu5F1rzYGVxyL6tmgBUAEPrEz92mBXjByMRiJdba9wpnN37RLLAXa
key is valid
```

However, this key is invalid.
```bash
bip32 validate 
Enter key: xpub661MyMwAqRbcEYS8w7XLSVeEsBXy79zSzH1J8vCdxAZningWLdN3zgtU6LBpB85b3D2yc8sfvZU521AAwdZafEz7mnzBBsz4wKY5fTtTQBm
Error: private key with public key version mismatch
Usage:
  bip32 validate [flags]

Flags:
  -h, --help   help for validate

Global Flags:
      --config string   config file (default is $HOME/.bip32.yaml)

Error: private key with public key version mismatch
```

See [test vector 5](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-5) for
more examples of invalid keys.

## tests
[Following](./test/test.sh) tests pass except for one at the time of writing this doc.
> One of the test cases in test vector 5 related to invalid public key is currently
> not getting detected.
```bash
./test.sh 
https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-1
ok ok ok ok ok ok ok ok ok ok ok ok 

https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-2
ok ok ok ok ok ok ok ok ok ok ok ok 

https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-3
ok ok ok ok 

https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-4
ok ok ok ok ok ok 

https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-5
ok ok ok ok ok ok ok ok ok ok ok ok ok ok ok validation failed
```

## references
* [BIP-32 Spec](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
* [Key generation online tool](https://iancoleman.io/bip39/#english)
* [Medium post on key generation](https://wolovim.medium.com/ethereum-201-hd-wallets-11d0c93c87f7)
