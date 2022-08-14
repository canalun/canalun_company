/////////////////////////////////////

//////////////////////////
/// init player oshiri ///
//////////////////////////

const initOshiri = (playSoundEffect) => {
	const height = 70
	const width = 70
	const onaraTime = 600
	const trackDelay = 0.1 //second
	const oshiriImgSrc = "../../materials/images/oshiri_katori/oshiri.svg"
	const onaraImgSrc = "../../materials/images/oshiri_katori/onara.svg"

	const oshiri = document.createElement("img")
	document.body.appendChild(oshiri)

	oshiri.src = oshiriImgSrc
	Object.assign(oshiri.style, {
		width: width + "px",
		height: height + "px",
	})

	window.addEventListener("mousedown", function () {
		oshiri.src = onaraImgSrc
		playSoundEffect()
		setTimeout(() => {
			oshiri.src = oshiriImgSrc
		}, onaraTime)
	})

	Object.assign(oshiri.style, {
		position: "fixed",
		transition:
			"left " +
			trackDelay +
			"s ease-in-out 0s, top " +
			trackDelay +
			"s ease-in-out 0s",
		"-webkit-transition":
			"left " +
			trackDelay +
			"s ease-in-out 0s, top " +
			trackDelay +
			"s ease-in-out 0s"
	})
	window.addEventListener("mousemove", function (e) {
		oshiri.style.left = e.clientX - width / 2 + "px"
		oshiri.style.top = e.clientY - height / 2 + "px"
	})

	return oshiri
}

///////////////////////////////////////////

/////////////////////////
///// enemy rendering ///
/////////////////////////

let speed = 90 // px/second
const increaseSpeed = () => {
	setInterval(() => { speed += 1 }, 300)
}

const generateMosquito = () => {
	const imgSrc =
		"../../materials/images/oshiri_katori/mosquito.png"
	const width = 30
	const height = 30
	const noMosquitoMargin = window.innerWidth * 0.1
	const fps = 60
	const moveDelay = 0.01 //second

	const mosquito = document.createElement("img")
	document.body.appendChild(mosquito)

	mosquito.className = "enemy"
	mosquito.src = imgSrc
	Object.assign(mosquito.style, {
		width: width + "px",
		height: height + "px"
	})

	// enemy appears from bottom.
	// set no-enemy margin at right and left edge. => margin < mosquitoLeft < window.innerWidth - margin
	Object.assign(mosquito.style, {
		top: window.innerHeight + "px",
		left:
			noMosquitoMargin +
			(window.innerWidth - 2 * noMosquitoMargin) * Math.random() +
			"px"
	})

	const stopMoveCalc = setInterval(() => {
		let top = mosquito.getBoundingClientRect().top
		let left = mosquito.getBoundingClientRect().left
		top -= speed / fps
		Object.assign(mosquito.style, {
			top: top + "px",
			left: left + "px"
		})
	}, Math.trunc(1000 / fps))

	Object.assign(mosquito.style, {
		position: "fixed",
		transition: "left " + moveDelay + "s 0s, top " + moveDelay + "s 0s",
		"-webkit-transition":
			"left " + moveDelay + "s 0s, top " + moveDelay + "s 0s"
	})

	return stopMoveCalc
}

const generateBee = () => {
	const beeThresholdScore = 30
	if (currentScore < beeThresholdScore) {
		return
	}

	const imgSrc =
		"../../materials/images/oshiri_katori/bee.png"
	const width = 30
	const height = 30
	const noBeeMargin = window.innerWidth * 0.1
	const fps = 60
	const moveDelay = 0.01 //second

	const bee = document.createElement("img")
	document.body.appendChild(bee)

	bee.className = "enemy"
	bee.src = imgSrc
	Object.assign(bee.style, {
		width: width + "px",
		height: height + "px"
	})

	// enemy appears from bottom.
	// set no-enemy margin at right and left edge. ==> margin < beeLeft < window.innerWidth - margin
	Object.assign(bee.style, {
		top: window.innerHeight + "px",
		left:
			noBeeMargin +
			(window.innerWidth - 2 * noBeeMargin) * Math.random() +
			"px"
	})

	const stopMoveCalc = setInterval(() => {
		let top = bee.getBoundingClientRect().top
		top -= speed * 1.2 / fps

		let left = bee.getBoundingClientRect().left
		const t = new Date().getTime() / 500
		const radius = noBeeMargin * 0.1
		left = left + (Math.sin(t) * radius)

		Object.assign(bee.style, {
			top: top + "px",
			left: left + "px"
		})
	}, Math.trunc(1000 / fps))

	Object.assign(bee.style, {
		position: "fixed",
		transition: "left " + moveDelay + "s 0s, top " + moveDelay + "s 0s",
		"-webkit-transition":
			"left " + moveDelay + "s 0s, top " + moveDelay + "s 0s"
	})

	return stopMoveCalc
}

