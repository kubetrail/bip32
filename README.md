# bip32
Bitcoin hierarchically deterministic (HD) key generation based on
[BIP-32 Spec](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)

Please review the spec in above link to learn about how parent and
child keys are generated and derived. In particular, the key generation
depends not only on the mnemonic but also on the derivation path
and network used.

## disclaimer
> The use of this tool does not guarantee security or usability for any
> particular purpose. Please review the code and use at your own risk.

> Don't trust, verify

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

All mnemonic and seed generation can be handled by
[bip39](https://github.com/kubetrail/bip39) CLI tool
```bash
bip39 gen
```
```text
client sustain stumble prosper pepper maze prison view omit gold organ youth vintage tattoo practice mutual budget excite bubble economy quick conduct spot end
```

The default is to generate a 24 word mnemonic sequence. A shorter sequence can be generated
using `--length` flag
```bash
bip39 gen --length=12
```
```text
high human season brick chimney spoil open butter better spice refuse obey
```

The high level sequence of events during a key generation is as follows:
* A random generator is used to generate so-called `entropy`
* `entropy` is used to generate a mnemonic sentence
* Mnemonic sentence is combined with optional passphrase to generate `seed`
* If a mnemonic language is different from english, it is first translated to english
* Seed is used to generate the master key pair (private and public keys)
* These are wrapped into so-called extended keys
* Private extended key is used to generate the private WIF key (more info on WIF below)
* Private WIF key is used to generate a public hex key
* Public hex key is used to generate the address

Last three steps are repeated for child key generation based on derivation path.

Let's generate some keys:
```bash
bip32 gen
```
```text
Enter mnemonic: client sustain stumble prosper pepper maze prison view omit gold organ youth vintage tattoo practice mutual budget excite bubble economy quick conduct spot end
seed: b66ad217f4bd1a7a9889cc86318ef218b72d4ef34d759facae297a65210325ce5de3b106f566be9be4f43a84298fc64cbd8b4419649ed59b3e4a0d5b8e706e7b
xPrv: xprv9s21ZrQH143K41PSWfXjMxph7HzErpb6yyNhtPWEpovRDuWgH8vqWGAndNz1oodj88J8JnaNyQMoL2yNKbYWCubfVTF9ux7aiNJCrF8thw7
xPub: xpub661MyMwAqRbcGVTuch4jj6mRfKpjGHJxMCJJgmurP9TQ6hqppgF644VGUevkyCztRpY4PjssirGR5LPpBSyr8BE8GGWev9qGihrfzGB7TpM
prvKeyWif: L4XzfA8957Wu6foyk1ZYcPJYxv444q6F7gZKKy9BHXMLTVzb1w2N
pubKeyHex: 03528f86fa0f2f0f90011a278a7f96800f9c2016c5ee257b3234af66420320a55e
addr: 19JduQ1dmpeb3V9d1rKfyb8Wkg814GXL83
network: mainnet
coinType: btc
```

`seed` is the hexadecimal seed used for generating master key.

> Please note that the seed is the raw cryptographic secret material
> that by itself is sufficient to generate master key. Unlike a
> mnemonic with passphrase, a seed is not protected using a passphrase

`xPrv` and `xPub` are extended private and public keys, respectively. Extended keys
contain not just the raw key material, but also metadata such as key version and
child derivation details. These are `base58` encoded. Learn more about
[extended keys](https://learnmeabitcoin.com/technical/extended-keys)

`prvKeyWif` is the wallet import format (WIF) compatible private key. WIF keys are
essentially private key material combined with network info, so-called compression
byte and a checksum. Learn more about [WIF private key](https://learnmeabitcoin.com/technical/wif)

`pubKeyHex` is a hexadecimal formatted public key bytes. These are raw bytes without
any network info or key version in it.

`addr` is the bitcoin address for receiving transactions that correspond to the private
key.

`network` and `coinType` indicate the blockchain on which these keys will work.

> Please note that it is not only necessary to store mnemonic, but also necessary
> to back up any passcode, derivation path and network info used for key generation.
> It is impossible to trace the path backwards and infer these inputs from the keys.

Use or `passphrase` is enabled using `--use-passphrase` flag. A `passphrase` is an
additional layer of security on top of mnemonic. It, therefore, protects the mnemonic.

A `passphrase` input is only allowed via STDIN and the user has to confirm it by entering
it again.

```bash
bip32 gen --use-passphrase
```
```text
Enter mnemonic: high human season brick chimney spoil open butter better spice refuse obey
Enter secret passphrase:
Enter secret passphrase again:
seed: 4d23f57c177dd24254d03fa7582cb108bea8201df311b6193922ce2099dd9a551fefbab99b273c23590ebab934ae85084457817b9a75954e9c0f922a1778b1fd
xPrv: xprv9s21ZrQH143K4QTDT4BBvTBefiFGi3sxNZYNpY4eXSfMo5AVDYVAq1jrJeHeDeqUHzMFafzsSdfGcL8mDxGF1bDDUtWpqbh1G8b4RzJYafJ
xPub: xpub661MyMwAqRbcGtXgZ5iCHb8PDk5m7WbojnTycvUG5nCLfsVdm5oRNp4L9uGbpPXeVJ2bpxTaayvaASGds1YNg6Y1zaPvKsSc4X6SEjfPKGs
prvKeyWif: KwyTBygSkeNGvCipEStfLxAtvk7w4vZXBhbBj7TjBjQfGyE5b1mo
pubKeyHex: 0210208edb2b764b3fb2f6d815280aa2e13b9bcc95b9b91cd7edd19be947de7785
addr: 1AXNUgE7avg43bKQejXfdtu5P3vkAGUonA
network: mainnet
coinType: btc
```

Alternatively, pass mnemonic via STDIN pipe when a passphrase is not being used:
```bash
echo client sustain stumble prosper pepper maze prison view omit gold organ youth vintage tattoo practice mutual budget excite bubble economy quick conduct spot end \
  | bip32 gen \
  | jq '.'
```
```json
{
  "seed": "b66ad217f4bd1a7a9889cc86318ef218b72d4ef34d759facae297a65210325ce5de3b106f566be9be4f43a84298fc64cbd8b4419649ed59b3e4a0d5b8e706e7b",
  "xPrv": "xprv9s21ZrQH143K41PSWfXjMxph7HzErpb6yyNhtPWEpovRDuWgH8vqWGAndNz1oodj88J8JnaNyQMoL2yNKbYWCubfVTF9ux7aiNJCrF8thw7",
  "xPub": "xpub661MyMwAqRbcGVTuch4jj6mRfKpjGHJxMCJJgmurP9TQ6hqppgF644VGUevkyCztRpY4PjssirGR5LPpBSyr8BE8GGWev9qGihrfzGB7TpM",
  "prvKeyWif": "L4XzfA8957Wu6foyk1ZYcPJYxv444q6F7gZKKy9BHXMLTVzb1w2N",
  "pubKeyHex": "03528f86fa0f2f0f90011a278a7f96800f9c2016c5ee257b3234af66420320a55e",
  "addr": "19JduQ1dmpeb3V9d1rKfyb8Wkg814GXL83",
  "network": "mainnet",
  "coinType": "btc"
}
```

Or skip mnemonic validation when using an invalid mnemonic
```bash
bip32 gen --skip-mnemonic-validation this is an invalid mnemonic
```
```text
seed: bb06e6570ed0b71ac71e4feefeb3a7e2e4cf04ba80a065408150800f86583add8d7ba2ed117444a00f95ca8966ea2e7ff5c8a84b0f5b35a43388d76f0eca043f
xPrv: xprv9s21ZrQH143K3w42Wvz4brLtC16GRNq2CbVSBWvj37HAiYPjHTv1dFD3oecPnmevt1oRsvNJc8pKspRVvq2yoehVwbzkek1nHVwDraPCjvc
xPub: xpub661MyMwAqRbcGR8VcxX4xzHck2vkpqYsZpR2yuLLbSp9bLisq1EGB3XXeuz4xhRG5P92Witd9Qefo6qLaPmAXv8JPfcYwYdQMWU9g1DCAk1
prvKeyWif: L3USrmAFQ1Bc7VoPH4hLAiDoBGJpMTVsWpa78yaiDq7EdG74jTPN
pubKeyHex: 02e7cfa509815cfd952b4df4ca677dc1f18551fef158b850703aafd4093c066c67
addr: 1CcNBXdpmBMRYaNz8Czgn14o584zjaZXeK
network: mainnet
coinType: btc
```

> Please note that the mnemonic is independent of extra white spaces.
> So, following two mnemonics will result in identical keys
```bash
bip32 gen "machine notice lobster hundred mutual creek earth upgrade sea price copy frost"
```
```text
seed: 6b4eb91a3562f7d557e9ba09daa42141d37a9cb55ce1efab3bb728e13a93bb9be963406b6845459644c9f008b3a28a482ed8aafbdf982a600dd2c35b0648031d
xPrv: xprv9s21ZrQH143K2krDy28uPZgHhFvcxpsrVzfkTBE8Ai5qqXPakRH1CTnqMWADFkzccesFYGoAnJgztWiojmELqxmqzwnHxsyV5PNckwnhqFR
xPub: xpub661MyMwAqRbcFEvh53fukhd2FHm7NHbhsDbMFZdjj3cpiKijHxbFkG7KCp7y89qHM2Zk5hfcEQBL8ETt25irqEcFSx3V3wZKh4qb9ckLFSp
prvKeyWif: KzvGFD6ofxw6rcYeeCZqut2y4Gq6oiA5YbQLJDPqydrc1Jn2XNM9
pubKeyHex: 03f05b80b3f8903d97cea3ebf7b7496ea368703a35ac512655a74a1432560577b3
addr: 12EDNPj2VeBoArRYqB1Fb9LpE8ufeXfAr3
network: mainnet
coinType: btc
```
```bash
bip32 gen "   machine    notice    lobster    hundred    mutual creek earth upgrade sea price copy frost   "
```
```text
seed: 6b4eb91a3562f7d557e9ba09daa42141d37a9cb55ce1efab3bb728e13a93bb9be963406b6845459644c9f008b3a28a482ed8aafbdf982a600dd2c35b0648031d
xPrv: xprv9s21ZrQH143K2krDy28uPZgHhFvcxpsrVzfkTBE8Ai5qqXPakRH1CTnqMWADFkzccesFYGoAnJgztWiojmELqxmqzwnHxsyV5PNckwnhqFR
xPub: xpub661MyMwAqRbcFEvh53fukhd2FHm7NHbhsDbMFZdjj3cpiKijHxbFkG7KCp7y89qHM2Zk5hfcEQBL8ETt25irqEcFSx3V3wZKh4qb9ckLFSp
prvKeyWif: KzvGFD6ofxw6rcYeeCZqut2y4Gq6oiA5YbQLJDPqydrc1Jn2XNM9
pubKeyHex: 03f05b80b3f8903d97cea3ebf7b7496ea368703a35ac512655a74a1432560577b3
addr: 12EDNPj2VeBoArRYqB1Fb9LpE8ufeXfAr3
network: mainnet
coinType: btc
```

Similarly, mnemonics from different languages are treated the same if their underlying `entropy` is the same.
To illustrate this let's translate a mnemonic from English to Japanese and use the Japanese version to generate
the keys
```bash
bip39 translate --from-language=english --to-language=japanese machine notice lobster hundred mutual creek earth upgrade sea price copy frost
```
```text
たいちょう ちょさくけん ぞんび すごい だんわ きそう けとばす やおや はこぶ とめる きおち さばく
```

Now generating the keys using Japanese version of mnemonic results in same keys as before
```bash
bip32 gen --mnemonic-language=Japanese たいちょう ちょさくけん ぞんび すごい だんわ きそう けとばす やおや はこぶ とめる きおち さばく
```
```text
seed: 6b4eb91a3562f7d557e9ba09daa42141d37a9cb55ce1efab3bb728e13a93bb9be963406b6845459644c9f008b3a28a482ed8aafbdf982a600dd2c35b0648031d
xPrv: xprv9s21ZrQH143K2krDy28uPZgHhFvcxpsrVzfkTBE8Ai5qqXPakRH1CTnqMWADFkzccesFYGoAnJgztWiojmELqxmqzwnHxsyV5PNckwnhqFR
xPub: xpub661MyMwAqRbcFEvh53fukhd2FHm7NHbhsDbMFZdjj3cpiKijHxbFkG7KCp7y89qHM2Zk5hfcEQBL8ETt25irqEcFSx3V3wZKh4qb9ckLFSp
prvKeyWif: KzvGFD6ofxw6rcYeeCZqut2y4Gq6oiA5YbQLJDPqydrc1Jn2XNM9
pubKeyHex: 03f05b80b3f8903d97cea3ebf7b7496ea368703a35ac512655a74a1432560577b3
addr: 12EDNPj2VeBoArRYqB1Fb9LpE8ufeXfAr3
network: mainnet
coinType: btc
```

Use hex seed instead of a mnemonic to generate keys:
```bash
bip32 gen --input-hex-seed
```
```text
Enter seed in hex: 6b4eb91a3562f7d557e9ba09daa42141d37a9cb55ce1efab3bb728e13a93bb9be963406b6845459644c9f008b3a28a482ed8aafbdf982a600dd2c35b0648031d
seed: 6b4eb91a3562f7d557e9ba09daa42141d37a9cb55ce1efab3bb728e13a93bb9be963406b6845459644c9f008b3a28a482ed8aafbdf982a600dd2c35b0648031d
xPrv: xprv9s21ZrQH143K2krDy28uPZgHhFvcxpsrVzfkTBE8Ai5qqXPakRH1CTnqMWADFkzccesFYGoAnJgztWiojmELqxmqzwnHxsyV5PNckwnhqFR
xPub: xpub661MyMwAqRbcFEvh53fukhd2FHm7NHbhsDbMFZdjj3cpiKijHxbFkG7KCp7y89qHM2Zk5hfcEQBL8ETt25irqEcFSx3V3wZKh4qb9ckLFSp
prvKeyWif: KzvGFD6ofxw6rcYeeCZqut2y4Gq6oiA5YbQLJDPqydrc1Jn2XNM9
pubKeyHex: 03f05b80b3f8903d97cea3ebf7b7496ea368703a35ac512655a74a1432560577b3
addr: 12EDNPj2VeBoArRYqB1Fb9LpE8ufeXfAr3
network: mainnet
coinType: btc
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
  "seed": "3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678",
  "xPrv": "xprv9s21ZrQH143K48vGoLGRPxgo2JNkJ3J3fqkirQC2zVdk5Dgd5w14S7fRDyHH4dWNHUgkvsvNDCkvAwcSHNAQwhwgNMgZhLtQC63zxwhQmRv",
  "xPub": "xpub661MyMwAqRbcGczjuMoRm6dXaLDEhW1u34gKenbeYqAix21mdUKJyuyu5F1rzYGVxyL6tmgBUAEPrEz92mBXjByMRiJdba9wpnN37RLLAXa",
  "prvKeyWif": "KwrAWXgyy1L75ZBRp1PzHj2aWBoYcddgrEMfF6iBJFuw8adwRNLu",
  "pubKeyHex": "026f6fedc9240f61daa9c7144b682a430a3a1366576f840bf2d070101fcbc9a02d",
  "addr": "1GpWFBBE37FQumRkrVUL6HB1bqSWCuYsKt",
  "network": "mainnet",
  "coinType": "btc"
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
  "seed": "3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678",
  "xPrv": "xprv9vB7xEWwNp9kh1wQRfCCQMnZUEG21LpbR9NPCNN1dwhiZkjjeGRnaALmPXCX7SgjFTiCTT6bXes17boXtjq3xLpcDjzEuGLQBM5ohqkao9G",
  "xPub": "xpub69AUMk3qDBi3uW1sXgjCmVjJ2G6WQoYSnNHyzkmdCHEhSZ4tBok37xfFEqHd2AddP56Tqp4o56AePAgCjYdvpW2PU2jbUPFKsav5ut6Ch1m",
  "prvKeyWif": "KwFMsuZ3pmk7ebtbTiPirTpdcPkS6wvnSazU3bvixwiCw1bNQLhG",
  "pubKeyHex": "039382d2b6003446792d2917f7ac4b3edf079a1a94dd4eb010dc25109dda680a9d",
  "addr": "1KvwpccVR6CsN3ve2LZpxkSZ5od5262b75",
  "network": "mainnet",
  "coinType": "btc"
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
  "seed": "3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678",
  "xPrv": "xprv9xenovaMSsLaNKX8Yz2K1TEZ1b8VymYyji1SL6URcvAMT4EXKQTQxySayFFk2CA6BrhVaBkXWuzTSfNHMEuu1a6gCxZhdc5t9afpx7YRdq4",
  "xPub": "xpub6Be9DS7FHEtsaobbf1ZKNbBHZcxzPEGq6vw38Ut3BFhLKrZfrwmfWmm4pWbqVMyPauABhiVdazRtW9ZBT7fpKR9Pbw5puUAsZaTSRhshGU4",
  "prvKeyWif": "KzeBeJFoxNctzErrKH9GqBS8VGakySW2bQ33sE43X64aRSjxh7Ei",
  "pubKeyHex": "028fb7e34b1d9f41c1b7c4a9f93d75a18b4bf0ef9537a270a4018e43214448ac0d",
  "addr": "19q3NgbmofP5eB62zNawU68pxVwopdHDZ2",
  "network": "mainnet",
  "coinType": "btc"
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
  "seed": "3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678",
  "xPrv": "xprv9xenovaMSsLaNKX8Yz2K1TEZ1b8VymYyji1SL6URcvAMT4EXKQTQxySayFFk2CA6BrhVaBkXWuzTSfNHMEuu1a6gCxZhdc5t9afpx7YRdq4",
  "xPub": "xpub6Be9DS7FHEtsaobbf1ZKNbBHZcxzPEGq6vw38Ut3BFhLKrZfrwmfWmm4pWbqVMyPauABhiVdazRtW9ZBT7fpKR9Pbw5puUAsZaTSRhshGU4",
  "prvKeyWif": "KzeBeJFoxNctzErrKH9GqBS8VGakySW2bQ33sE43X64aRSjxh7Ei",
  "pubKeyHex": "028fb7e34b1d9f41c1b7c4a9f93d75a18b4bf0ef9537a270a4018e43214448ac0d",
  "addr": "19q3NgbmofP5eB62zNawU68pxVwopdHDZ2",
  "network": "mainnet",
  "coinType": "btc"
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
  "seed": "3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678",
  "xPrv": "tprv8fKjbFtgr9Aey8kfDYspB6rYKiYiDHaz5FvZCWtt6teqEeycJmoAUip2tRRQ2ZYQZJEGaHNHgGaFuWv2UTFqpdNGjbn1Hxow4gRFPpNSMdh",
  "xPub": "tpubDC1mjfvvzWrKrbnT7CYQaWWetk4eNcmteZXLV2wBXATE59ENwAckfDRu4Ygd1rvdNoh6KJ1W86FwCL3nqyX4CrbgfXq8azq79Y3rBjRaSFE",
  "prvKeyWif": "cR1B7DFfPSKA9gL7hgxQCVwC7VtAdtbifSBWyeWZ2CiagBqB5jAk",
  "pubKeyHex": "028fb7e34b1d9f41c1b7c4a9f93d75a18b4bf0ef9537a270a4018e43214448ac0d",
  "addr": "mpLzfjgkcgpLRHZehwZKJ1M9pVYWjExGEV",
  "network": "testnet",
  "coinType": "btc"
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
pubKeyHex: 024fe10e81436925358a757f41250a3ea69ccdfcb80a15a6e90935ec7ac083fab8
addr: 1J2cWcYafpAYkPGcR6Wu8ybjEUZL4G1Ukw
network: mainnet
coinType: btc
```

Alternatively, when no derivation path is provided the default behavior is to parse the key as it without
any child key derivation. This helps to obtain address if needed
```bash
bip32 derive xpub661MyMwAqRbcGaPMRtcCZ91tqAFfLKWdoSbr3PLUAFjVB6sTksegmK4NeEjVWiaYG3e3WgEDNsyGGVkghhjjbsksC7z9R3ZoFYWE3oo2tuG
```
```text
xPub: xpub661MyMwAqRbcGaPMRtcCZ91tqAFfLKWdoSbr3PLUAFjVB6sTksegmK4NeEjVWiaYG3e3WgEDNsyGGVkghhjjbsksC7z9R3ZoFYWE3oo2tuG
pubKeyHex: 0367117eab2ef405c130c44f98b1e65be9047dd152811fb06550a66fd4889e2b6a
addr: 14nATe3WojN6ojChza1TnzHWhmPUUs7vkX
network: mainnet
coinType: btc
```

Similarly, a private key can be used as input which further allows us to generate hardened keys:
```bash
bip32 derive --derivation-path=m/0h
```
```text
Enter key: xprv9s21ZrQH143K46JtKs5CC15AH8RAvrnnSDgFEzvrbvCWJJYKDLLSDWjtnxhzPX3A1MMH4i2woK1JZRLzWof4MBVndpUNuWTqJGuMApJNLfN
xPrv: xprv9uCpJGca4rNA4C6psKCLDRNFg9X2HHpXcwxNwSStMh4M7zik3cueWfjGKZFUQoyeMsikVnsrJVXqFA7VmoaT7cFKfjJ7xFMNCXtjaFkcyZ9
xPub: xpub68CAhn9TuDvTGgBHyLjLaZJzEBMWgkYNzAsyjprVv2bKzo3tbADu4U3kApixKhaLxML26dDayM37WFZiAb4hbiJueGakp2BpXxTmAip5TTR
prvKeyWif: L4YxuN6KWdMpL5reFZ4JmyNX4oDYjR8ztPnZti31b25kEsELpraQ
pubKeyHex: 031527dba89db3032ccc4ae594aa4b775c0b99c743069f08b0a5f9e91fcd9d1180
addr: 1NwX8P2KBAhJnMDpicqjibvYS63a2L3gcD
network: mainnet
coinType: btc
```

Generation of hardened keys is only allowed for parent private keys.

## decode keys
While `derive` command is used for deriving child keys, `decode` works with a variety of key inputs:
* Extended keys (both private and public)
* Private WIF keys
* Public HEX keys
* 
```bash
echo xprv9xenovaMSsLaNKX8Yz2K1TEZ1b8VymYyji1SL6URcvAMT4EXKQTQxySayFFk2CA6BrhVaBkXWuzTSfNHMEuu1a6gCxZhdc5t9afpx7YRdq4 \
  | bip32 decode \
  | jq '.'
```
```json
{
  "xPrv": "xprv9xenovaMSsLaNKX8Yz2K1TEZ1b8VymYyji1SL6URcvAMT4EXKQTQxySayFFk2CA6BrhVaBkXWuzTSfNHMEuu1a6gCxZhdc5t9afpx7YRdq4",
  "xPub": "xpub6Be9DS7FHEtsaobbf1ZKNbBHZcxzPEGq6vw38Ut3BFhLKrZfrwmfWmm4pWbqVMyPauABhiVdazRtW9ZBT7fpKR9Pbw5puUAsZaTSRhshGU4",
  "prvKeyWif": "KzeBeJFoxNctzErrKH9GqBS8VGakySW2bQ33sE43X64aRSjxh7Ei",
  "pubKeyHex": "028fb7e34b1d9f41c1b7c4a9f93d75a18b4bf0ef9537a270a4018e43214448ac0d",
  "addr": "19q3NgbmofP5eB62zNawU68pxVwopdHDZ2",
  "network": "mainnet",
  "coinType": "btc"
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
  "xPub": "xpub661MyMwAqRbcGczjuMoRm6dXaLDEhW1u34gKenbeYqAix21mdUKJyuyu5F1rzYGVxyL6tmgBUAEPrEz92mBXjByMRiJdba9wpnN37RLLAXa",
  "pubKeyHex": "026f6fedc9240f61daa9c7144b682a430a3a1366576f840bf2d070101fcbc9a02d",
  "addr": "1GpWFBBE37FQumRkrVUL6HB1bqSWCuYsKt",
  "network": "mainnet",
  "coinType": "btc"
}
```

Decode private WIF key:
```bash
bip32 decode KzeBeJFoxNctzErrKH9GqBS8VGakySW2bQ33sE43X64aRSjxh7Ei \
  | jq '.'
```
```json
{
  "prvKeyWif": "KzeBeJFoxNctzErrKH9GqBS8VGakySW2bQ33sE43X64aRSjxh7Ei",
  "pubKeyHex": "028fb7e34b1d9f41c1b7c4a9f93d75a18b4bf0ef9537a270a4018e43214448ac0d",
  "addr": "19q3NgbmofP5eB62zNawU68pxVwopdHDZ2",
  "network": "mainnet",
  "coinType": "btc"
}
```

Decode public HEX key.

> Please note that assumption is made about the network to be
> `mainnet` for cointype of BTC since public hex keys do not
> encode key versions and therefore it is not possible to
> decode which network these keys were originally meant for.

```bash
bip32 decode 028fb7e34b1d9f41c1b7c4a9f93d75a18b4bf0ef9537a270a4018e43214448ac0d \
  | jq '.'
```
```json
{
  "pubKeyHex": "028fb7e34b1d9f41c1b7c4a9f93d75a18b4bf0ef9537a270a4018e43214448ac0d",
  "addr": "19q3NgbmofP5eB62zNawU68pxVwopdHDZ2",
  "network": "mainnet",
  "coinType": "btc"
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
* [Extended keys](https://learnmeabitcoin.com/technical/extended-keys)
* [Checksum](https://learnmeabitcoin.com/technical/checksum)
* [WIF private key](https://learnmeabitcoin.com/technical/wif)
* [Key explorer](http://bip32.org/)
* [Key generator](https://www.bitaddress.org/)
* [Bitcoin the hard way](http://www.righto.com/2014/02/bitcoins-hard-way-using-raw-bitcoin.html)
* [btckeygen, a bitcoin key generator CLI](https://github.com/modood/btckeygen)
