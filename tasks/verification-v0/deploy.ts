// WARNING: deprecated in preference of hardhat ignition
import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";

task("deployVerificationV0", "Deploy a VerificationV2 contract")
	.addPositionalParam("symbol")
	.addPositionalParam("name")
	.addPositionalParam("typename")
	.setAction(async (taskArgs, hre) => {
		await deploy(hre, taskArgs.symbol, taskArgs.name, taskArgs.typeName)
	})

async function deploy(hre: HardhatRuntimeEnvironment, symbol: string, name: string, typeName: string) {

	const [deployer, recevier] = await hre.ethers.getSigners()
	console.log('Deploying cosigned pass nft with the account:', deployer.address)
	console.log('Deployer account balance:', (await deployer.provider.getBalance(deployer.address)))

	const factory = await hre.ethers.getContractFactory('VerificationV0')
	const deployed = await factory.deploy(symbol, name, typeName)
	console.log('deployed VerificationV0 NFT at:', deployed.getAddress())
}