///////////////////////////////////

///////////////////////////
//// stage rendering //////
///////////////////////////

const renderingStage = () => {
	const oshiriKuni = document.createElement("div")
	Object.assign(oshiriKuni.style, {
		border: "4px solid blue",
		fontSize: "40px",
		height: "50px",
		"line-height": "50px",
		"text-align": "center",
		padding: "4px",
		color: "blue"
	})
	oshiriKuni.innerText = "お し り の く に"
	document.body.appendChild(oshiriKuni)
	return oshiriKuni.getBoundingClientRect().bottom
}

///////////////////////////////////

//////////////////////////
//// onara detection /////
//////////////////////////

// attack is determined by checking the rectangular area where the player exists in the center of the top side
const rectDetect = (ex, ey, px, py, xRange, yRange) => {
	if ((px - xRange < ex &&
		ex < px + xRange) &&
		(py < ey &&
			ey < py + yRange)) {
		return true
	} else {
		return false
	}
}

// attack is determined by checking the downward sector area with the central angle of 120 degrees centered on the player, with the radius equal to xrange
const circleDetect = (ex, ey, px, py, xRange, yRange) => {
	const conditionalOne = () => {
		if (ey >= (ex - px) / Math.sqrt(3.0) + py) {
			return true
		} else {
			return false
		}
	}

	const conditionalTwo = () => {
		if (ey >= -1 * (ex - px) / Math.sqrt(3.0) + py) {
			return true
		} else {
			return false
		}
	}

	const conditionalThree = () => {
		if ((ex - px) ** 2 + (ey - py) ** 2 <= xRange * yRange) {
			return true
		} else {
			return false
		}
	}

	if (conditionalOne() && conditionalTwo() && conditionalThree()) {
		return true
	} else {
		return false
	}
}

// initial detection method is rectDetect
let detect = rectDetect
const originalXrange = 60
const originalYrange = 80
let xRange = originalXrange
let yRange = originalYrange

const onaraDetector = (target, player, callback) => {
	const ex = target.getBoundingClientRect().left
	const ey = target.getBoundingClientRect().top
	const px = player.getBoundingClientRect().left
	const py = player.getBoundingClientRect().top

	if (detect(ex, ey, px, py, xRange, yRange)) {
		callback()
	}
}

///////////////////////////////////

///////////////////////////
//// attack detection /////
///////////////////////////

const attackDetector = (player) => {
	const attackDetectCalc = function () {
		const enemyList = document.getElementsByClassName("enemy")
		for (let i = 0; i < enemyList.length; i++) {
			onaraDetector(enemyList[i], player, () => {
				currentScore++
				scoreUpdate()
				enemyList[i].remove()
			})
		}
	}
	window.addEventListener("mousedown", attackDetectCalc)
	return () => window.removeEventListener("mousedown", attackDetectCalc)
}

/////////////////////////////////////////

////////////////
//// item  /////
////////////////

const sprayGetDetector = (player) => {
	const sprayGetDetectCalc = function () {
		const sprayList = document.getElementsByClassName("spray")
		for (let i = 0; i < sprayList.length; i++) {
			onaraDetector(sprayList[i], player, () => {
				causeSprayEffect()
				sprayList[i].remove()
			})
		}
	}
	window.addEventListener("mousedown", sprayGetDetectCalc)
	return () => window.removeEventListener("mousedown", sprayGetDetectCalc)
}

