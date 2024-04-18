// Hadhat task
// Calls verification for a given address on a deployed VerificationV0 contract

import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";

task("v0Verifications", "Check VerificationV0 NFT's for an address")
	.addPositionalParam("address")
	.addPositionalParam("owner")
	.setAction(async (taskArgs, hre) => {
		await mintVerificationV0(hre, taskArgs.address, taskArgs.owner)
	})

async function mintVerificationV0(hre: HardhatRuntimeEnvironment, address: string, owner: string) {

	const [deployer] = await hre.ethers.getSigners()
	const deployerAddress = await deployer.getAddress()
	console.log('Operating contract with the account: ', deployerAddress)
	console.log('Deployer account balance: ', (await deployer.provider.getBalance(deployer.address)))

	const deployed = await hre.ethers.getContractAt('VerificationV0', address, deployer)
	console.log(`${deployer}`)
	console.log(`checking owner: ${JSON.stringify(owner)}`)
	console.log(await deployed.verifications(owner))
}
