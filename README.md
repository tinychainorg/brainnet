---
title: blockchain that runs brainfuck
---
# get ready for Brainnet

A blockchain that runs Brainfuck smart contracts.

It's based on Bitcoin proof-of-work consensus with the ZK friendly hash function, Poseidon.

The entire thing is proven using [Halo 2 SNARK's](https://github.com/cryptape/ckb-bf-zkvm)

> "This is plausible the most perverted blockchain project in history." - Vinay Gupta, co-ran the Ethereum launch

> "cannot wait to write a constant product AMM in Brainfuck" - Sam, Robot Labs.

![fe5f2450-7b82-49d0-8a66-414f96adcc35](https://github.com/tinychainorg/tinychainorg.github.io/assets/584141/2f8d97f9-6883-49c4-9982-06ec5d7c215a)

## Get ready 

we need devs and people willing to run the brainfuck nodes

https://t.me/+t_OqbV1y2Nw1YjFl

## Status

| **Area**     | **Description**                                                                                                                                                                                                                                                                                            | **Status**  |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| VM           | [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) smart contracts                                                                                                                                                                                                                                       | ✅⚠️ 60% Done |
| Consensus    | Bitcoin / Nakamoto / POW with ZK-friendly hash function                                                                                                                                                                                                                                                    | ⚠️ WIP       |
| Tokenomics   | Ethereum-like - native token + fixed exchange rate to gas                                                                                                                                                                                                                                                  | ✅ Done      |
| Cryptography | ECDSA wallets, SECP256k1 curve (same as BTC), SHA-2 256 bit hash function                                                                                                                                                                                                                                  | ✅ Done      |
| Networking   | P2P and RPC servers both use HTTP, gossip network architecture                                                                                                                                                                                                                                             | ✅ Done       |
| ZK proofs    | ZK for compression of tinychain. Use either [groth16](https://github.com/erhant/zkbrainfuck) or [halo2](https://github.com/cryptape/ckb-bf-zkvm) SNARK proofs for brainfuck. TBA we will rework consensus/crypto to use SNARK-friendly tech (MiMC/Poseidon hash function, SNARK-friendly signature scheme) |             |

