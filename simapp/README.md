---
sidebar_position: 1
---

# `Brainfuck Mainnet`

This is the official sourcecode of the Brainfuck blockchain. Its mainnet is Brainnet. Its testnet is Fucknet (because you are fucked if you don't have a brain).

## Running testnets with `brainnetd`

```sh
# From the root directory of the Cosmos SDK repository
# This will build the `brainnetd` binary inside a new `build` directory
make build

# If you've run `brainnetd` before, you may need to reset your database before starting a new
# testnet. You can reset your database with the following command: `$ ./brainnetd comet unsafe-reset-all`.

# 3.
# This will initialize a new working directory
# at the default location `~/.simapp`. You need to provide a "moniker" and a "chain id". These
# two names can be anything, but you will need to use the same "chain id" in the following steps.
./brainnetd init [moniker] --chain-id [chain-id]

# This will create a new key, with a name of your choosing.
# Save the output of this command somewhere; you'll need the address generated here later.
./brainnetd keys add [key_name]

# 5. 
# where `key_name` is the same key name as before; and `amount` is something like `10000000000000000000000000stake`.
./brainnetd genesis add-genesis-account [key_name] [amount]

# 6. This will create the genesis
# transaction for your new chain. Here `amount` should be at least `1000000000stake`. If you
# provide too much or too little, you will encounter an error when starting your node.
./brainnetd genesis gentx [key_name] [amount] --chain-id [chain-id]

# 7. Now, one person needs to create the genesis file `genesis.json` using the genesis transactions
# from every participant, by gathering all the genesis transactions under `config/gentx` and then
# calling `$ ./brainnetd genesis collect-gentxs`. This will create a new `genesis.json` file that includes data
# from all the validators (we sometimes call it the "super genesis file" to distinguish it from
# single-validator genesis files).

# 8. Once you've received the super genesis file, overwrite your original `genesis.json` file with
# the new super `genesis.json`.

# 9. Modify your `config/config.toml` (in the simapp working directory) to include the other participants as
# persistent peers:

# ```text
# # Comma separated list of nodes to keep persistent connections to
# persistent_peers = "[validator_address]@[ip_address]:[port],[validator_address]@[ip_address]:[port]"
# ```

# You can find `validator_address` by running `$ ./brainnetd comet show-node-id`. The output will
# be the hex-encoded `validator_address`. The default `port` is 26656.
# 10. Now you can start your nodes: `$ ./brainnetd start`.

# Now you have a small testnet that you can use to try out changes to the Cosmos SDK or CometBFT!
```