# bip32
Generate Bitcoin wallet keys using mnemonic phrase.

Bitcoin key generation has changed over the years and has nuances that
are covered in following specs:
* [BIP-32 Spec](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
* [BIP-39 Spec](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki)
* [BIP-44 Spec](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki)
* [BIP-49 Spec](https://github.com/bitcoin/bips/blob/master/bip-0049.mediawiki)
* [BIP-84 Spec](https://github.com/bitcoin/bips/blob/master/bip-0084.mediawiki)

In particular, you may have seen bitcoin addresses in one of the following
forms:
* 1H1RmnHvTsgLBxvGke9XDPP74w7K9uVT9c, i.e., starting with `1`
* 3LZoyoAc9TDXLGrGoWsmv3dtJNbBwm1HKz, i.e., starting with `3`
* bc1q87e8lc7523q67qqjdnah6yk4g547nh50vavhlc, i.e. starting with `bc1`

These addresses can all be traced back to single mnemonic sentence and are
therefore, "derived" from a master key. In particular, addresses starting with
`1` correspond to the spec `BIP-44`, whereas addresses starting with `3`
correspond to the spec `BIP-49` and the ones starting with `bc1` correspond
to the spec `BIP84`.

More info on how to go about generating these keys is below.

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

## overview of key generation steps
Keys can be generated using a mnemonic sentence, which is collection,
typically of 12 to 24, random words.

All mnemonic and seed generation can be handled by
[bip39](https://github.com/kubetrail/bip39) CLI tool

For example, a 12 word mnemonic can be generated as follows:
```bash
bip39 gen --length=12
```
```text
surround disagree build occur pluck main ignore define hurdle excess chicken gold
```

Export mnemonic as an environment variable for use in next steps.
```bash
export MNEMONIC="surround disagree build occur pluck main ignore define hurdle excess chicken gold"
```

> Please note that mnemonic generation occurs using a cryptographic source
> of randomness that is internal to the tool. Hence, security of such
> mnemonic is inherently limited by the security of the random generator used.

> All keys are derived from mnemonic in a deterministic way, therefore, the
> security of Bitcoin keys is essentially limited by security of mnemonic.

> Needless to say, please only use the mnemonic shown above for experimental
> purposes and not for performing actual transactions. Generate a new mnemonic
> and do not share with anyone!

The high level sequence of events during a key generation is as follows:
* A random generator is used to generate so-called `entropy`
* `entropy` is used to generate a mnemonic sentence
* Mnemonic sentence is combined with optional passphrase to generate a `seed`
* If a mnemonic language is different from english, it is first translated to english
* Seed is used to generate the master `ECDSA` key pair (private and public keys)
* These are wrapped into so-called extended keys
* Private extended key is used to generate the private WIF key (more info on WIF below)
* Private WIF key is used to generate a public hex key
* Public hex key is used to generate the address
* Address type governs the type of address generated, i.e., either staring with `1`, `3`,
or `bc1`. It also decided the prefix of extended keys, which can be either `xprv`, `yprv`
or `zprv`
* Similarly, choice of network, i.e., either `mainnet` or `testnet` also governs the
form of extended keys.

## key generation
Assuming you stored mnemonic in an environment variable, private key and
public address can be generated using a few defaults as follows:
```bash
bip32 gen ${MNEMONIC}
```
```yaml
prvKeyWif: Ky7kJQEFQDCRhShHaZs7TSCEa1UqGhVZB6BhXUHf3T9pAG6q7987
addr: 1MJ9PojuE1rA1E8wtrdQXjxaqZdsgddhoh
```

In order to explore the keys better, let's output all keys that are
part of this step. This is done using `--show-all-keys` flag

```bash
bip32 gen --show-all-keys ${MNEMONIC}
```
```yaml
seed: bf4848ce9688a8fc5e149aa038c454885a6727e0c7de50754cef7e506fb3d8a80d4c92b7cd77da542b1b6764a2513899c311ab2c7d9d192546e2d6442ac99e11
xPrv: xprvA4HQhA4Br6afRrswRBRydnSPqHGAVFptd3iiMp8kuMMsuso4vUuPSb91AizescB7NwS8uHijxBK4L1J5nJj98be4cSTKthrPraRZiFttSPj
xPub: xpub6HGm6fb5gU8xeLxQXCxyzvP8PK6etiYjzGeKACYNTgtrng8DU2DdzPTV1yvcwsjZ9o1UAUYq39RhLXBmJu66xMmAGGnDRGb4iLkP99rLmL5
pubKeyHex: 022b70459564b65102394e088bdb68f8a80d386939e39319363377b00f23b21cc4
prvKeyWif: Ky7kJQEFQDCRhShHaZs7TSCEa1UqGhVZB6BhXUHf3T9pAG6q7987
addr: 1MJ9PojuE1rA1E8wtrdQXjxaqZdsgddhoh
addrType: legacy
derivationPath: m/44h/0h/0h/0/0
coinType: btc
network: mainnet
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

`addr` is the legacy bitcoin address for receiving transactions that correspond to the private
key. `addrType` indicates type of the address, i.e., either `legacy`, `segwit-compatibe` or
`segwit-native` etc.

`network` and `coinType` indicate the blockchain on which these keys will work.

A secret passphrase can be additionally used in conjunction with mnemonic to apply an
additional layer of security. Use of `passphrase` is enabled using `--use-passphrase` 
flag.

A `passphrase` input is only allowed via `STDIN` and the user has to confirm it by entering
it again.

```bash
bip32 gen --use-passphrase ${MNEMONIC}
```
```yaml
Enter secret passphrase:
Enter secret passphrase again:
prvKeyWif: KzjwJtryjybsJjVvse22oG7f1pfQbYcj2E8gxbM6bivai3pLrczh
addr: 181mk5LncqeKPEz1Z1KihzykXD5XKLTTr2
```

### address types
As mentioned before, three types of addresses can be derived from mnemonic.

For instance `legacy` `P2PKH` addresses can be generated as follows, which is also the
default:
```bash
bip32 gen --addr-type=legacy ${MNEMONIC}
```
```yaml
prvKeyWif: Ky7kJQEFQDCRhShHaZs7TSCEa1UqGhVZB6BhXUHf3T9pAG6q7987
addr: 1MJ9PojuE1rA1E8wtrdQXjxaqZdsgddhoh
```

`segwit` compatible `P2SH` addresses can be generated as follows:
```bash
bip32 gen --addr-type=segwit-compatible ${MNEMONIC}
```
```yaml
prvKeyWif: L5Nx5ePGYjN2TzoadVXorDQXrchWEtwuDQEnsC4SSztz9b4tcCWQ
addr: 37vznvAgCmaKERDZmYaw3X4ArHracgVUfa
```

And finally, `segwit` native `Bech32` addresses can be generated as follows:
```bash
bip32 gen --addr-type=segwit-native ${MNEMONIC}
```
```yaml
prvKeyWif: L125eeMMWH93ZVb6NXEpQ5qXYPiKDF6Qy3AHE8UcBmh6SdcBvEfQ
addr: bc1qsah54m5u94ktfymcv4jf656rqnu9dxnuhcjvx8
```

### verify public addresses using external wallet app
At this point, it might be a good idea to verify these addresses
match those produced by external wallet apps such as
[Mycelium](https://play.google.com/store/apps/details?id=com.mycelium.wallet)

After installing the app, select the option to restore the wallet from backup
instead of creating a new one.

Select 12 words as mnemonic length option and then enter your mnemonic that was
used so far.

Once the master key generation is done, go to the `Balance` tab and tap on
qrcode to rotate through different key types and verify they are the same
as generated in previous step.

## chain derivation path
Keys generated so far used defaults for the so-called `derivation-path`, which
governs how child keys are to be derived from the master key. It is a commandline
flag `--derivation-path` and a few examples of such paths are as follows:
* m
* m/44h/0h/0h
* m/49h/0h/0h/0/0
* m/49'/0'/0'/0/0 and so on

`m` stands for master key and when there is nothing else after `m`, no key derivation
is performed and master key itself is used to derive the public address.

The default value of `derivation-path` is `auto`, which signifies that the path will be
chosen according to the `addr-type` flag. In particular, below is the mapping between
`addr-type` and the `derivation-path`:
```text
addr type                         derivation path
--------------------------------------------------
legacy                            m/44'/0'/0'/0/0
segwit compatible P2SH            m/49'/0'/0'/0/0
segwit native Bech32              m/84'/0'/0'/0/0
--------------------------------------------------
```

The way to read these derivation paths is as follows:
> m / purpose / coin type / account / change / index

where, `'` or `h` or `H` denotes so-called `hardened` value. Read more about it
on BIP-44 spec.

Thus, master private extended keys correspond to the derivation paths that do not
account `change` and `index` values, such as `m/44'/0'/0'` or `m/49'/0'/0'` etc.

Let's generate private extended keys for different address types:

Make sure to use purpose value of `44` with address type of `legacy`:
```bash
bip32 gen \
  --output-format=json \
  --addr-type=legacy \
  --show-all-keys \
  --derivation-path=m/44h/0h/0h \
  ${MNEMONIC} \
  | jq '{xPub: .xPub, xPrv: .xPrv}'
```
```json
{
  "xPub": "xpub6D2evtM5oHGd4MxbT4oDhrgKAVXWgoRBLPQgPScQYS1LN1NUqaQ5jJ4azZxfbiUh9EUDurfuaZkFewCoNYyzXm84BDMp2PbS9mvcFwtvfLm",
  "xPrv": "xprv9z3JXNpBxuiKqst8M3GDLijacTh2HLhKyAV5b4Cnz6UMVD3LJ35qBVk79JpMs6XoygJoaEmVd4vDNQrsfdBaXngkCrvvmSSLiExb2usfLG1"
}
```

Please use purpose of `49` when using address type `segwit-compatible`
```bash
bip32 gen \
  --output-format=json \
  --addr-type=segwit-compatible \
  --show-all-keys \
  --derivation-path=m/49h/0h/0h \
  ${MNEMONIC} \
  | jq '{xPub: .xPub, xPrv: .xPrv}'
```
```json
{
  "xPub": "ypub6X1LhQqze25nBqKDhxpMUYueRM6VdALocQomFPPEtryESNvQ9d1mo8VsqyWMsZiHszffJXHPfgdBv2tS38Sk2FB7hbjYEMCp3TvZSr2CDaj",
  "xPrv": "yprvAJ1zHuK6oeXUyMEkbwHM7QxusKG1DhcxFBtASzydLXSFZabFc5hXFLBPzij6CSnXoiPawEqBxYV8ZKSVuBoCzCEhZ3Rksjrc6WSJggMaTpa"
}
```

Similarly, please use purpose of `84` when using address type of `segwit-native`
```bash
bip32 gen \
  --output-format=json \
  --addr-type=segwit-native \
  --show-all-keys \
  --derivation-path=m/84h/0h/0h \
  ${MNEMONIC} \
  | jq '{xPub: .xPub, xPrv: .xPrv}'
```
```json
{
  "xPub": "zpub6rWTAsb9uWGq6RBefNvtyMtXdscHxQb5xLPPmyPN9p43978FzZqmxsPQ7eYu8qJb6xkPnDfP9ciZC5hoNv6ZoFKspuDgHuR1ad5SmSZkw7R",
  "xPrv": "zprvAdX6mN4G58iXsw7BZMPtcDwo5qmoYwsEb7TnyaykbUX4GJo7T2XXR54vGM6DbnNMM9f9WEpsGhXABP2eyvNwQXzuwywQJH4opQNEwC1ZAmX"
}
```

### verify root extended keys using external wallet app
Again, please ensure that these are indeed the private key values that you can export out of
an external wallet app such as `Mycelium`.

## network selection
Bitcoin networks `mainnet` (default) and `testnet` can be selected using `--network` flag.

Example below shows generation for `mainnet` using a hex seed
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --show-all-keys --network=mainnet --output-format=json \
  | jq '.'
```
```json
{
  "seed": "3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678",
  "xPrv": "xprvA2X4DRgATz4wyfc3rjYT2o24MDk2s6hDsU6Xpzj7N5zqi2mGwkcojjDhmz8acg2zHDHWkL4rdavReSrJTWJNtsnmcU23qKCJeFixAdy27Y8",
  "xPub": "xpub6FWQcwD4JMdFC9gWxm5TPvxnuFaXGZR5Eh28dP8ivRXpaq6RVHw4HXYBdFL69c4YbLETQnxh2Qo4j4VKDdYsPuttZNMC3dauvKKhqNLzWRb",
  "pubKeyHex": "024bf0776f0b553c9acfe6d47206849012eae6c81bd8e9fc645b4e84cabcfddc76",
  "prvKeyWif": "Ky2KivUN8WQ9Rdd9FyacLPmfqWy3mUMq4uqJPnXNoctFu15A79BH",
  "addr": "16XvtkkJio36xFbuXjLMvEGRuLFB15ScKY",
  "addrType": "legacy",
  "derivationPath": "m/44h/0h/0h/0/0",
  "coinType": "btc",
  "network": "mainnet"
}
```

Same hex seed used for generation on `testnet` results in different keys.
```bash
echo 3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678 \
  | bip32 gen --input-hex-seed --show-all-keys --network=testnet --output-format=json \
  | jq '.'
```
```json
{
  "seed": "3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678",
  "xPrv": "tprv8jBzzkzVsFu2aUqaXJPxCSe3fMAF6cjED21ehR9Zr4VKVdWMw7xZFUb9hAJEd3RJeepHkRgcnwWE7JQ3aieKhw4N97EMVfvMZMUNcLcSGf3",
  "xPub": "tpubDFt39B2k1dahTwsNQx4YbrJAENgBFwv8nKcRywBsGLHiL7m8ZWn9RyD1sHQsg71nPEmN2NUZZWd7REyvcVQ7HMMBcy6VjAF9WGv7bPjpQuK",
  "pubKeyHex": "024bf0776f0b553c9acfe6d47206849012eae6c81bd8e9fc645b4e84cabcfddc76",
  "prvKeyWif": "cPPKBqUDZa6Qb56QePPjhiGjTkGTRvTX8wymWCytJjYG9kCHuwY8",
  "addr": "mm3tBoqHXpUMjN5XFJJjk9UkmKqssxjVj1",
  "addrType": "legacy",
  "derivationPath": "m/44h/0h/0h/0/0",
  "coinType": "btc",
  "network": "testnet"
}
```

> Please note that `testnet` keys are currently not tested via any test vector

## derived keys
Child keys can be derived using parent private or public keys and derivation paths. 

For instance, if we start with private extended key that was compared against
exported key from `Mycelium` in previous steps, and derive the child key for
`change=0` and `index=0`, it would correspond to the derivation path of `m/0/0`
since the private extended key was already derived from the master key using
derivatation path of `m/44h/0h/0h`:
```bash
bip32 derive \
  --derivation-path=m/0/0 \
  xprv9z3JXNpBxuiKqst8M3GDLijacTh2HLhKyAV5b4Cnz6UMVD3LJ35qBVk79JpMs6XoygJoaEmVd4vDNQrsfdBaXngkCrvvmSSLiExb2usfLG1
```
```yaml
xPrv: xprvA4HQhA4Br6afRrswRBRydnSPqHGAVFptd3iiMp8kuMMsuso4vUuPSb91AizescB7NwS8uHijxBK4L1J5nJj98be4cSTKthrPraRZiFttSPj
xPub: xpub6HGm6fb5gU8xeLxQXCxyzvP8PK6etiYjzGeKACYNTgtrng8DU2DdzPTV1yvcwsjZ9o1UAUYq39RhLXBmJu66xMmAGGnDRGb4iLkP99rLmL5
pubKeyHex: 022b70459564b65102394e088bdb68f8a80d386939e39319363377b00f23b21cc4
prvKeyWif: Ky7kJQEFQDCRhShHaZs7TSCEa1UqGhVZB6BhXUHf3T9pAG6q7987
addr: 1MJ9PojuE1rA1E8wtrdQXjxaqZdsgddhoh
coinType: btc
network: mainnet
```

As you can see the address `1MJ9PojuE1rA1E8wtrdQXjxaqZdsgddhoh` matches the one
produced above using command:
```bash
bip32 gen --addr-type=legacy ${MNEMONIC}
```
```yaml
prvKeyWif: Ky7kJQEFQDCRhShHaZs7TSCEa1UqGhVZB6BhXUHf3T9pAG6q7987
addr: 1MJ9PojuE1rA1E8wtrdQXjxaqZdsgddhoh
```

Similarly, extended keys from other address types can be used:
```bash
bip32 derive \
  --derivation-path=m/0/0 \
  yprvAJ1zHuK6oeXUyMEkbwHM7QxusKG1DhcxFBtASzydLXSFZabFc5hXFLBPzij6CSnXoiPawEqBxYV8ZKSVuBoCzCEhZ3Rksjrc6WSJggMaTpa
```
```yaml
xPrv: yprvAMnfCAc35Zzftqq3WA8yhmD7juNk1UykJBUtc6QdWJkAoQJUajnCgxEc6KJwwMYTYzETw986nt6qVLmEMAbUYMiRcWuS8kvgAZKUdwHT5aC
xPub: ypub6an1bg8vuwYy7KuWcBfz4u9rHwDEQwhbfQQVQUpF4eH9gCdd8H6TEkZ5wasZcRfWY1U8ZYBTHquqSJbxMg1r6WNC3Zwxmwnt8SCxmirsu2G
pubKeyHex: 03397f9677279f78472b1c9528a760eb84ebb3c6019a5dfa6bbeb971cf58ae173b
prvKeyWif: L5Nx5ePGYjN2TzoadVXorDQXrchWEtwuDQEnsC4SSztz9b4tcCWQ
addr: 37vznvAgCmaKERDZmYaw3X4ArHracgVUfa
addrType: segwit-compatible, p2sh
coinType: btc
network: mainnet
```

As you can see the address `37vznvAgCmaKERDZmYaw3X4ArHracgVUfa` is the same as that generated 
previously using following command for address type of `segwit-compatible`
```bash
bip32 gen --addr-type=segwit-compatible ${MNEMONIC}
```
```yaml
prvKeyWif: L5Nx5ePGYjN2TzoadVXorDQXrchWEtwuDQEnsC4SSztz9b4tcCWQ
addr: 37vznvAgCmaKERDZmYaw3X4ArHracgVUfa
```

> Generation of hardened keys is only allowed for parent private keys.

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
bip32 validate xpub661MyMwAqRbcGczjuMoRm6dXaLDEhW1u34gKenbeYqAix21mdUKJyuyu5F1rzYGVxyL6tmgBUAEPrEz92mBXjByMRiJdba9wpnN37RLLAXa
```
```text
key is valid
```

However, this key is invalid.
```bash
bip32 validate xpub661MyMwAqRbcEYS8w7XLSVeEsBXy79zSzH1J8vCdxAZningWLdN3zgtU6LBpB85b3D2yc8sfvZU521AAwdZafEz7mnzBBsz4wKY5fTtTQBm
```
```text
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

passed bip84 test matrix
passed bip49 test matrix
passed bip44 test matrix
passed bip32 test matrix
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
