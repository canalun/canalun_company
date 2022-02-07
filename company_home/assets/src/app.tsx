import React from "react";
import { render } from "react-dom";
import CeoImage from "./components/ceoImage";

const ceoImageDom = document.getElementById("ceoImage");
if (ceoImageDom !== null) render(<CeoImage />, ceoImageDom);
