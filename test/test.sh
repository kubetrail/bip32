#!/usr/bin/env bash

# test vector 1
echo https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-1
seed="000102030405060708090a0b0c0d0e0f"
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m --input-hex-seed \
	| jq -r '.pub')" == "xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m --input-hex-seed \
	| jq -r '.prv')" == "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H --input-hex-seed \
	| jq -r '.pub')" == "xpub68Gmy5EdvgibQVfPdqkBBCHxA5htiqg55crXYuXoQRKfDBFA1WEjWgP6LHhwBZeNK1VTsfTFUHCdrfp1bgwQ9xv5ski8PX9rL2dZXvgGDnw" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H --input-hex-seed \
	| jq -r '.prv')" == "xprv9uHRZZhk6KAJC1avXpDAp4MDc3sQKNxDiPvvkX8Br5ngLNv1TxvUxt4cV1rGL5hj6KCesnDYUhd7oWgT11eZG7XnxHrnYeSvkzY7d2bhkJ7" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1 --input-hex-seed \
	| jq -r '.pub')" == "xpub6ASuArnXKPbfEwhqN6e3mwBcDTgzisQN1wXN9BJcM47sSikHjJf3UFHKkNAWbWMiGj7Wf5uMash7SyYq527Hqck2AxYysAA7xmALppuCkwQ" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1 --input-hex-seed \
	| jq -r '.prv')" == "xprv9wTYmMFdV23N2TdNG573QoEsfRrWKQgWeibmLntzniatZvR9BmLnvSxqu53Kw1UmYPxLgboyZQaXwTCg8MSY3H2EU4pWcQDnRnrVA1xe8fs" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1/2H --input-hex-seed \
	| jq -r '.pub')" == "xpub6D4BDPcP2GT577Vvch3R8wDkScZWzQzMMUm3PWbmWvVJrZwQY4VUNgqFJPMM3No2dFDFGTsxxpG5uJh7n7epu4trkrX7x7DogT5Uv6fcLW5" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1/2H --input-hex-seed \
	| jq -r '.prv')" == "xprv9z4pot5VBttmtdRTWfWQmoH1taj2axGVzFqSb8C9xaxKymcFzXBDptWmT7FwuEzG3ryjH4ktypQSAewRiNMjANTtpgP4mLTj34bhnZX7UiM" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1/2H/2 --input-hex-seed \
	| jq -r '.pub')" == "xpub6FHa3pjLCk84BayeJxFW2SP4XRrFd1JYnxeLeU8EqN3vDfZmbqBqaGJAyiLjTAwm6ZLRQUMv1ZACTj37sR62cfN7fe5JnJ7dh8zL4fiyLHV" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1/2H/2 --input-hex-seed \
	| jq -r '.prv')" == "xprvA2JDeKCSNNZky6uBCviVfJSKyQ1mDYahRjijr5idH2WwLsEd4Hsb2Tyh8RfQMuPh7f7RtyzTtdrbdqqsunu5Mm3wDvUAKRHSC34sJ7in334" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1/2H/2/1000000000 --input-hex-seed \
	| jq -r '.pub')" == "xpub6H1LXWLaKsWFhvm6RVpEL9P4KfRZSW7abD2ttkWP3SSQvnyA8FSVqNTEcYFgJS2UaFcxupHiYkro49S8yGasTvXEYBVPamhGW6cFJodrTHy" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1/2H/2/1000000000 --input-hex-seed \
	| jq -r '.prv')" == "xprvA41z7zogVVwxVSgdKUHDy1SKmdb533PjDz7J6N6mV6uS3ze1ai8FHa8kmHScGpWmj4WggLyQjgPie1rFSruoUihUZREPSL39UNdE3BBDu76" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

