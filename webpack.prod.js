/* eslint-disable no-undef */
/* eslint-disable @typescript-eslint/no-var-requires */
const { merge } = require("webpack-merge")
const common = require("./webpack.common.js")
const path = require("path")


module.exports = merge(common, {
	mode: "production",
	output: {
		path: path.resolve(__dirname, "./company_home/scripts/dist"),
		filename: "app.js",
	},
})