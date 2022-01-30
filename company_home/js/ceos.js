const images = ['./images/CEOs/boy.jpeg',
                './images/CEOs/dragon.png',
                './images/CEOs/alpaca.jpeg',
                './images/CEOs/girrafe.jpeg',
                './images/CEOs/goat.jpeg'];

const imageNo = Math.floor( Math.random() * images.length);
imageSrc = images[imageNo];
console.log(imageNo, imageSrc)
document.write('<img src="'+imageSrc+'" width="395" height="395">');