const path = require("path")
const assetsDir = path.resolve(__dirname, "./company_home/assets")

module.exports = {
	entry: assetsDir + "/src/app.tsx",
	output: {
		path: assetsDir + "/dist",
		filename: "app.js",
	},
	module: {
		rules: [
			{
				test: /\.tsx?$/,
				use: "ts-loader",
			},
			{
				test: /\.(png|jpg|jpeg|gif)$/i,
				type: "asset/resource",
			},
		],
	},
	resolve: {
		extensions: [".ts", ".tsx", ".js"],
		alias: {
			"@contents": path.resolve(__dirname, "./company_home/contents"),
		},
	},
	devtool: "cheap-module-source-map",
}
