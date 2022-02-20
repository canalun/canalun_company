import React from "react"
import { render } from "react-dom"
import CeoImage from "./components/ceoImage"
import { JapaneseEntryList } from "./components/Japaneseentries"

const ceoImageDom = document.getElementById("ceoImage")
if (ceoImageDom !== null) render(<CeoImage />, ceoImageDom)

const japaneseEntryListDom = document.getElementById("japaneseEntryList")
if (japaneseEntryListDom !== null) render(<JapaneseEntryList />, japaneseEntryListDom)