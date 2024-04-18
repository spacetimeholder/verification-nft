#!/bin/bash

# call from the root directory, not /scripts.

# Script to generate Golang bindings for whatever contracts are declared here.

npx hardhat clean; npx hardhat compile
(cd extract-from-artifact && npx .  --artifact ../artifacts/contracts/VerificationV0.sol/VerificationV0.json --bin ../bytecode/VerificationV0.bin --abi ../abi/VerificationV0.json)
abigen --abi abi/VerificationV0.json --bin bytecode/VerificationV0.bin --pkg bindings --type VerificationV0 --out bindings/verification_v0.go
# call using an abi file copied directly into /abi
abigen --abi abi/USDC.json --pkg bindings --type USDC --out bindings/usdc.go
