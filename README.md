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
Enter mnemonic:  client sustain stumble prosper pepper maze prison view omit gold organ youth vintage tattoo practice mutual budget excite bubble economy quick conduct spot end
seed: b66ad217f4bd1a7a9889cc86318ef218b72d4ef34d759facae297a65210325ce5de3b106f566be9be4f43a84298fc64cbd8b4419649ed59b3e4a0d5b8e706e7b
xPrv: xprvA3MV9gbvFKuzyqBEfME4SGgybXA451htHMVR8eVwotHPAnpHYenLWirLRBrcK5NDX4Riq89v8weeNY7QNhfY5ZcDkCHqkJHzc5attHZie2Q
xPub: xpub6GLqZC8p5hUJCKFhmNm4oQdi9YzYUURjeaR1w2uZNDpN3b9S6C6b4XApGTLjRxKwbQ8JJP4Hw1GhjLTCpyT1cwnYSozUVmyexLrUfHjBssR
prvKeyWif: L1SbuQzhGFkGDuWQMfyUi71PXtH6BbgVFZjASh3hXiTAH3hrywLP
pubKeyHex: 02b9d3c47c4523de30e0c228586277687b5ed5a49626b5f660c747e82bf936b203
addr: 12u3DEDrEyY7vyRmfVpCjSvNZJeXXKC7Bf
network: mainnet
derivationPath: m/44h/0h/0h/0/0
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
seed: 7ad1ceebbbcef39a1ceee573e6446c32e0bf047ee90daf8f84d0249ea160530fb9a5cfe17fe83cf835707e4e92cb327e53c971362e1b9020b90a3c14f2bf18db
xPrv: xprvA3f5HGHwNfEtudV4P7Djj1rfdM3MTGxyDVzBE8QVv9T1i539j71MyfKUmnw2dSgFFeExPRDYWgvnUnqZJTuCoAAqAf9HnxZAR62rh2bxFSA
xPub: xpub6GeRgmpqD2oC87ZXV8kk69oQBNsqrjgpaiun2Wp7UUyzasNJGeKcXTdxd3wnax6h64jUPovvTZUYcSvVDYRcwMSPXv3jQZiWMM5TmLTfKZ7
prvKeyWif: L3Hk4ckie6rktaBfZEj3eY2czmD67rZd7RjyBTFtnhXrBxha22ZP
pubKeyHex: 02b2cf26a0b6ab3ad6cee283f954b6952ac44da6c04232e0b381e80fa4cbf9a985
addr: 1KtGCn7vzXCbj7kHKU61CxmctDvbwBJZNt
network: mainnet
derivationPath: m/44h/0h/0h/0/0
coinType: btc
```

Alternatively, pass mnemonic via STDIN pipe when a passphrase is not being used:
```bash
echo client sustain stumble prosper pepper maze prison view omit gold organ youth vintage tattoo practice mutual budget excite bubble economy quick conduct spot end \
  | bip32 gen --output-format=json \
  | jq '.'
