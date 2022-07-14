---
title: "Intiting Contract"
---

# Setup Secret Contracts (CosmWasm)

Secret Contracts are based on [CosmWasm](https://www.cosmwasm.com) which is implementated on various Cosmos SDK blockchains. The CosmWasm smart contracts are like Ethereum's smart contracts except they can be used on other networks using the [Inter-Blockchain Protocol](https://cosmos.network/ibc) (IBC). CosmWasm smart contracts are written in the Rust language.

The Secret Network has a _compute_ module that we'll use to store, query and instantiate the smart contract. Once stored on the blockchain the smart contract has to be created (or instantiated) in order to execute its methods. This is similar to doing an Ethereum `migrate` using truffle which handles the deployment and creation of a smart contract.

Eventually the smart contracts will become Secret Contracts (in a future blockchain upgrade) running in an SGX enclave (Trusted Execution Environment) where computations are performed on the encrypted contract data (i.e. inputs, state).

Next we'll walkthrough steps to:

- install Rust
- install the Rust dependencies
- create your first project

The Rust dependencies include the Rust compiler, cargo (_package manager_), toolchain and a package to generate projects. You can check out the Rust book, rustlings course, examples and more [here](https://www.rust-lang.org/learn).

1. Install Rust

More information about installing Rust can be found here: https://www.rust-lang.org/tools/install.

```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
source $HOME/.cargo/env
```

2. Add rustup target wasm32 for both stable and nightly

```
rustup default stable
rustup target list --installed
rustup target add wasm32-unknown-unknown

rustup install nightly
rustup target add wasm32-unknown-unknown --toolchain nightly
```

3. If using linux, install the standard build tools:
```
apt install build-essential
```

4. Run cargo install cargo-generate

[Cargo generate](https://doc.rust-lang.org/cargo) is the tool you'll use to create a smart contract project.

```
cargo install cargo-generate --features vendored-openssl
```

### Create Initial Smart Contract

To create the smart contract you'll:
- generate the initial project
- compile the smart contract
- run unit tests
- optimize the wasm contract bytecode to prepare for deployment
- deploy the smart contract to your local SecretNetwork
- instantiate it with contract parameters

Generate the smart contract project

```
cargo generate --git https://github.com/enigmampc/secret-template --name mysimplecounter
```

The git project above is a cosmwasm smart contract template that implements a simple counter. The contract is created with a parameter for the initial count and allows subsequent incrementing.

Change directory to the project you created and view the structure and files that were created.

```
cd mysimplecounter
```

The generate creates a directory with the project name and has this structure:

```
Cargo.lock	Developing.md	LICENSE		Publishing.md	examples	schema		tests
Cargo.toml	Importing.md	NOTICE		README.md	rustfmt.toml	src
```

### Compile

Use the following command to compile the smart contract which produces the wasm contract file.

```
cargo wasm
```

### Unit Tests (NB Tests in this template currently fail unless you have SGX enabled)

Run unit tests

```
RUST_BACKTRACE=1 cargo unit-test
```

### Integration Tests

The integration tests are under the `tests/` directory and run as:

```
cargo integration-test
```

### Generate Msg Schemas

We can also generate JSON Schemas that serve as a guide for anyone trying to use the contract, to specify which arguments they need.

Auto-generate msg schemas (when changed):

```
cargo schema
```
