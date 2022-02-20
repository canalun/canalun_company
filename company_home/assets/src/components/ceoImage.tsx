import React from "react"
import Boy from "@contents/images/CEOs/boy.jpeg"
import Dragon from "@contents/images/CEOs/dragon.png"
import Alpaca from "@contents/images/CEOs/alpaca.jpeg"
import Goat from "@contents/images/CEOs/goat.jpeg"
import Girrafe from "@contents/images/CEOs/girrafe.jpeg"
import Doge from "@contents/images/CEOs/doge.jpeg"
import BoxCat from "@contents/images/CEOs/cat_in_the_box.jpeg"
import DjCat from "@contents/images/CEOs/dj_cat.jpg"

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
