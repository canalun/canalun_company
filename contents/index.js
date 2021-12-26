const imageArea = document.getElementById('imageArea');
const images = ['./images/CEOs/boy.jpeg', './images/CEOs/dragon.png'];

const imageNo = Math.floor( Math.random() * images.length)
imageArea.src = images[imageNo];