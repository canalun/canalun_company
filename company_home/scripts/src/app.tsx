import React from "react"
import { render } from "react-dom"
import CeoImage from "./components/ceoImage"
import { JapaneseEntries } from "./components/japaneseEntries"

const ceoImageDom = document.getElementById("ceoImage")
if (ceoImageDom !== null) render(<CeoImage />, ceoImageDom)

const japaneseEntryListDom = document.getElementById("japaneseEntryList")
if (japaneseEntryListDom !== null) render(<JapaneseEntries />, japaneseEntryListDom)