# test vector 2
echo; echo; echo https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-2
seed="fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542"
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m --input-hex-seed \
	| jq -r '.pub')" == "xpub661MyMwAqRbcFW31YEwpkMuc5THy2PSt5bDMsktWQcFF8syAmRUapSCGu8ED9W6oDMSgv6Zz8idoc4a6mr8BDzTJY47LJhkJ8UB7WEGuduB" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m --input-hex-seed \
	| jq -r '.prv')" == "xprv9s21ZrQH143K31xYSDQpPDxsXRTUcvj2iNHm5NUtrGiGG5e2DtALGdso3pGz6ssrdK4PFmM8NSpSBHNqPqm55Qn3LqFtT2emdEXVYsCzC2U" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0 --input-hex-seed \
	| jq -r '.pub')" == "xpub69H7F5d8KSRgmmdJg2KhpAK8SR3DjMwAdkxj3ZuxV27CprR9LgpeyGmXUbC6wb7ERfvrnKZjXoUmmDznezpbZb7ap6r1D3tgFxHmwMkQTPH" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0 --input-hex-seed \
	| jq -r '.prv')" == "xprv9vHkqa6EV4sPZHYqZznhT2NPtPCjKuDKGY38FBWLvgaDx45zo9WQRUT3dKYnjwih2yJD9mkrocEZXo1ex8G81dwSM1fwqWpWkeS3v86pgKt" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0/2147483647H --input-hex-seed \
	| jq -r '.pub')" == "xpub6ASAVgeehLbnwdqV6UKMHVzgqAG8Gr6riv3Fxxpj8ksbH9ebxaEyBLZ85ySDhKiLDBrQSARLq1uNRts8RuJiHjaDMBU4Zn9h8LZNnBC5y4a" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0/2147483647H --input-hex-seed \
	| jq -r '.prv')" == "xprv9wSp6B7kry3Vj9m1zSnLvN3xH8RdsPP1Mh7fAaR7aRLcQMKTR2vidYEeEg2mUCTAwCd6vnxVrcjfy2kRgVsFawNzmjuHc2YmYRmagcEPdU9" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0/2147483647H/1 --input-hex-seed \
	| jq -r '.pub')" == "xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0/2147483647H/1 --input-hex-seed \
	| jq -r '.prv')" == "xprv9zFnWC6h2cLgpmSA46vutJzBcfJ8yaJGg8cX1e5StJh45BBciYTRXSd25UEPVuesF9yog62tGAQtHjXajPPdbRCHuWS6T8XA2ECKADdw4Ef" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0/2147483647H/1/2147483646H --input-hex-seed \
	| jq -r '.pub')" == "xpub6ERApfZwUNrhLCkDtcHTcxd75RbzS1ed54G1LkBUHQVHQKqhMkhgbmJbZRkrgZw4koxb5JaHWkY4ALHY2grBGRjaDMzQLcgJvLJuZZvRcEL" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0/2147483647H/1/2147483646H --input-hex-seed \
	| jq -r '.prv')" == "xprvA1RpRA33e1JQ7ifknakTFpgNXPmW2YvmhqLQYMmrj4xJXXWYpDPS3xz7iAxn8L39njGVyuoseXzU6rcxFLJ8HFsTjSyQbLYnMpCqE2VbFWc" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0/2147483647H/1/2147483646H/2 --input-hex-seed \
	| jq -r '.pub')" == "xpub6FnCn6nSzZAw5Tw7cgR9bi15UV96gLZhjDstkXXxvCLsUXBGXPdSnLFbdpq8p9HmGsApME5hQTZ3emM2rnY5agb9rXpVGyy3bdW6EEgAtqt" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0/2147483647H/1/2147483646H/2 --input-hex-seed \
	| jq -r '.prv')" == "xprvA2nrNbFZABcdryreWet9Ea4LvTJcGsqrMzxHx98MMrotbir7yrKCEXw7nadnHM8Dq38EGfSh6dqA9QWTyefMLEcBYJUuekgW4BYPJcr9E7j" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

