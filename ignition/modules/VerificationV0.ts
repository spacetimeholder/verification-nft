import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

export default buildModule("VerificationV0", (m) => {
	const ver2 = m.contract("VerificationV0", ["VER2TEST", 'VerificationV0 Test', 'test type']);

	// m.call(ver2, "launch", []);

	return { ver2 };
});