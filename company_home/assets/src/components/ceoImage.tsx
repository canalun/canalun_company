import React from "react"
import Boy from "@materials/images/CEOs/boy.jpeg"
import Dragon from "@materials/images/CEOs/dragon.png"
import Alpaca from "@materials/images/CEOs/alpaca.jpeg"
import Goat from "@materials/images/CEOs/goat.jpeg"
import Girrafe from "@materials/images/CEOs/girrafe.jpeg"
import Doge from "@materials/images/CEOs/doge.jpeg"
import BoxCat from "@materials/images/CEOs/cat_in_the_box.jpeg"
import DjCat from "@materials/images/CEOs/dj_cat.jpg"

const CeoImage = () => {
	const images = [Boy, Dragon, Alpaca, Goat, Girrafe, Doge, BoxCat, DjCat]
	const imageNo = Math.floor(Math.random() * images.length)
	const width = 395
	const height = 395

	return (
		<>
			<img
				src={images[imageNo]}
				alt="CEO is at home today"
				width={width}
				height={height}
			/>
			<p>wanna see more? please reload...</p>
		</>
	)
}

export default CeoImage