# test vector 3
echo; echo; echo https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-3
seed="4b381541583be4423346c643850da4b320e46a87ae3d2a4e6da11eba819cd4acba45d239319ac14f863b8d5ab5a0d0c64d2e8a1e7d1457df2e5a3c51c73235be"
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m --input-hex-seed \
	| jq -r '.pub')" == "xpub661MyMwAqRbcEZVB4dScxMAdx6d4nFc9nvyvH3v4gJL378CSRZiYmhRoP7mBy6gSPSCYk6SzXPTf3ND1cZAceL7SfJ1Z3GC8vBgp2epUt13" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m --input-hex-seed \
	| jq -r '.prv')" == "xprv9s21ZrQH143K25QhxbucbDDuQ4naNntJRi4KUfWT7xo4EKsHt2QJDu7KXp1A3u7Bi1j8ph3EGsZ9Xvz9dGuVrtHHs7pXeTzjuxBrCmmhgC6" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H --input-hex-seed \
	| jq -r '.pub')" == "xpub68NZiKmJWnxxS6aaHmn81bvJeTESw724CRDs6HbuccFQN9Ku14VQrADWgqbhhTHBaohPX4CjNLf9fq9MYo6oDaPPLPxSb7gwQN3ih19Zm4Y" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H --input-hex-seed \
	| jq -r '.prv')" == "xprv9uPDJpEQgRQfDcW7BkF7eTya6RPxXeJCqCJGHuCJ4GiRVLzkTXBAJMu2qaMWPrS7AANYqdq6vcBcBUdJCVVFceUvJFjaPdGZ2y9WACViL4L" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

# test vector 4
echo; echo; echo https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-4
seed="3ddd5602285899a946114506157c7997e5444528f3003f6134712147db19b678"
if [[ "$(echo ${seed} \
	| bip32 gen --chain=m --input-hex-seed \
	| jq -r '.pub')" == "xpub661MyMwAqRbcGczjuMoRm6dXaLDEhW1u34gKenbeYqAix21mdUKJyuyu5F1rzYGVxyL6tmgBUAEPrEz92mBXjByMRiJdba9wpnN37RLLAXa" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m --input-hex-seed \
	| jq -r '.prv')" == "xprv9s21ZrQH143K48vGoLGRPxgo2JNkJ3J3fqkirQC2zVdk5Dgd5w14S7fRDyHH4dWNHUgkvsvNDCkvAwcSHNAQwhwgNMgZhLtQC63zxwhQmRv" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H --input-hex-seed \
	| jq -r '.pub')" == "xpub69AUMk3qDBi3uW1sXgjCmVjJ2G6WQoYSnNHyzkmdCHEhSZ4tBok37xfFEqHd2AddP56Tqp4o56AePAgCjYdvpW2PU2jbUPFKsav5ut6Ch1m" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H --input-hex-seed \
	| jq -r '.prv')" == "xprv9vB7xEWwNp9kh1wQRfCCQMnZUEG21LpbR9NPCNN1dwhiZkjjeGRnaALmPXCX7SgjFTiCTT6bXes17boXtjq3xLpcDjzEuGLQBM5ohqkao9G" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1H --input-hex-seed \
	| jq -r '.pub')" == "xpub6BJA1jSqiukeaesWfxe6sNK9CCGaujFFSJLomWHprUL9DePQ4JDkM5d88n49sMGJxrhpjazuXYWdMf17C9T5XnxkopaeS7jGk1GyyVziaMt" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

if [[ "$(echo ${seed} \
	| bip32 gen --chain=m/0H/1H --input-hex-seed \
	| jq -r '.prv')" == "xprv9xJocDuwtYCMNAo3Zw76WENQeAS6WGXQ55RCy7tDJ8oALr4FWkuVoHJeHVAcAqiZLE7Je3vZJHxspZdFHfnBEjHqU5hG1Jaj32dVoS6XLT1" ]]; \
then echo -n "ok "; else echo failed for ${seed}; fi

