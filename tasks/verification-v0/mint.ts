// Hardhat task
// Mints a new VerificationV0 NFT

import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";

task("mintVerificationV0", "mint VerificationV0 NFT")
	.addPositionalParam("address")
	.addPositionalParam("to")
	.setAction(async (taskArgs, hre) => {
		await mintVerificationV0(hre, taskArgs.address, taskArgs.to)
	})

async function mintVerificationV0(hre: HardhatRuntimeEnvironment, address: string, to: string) {

	const [deployer] = await hre.ethers.getSigners()
	const deployerAddress = await deployer.getAddress()
	console.log('Operating contract with the account: ', deployerAddress)
	console.log('Deployer account balance: ', (await deployer.provider.getBalance(deployer.address)))

	const deployed = await hre.ethers.getContractAt('VerificationV0', address, deployer)
	console.log(await deployed.mint(to, 'test service', new Date().getTime()))
}