const generateSpray = () => {
	// only one spray appears
	if (document.getElementsByClassName("spray").length > 0) {
		return
	}

	const imgSrc =
		"../../materials/images/oshiri_katori/spray.png"
	const width = 30
	const height = 30
	const noSprayMargin = window.innerWidth * 0.1
	const spraySpeed = 90
	const fps = 60
	const moveDelay = 0.01 //second

	const spray = document.createElement("img")
	document.body.appendChild(spray)

	spray.className = "spray"
	spray.src = imgSrc
	Object.assign(spray.style, {
		width: width + "px",
		height: height + "px"
	})

	// enemy appears from bottom.
	// set no-enemy margin at right and left edge. => margin < sprayLeft < window.innerWidth - margin
	Object.assign(spray.style, {
		top: window.innerHeight + "px",
		left:
			noSprayMargin +
			(window.innerWidth - 2 * noSprayMargin) * Math.random() +
			"px"
	})

	const stopMoveCalc = setInterval(() => {
		let top = spray.getBoundingClientRect().top
		let left = spray.getBoundingClientRect().left
		top -= spraySpeed / fps
		Object.assign(spray.style, {
			top: top + "px",
			left: left + "px"
		})
	}, Math.trunc(1000 / fps))

	Object.assign(spray.style, {
		position: "fixed",
		transition: "left " + moveDelay + "s 0s, top " + moveDelay + "s 0s",
		"-webkit-transition":
			"left " + moveDelay + "s 0s, top " + moveDelay + "s 0s"
	})

	return stopMoveCalc
}

const causeSprayEffect = () => {
	const currentSpeed = speed
	speed = 5
	// set speed as 95% of the one of before
	setTimeout(() => { speed = currentSpeed * 0.95 }, 3000)
}

const yakiimoGetDetector = (player) => {
	const yakiimoGetDetectCalc = function () {
		const yakiimoList = document.getElementsByClassName("yakiimo")
		for (let i = 0; i < yakiimoList.length; i++) {
			onaraDetector(yakiimoList[i], player, () => {
				causeYakiimoEffect()
				yakiimoList[i].remove()
			})
		}
	}
	window.addEventListener("mousedown", yakiimoGetDetectCalc)
	return () => window.removeEventListener("mousedown", yakiimoGetDetectCalc)
}

const generateYakiimo = () => {
	// only one yakiimo appears
	if (document.getElementsByClassName("yakiimo").length > 0) {
		return
	}

	const imgSrc =
		"../../materials/images/oshiri_katori/yakiimo.png"
	const width = 30
	const height = 30
	const noYakiimoMargin = window.innerWidth * 0.1
	const yakiimoSpeed = 90
	const fps = 60
	const moveDelay = 0.01 //second

	const yakiimo = document.createElement("img")
	document.body.appendChild(yakiimo)

	yakiimo.className = "yakiimo"
	yakiimo.src = imgSrc
	Object.assign(yakiimo.style, {
		width: width + "px",
		height: height + "px"
	})

	// enemy appears from bottom.
	// set no-enemy margin at right and left edge. => margin < yakiimoLeft < window.innerWidth - margin
	Object.assign(yakiimo.style, {
		top: window.innerHeight + "px",
		left:
			noYakiimoMargin +
			(window.innerWidth - 2 * noYakiimoMargin) * Math.random() +
			"px"
	})

	const stopMoveCalc = setInterval(() => {
		let top = yakiimo.getBoundingClientRect().top
		let left = yakiimo.getBoundingClientRect().left
		top -= yakiimoSpeed / fps
		Object.assign(yakiimo.style, {
			top: top + "px",
			left: left + "px"
		})
	}, Math.trunc(1000 / fps))

	Object.assign(yakiimo.style, {
		position: "fixed",
		transition: "left " + moveDelay + "s 0s, top " + moveDelay + "s 0s",
		"-webkit-transition":
			"left " + moveDelay + "s 0s, top " + moveDelay + "s 0s"
	})

	return stopMoveCalc
}

const causeYakiimoEffect = () => {
	xRange = originalYrange * 2.5
	yRange = originalYrange * 2.5
	detect = circleDetect

	setTimeout(() => {
		xRange = originalXrange
		yRange = originalYrange
		detect = rectDetect
	}, 5000)

	const height = 900
	const width = 900
	const trackDelay = 0.1 //second
	const ougiImgSrc = "../../materials/images/oshiri_katori/ougi.svg"

	const ougi = document.createElement("img")
	document.body.appendChild(ougi)
	setTimeout(() => { ougi.remove() }, 5000)

	ougi.src = ougiImgSrc
	Object.assign(ougi.style, {
		// prevent flash
		top: window.innerHeight + 10 + "px",
		left: window.innerWidth + 10 + "px",
		width: width + "px",
		height: height + "px",
		"z-index": "-1",
	})

	Object.assign(ougi.style, {
		position: "fixed",
		transition:
			"left " +
			trackDelay +
			"s ease-in-out 0s, top " +
			trackDelay +
			"s ease-in-out 0s",
		"-webkit-transition":
			"left " +
			trackDelay +
			"s ease-in-out 0s, top " +
			trackDelay +
			"s ease-in-out 0s"
	})
	window.addEventListener("mousemove", function (e) {
		ougi.style.left = e.clientX - width / 2 + "px"
		ougi.style.top = e.clientY - height / 2 + "px"
	})
}


