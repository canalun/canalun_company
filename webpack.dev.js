/* eslint-disable no-undef */
/* eslint-disable @typescript-eslint/no-var-requires */
const path = require("path")
const { merge } = require("webpack-merge")
const common = require("./webpack.common.js")

module.exports = merge(common, {
	mode: "development",
	devtool: "cheap-module-source-map",
	output: {
		path: path.resolve(__dirname, "./company_home/scripts/dev_dist"),
		publicPath: "/",
		filename: "app.js",
	},
	devServer: {
		static: "./company_home/scripts/dev_dist",
		port: 8080,
		hot: true,
	},
})