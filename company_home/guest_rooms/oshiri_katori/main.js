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
		height: height + "px"
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

const onaraDetector = (target, player, callback) => {
	const XRange = 60
	const YRange = 80

	const ex = target.getBoundingClientRect().left
	const ey = target.getBoundingClientRect().top
	const px = player.getBoundingClientRect().left
	const py = player.getBoundingClientRect().top

	if (
		(px - XRange < ex &&
			ex < px + XRange) &&
		(py < ey &&
			ey < py + YRange)
	) {
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

const itemGetDetector = (player) => {
	const itemGetDetectCalc = function () {
		const itemList = document.getElementsByClassName("item")
		for (let i = 0; i < itemList.length; i++) {
			onaraDetector(itemList[i], player, () => {
				causeItemEffect()
				itemList[i].remove()
			})
		}
	}
	window.addEventListener("mousedown", itemGetDetectCalc)
	return () => window.removeEventListener("mousedown", itemGetDetectCalc)
}

const generateItem = () => {
	// only one item appears
	if (document.getElementsByClassName("item").length > 0) {
		return
	}

	const imgSrc =
		"../../materials/images/oshiri_katori/spray.png"
	const width = 30
	const height = 30
	const noItemMargin = window.innerWidth * 0.1
	const itemSpeed = 90
	const fps = 60
	const moveDelay = 0.01 //second

	const item = document.createElement("img")
	document.body.appendChild(item)

	item.className = "item"
	item.src = imgSrc
	Object.assign(item.style, {
		width: width + "px",
		height: height + "px"
	})

	// enemy appears from bottom.
	// set no-enemy margin at right and left edge. => margin < itemLeft < window.innerWidth - margin
	Object.assign(item.style, {
		top: window.innerHeight + "px",
		left:
			noItemMargin +
			(window.innerWidth - 2 * noItemMargin) * Math.random() +
			"px"
	})

	const stopMoveCalc = setInterval(() => {
		let top = item.getBoundingClientRect().top
		let left = item.getBoundingClientRect().left
		top -= itemSpeed / fps
		Object.assign(item.style, {
			top: top + "px",
			left: left + "px"
		})
	}, Math.trunc(1000 / fps))

	Object.assign(item.style, {
		position: "fixed",
		transition: "left " + moveDelay + "s 0s, top " + moveDelay + "s 0s",
		"-webkit-transition":
			"left " + moveDelay + "s 0s, top " + moveDelay + "s 0s"
	})

	return stopMoveCalc
}

const causeItemEffect = () => {
	speed = 5
	setTimeout(() => { speed = 80 }, 3000)
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
		alignItems: "center",
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

	const cancelAttackDetector = attackDetector(oshiri)
	functionsToClean.push(() => cancelAttackDetector())

	const stopItemGenerator = setInterval(() => {
		const clearItemMoveCalc = generateItem()
		functionsToClean.push(() => clearInterval(clearItemMoveCalc))
	}, 8500)
	functionsToClean.push(() => clearInterval(stopItemGenerator))

	const cancelItemGetDetector = itemGetDetector(oshiri)
	functionsToClean.push(() => cancelItemGetDetector())

	gameOverDetector(safeZoneBorder, functionsToClean, playBombSoundEffect)
}

/////////////////////////////////////////////

window.onload = () => {
	banScroll()

	getScoreFromCookie()
	setScoreBoard()
	displayStartDialog()
}
