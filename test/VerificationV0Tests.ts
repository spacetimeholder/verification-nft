// Hardhat tests for contract VerificationV0.sol

import { ethers } from 'hardhat'
import { expect } from 'chai'

let deployer: any
let receiver1: any
let receiver2: any
let factory: any
let deployed: any

const tokenUri = 'https://arweave.net/QhsCULISvUp0C2TgVLaQlBPbn00l7CZoTNGs0z59BzU'

describe('Testing harness for Verification V0 Contract', () => {

	before(async () => {
		[deployer, receiver1, receiver2] = await ethers.getSigners()
	})

	describe('initialization', async () => {

		it('deploys Verification', async () => {
			factory = await ethers.getContractFactory('VerificationV0')
			deployed = await factory.deploy('VerificationV0', 'VER0', 'testType')
			expect(deployed).not.to.be.null;
		})

		it('has correct type set', async () => {
			expect(await deployed.getTypeName()).to.equal('testType')
		})
	})

	describe('minting', async () => {
		it('mints tokens', async () => {
			await expect(deployed.mint(receiver1.address, 'service1', 1))
				.to.emit(deployed, 'Transfer')
				.withArgs(ethers.ZeroAddress, receiver1.address, 1)

			expect(await deployed.tokenURI(1)).to.equal(`${tokenUri}`)
		})

		it('prevents mints by non-minter', async () => {
			await expect(deployed.connect(receiver1).mint(receiver1.address, 'service2', 2))
				.to.be.reverted
		})
	})

	describe('verifications', async () => {
		it('returns verifications array', async () => {
			// chai executes in order so there is already one token owned
			await expect(deployed.mint(receiver1.address, 'service2', 2))
				.not.to.be.reverted
			await expect(deployed.mint(receiver2.address, 'service2', 3))
				.not.to.be.reverted
			await expect(deployed.mint(receiver1.address, 'service3', 4))
				.not.to.be.reverted

			let vers = await deployed.verifications(receiver1.address)
			expect(vers.length).to.equal(3)
			// tokens are returned by owner index, e.g. increasing in order
			expect(vers[0].length).to.equal(3)
			expect(vers[0][0]).to.equal(1)
			expect(vers[0][1]).to.equal('service1')
			expect(vers[0][2]).to.equal(1)
			expect(vers[1].length).to.equal(3)
			expect(vers[1][0]).to.equal(2)
			expect(vers[1][1]).to.equal('service2')
			expect(vers[1][2]).to.equal(2)
			expect(vers[2].length).to.equal(3)
			expect(vers[2][0]).to.equal(4)
			expect(vers[2][1]).to.equal('service3')
			expect(vers[2][2]).to.equal(4)

			vers = await deployed.verifications(receiver2.address)
			expect(vers.length).to.equal(1)
			expect(vers[0].length).to.equal(3)
			expect(vers[0][0]).to.equal(3)
			expect(vers[0][1]).to.equal('service2')
			expect(vers[0][2]).to.equal(3)
		})
	})

	describe('transfers', async () => {
		it('prevents transfers', async () => {
			await expect(deployed.transferFrom(receiver1.address, receiver2.address, 1))
				.to.be.revertedWith('tokens non-transferrable')
		})
	})

	describe('burning', async () => {

		it('prevents burning by non-minter', async () => {
			await expect(deployed.connect(receiver1).burn(4))
				.to.be.reverted
		})

		it('burns tokens', async () => {
			let vers = await deployed.verifications(receiver1.address)
			expect(vers.length).to.equal(3)

			// receiver1's tokens have not changed from above
			expect(vers[0].length).to.equal(3)
			expect(vers[0][0]).to.equal(1)
			expect(vers[0][1]).to.equal('service1')
			expect(vers[0][2]).to.equal(1)
			expect(vers[1].length).to.equal(3)
			expect(vers[1][0]).to.equal(2)
			expect(vers[1][1]).to.equal('service2')
			expect(vers[1][2]).to.equal(2)
			expect(vers[2].length).to.equal(3)
			expect(vers[2][0]).to.equal(4)
			expect(vers[2][1]).to.equal('service3')
			expect(vers[2][2]).to.equal(4)

			await expect(deployed.burn(4))
				.not.to.be.reverted

			vers = await deployed.verifications(receiver1.address)
			expect(vers.length).to.equal(2)
			expect(vers[0].length).to.equal(3)
			expect(vers[0][0]).to.equal(1)
			expect(vers[0][1]).to.equal('service1')
			expect(vers[0][2]).to.equal(1)
			expect(vers[1].length).to.equal(3)
			expect(vers[1][0]).to.equal(2)
			expect(vers[1][1]).to.equal('service2')
			expect(vers[1][2]).to.equal(2)
		})
	})
})