/////////////////////////////////////////

///////////////////////
// score calculation //
///////////////////////

var currentScore = 0
var lastScore = 0
var bestScore = 0

const setScoreBoard = () => {
	const scoreBoard = document.createElement("div")
	document.body.appendChild(scoreBoard)

	scoreBoard.className = "scoreBoard"
	Object.assign(scoreBoard.style, {
		backgroundColor: "black",
		color: "white",
		fontSize: "20px"
	})

	scoreUpdate()
}

const scoreUpdate = () => {
	const scoreBoard = document.getElementsByClassName("scoreBoard")[0]

	const currentScoreMessage = "いまのてんすう: " + currentScore + "pt "
	const lastScoreMessage = "さいごのてんすう: " + lastScore + "pt "
	const bestScoreMessage = "さいこうのてんすう: " + bestScore + "pt "
	scoreBoard.innerText =
		currentScoreMessage + lastScoreMessage + bestScoreMessage
}

const getScoreFromCookie = () => {
	const rawCookie = document.cookie

	if (rawCookie !== "") {
		const cookies = rawCookie.split(";")
		for (let i = 0; i < cookies.length; i++) {
			const [key, value] = cookies[i].split("=")
			if (key.trim() === "best_score") {
				bestScore = parseInt(value, 10)
			} else if (key.trim() === "last_score") {
				lastScore = parseInt(value, 10)
			}
		}
	} else {
		bestScore = 0
		lastScore = 0
	}
}

const saveScore = () => {
	const _bestScore = Math.max(bestScore, currentScore)
	document.cookie = "last_score=" + currentScore
	document.cookie = "best_score=" + _bestScore
}

/////////////////////////////////////////

//////////////////////////
// gameover calculation //
//////////////////////////

const gameOverDetector = (safeZoneBorder, functionsToClean, playGameOverSoundEffect) => {
	const stopGameOverDetector = setInterval(() => {
		const enemyList = document.getElementsByClassName("enemy")
		for (let i = 0; i < enemyList.length; i++) {
			if (enemyList[i].getBoundingClientRect().top < safeZoneBorder) {
				playGameOverSoundEffect()
				for (let i = 0; i < functionsToClean.length; i++) {
					functionsToClean[i]()
				}
				displayGameOverMessage()
				clearInterval(stopGameOverDetector)
				saveScore()
				break
			}
		}
	}, 100)
}

const displayGameOverMessage = () => {
	const messagePadding = 10

	const gameOverMessage = document.createElement("div")
	Object.assign(gameOverMessage.style, {
		"text-align": "center",
		padding: messagePadding + "px",
		fontSize: "30px",
		backgroundColor: "red",
		color: "white"
	})
	gameOverMessage.innerText = "GAME OVER"
	document.body.appendChild(gameOverMessage)
	Object.assign(gameOverMessage.style, {
		position: "absolute",
		left: window.innerWidth / 2 - 130 + "px",
		top: (window.innerHeight - gameOverMessage.clientHeight) / 2 + "px"
	})

	const retryButton = document.createElement("button")
	Object.assign(retryButton.style, {
		border: "2pt solid orange",
		fontSize: "20px",
		display: "block"
	})
	retryButton.innerText = "もういちどたたかう……！"
	retryButton.onclick = () => window.location.reload()
	gameOverMessage.appendChild(retryButton)
}

///////////////////////////////////////////

///////////////////////
/// start dialog //////
///////////////////////