```
```json
{
  "seed": "b66ad217f4bd1a7a9889cc86318ef218b72d4ef34d759facae297a65210325ce5de3b106f566be9be4f43a84298fc64cbd8b4419649ed59b3e4a0d5b8e706e7b",
  "xPrv": "xprvA3MV9gbvFKuzyqBEfME4SGgybXA451htHMVR8eVwotHPAnpHYenLWirLRBrcK5NDX4Riq89v8weeNY7QNhfY5ZcDkCHqkJHzc5attHZie2Q",
  "xPub": "xpub6GLqZC8p5hUJCKFhmNm4oQdi9YzYUURjeaR1w2uZNDpN3b9S6C6b4XApGTLjRxKwbQ8JJP4Hw1GhjLTCpyT1cwnYSozUVmyexLrUfHjBssR",
  "prvKeyWif": "L1SbuQzhGFkGDuWQMfyUi71PXtH6BbgVFZjASh3hXiTAH3hrywLP",
  "pubKeyHex": "02b9d3c47c4523de30e0c228586277687b5ed5a49626b5f660c747e82bf936b203",
  "addr": "12u3DEDrEyY7vyRmfVpCjSvNZJeXXKC7Bf",
  "network": "mainnet",
  "derivationPath": "m/44h/0h/0h/0/0",
  "coinType": "btc"
}
```

Or skip mnemonic validation when using an invalid mnemonic
```bash
bip32 gen --skip-mnemonic-validation this is an invalid mnemonic
```
```text
seed: bb06e6570ed0b71ac71e4feefeb3a7e2e4cf04ba80a065408150800f86583add8d7ba2ed117444a00f95ca8966ea2e7ff5c8a84b0f5b35a43388d76f0eca043f
xPrv: xprvA3UmZrJ2YeNBdbx2oeNgxNTjhbRrpqAQZr1aeTmToDJLKU2eKCvo4jxScANisRUGWZnFz6Q2aw3fLFwfKSUkNcMvk3sYSqNuaouccauM6zb
xPub: xpub6GU7yMpvP1vUr62VufuhKWQUFdGMEHtFw4wBSrB5MYqKCGMnrkF3cYGvTQG9mi7hNDGo7pfGhksNRCTuBwhsHTRmJdnLKP4ARLtmCAH2J6e
prvKeyWif: L46QdcLXwjJAJrDcot65v1JkruUMfS4aY2jXwz52KR38Ki6zG1FB
pubKeyHex: 02367a9d8a72dc738402cdbc6e0ba426f90b76b895322c3ff4a0c412d9bf1962f2
addr: 1G7VNts7A9GGCus9whX8McPexcfxTCjuUq
network: mainnet
derivationPath: m/44h/0h/0h/0/0
coinType: btc
```

> Please note that the mnemonic is independent of extra white spaces.
> So, following two mnemonics will result in identical keys
```bash
bip32 gen "machine notice lobster hundred mutual creek earth upgrade sea price copy frost"
```
```text
seed: 6b4eb91a3562f7d557e9ba09daa42141d37a9cb55ce1efab3bb728e13a93bb9be963406b6845459644c9f008b3a28a482ed8aafbdf982a600dd2c35b0648031d
xPrv: xprvA3WunfZRAhVjKA3fFFJnPYbjvTraEuEuFyXejamSsbH1QQWxV7z6BKLicGGYhEuHpDTsnFH2NmC53zCLb49qxBZ6wUwFon36JwV9n2PcKjp
xPub: xpub6GWGCB6K1542Xe88MGqnkgYUUVh4eMxkdCTFXyB4RvozHCr72fJLj7fCTYsdk3GxGTjVFeW9SZFAtykc2TV31gT6fwACFrTgzdgg3rJtTrV
prvKeyWif: KwoacAwjNMGs8HdfjRh25D9edqFYuzMZhA9z6vaDGhvHCr9q4YoM
pubKeyHex: 02e0c75a3986eda7c7da1abf844b9042ab5e987cdd9af76325073c638a563544e1
addr: 1CnSnUkwcXSD26nRmHepDaWLJFrCVJgtsK
network: mainnet
derivationPath: m/44h/0h/0h/0/0
coinType: btc
```
```bash
bip32 gen "   machine    notice    lobster    hundred    mutual creek earth upgrade sea price copy frost   "
```
```text
seed: 6b4eb91a3562f7d557e9ba09daa42141d37a9cb55ce1efab3bb728e13a93bb9be963406b6845459644c9f008b3a28a482ed8aafbdf982a600dd2c35b0648031d
xPrv: xprvA3WunfZRAhVjKA3fFFJnPYbjvTraEuEuFyXejamSsbH1QQWxV7z6BKLicGGYhEuHpDTsnFH2NmC53zCLb49qxBZ6wUwFon36JwV9n2PcKjp
xPub: xpub6GWGCB6K1542Xe88MGqnkgYUUVh4eMxkdCTFXyB4RvozHCr72fJLj7fCTYsdk3GxGTjVFeW9SZFAtykc2TV31gT6fwACFrTgzdgg3rJtTrV
prvKeyWif: KwoacAwjNMGs8HdfjRh25D9edqFYuzMZhA9z6vaDGhvHCr9q4YoM
pubKeyHex: 02e0c75a3986eda7c7da1abf844b9042ab5e987cdd9af76325073c638a563544e1
addr: 1CnSnUkwcXSD26nRmHepDaWLJFrCVJgtsK
network: mainnet
derivationPath: m/44h/0h/0h/0/0
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
xPrv: xprvA3WunfZRAhVjKA3fFFJnPYbjvTraEuEuFyXejamSsbH1QQWxV7z6BKLicGGYhEuHpDTsnFH2NmC53zCLb49qxBZ6wUwFon36JwV9n2PcKjp
xPub: xpub6GWGCB6K1542Xe88MGqnkgYUUVh4eMxkdCTFXyB4RvozHCr72fJLj7fCTYsdk3GxGTjVFeW9SZFAtykc2TV31gT6fwACFrTgzdgg3rJtTrV
prvKeyWif: KwoacAwjNMGs8HdfjRh25D9edqFYuzMZhA9z6vaDGhvHCr9q4YoM
pubKeyHex: 02e0c75a3986eda7c7da1abf844b9042ab5e987cdd9af76325073c638a563544e1
addr: 1CnSnUkwcXSD26nRmHepDaWLJFrCVJgtsK
network: mainnet
derivationPath: m/44h/0h/0h/0/0
coinType: btc
```

Use hex seed instead of a mnemonic to generate keys:
```bash
bip32 gen --input-hex-seed
```
```text
Enter seed in hex: 6b4eb91a3562f7d557e9ba09daa42141d37a9cb55ce1efab3bb728e13a93bb9be963406b6845459644c9f008b3a28a482ed8aafbdf982a600dd2c35b0648031d
seed: 6b4eb91a3562f7d557e9ba09daa42141d37a9cb55ce1efab3bb728e13a93bb9be963406b6845459644c9f008b3a28a482ed8aafbdf982a600dd2c35b0648031d
xPrv: xprvA3WunfZRAhVjKA3fFFJnPYbjvTraEuEuFyXejamSsbH1QQWxV7z6BKLicGGYhEuHpDTsnFH2NmC53zCLb49qxBZ6wUwFon36JwV9n2PcKjp
xPub: xpub6GWGCB6K1542Xe88MGqnkgYUUVh4eMxkdCTFXyB4RvozHCr72fJLj7fCTYsdk3GxGTjVFeW9SZFAtykc2TV31gT6fwACFrTgzdgg3rJtTrV
prvKeyWif: KwoacAwjNMGs8HdfjRh25D9edqFYuzMZhA9z6vaDGhvHCr9q4YoM
pubKeyHex: 02e0c75a3986eda7c7da1abf844b9042ab5e987cdd9af76325073c638a563544e1
addr: 1CnSnUkwcXSD26nRmHepDaWLJFrCVJgtsK
network: mainnet
derivationPath: m/44h/0h/0h/0/0
coinType: btc
```

## chain derivation path
A chain derivation path can be provided such as `m`, `m/0`, `m/0H`, `m/0/234` etc.

Generate root keys
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m  --output-format=json \
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
  "derivationPath": "m",
  "coinType": "btc"
}
```

Generate first hardened child of root key
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m/0h --output-format=json \
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
  "derivationPath": "m/0h",
  "coinType": "btc"
}
```

Generate fourth child of third hardened child of root key
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m/2h/3 --output-format=json \
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
  "derivationPath": "m/2h/3",
  "coinType": "btc"
}
```

## network selection
Bitcoin networks `mainnet` (default) and `testnet` can be selected using `--network` flag.

Example below shows generation for `mainnet` using a hex seed
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m/2h/3 --network=mainnet --output-format=json \
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
  "derivationPath": "m/2h/3",
  "coinType": "btc"
}
```

Same hex seed used for generation on `testnet` results in different keys.
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --derivation-path=m/2h/3 --network=testnet --output-format=json \
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
  "derivationPath": "m/2h/3",
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
  | bip32 decode --output-format=json \
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
  | bip32 decode --output-format=json \
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
bip32 decode KzeBeJFoxNctzErrKH9GqBS8VGakySW2bQ33sE43X64aRSjxh7Ei --output-format=json \
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
bip32 decode 028fb7e34b1d9f41c1b7c4a9f93d75a18b4bf0ef9537a270a4018e43214448ac0d --output-format=json \
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