# test vector 5
echo; echo; echo https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-5
if $(echo xpub661MyMwAqRbcEYS8w7XLSVeEsBXy79zSzH1J8vCdxAZningWLdN3zgtU6LBpB85b3D2yc8sfvZU521AAwdZafEz7mnzBBsz4wKY5fTtTQBm | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xprv9s21ZrQH143K24Mfq5zL5MhWK9hUhhGbd45hLXo2Pq2oqzMMo63oStZzFGTQQD3dC4H2D5GBj7vWvSQaaBv5cxi9gafk7NF3pnBju6dwKvH | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xpub661MyMwAqRbcEYS8w7XLSVeEsBXy79zSzH1J8vCdxAZningWLdN3zgtU6Txnt3siSujt9RCVYsx4qHZGc62TG4McvMGcAUjeuwZdduYEvFn | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xprv9s21ZrQH143K24Mfq5zL5MhWK9hUhhGbd45hLXo2Pq2oqzMMo63oStZzFGpWnsj83BHtEy5Zt8CcDr1UiRXuWCmTQLxEK9vbz5gPstX92JQ | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xpub661MyMwAqRbcEYS8w7XLSVeEsBXy79zSzH1J8vCdxAZningWLdN3zgtU6N8ZMMXctdiCjxTNq964yKkwrkBJJwpzZS4HS2fxvyYUA4q2Xe4 | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xprv9s21ZrQH143K24Mfq5zL5MhWK9hUhhGbd45hLXo2Pq2oqzMMo63oStZzFAzHGBP2UuGCqWLTAPLcMtD9y5gkZ6Eq3Rjuahrv17fEQ3Qen6J | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xprv9s2SPatNQ9Vc6GTbVMFPFo7jsaZySyzk7L8n2uqKXJen3KUmvQNTuLh3fhZMBoG3G4ZW1N2kZuHEPY53qmbZzCHshoQnNf4GvELZfqTUrcv | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xpub661no6RGEX3uJkY4bNnPcw4URcQTrSibUZ4NqJEw5eBkv7ovTwgiT91XX27VbEXGENhYRCf7hyEbWrR3FewATdCEebj6znwMfQkhRYHRLpJ | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xprv9s21ZrQH4r4TsiLvyLXqM9P7k1K3EYhA1kkD6xuquB5i39AU8KF42acDyL3qsDbU9NmZn6MsGSUYZEsuoePmjzsB3eFKSUEh3Gu1N3cqVUN | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xpub661MyMwAuDcm6CRQ5N4qiHKrJ39Xe1R1NyfouMKTTWcguwVcfrZJaNvhpebzGerh7gucBvzEQWRugZDuDXjNDRmXzSZe4c7mnTK97pTvGS8 | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo DMwo58pR1QLEFihHiXPVykYB6fJmsTeHvyTp7hRThAtCX8CvYzgPcn8XnmdfHGMQzT7ayAmfo4z3gY5KfbrZWZ6St24UVf2Qgo6oujFktLHdHY4 | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo DMwo58pR1QLEFihHiXPVykYB6fJmsTeHvyTp7hRThAtCX8CvYzgPcn8XnmdfHPmHJiEDXkTiJTVV9rHEBUem2mwVbbNfvT2MTcAqj3nesx8uBf9 | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xprv9s21ZrQH143K24Mfq5zL5MhWK9hUhhGbd45hLXo2Pq2oqzMMo63oStZzF93Y5wvzdUayhgkkFoicQZcP3y52uPPxFnfoLZB21Teqt1VvEHx | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xprv9s21ZrQH143K24Mfq5zL5MhWK9hUhhGbd45hLXo2Pq2oqzMMo63oStZzFAzHGBP2UuGCqWLTAPLcMtD5SDKr24z3aiUvKr9bJpdrcLg1y3G | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHL | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

if $(echo xpub661MyMwAqRbcEYS8w7XLSVeEsBXy79zSzH1J8vCdxAZningWLdN3zgtU6Q5JXayek4PRsn35jii4veMimro1xefsM58PgBMrvdYre8QyULY | bip32 validate 2> /dev/null); \
then echo validation failed; else echo -n "ok "; fi

echo
