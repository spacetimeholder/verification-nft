import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import 'hardhat-abi-exporter';
import 'hardhat-contract-sizer';

import './tasks/verification-v0/mint'
import './tasks/verification-v0/v0-verifications'

const config: HardhatUserConfig = {
	solidity: {
		version: "0.8.24",
		settings: {
			optimizer: {
				enabled: true,
				runs: 10,
			},
		},
	},
	networks: {
		sepolia: {
			url: 'https://sepolia.infura.io/v3/c06d50e6f7c542a886116aecb42e9dc3',
			accounts: [],
			gasPrice: 'auto',
			gas: 'auto',
		},
		optimismsepolia: {
			url: 'https://optimism-sepolia.infura.io/v3/c06d50e6f7c542a886116aecb42e9dc3',
			accounts: [],
			gasPrice: 'auto',
			gas: 'auto',
		},
		arbitrumsepolia: {
			url: 'https://arbitrum-sepolia.infura.io/v3/c06d50e6f7c542a886116aecb42e9dc3',
			accounts: [],
			gasPrice: 'auto',
			gas: 'auto',
		},
		polygonmumbai: {
			url: 'https://polygon-mumbai.infura.io/v3/c06d50e6f7c542a886116aecb42e9dc3',
			accounts: [],
			gasPrice: 'auto',
			gas: 'auto',
		},
		basesepolia: {
			url: 'https://sepolia.base.org',
			accounts: [],
			gasPrice: 'auto',
			gas: 'auto',
		},
		avalanchefuji: {
			url: 'https://avalanche-fuji.infura.io/v3/c06d50e6f7c542a886116aecb42e9dc3',
			accounts: [],
			gasPrice: 'auto',
			gas: 'auto',
		}
	},
	contractSizer: {
		alphaSort: true,
		runOnCompile: true,
		disambiguatePaths: false,
	}
};

export default config;
