/*
 * For a detailed explanation regarding each configuration property, visit:
 * https://jestjs.io/docs/configuration
 */

module.exports = {
	// A list of paths to directories that Jest should use to search for files in
	roots: [
		"./company_home"
	],

	// The glob patterns Jest uses to detect test files
	testMatch: [
		"**/__tests__/**/*.[jt]s?(x)",
		"**/?(*.)+(spec|test).[tj]s?(x)"
	],

	// A map from regular expressions to paths to transformers
	"transform": {
		"^.+\\.(ts|tsx)$": "ts-jest"
	},
}