const displayStartDialog = () => {
	const dialog = document.createElement("div")
	document.body.appendChild(dialog)
	dialog.id = "dialog"
	Object.assign(dialog.style, {
		border: "2px solid brown",
		fontSize: "14px",
		width: "200px",
		padding: "4px",
		position: "absolute",
		"text-align": "center",
		left: window.innerWidth / 2 - 100 + "px",
		top: window.innerHeight / 2 - 80 + "px"
	})
	dialog.innerText =
		"おしりのくにに、おそろしい「か」たちがやってきた！！\nおしりのくにをまもれ！\n\nクリック:「か」をやっつける\n\n"

	const startButton = document.createElement("button")
	Object.assign(startButton.style, {
		borderColor: "brown",
		borderWidth: "2px",
		fontSize: "14px",
		diplay: "flex",
		alignSprays: "center",
		justifyContent: "center",
		padding: "2px"
	})
	startButton.innerText =
		"たたかいをはじめる……！\nおとがなるよ！\n(おんがく:まおうだましい)"
	dialog.appendChild(startButton)
	startButton.onclick = () => gameStart()
}

/////////////////////////////////////////////

/////////////
/// music ///
/////////////

const playBackgroundMusic = () => {
	const backgroundMusic = document.createElement("audio")
	backgroundMusic.src = "../../materials/audio/oshiri-katori.mp3"
	backgroundMusic.id = "backgroundMusic"
	backgroundMusic.autoplay = true
	document.body.appendChild(backgroundMusic)
	document.getElementById("backgroundMusic").play()
}

const setOnaraSoundEffect = () => {
	const onaraSoundEffect = document.createElement("audio")
	onaraSoundEffect.src = "../../materials/audio/onara.mp3"
	onaraSoundEffect.id = "onaraSoundEffect"
	onaraSoundEffect.autoplay = true
	document.body.appendChild(onaraSoundEffect)
	return () => document.getElementById("onaraSoundEffect").play()
}

const setBombSoundEffect = () => {
	const bombSoundEffect = document.createElement("audio")
	bombSoundEffect.src = "../../materials/audio/bomb.mp3"
	bombSoundEffect.id = "bombSoundEffect"
	bombSoundEffect.autoplay = true
	document.body.appendChild(bombSoundEffect)
	return () => document.getElementById("bombSoundEffect").play()
}


/////////////////////////////////////////////

/////////////////////////////
//// control user action ////
/////////////////////////////

// ban scrolling
const banScroll = () => {
	const prevent = function (e) {
		e.preventDefault()
	}
	document.addEventListener("touchmove", prevent, { passive: false })
	document.addEventListener("mousewheel", prevent, { passive: false })
}

/////////////////////////////////////////////

const gameStart = () => {
	document.getElementById("dialog").remove()

	playBackgroundMusic()
	const playOnaraSoundEffect = setOnaraSoundEffect()
	const playBombSoundEffect = setBombSoundEffect()

	const safeZoneBorder = renderingStage()
	const oshiri = initOshiri(playOnaraSoundEffect)

	const functionsToClean = []

	const stopMosquitoGenerator = setInterval(() => {
		const clearMosquitoMoveCalc = generateMosquito()
		functionsToClean.push(() => clearInterval(clearMosquitoMoveCalc))
	}, 500)
	increaseSpeed()
	functionsToClean.push(() => clearInterval(stopMosquitoGenerator))

	const stopBeeGenerator = setInterval(() => {
		const clearBeeMoveCalc = generateBee()
		functionsToClean.push(() => clearInterval(clearBeeMoveCalc))
	}, 1000)
	functionsToClean.push(() => clearInterval(stopBeeGenerator))

	const cancelAttackDetector = attackDetector(oshiri)
	functionsToClean.push(() => cancelAttackDetector())

	const stopSprayGenerator = setInterval(() => {
		const clearSprayMoveCalc = generateSpray()
		functionsToClean.push(() => clearInterval(clearSprayMoveCalc))
	}, 8500)
	functionsToClean.push(() => clearInterval(stopSprayGenerator))

	const cancelSprayGetDetector = sprayGetDetector(oshiri)
	functionsToClean.push(() => cancelSprayGetDetector())

	const stopYakiimoGenerator = setInterval(() => {
		const clearYakiimoMoveCalc = generateYakiimo()
		functionsToClean.push(() => clearInterval(clearYakiimoMoveCalc))
	}, 8500)
	functionsToClean.push(() => clearInterval(stopYakiimoGenerator))

	const cancelYakiimoGetDetector = yakiimoGetDetector(oshiri)
	functionsToClean.push(() => cancelYakiimoGetDetector())

	gameOverDetector(safeZoneBorder, functionsToClean, playBombSoundEffect)
}

/////////////////////////////////////////////

window.onload = () => {
	banScroll()

	getScoreFromCookie()
	setScoreBoard()
	displayStartDialog()
}
