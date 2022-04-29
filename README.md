# bip32
Key generation and validation based on BIP-32 spec

## disclaimer
> The use of this tool does not guarantee security or usability for any
> particular purpose. Please review the code and use at your own risk.

## installation
This step assumes you have [Go compiler toolchain](https://go.dev/dl/)
installed on your system.

```bash
go install github.com/kubetrail/bip32@latest
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
```
```text
Enter mnemonic: client sustain stumble prosper pepper maze prison view omit gold organ youth vintage tattoo practice mutual budget excite bubble economy quick conduct spot end
addr: 15MNqcoicDEFkbTLLPYpDmxo5TDTPcNsQv
xpub: xpub661MyMwAqRbcGVTuch4jj6mRfKpjGHJxMCJJgmurP9TQ6hqppgF644VGUevkyCztRpY4PjssirGR5LPpBSyr8BE8GGWev9qGihrfzGB7TpM
xprv: xprv9s21ZrQH143K41PSWfXjMxph7HzErpb6yyNhtPWEpovRDuWgH8vqWGAndNz1oodj88J8JnaNyQMoL2yNKbYWCubfVTF9ux7aiNJCrF8thw7
```

Alternatively, pass mnemonic via STDIN pipe:
```bash
echo client sustain stumble prosper pepper maze prison view omit gold organ youth vintage tattoo practice mutual budget excite bubble economy quick conduct spot end \
  | bip32 gen \
  | jq '.'
```
```json
{
  "addr": "15MNqcoicDEFkbTLLPYpDmxo5TDTPcNsQv",
  "xprv": "xprv9s21ZrQH143K41PSWfXjMxph7HzErpb6yyNhtPWEpovRDuWgH8vqWGAndNz1oodj88J8JnaNyQMoL2yNKbYWCubfVTF9ux7aiNJCrF8thw7",
  "xpub": "xpub661MyMwAqRbcGVTuch4jj6mRfKpjGHJxMCJJgmurP9TQ6hqppgF644VGUevkyCztRpY4PjssirGR5LPpBSyr8BE8GGWev9qGihrfzGB7TpM"
}
```

Or skip mnemonic validation when using an invalid mnemonic
```bash
bip32 gen
```
```text
Enter mnemonic: this is an invalid mnemonic
Error: mnemonic is invalid or please use --skip-mnemonic-validation flag
Usage:
  bip32 gen [flags]

Flags:
      --derivation-path string               Chain Derivation path (default "m")
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
```
```text
Enter mnemonic: this is an invalid mnemonic
addr: 18P6a9ciMaVKg7EizNijgzdHbvHY7mQt7x
xpub: xpub661MyMwAqRbcGR8VcxX4xzHck2vkpqYsZpR2yuLLbSp9bLisq1EGB3XXeuz4xhRG5P92Witd9Qefo6qLaPmAXv8JPfcYwYdQMWU9g1DCAk1
xprv: xprv9s21ZrQH143K3w42Wvz4brLtC16GRNq2CbVSBWvj37HAiYPjHTv1dFD3oecPnmevt1oRsvNJc8pKspRVvq2yoehVwbzkek1nHVwDraPCjvc
```

Use hex seed instead of a mnemonic to generate keys:
```bash
bip32 gen --input-hex-seed
```
```text
Enter seed in hex: 000102030405060708090a0b0c0d0e0f
addr: 1Lq9EYuAopf99uMbbCJoyyMpac5iGSn5cZ
xpub: xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8
xprv: xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi
```

## chain derivation path
A chain derivation path can be provided such as `m`, `m/0`, `m/0H`, `m/0/234` etc.

Generate root keys
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m \
  | jq '.'
```
```json
{
  "addr": "14Zm6hiqcKz4JgBgYzR7yMoZTi7QjRVcfm",
  "xprv": "xprv9s21ZrQH143K48vGoLGRPxgo2JNkJ3J3fqkirQC2zVdk5Dgd5w14S7fRDyHH4dWNHUgkvsvNDCkvAwcSHNAQwhwgNMgZhLtQC63zxwhQmRv",
  "xpub": "xpub661MyMwAqRbcGczjuMoRm6dXaLDEhW1u34gKenbeYqAix21mdUKJyuyu5F1rzYGVxyL6tmgBUAEPrEz92mBXjByMRiJdba9wpnN37RLLAXa"
}
```

Generate first hardened child of root key
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m/0h \
  | jq '.'
```
```json
{
  "addr": "1MRgW1SSkSZVnGwWGm48Q3VYpFhcu7HXey",
  "xprv": "xprv9vB7xEWwNp9kh1wQRfCCQMnZUEG21LpbR9NPCNN1dwhiZkjjeGRnaALmPXCX7SgjFTiCTT6bXes17boXtjq3xLpcDjzEuGLQBM5ohqkao9G",
  "xpub": "xpub69AUMk3qDBi3uW1sXgjCmVjJ2G6WQoYSnNHyzkmdCHEhSZ4tBok37xfFEqHd2AddP56Tqp4o56AePAgCjYdvpW2PU2jbUPFKsav5ut6Ch1m"
}
```

Generate fourth child of third hardened child of root key
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m/2h/3 \
  | jq '.'
```
```json
{
  "addr": "18NvpJbsjLYV3aMGADkD7UVjtyKEG52fW4",
  "xprv": "xprv9xenovaMSsLaNKX8Yz2K1TEZ1b8VymYyji1SL6URcvAMT4EXKQTQxySayFFk2CA6BrhVaBkXWuzTSfNHMEuu1a6gCxZhdc5t9afpx7YRdq4",
  "xpub": "xpub6Be9DS7FHEtsaobbf1ZKNbBHZcxzPEGq6vw38Ut3BFhLKrZfrwmfWmm4pWbqVMyPauABhiVdazRtW9ZBT7fpKR9Pbw5puUAsZaTSRhshGU4"
}
```

## network selection
Bitcoin networks `mainnet` (default) and `testnet` can be selected using `--network` flag.

Example below shows generation for `mainnet` using a hex seed
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m/2h/3 --network=mainnet \
  | jq '.'
```
```json
{
  "addr": "18NvpJbsjLYV3aMGADkD7UVjtyKEG52fW4",
  "xprv": "xprv9xenovaMSsLaNKX8Yz2K1TEZ1b8VymYyji1SL6URcvAMT4EXKQTQxySayFFk2CA6BrhVaBkXWuzTSfNHMEuu1a6gCxZhdc5t9afpx7YRdq4",
  "xpub": "xpub6Be9DS7FHEtsaobbf1ZKNbBHZcxzPEGq6vw38Ut3BFhLKrZfrwmfWmm4pWbqVMyPauABhiVdazRtW9ZBT7fpKR9Pbw5puUAsZaTSRhshGU4"
}
```

Same hex seed used for generation on `testnet` results in different keys.
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m/2h/3 --network=testnet \
  | jq '.'
```
```json
{
  "addr": "mntt7MgrYMyjpgpssniawPi4kxuwCUphN3",
  "xprv": "tprv8fKjbFtgr9Aey8kfDYspB6rYKiYiDHaz5FvZCWtt6teqEeycJmoAUip2tRRQ2ZYQZJEGaHNHgGaFuWv2UTFqpdNGjbn1Hxow4gRFPpNSMdh",
  "xpub": "tpubDC1mjfvvzWrKrbnT7CYQaWWetk4eNcmteZXLV2wBXATE59ENwAckfDRu4Ygd1rvdNoh6KJ1W86FwCL3nqyX4CrbgfXq8azq79Y3rBjRaSFE"
}
```

> Please note that `testnet` keys are currently not tested via any test vector

## derived keys
Child keys can be derived using parent private or public keys and derivation paths. 

For instance first child key can be derived using an extended public key:
```bash
bip32 derive --derivation-path=m/0 xpub661MyMwAqRbcGaPMRtcCZ91tqAFfLKWdoSbr3PLUAFjVB6sTksegmK4NeEjVWiaYG3e3WgEDNsyGGVkghhjjbsksC7z9R3ZoFYWE3oo2tuG
```
```text
xPub: xpub68CAhn9KZZPV6xY9VgRxazetHmch9k7JbERL1G6ZsDdgsCnNgc8bRVt9cah3WmKzYCAZzHqErz8H7amaQwUjj524BejztnhdHyjsYcCFNCL
addr: 1J2cWcYafpAYkPGcR6Wu8ybjEUZL4G1Ukw
pubKeyHex: 024fe10e81436925358a757f41250a3ea69ccdfcb80a15a6e90935ec7ac083fab8
```

Alternatively, when no derivation path is provided the default behavior is to parse the key as it without
any child key derivation. This helps to obtain address if needed
```bash
bip32 derive xpub661MyMwAqRbcGaPMRtcCZ91tqAFfLKWdoSbr3PLUAFjVB6sTksegmK4NeEjVWiaYG3e3WgEDNsyGGVkghhjjbsksC7z9R3ZoFYWE3oo2tuG
```
```text
xPub: xpub661MyMwAqRbcGaPMRtcCZ91tqAFfLKWdoSbr3PLUAFjVB6sTksegmK4NeEjVWiaYG3e3WgEDNsyGGVkghhjjbsksC7z9R3ZoFYWE3oo2tuG
addr: 14nATe3WojN6ojChza1TnzHWhmPUUs7vkX
pubKeyHex: 0367117eab2ef405c130c44f98b1e65be9047dd152811fb06550a66fd4889e2b6a
```

Similarly, a private key can be used as input which further allows us to generate hardened keys:
```bash
bip32 derive --derivation-path=m/0h
```
```text
Enter key: xprv9s21ZrQH143K46JtKs5CC15AH8RAvrnnSDgFEzvrbvCWJJYKDLLSDWjtnxhzPX3A1MMH4i2woK1JZRLzWof4MBVndpUNuWTqJGuMApJNLfN
xPrv: xprv9uCpJGca4rNA4C6psKCLDRNFg9X2HHpXcwxNwSStMh4M7zik3cueWfjGKZFUQoyeMsikVnsrJVXqFA7VmoaT7cFKfjJ7xFMNCXtjaFkcyZ9
xPub: xpub68CAhn9TuDvTGgBHyLjLaZJzEBMWgkYNzAsyjprVv2bKzo3tbADu4U3kApixKhaLxML26dDayM37WFZiAb4hbiJueGakp2BpXxTmAip5TTR
addr: 1NwX8P2KBAhJnMDpicqjibvYS63a2L3gcD
prvKeyWif: L4YxuN6KWdMpL5reFZ4JmyNX4oDYjR8ztPnZti31b25kEsELpraQ
pubKeyHex: 031527dba89db3032ccc4ae594aa4b775c0b99c743069f08b0a5f9e91fcd9d1180
```

Generation of hardened keys is only allowed for parent private keys.

## decode keys
Keys have internal structure such as value, child index, parent signature, fingerprints etc.
```bash
echo xprv9xenovaMSsLaNKX8Yz2K1TEZ1b8VymYyji1SL6URcvAMT4EXKQTQxySayFFk2CA6BrhVaBkXWuzTSfNHMEuu1a6gCxZhdc5t9afpx7YRdq4 \
  | bip32 decode \
  | jq '.'
```
```json
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
  | bip32 decode \
  | jq '.'
```
```json
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
```
```text
Enter key: xpub661MyMwAqRbcGczjuMoRm6dXaLDEhW1u34gKenbeYqAix21mdUKJyuyu5F1rzYGVxyL6tmgBUAEPrEz92mBXjByMRiJdba9wpnN37RLLAXa
key is valid
```

However, this key is invalid.
```bash
bip32 validate
```
```text
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
./test/test.sh
```
```text
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
* [BIP-44 Spec](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki)
* [Key generation online tool](https://iancoleman.io/bip39/#english)
* [Medium post on key generation](https://wolovim.medium.com/ethereum-201-hd-wallets-11d0c93c87f7)
